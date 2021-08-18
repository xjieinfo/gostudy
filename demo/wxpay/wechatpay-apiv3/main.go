package main

/*
基于微信官方提供的库（微信支付API v3）
https://github.com/wechatpay-apiv3/wechatpay-go
*/
import (
	"context"
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
)

func main() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                      string          // 商户号
		mchCertificateSerialNumber string          // 商户证书序列号
		mchPrivateKey              *rsa.PrivateKey // 商户私钥
		mchAPIv3Key                string          // 商户APIv3密钥
		customHTTPClient           *http.Client    // 可选，自定义客户端实例
	)

	ctx := context.Background()
	opts := []core.ClientOption{
		// 一次性设置 签名/验签/敏感字段加解密，并注册 平台证书下载器，自动定时获取最新的平台证书
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
		// 设置自定义 HTTPClient 实例，不设置时默认使用 http.Client{}，并设置超时时间为 30s
		option.WithHTTPClient(customHTTPClient),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}
