package ecs

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

// JoinResourceGroup invokes the ecs.JoinResourceGroup API synchronously
// api document: https://help.aliyun.com/api/ecs/joinresourcegroup.html
func (client *Client) JoinResourceGroup(request *JoinResourceGroupRequest) (response *JoinResourceGroupResponse, err error) {
	response = CreateJoinResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// JoinResourceGroupWithChan invokes the ecs.JoinResourceGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/joinresourcegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinResourceGroupWithChan(request *JoinResourceGroupRequest) (<-chan *JoinResourceGroupResponse, <-chan error) {
	responseChan := make(chan *JoinResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinResourceGroup(request)
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

// JoinResourceGroupWithCallback invokes the ecs.JoinResourceGroup API asynchronously
// api document: https://help.aliyun.com/api/ecs/joinresourcegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinResourceGroupWithCallback(request *JoinResourceGroupRequest, callback func(response *JoinResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.JoinResourceGroup(request)
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

// JoinResourceGroupRequest is the request struct for api JoinResourceGroup
type JoinResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
}

// JoinResourceGroupResponse is the response struct for api JoinResourceGroup
type JoinResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateJoinResourceGroupRequest creates a request to invoke JoinResourceGroup API
func CreateJoinResourceGroupRequest() (request *JoinResourceGroupRequest) {
	request = &JoinResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "JoinResourceGroup", "ecs", "openAPI")
	return
}

// CreateJoinResourceGroupResponse creates a response to parse from JoinResourceGroup response
func CreateJoinResourceGroupResponse() (response *JoinResourceGroupResponse) {
	response = &JoinResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
