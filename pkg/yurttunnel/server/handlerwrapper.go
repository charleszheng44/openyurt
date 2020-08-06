/*
Copyright 2020 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"net/http"
	"time"

	"k8s.io/klog"
)

// Middleware takes in one Handler and wrap it within another
type Middleware func(http.Handler) http.Handler

// PrintReqInfoMiddleware prints request information when start/stop
// handling the request
var PrintReqInfoMiddleware = func(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		klog.V(4).Infof("start handling request %s %s, from %s to %s",
			req.Method, req.URL.String(), req.Host, req.RemoteAddr)
		start := time.Now()
		handler.ServeHTTP(w, req)
		klog.V(4).Infof("stop handling request %s %s, request handling lasts %v",
			req.Method, req.URL.String(), time.Now().Sub(start))
	})
}

// WrapWrapHandler wraps the coreHandler with chains
func WrapHandler(coreHandler http.Handler, middlewares ...Middleware) http.Handler {
	handler := coreHandler
	if len(middlewares) == 0 {
		return handler
	}
	for i := len(middlewares) - 1; i > 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
