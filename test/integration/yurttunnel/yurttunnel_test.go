package integration

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	// utildbus "k8s.io/kubernetes/pkg/util/dbus"
	// "k8s.io/kubernetes/pkg/util/iptables"
	// "k8s.io/utils/exec"
	anpserver "sigs.k8s.io/apiserver-network-proxy/pkg/server"

	ta "github.com/alibaba/openyurt/pkg/yurttunnel/agent"
	hw "github.com/alibaba/openyurt/pkg/yurttunnel/handlerwrapper"
	tr "github.com/alibaba/openyurt/pkg/yurttunnel/handlerwrapper/tracerequest"
	ts "github.com/alibaba/openyurt/pkg/yurttunnel/server"
)

const (
	ResponseForRegularRequest = "Fake Server"
	FakeServerPort            = 9515
	ServerMasterPort          = 9516
	ServerMasterInsecurePort  = 9517
	ServerAgentPort           = 9518
	RootCAFile                = "pki/ca.pem"
	ServerCertFile            = "pki/fake-server.pem"
	ServerKeyFile             = "pki/fake-server-key.pem"
	TunnelServerCertFile      = "pki/fake-server.pem"
	TunnelServerCertKeyFile   = "pki/fake-server-key.pem"
	TunnelAgentCertFile       = "pki/fake-client.pem"
	TunnelAgentCertKeyFile    = "pki/fake-client-key.pem"
	ClientCertFile            = "pki/fake-client.pem"
	ClientKeyFile             = "pki/fake-client-key.pem"
	InterceptorServerUDSFile  = "interceptor-proxier.sock"
)

func genCAPool(t *testing.T, CAPath string) *x509.CertPool {
	caCertPEM, err := ioutil.ReadFile(CAPath)
	if err != nil {
		t.Fatalf("fail to load the CA: %v", err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		t.Fatal("fail to append the ca PEM to cert pool")
	}
	return roots
}

func genCert(t *testing.T, certPath, keyPath string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair("pki/fake-server.pem", "pki/fake-server-key.pem")
	if err != nil {
		t.Fatalf("fail to load the fake server cert: %v", err)
	}
	return cert
}

func startFakeServer(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Logf("++++++++++++++++ receive request from %s with URL %s", r.RemoteAddr, r.URL.String())
		w.Write([]byte(ResponseForRegularRequest))
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", FakeServerPort),
		Handler: m,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{
				genCert(t, ServerCertFile, ServerKeyFile),
			},
			ClientCAs: genCAPool(t, RootCAFile),
		},
	}

	t.Log("fake-server is listening on :9515")
	err := s.ListenAndServeTLS("", "")
	if err != nil {
		t.Fatalf("the fake-server failed: %v", err)
	}
}

func startFakeClient(t *testing.T, wg *sync.WaitGroup) {
	defer wg.Done()
	tsURL, err := url.Parse(fmt.Sprintf("https://127.0.0.1:%d", ServerMasterPort))
	if err != nil {
		t.Fatalf("fail to parse the tunnel server url: %v", err)
	}
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{
					genCert(t, ClientCertFile, ClientKeyFile),
				},
				RootCAs: genCAPool(t, RootCAFile),
			},
			Proxy: http.ProxyURL(tsURL),
		},
	}

	r, err := http.NewRequest(http.MethodGet, "https://127.0.0.1:9515", nil)
	if err != nil {
		t.Fatalf("fail to generate http request: %v", err)
	}

	rep, err := c.Do(r)
	if err != nil {
		t.Fatalf("fail to send request to the server: %v", err)
	}
	defer rep.Body.Close()

	bodyByts, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		t.Fatalf("fail to read from the response body: %v", err)
	}
	if string(bodyByts) != ResponseForRegularRequest {
		t.Fatalf("invalid response content, expect: Fake Server, got: %s", string(bodyByts))
	}
	t.Log("successfully send request to the fake server")
}

func startTunnelServer(t *testing.T) {
	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{
			genCert(t, TunnelServerCertFile, TunnelServerCertKeyFile),
		},
		ClientCAs: genCAPool(t, RootCAFile),
	}
	wrappers := hw.HandlerWrappers([]hw.Middleware{
		tr.NewTraceReqMiddleware(),
	})

	_, err := os.Stat(InterceptorServerUDSFile)
	if !os.IsNotExist(err) {
		os.Remove(InterceptorServerUDSFile)
	}

	tunnelServer := ts.NewTunnelServer(
		false,                                /*egressSelectorEnabled*/
		InterceptorServerUDSFile,             /* interceptorServerUDSFile*/
		fmt.Sprintf(":%d", ServerMasterPort), /* serverMasterAddr */
		fmt.Sprintf(":%d", ServerMasterInsecurePort), /* serverMasterInsecureAddr */
		fmt.Sprintf(":%d", ServerAgentPort),          /* serverAgentAddr */
		1,                                            /* serverCount */
		&tlsCfg,                                      /* tlsCfg */
		wrappers,                                     /* hw.HandlerWrappers */
		string(anpserver.ProxyStrategyDestHost),      /* proxyStrategy */
	)
	tunnelServer.Run()
	t.Log("Yurttunnel Server is running")
}

func startTunnelAgent(t *testing.T) {
	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{
			genCert(t, TunnelServerCertFile, TunnelServerCertFile),
		},
		RootCAs:    genCAPool(t, RootCAFile),
		ServerName: "127.0.0.1",
	}
	tunnelAgent := ta.NewTunnelAgent(
		&tlsCfg,                             /* tlsCfg */
		fmt.Sprintf(":%d", ServerAgentPort), /* tunnelServerAddr */
		"fake-agent",                        /* nodeName */
		"127.0.0.1",                         /* agentIdentifiers */
	)
	tunnelAgent.Run(wait.NeverStop)
	t.Log("Yurttunnel Agent is running")
}

func TestYurttunnel(t *testing.T) {
	// exec := exec.New()
	// dbus := utildbus.New()
	// iptInf := iptables.New(exec, dbus, iptables.ProtocolIpv4)
	// iptInf.EnsureRule
	var wg sync.WaitGroup
	// 1. start a fake serve
	go startFakeServer(t)

	// 2. setup a tunnel server
	go startTunnelServer(t)
	time.Sleep(1 * time.Second)

	// // 3. setup a tunnel agent
	go startTunnelAgent(t)
	time.Sleep(1 * time.Second)

	// 4. create a fake client and send requests to the tunnel server
	wg.Add(1)
	go startFakeClient(t, &wg)
	wg.Wait()
}
