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

// DescribeAvailableResourceInfo invokes the ens.DescribeAvailableResourceInfo API synchronously
// api document: https://help.aliyun.com/api/ens/describeavailableresourceinfo.html
func (client *Client) DescribeAvailableResourceInfo(request *DescribeAvailableResourceInfoRequest) (response *DescribeAvailableResourceInfoResponse, err error) {
	response = CreateDescribeAvailableResourceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableResourceInfoWithChan invokes the ens.DescribeAvailableResourceInfo API asynchronously
// api document: https://help.aliyun.com/api/ens/describeavailableresourceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableResourceInfoWithChan(request *DescribeAvailableResourceInfoRequest) (<-chan *DescribeAvailableResourceInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableResourceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableResourceInfo(request)
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

// DescribeAvailableResourceInfoWithCallback invokes the ens.DescribeAvailableResourceInfo API asynchronously
// api document: https://help.aliyun.com/api/ens/describeavailableresourceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableResourceInfoWithCallback(request *DescribeAvailableResourceInfoRequest, callback func(response *DescribeAvailableResourceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableResourceInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableResourceInfo(request)
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

// DescribeAvailableResourceInfoRequest is the request struct for api DescribeAvailableResourceInfo
type DescribeAvailableResourceInfoRequest struct {
	*requests.RpcRequest
	Version string `position:"Query" name:"Version"`
}

// DescribeAvailableResourceInfoResponse is the response struct for api DescribeAvailableResourceInfo
type DescribeAvailableResourceInfoResponse struct {
	*responses.BaseResponse
	RequestId        string                                          `json:"RequestId" xml:"RequestId"`
	SupportResources SupportResourcesInDescribeAvailableResourceInfo `json:"SupportResources" xml:"SupportResources"`
	Images           ImagesInDescribeAvailableResourceInfo           `json:"Images" xml:"Images"`
}

// CreateDescribeAvailableResourceInfoRequest creates a request to invoke DescribeAvailableResourceInfo API
func CreateDescribeAvailableResourceInfoRequest() (request *DescribeAvailableResourceInfoRequest) {
	request = &DescribeAvailableResourceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeAvailableResourceInfo", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableResourceInfoResponse creates a response to parse from DescribeAvailableResourceInfo response
func CreateDescribeAvailableResourceInfoResponse() (response *DescribeAvailableResourceInfoResponse) {
	response = &DescribeAvailableResourceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
