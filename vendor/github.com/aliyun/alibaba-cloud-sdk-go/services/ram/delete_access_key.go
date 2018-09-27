package ram

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

// DeleteAccessKey invokes the ram.DeleteAccessKey API synchronously
// api document: https://help.aliyun.com/api/ram/deleteaccesskey.html
func (client *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (response *DeleteAccessKeyResponse, err error) {
	response = CreateDeleteAccessKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAccessKeyWithChan invokes the ram.DeleteAccessKey API asynchronously
// api document: https://help.aliyun.com/api/ram/deleteaccesskey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAccessKeyWithChan(request *DeleteAccessKeyRequest) (<-chan *DeleteAccessKeyResponse, <-chan error) {
	responseChan := make(chan *DeleteAccessKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAccessKey(request)
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

// DeleteAccessKeyWithCallback invokes the ram.DeleteAccessKey API asynchronously
// api document: https://help.aliyun.com/api/ram/deleteaccesskey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAccessKeyWithCallback(request *DeleteAccessKeyRequest, callback func(response *DeleteAccessKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAccessKeyResponse
		var err error
		defer close(result)
		response, err = client.DeleteAccessKey(request)
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

// DeleteAccessKeyRequest is the request struct for api DeleteAccessKey
type DeleteAccessKeyRequest struct {
	*requests.RpcRequest
	UserName        string `position:"Query" name:"UserName"`
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
}

// DeleteAccessKeyResponse is the response struct for api DeleteAccessKey
type DeleteAccessKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAccessKeyRequest creates a request to invoke DeleteAccessKey API
func CreateDeleteAccessKeyRequest() (request *DeleteAccessKeyRequest) {
	request = &DeleteAccessKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeleteAccessKey", "", "")
	return
}

// CreateDeleteAccessKeyResponse creates a response to parse from DeleteAccessKey response
func CreateDeleteAccessKeyResponse() (response *DeleteAccessKeyResponse) {
	response = &DeleteAccessKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}