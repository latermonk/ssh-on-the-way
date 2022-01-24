package main

import (
	"cloud/utils"
	"fmt"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {

	var (
		secretId  = "AKIDJRM2Tf1Mr78A4jpIp8F7pLFuKVRAkvxk"
		secretKey = "iBQTHqDLEsHQL0QdbLr9CzXqzta2FCuM"

		host = "cvm.tencentcloudapi.com"

		region = "ap-chengdu"

		service = "cvm"

		action = "RebootInstances"
		body   = `{"InstanceIds" : ["ins-a3llvb55"]}`

		now     = time.Now()
		version = "2017-03-12"
	)

	authorization := utils.TencentAPISignature(secretId, secretKey, host, service, body, now)

	fmt.Println("curl cmd: ")
	curl := `curl -XPOST https://%s \
	-H "Authorization: %s" \
	-H "Content-Type: application/json; charset=utf-8" \
	-H "Host: %s" \
	-H "X-TC-Action: %s" \
	-H "X-TC-Timestamp: %d" \
	-H "X-TC-Version: %s" \
	-H "X-TC-Region: %s" \
	-d '%s'`
	fmt.Printf(curl, host, authorization, host, action, now.Unix(), version, region, body)
	fmt.Println()

	// 	fmt.Println("go req:")
	// 	header := req.Header{}
	// 	header["Authorization"] = authorization
	// 	header["Content-Type"] = "application/json; charset=utf-8"
	// 	header["Host"] = host
	// 	header["X-TC-Action"] = action
	// 	header["X-TC-Timestamp"] = strconv.FormatInt(now.Unix(), 10)
	// 	header["X-TC-Version"] = version
	// 	header["X-TC-Region"] = region

	// 	response, _ := req.Post(fmt.Sprintf("https://%s", host), body, header)
	// 	fmt.Println(response.ToString())

	fmt.Println("go sdk:")
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = host

	client, _ := cvm.NewClient(credential, region, cpf)

	request := cvm.NewDescribeRegionsRequest()
	request.FromJsonString(body)

	response2, _ := client.DescribeRegions(request)
	fmt.Println(response2.ToJsonString())
}
