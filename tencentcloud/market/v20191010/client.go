// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20191010

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-10"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewGetCateTreeRequest() (request *GetCateTreeRequest) {
    request = &GetCateTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("market", APIVersion, "GetCateTree")
    return
}

func NewGetCateTreeResponse() (response *GetCateTreeResponse) {
    response = &GetCateTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCateTree
// 获取分类名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetCateTree(request *GetCateTreeRequest) (response *GetCateTreeResponse, err error) {
    if request == nil {
        request = NewGetCateTreeRequest()
    }
    response = NewGetCateTreeResponse()
    err = c.Send(request, response)
    return
}

func NewGetUsagePlanUsageAmountRequest() (request *GetUsagePlanUsageAmountRequest) {
    request = &GetUsagePlanUsageAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("market", APIVersion, "GetUsagePlanUsageAmount")
    return
}

func NewGetUsagePlanUsageAmountResponse() (response *GetUsagePlanUsageAmountResponse) {
    response = &GetUsagePlanUsageAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUsagePlanUsageAmount
// 该接口可以根据InstanceId查询实例的api的使用情况。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetUsagePlanUsageAmount(request *GetUsagePlanUsageAmountRequest) (response *GetUsagePlanUsageAmountResponse, err error) {
    if request == nil {
        request = NewGetUsagePlanUsageAmountRequest()
    }
    response = NewGetUsagePlanUsageAmountResponse()
    err = c.Send(request, response)
    return
}
