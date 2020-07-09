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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type GetComponentAndVersionsRequest struct {
	*tchttp.BaseRequest

	// Component Name
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

func (r *GetComponentAndVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetComponentAndVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetComponentAndVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComponentAndVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetComponentAndVersionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetComponentVersionRequest struct {
	*tchttp.BaseRequest

	// Component Name
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// Component Version
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
}

func (r *GetComponentVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetComponentVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetComponentVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetComponentVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetComponentVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// App Name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stage Name
	StageName *string `json:"StageName,omitempty" name:"StageName"`

	// Instance Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *GetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPackageRequest struct {
	*tchttp.BaseRequest

	// Package Name
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// Package Version
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

func (r *GetPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadUrlsRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`
}

func (r *GetUploadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadUrlsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUploadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadUrlsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListComponentsRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`
}

func (r *ListComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListComponentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListComponentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListComponentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListInstancesRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// App Name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stage Name
	StageName *string `json:"StageName,omitempty" name:"StageName"`
}

func (r *ListInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPackagesRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`
}

func (r *ListPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPackagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPackagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostPublishComponentRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// Component Name
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// Component Version
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
}

func (r *PostPublishComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostPublishComponentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostPublishComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostPublishComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostPublishComponentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostPublishPackageRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`
}

func (r *PostPublishPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostPublishPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostPublishPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PostPublishPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PostPublishPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PrePublishComponentRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// Component Name
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// Component Version
	ComponentVersion *string `json:"ComponentVersion,omitempty" name:"ComponentVersion"`
}

func (r *PrePublishComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PrePublishComponentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PrePublishComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PrePublishComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PrePublishComponentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreparePublishPackageRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`
}

func (r *PreparePublishPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreparePublishPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreparePublishPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreparePublishPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreparePublishPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunComponentRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// App Name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stage Name
	StageName *string `json:"StageName,omitempty" name:"StageName"`

	// Instance Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Channel URL
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// Role Name
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *RunComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunComponentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunComponentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunFinishComponentRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// App Name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stage Name
	StageName *string `json:"StageName,omitempty" name:"StageName"`

	// Instance Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *RunFinishComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunFinishComponentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunFinishComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunFinishComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunFinishComponentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaveInstanceRequest struct {
	*tchttp.BaseRequest

	// JSON stringified object
	Body *string `json:"Body,omitempty" name:"Body"`

	// App Name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Stage  Name
	StageName *string `json:"StageName,omitempty" name:"StageName"`

	// Instance Name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *SaveInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SaveInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON stringified object containing response payload.
		Body *string `json:"Body,omitempty" name:"Body"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SaveInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SaveInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendCouponRequest struct {
	*tchttp.BaseRequest

	// 发送代金券类型(活动tag)
	Type []*string `json:"Type,omitempty" name:"Type" list`
}

func (r *SendCouponRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCouponRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendCouponResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 错误代码,为0成功
		ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCouponResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendCouponResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
