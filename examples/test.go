package main

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sls "github.com/serverless-tencent/servereless-sdk-go/v20200205"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	client, _ := sls.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	req := sls.NewGetComponentAndVersionsRequest()
	name := "component-name"
	req.ComponentName = &name
	
	resp, err := client.GetComponentAndVersions(req)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %+v\n", resp)
}
