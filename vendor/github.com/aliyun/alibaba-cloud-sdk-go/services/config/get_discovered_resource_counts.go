package config

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

// GetDiscoveredResourceCounts invokes the config.GetDiscoveredResourceCounts API synchronously
// api document: https://help.aliyun.com/api/config/getdiscoveredresourcecounts.html
func (client *Client) GetDiscoveredResourceCounts(request *GetDiscoveredResourceCountsRequest) (response *GetDiscoveredResourceCountsResponse, err error) {
	response = CreateGetDiscoveredResourceCountsResponse()
	err = client.DoAction(request, response)
	return
}

// GetDiscoveredResourceCountsWithChan invokes the config.GetDiscoveredResourceCounts API asynchronously
// api document: https://help.aliyun.com/api/config/getdiscoveredresourcecounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDiscoveredResourceCountsWithChan(request *GetDiscoveredResourceCountsRequest) (<-chan *GetDiscoveredResourceCountsResponse, <-chan error) {
	responseChan := make(chan *GetDiscoveredResourceCountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDiscoveredResourceCounts(request)
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

// GetDiscoveredResourceCountsWithCallback invokes the config.GetDiscoveredResourceCounts API asynchronously
// api document: https://help.aliyun.com/api/config/getdiscoveredresourcecounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDiscoveredResourceCountsWithCallback(request *GetDiscoveredResourceCountsRequest, callback func(response *GetDiscoveredResourceCountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDiscoveredResourceCountsResponse
		var err error
		defer close(result)
		response, err = client.GetDiscoveredResourceCounts(request)
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

// GetDiscoveredResourceCountsRequest is the request struct for api GetDiscoveredResourceCounts
type GetDiscoveredResourceCountsRequest struct {
	*requests.RpcRequest
	MultiAccount requests.Boolean `position:"Query" name:"MultiAccount"`
	GroupByKey   string           `position:"Query" name:"GroupByKey"`
	MemberId     requests.Integer `position:"Query" name:"MemberId"`
}

// GetDiscoveredResourceCountsResponse is the response struct for api GetDiscoveredResourceCounts
type GetDiscoveredResourceCountsResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	GroupedResourceCounts GroupedResourceCounts `json:"GroupedResourceCounts" xml:"GroupedResourceCounts"`
}

// CreateGetDiscoveredResourceCountsRequest creates a request to invoke GetDiscoveredResourceCounts API
func CreateGetDiscoveredResourceCountsRequest() (request *GetDiscoveredResourceCountsRequest) {
	request = &GetDiscoveredResourceCountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "GetDiscoveredResourceCounts", "config", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetDiscoveredResourceCountsResponse creates a response to parse from GetDiscoveredResourceCounts response
func CreateGetDiscoveredResourceCountsResponse() (response *GetDiscoveredResourceCountsResponse) {
	response = &GetDiscoveredResourceCountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}