package cdn

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

// AddLivePullStreamInfo invokes the cdn.AddLivePullStreamInfo API synchronously
// api document: https://help.aliyun.com/api/cdn/addlivepullstreaminfo.html
func (client *Client) AddLivePullStreamInfo(request *AddLivePullStreamInfoRequest) (response *AddLivePullStreamInfoResponse, err error) {
	response = CreateAddLivePullStreamInfoResponse()
	err = client.DoAction(request, response)
	return
}

// AddLivePullStreamInfoWithChan invokes the cdn.AddLivePullStreamInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/addlivepullstreaminfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLivePullStreamInfoWithChan(request *AddLivePullStreamInfoRequest) (<-chan *AddLivePullStreamInfoResponse, <-chan error) {
	responseChan := make(chan *AddLivePullStreamInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLivePullStreamInfo(request)
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

// AddLivePullStreamInfoWithCallback invokes the cdn.AddLivePullStreamInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/addlivepullstreaminfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLivePullStreamInfoWithCallback(request *AddLivePullStreamInfoRequest, callback func(response *AddLivePullStreamInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLivePullStreamInfoResponse
		var err error
		defer close(result)
		response, err = client.AddLivePullStreamInfo(request)
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

// AddLivePullStreamInfoRequest is the request struct for api AddLivePullStreamInfo
type AddLivePullStreamInfoRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
	SourceUrl     string           `position:"Query" name:"SourceUrl"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

// AddLivePullStreamInfoResponse is the response struct for api AddLivePullStreamInfo
type AddLivePullStreamInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLivePullStreamInfoRequest creates a request to invoke AddLivePullStreamInfo API
func CreateAddLivePullStreamInfoRequest() (request *AddLivePullStreamInfoRequest) {
	request = &AddLivePullStreamInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "AddLivePullStreamInfo", "", "")
	return
}

// CreateAddLivePullStreamInfoResponse creates a response to parse from AddLivePullStreamInfo response
func CreateAddLivePullStreamInfoResponse() (response *AddLivePullStreamInfoResponse) {
	response = &AddLivePullStreamInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
