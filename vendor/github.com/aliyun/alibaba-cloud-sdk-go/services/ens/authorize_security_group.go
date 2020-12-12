package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AuthorizeSecurityGroup invokes the ens.AuthorizeSecurityGroup API synchronously
// api document: https://help.aliyun.com/api/ens/authorizesecuritygroup.html
func (client *Client) AuthorizeSecurityGroup(request *AuthorizeSecurityGroupRequest) (response *AuthorizeSecurityGroupResponse, err error) {
	response = CreateAuthorizeSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeSecurityGroupWithChan invokes the ens.AuthorizeSecurityGroup API asynchronously
// api document: https://help.aliyun.com/api/ens/authorizesecuritygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeSecurityGroupWithChan(request *AuthorizeSecurityGroupRequest) (<-chan *AuthorizeSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *AuthorizeSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeSecurityGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AuthorizeSecurityGroupWithCallback invokes the ens.AuthorizeSecurityGroup API asynchronously
// api document: https://help.aliyun.com/api/ens/authorizesecuritygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeSecurityGroupWithCallback(request *AuthorizeSecurityGroupRequest, callback func(response *AuthorizeSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeSecurityGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AuthorizeSecurityGroupRequest is the request struct for api AuthorizeSecurityGroup
type AuthorizeSecurityGroupRequest struct {
	*requests.RpcRequest
	SourcePortRange string           `position:"Query" name:"SourcePortRange"`
	SecurityGroupId string           `position:"Query" name:"SecurityGroupId"`
	Policy          string           `position:"Query" name:"Policy"`
	PortRange       string           `position:"Query" name:"PortRange"`
	IpProtocol      string           `position:"Query" name:"IpProtocol"`
	SourceCidrIp    string           `position:"Query" name:"SourceCidrIp"`
	Priority        requests.Integer `position:"Query" name:"Priority"`
	Version         string           `position:"Query" name:"Version"`
}

// AuthorizeSecurityGroupResponse is the response struct for api AuthorizeSecurityGroup
type AuthorizeSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAuthorizeSecurityGroupRequest creates a request to invoke AuthorizeSecurityGroup API
func CreateAuthorizeSecurityGroupRequest() (request *AuthorizeSecurityGroupRequest) {
	request = &AuthorizeSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "AuthorizeSecurityGroup", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAuthorizeSecurityGroupResponse creates a response to parse from AuthorizeSecurityGroup response
func CreateAuthorizeSecurityGroupResponse() (response *AuthorizeSecurityGroupResponse) {
	response = &AuthorizeSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
