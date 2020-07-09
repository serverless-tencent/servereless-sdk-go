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

package v20200205

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-05"

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


func NewGetComponentAndVersionsRequest() (request *GetComponentAndVersionsRequest) {
    request = &GetComponentAndVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "GetComponentAndVersions")
    return
}

func NewGetComponentAndVersionsResponse() (response *GetComponentAndVersionsResponse) {
    response = &GetComponentAndVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口获取指定Component的所有版本信息
func (c *Client) GetComponentAndVersions(request *GetComponentAndVersionsRequest) (response *GetComponentAndVersionsResponse, err error) {
    if request == nil {
        request = NewGetComponentAndVersionsRequest()
    }
    response = NewGetComponentAndVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetComponentVersionRequest() (request *GetComponentVersionRequest) {
    request = &GetComponentVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "GetComponentVersion")
    return
}

func NewGetComponentVersionResponse() (response *GetComponentVersionResponse) {
    response = &GetComponentVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定name和版本的Component信息
func (c *Client) GetComponentVersion(request *GetComponentVersionRequest) (response *GetComponentVersionResponse, err error) {
    if request == nil {
        request = NewGetComponentVersionRequest()
    }
    response = NewGetComponentVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceRequest() (request *GetInstanceRequest) {
    request = &GetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "GetInstance")
    return
}

func NewGetInstanceResponse() (response *GetInstanceResponse) {
    response = &GetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  用户获取一个已部署Component的Instance
func (c *Client) GetInstance(request *GetInstanceRequest) (response *GetInstanceResponse, err error) {
    if request == nil {
        request = NewGetInstanceRequest()
    }
    response = NewGetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetPackageRequest() (request *GetPackageRequest) {
    request = &GetPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "GetPackage")
    return
}

func NewGetPackageResponse() (response *GetPackageResponse) {
    response = &GetPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Package的详细信息
func (c *Client) GetPackage(request *GetPackageRequest) (response *GetPackageResponse, err error) {
    if request == nil {
        request = NewGetPackageRequest()
    }
    response = NewGetPackageResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadUrlsRequest() (request *GetUploadUrlsRequest) {
    request = &GetUploadUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "GetUploadUrls")
    return
}

func NewGetUploadUrlsResponse() (response *GetUploadUrlsResponse) {
    response = &GetUploadUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户获取Component Instance的预签名URL链接
func (c *Client) GetUploadUrls(request *GetUploadUrlsRequest) (response *GetUploadUrlsResponse, err error) {
    if request == nil {
        request = NewGetUploadUrlsRequest()
    }
    response = NewGetUploadUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewListComponentsRequest() (request *ListComponentsRequest) {
    request = &ListComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "ListComponents")
    return
}

func NewListComponentsResponse() (response *ListComponentsResponse) {
    response = &ListComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户获取Component的列表信息
func (c *Client) ListComponents(request *ListComponentsRequest) (response *ListComponentsResponse, err error) {
    if request == nil {
        request = NewListComponentsRequest()
    }
    response = NewListComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewListInstancesRequest() (request *ListInstancesRequest) {
    request = &ListInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "ListInstances")
    return
}

func NewListInstancesResponse() (response *ListInstancesResponse) {
    response = &ListInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户获取一个已部署Component的Instance列表
func (c *Client) ListInstances(request *ListInstancesRequest) (response *ListInstancesResponse, err error) {
    if request == nil {
        request = NewListInstancesRequest()
    }
    response = NewListInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListPackagesRequest() (request *ListPackagesRequest) {
    request = &ListPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "ListPackages")
    return
}

func NewListPackagesResponse() (response *ListPackagesResponse) {
    response = &ListPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取Package的列表信息
func (c *Client) ListPackages(request *ListPackagesRequest) (response *ListPackagesResponse, err error) {
    if request == nil {
        request = NewListPackagesRequest()
    }
    response = NewListPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewPostPublishComponentRequest() (request *PostPublishComponentRequest) {
    request = &PostPublishComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "PostPublishComponent")
    return
}

func NewPostPublishComponentResponse() (response *PostPublishComponentResponse) {
    response = &PostPublishComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布一个指定name和version的Component
func (c *Client) PostPublishComponent(request *PostPublishComponentRequest) (response *PostPublishComponentResponse, err error) {
    if request == nil {
        request = NewPostPublishComponentRequest()
    }
    response = NewPostPublishComponentResponse()
    err = c.Send(request, response)
    return
}

func NewPostPublishPackageRequest() (request *PostPublishPackageRequest) {
    request = &PostPublishPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "PostPublishPackage")
    return
}

func NewPostPublishPackageResponse() (response *PostPublishPackageResponse) {
    response = &PostPublishPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发布一个指定name和version的Package
func (c *Client) PostPublishPackage(request *PostPublishPackageRequest) (response *PostPublishPackageResponse, err error) {
    if request == nil {
        request = NewPostPublishPackageRequest()
    }
    response = NewPostPublishPackageResponse()
    err = c.Send(request, response)
    return
}

func NewPrePublishComponentRequest() (request *PrePublishComponentRequest) {
    request = &PrePublishComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "PrePublishComponent")
    return
}

func NewPrePublishComponentResponse() (response *PrePublishComponentResponse) {
    response = &PrePublishComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预发布一个指定name和version的Component
func (c *Client) PrePublishComponent(request *PrePublishComponentRequest) (response *PrePublishComponentResponse, err error) {
    if request == nil {
        request = NewPrePublishComponentRequest()
    }
    response = NewPrePublishComponentResponse()
    err = c.Send(request, response)
    return
}

func NewPreparePublishPackageRequest() (request *PreparePublishPackageRequest) {
    request = &PreparePublishPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "PreparePublishPackage")
    return
}

func NewPreparePublishPackageResponse() (response *PreparePublishPackageResponse) {
    response = &PreparePublishPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 预发布一个指定name和version的Package
func (c *Client) PreparePublishPackage(request *PreparePublishPackageRequest) (response *PreparePublishPackageResponse, err error) {
    if request == nil {
        request = NewPreparePublishPackageRequest()
    }
    response = NewPreparePublishPackageResponse()
    err = c.Send(request, response)
    return
}

func NewRunComponentRequest() (request *RunComponentRequest) {
    request = &RunComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "RunComponent")
    return
}

func NewRunComponentResponse() (response *RunComponentResponse) {
    response = &RunComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运行一个Component
func (c *Client) RunComponent(request *RunComponentRequest) (response *RunComponentResponse, err error) {
    if request == nil {
        request = NewRunComponentRequest()
    }
    response = NewRunComponentResponse()
    err = c.Send(request, response)
    return
}

func NewRunFinishComponentRequest() (request *RunFinishComponentRequest) {
    request = &RunFinishComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "RunFinishComponent")
    return
}

func NewRunFinishComponentResponse() (response *RunFinishComponentResponse) {
    response = &RunFinishComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 运行完成Component，更新Component Instance信息
func (c *Client) RunFinishComponent(request *RunFinishComponentRequest) (response *RunFinishComponentResponse, err error) {
    if request == nil {
        request = NewRunFinishComponentRequest()
    }
    response = NewRunFinishComponentResponse()
    err = c.Send(request, response)
    return
}

func NewSaveInstanceRequest() (request *SaveInstanceRequest) {
    request = &SaveInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "SaveInstance")
    return
}

func NewSaveInstanceResponse() (response *SaveInstanceResponse) {
    response = &SaveInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  用户保存一个已部署Component的Instance
func (c *Client) SaveInstance(request *SaveInstanceRequest) (response *SaveInstanceResponse, err error) {
    if request == nil {
        request = NewSaveInstanceRequest()
    }
    response = NewSaveInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSendCouponRequest() (request *SendCouponRequest) {
    request = &SendCouponRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sls", APIVersion, "SendCoupon")
    return
}

func NewSendCouponResponse() (response *SendCouponResponse) {
    response = &SendCouponResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 发送代金券
func (c *Client) SendCoupon(request *SendCouponRequest) (response *SendCouponResponse, err error) {
    if request == nil {
        request = NewSendCouponRequest()
    }
    response = NewSendCouponResponse()
    err = c.Send(request, response)
    return
}
