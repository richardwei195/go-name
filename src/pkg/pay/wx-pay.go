package pay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"time"
)

var (
	appid                      string = "wxd678efh567hg6787"                       // 小程序appid
	mchID                      string = "190000****"                               // 商户号
	mchCertificateSerialNumber string = "3775B6A45ACD588826D15E583A95F5DD********" // 商户证书序列号
	mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
)

// PrePay 微信支付生成预支付订单
func PrePay(amount int64) (rsp *jsapi.PrepayResponse, err error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(appid),
			Mchid:       core.String(mchID),
			Description: core.String("品质起名-增值服务"),
			OutTradeNo:  core.String("1217752501201407033233368018"),
			TimeExpire:  core.Time(time.Now()),
			//Attach:      core.String("name-pay"),
			NotifyUrl: core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			//GoodsTag:      core.String("WXG"),
			//LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(amount),
			},
			//支付者信息
			Payer: &jsapi.Payer{
				Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
			},
			//优惠功能
			//Detail: &jsapi.Detail{
			//	CostPrice: core.Int64(608800),
			//	GoodsDetail: []jsapi.GoodsDetail{jsapi.GoodsDetail{
			//		GoodsName:        core.String("iPhoneX 256G"),
			//		MerchantGoodsId:  core.String("ABC"),
			//		Quantity:         core.Int64(1),
			//		UnitPrice:        core.Int64(828800),
			//		WechatpayGoodsId: core.String("1001"),
			//	}},
			//	InvoiceId: core.String("wx123"),
			//},
			//支付场景
			//SceneInfo: &jsapi.SceneInfo{
			//	DeviceId:      core.String("013467007045764"),
			//	PayerClientIp: core.String("14.23.150.211"),
			//	StoreInfo: &jsapi.StoreInfo{
			//		Address:  core.String("广东省深圳市南山区科技中一道10000号"),
			//		AreaCode: core.String("440305"),
			//		Id:       core.String("0001"),
			//		Name:     core.String("腾讯大厦分店"),
			//	},
			//},
			SettleInfo: &jsapi.SettleInfo{
				ProfitSharing: core.Bool(false),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)

	}
	return resp, err
}
