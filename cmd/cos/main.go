package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, err := url.Parse("https://coolcar-1258228056.cos.ap-beijing.myqcloud.com");

	if err != nil {
		fmt.Errorf("ss", err);
	}
	secId := "AKIDDHIJ8gVBetsdKsGWbR9lDVas0xhBRWQ7";
	secKey := "Ws22yobGfx51wjNwOit13xnZGwdVEm6L";
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(secId),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv(secKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	name := "abc.png";
	// 获取预签名URL
	presignedURL, err := client.Object.GetPresignedURL(context.Background(), http.MethodPut, name, secId, secKey, 20 * time.Second, nil)
	if err != nil {
					panic(err)
	}
	fmt.Printf("sdds %s", presignedURL)
}