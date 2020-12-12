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

// DescribeInstanceTypes invokes the ens.DescribeInstanceTypes API synchronously
// api document: https://help.aliyun.com/api/ens/describeinstancetypes.html
func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
	response = CreateDescribeInstanceTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTypesWithChan invokes the ens.DescribeInstanceTypes API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstancetypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTypesWithChan(request *DescribeInstanceTypesRequest) (<-chan *DescribeInstanceTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTypes(request)
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

// DescribeInstanceTypesWithCallback invokes the ens.DescribeInstanceTypes API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstancetypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTypesWithCallback(request *DescribeInstanceTypesRequest, callback func(response *DescribeInstanceTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTypes(request)
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

// DescribeInstanceTypesRequest is the request struct for api DescribeInstanceTypes
type DescribeInstanceTypesRequest struct {
	*requests.RpcRequest
	Version string `position:"Query" name:"Version"`
}

// DescribeInstanceTypesResponse is the response struct for api DescribeInstanceTypes
type DescribeInstanceTypesResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Code          int           `json:"Code" xml:"Code"`
	InstanceTypes InstanceTypes `json:"InstanceTypes" xml:"InstanceTypes"`
}

// CreateDescribeInstanceTypesRequest creates a request to invoke DescribeInstanceTypes API
func CreateDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
	request = &DescribeInstanceTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeInstanceTypes", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceTypesResponse creates a response to parse from DescribeInstanceTypes response
func CreateDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
	response = &DescribeInstanceTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
