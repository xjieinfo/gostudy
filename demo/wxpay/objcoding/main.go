package main

import (
	"fmt"
	"github.com/objcoding/wxpay"
	"github.com/tidwall/gjson"
	"github.com/xjieinfo/xjgo/xjcore/xjgorm"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gostudy/demo/pub"
)

var (
	client *wxpay.Client
)

func init() {
	//连接数据库
	Datasource := xjtypes.Datasource{
		Drivername: "mysql",
		Dsn:        "root:root@tcp(localhost:3306)/gostudy?charset=utf8mb4&parseTime=true",
	}
	pub.Gorm = xjgorm.GormInit(Datasource)
	//从数据库中获取配置
	sql := "select * from pay_channel where channel_id='WX_APP'"
	var payChannel = make(map[string]interface{})
	pub.Gorm.Raw(sql).First(&payChannel)
	config := payChannel["config"].(string)
	appId := gjson.Get(config, "appId").String()
	mchId := gjson.Get(config, "mchId").String()
	key := gjson.Get(config, "key").String()
	// 创建支付账户
	account := wxpay.NewAccount(appId, mchId, key, false)
	// 新建微信支付客户端
	client = wxpay.NewClient(account)
}

func main() {
	// 统一下单
	//UnifiedOrder("436577857")
	//订单查询
	OrderQuery("436577857")

	//
	//// 退款
	//params := make(wxpay.Params)
	//params.SetString("out_trade_no", "3568785").
	//	SetString("out_refund_no", "19374568").
	//	SetInt64("total_fee", 1).
	//	SetInt64("refund_fee", 1)
	//p, _ := client.Refund(params)
	//
	//// 退款查询
	//params := make(wxpay.Params)
	//params.SetString("out_refund_no", "3568785")
	//p, _ := client.RefundQuery(params)
	//// 签名
	//signStr := client.Sign(params)
	//
	//// 校验签名
	//b := client.ValidSign(params)
	//// xml解析
	//params := wxpay.XmlToMap(xmlStr)
	//
	//// map封装xml请求参数
	//b := wxpay.MapToXml(params)
	//// 支付或退款返回成功信息
	//return wxpay.Notifies{}.OK()
	//
	//// 支付或退款返回失败信息
	//return wxpay.Notifies{}.NotOK("支付失败或退款失败了")

	println("ok")
}

// 统一下单
func UnifiedOrder(out_trade_no string) {
	// 统一下单
	params := make(wxpay.Params)
	params.SetString("body", "test").
		SetString("out_trade_no", out_trade_no).
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "APP")
	p, _ := client.UnifiedOrder(params)
	fmt.Println(p)
	/*
		map[appid:wxff781548f4e076b1 mch_id:1489559842 nonce_str:1mAsEggu9bBuonx2 prepay_id:wx171623381226911dba79d17569e27b0000 result_code:SUCCESS return_code:SUCCESS return_msg:OK sign:579358C536ABD65C18B2B4EF5AECC842 trade_type:APP]
	*/
}

//订单查询
func OrderQuery(out_trade_no string) {
	// 订单查询
	params := make(wxpay.Params)
	params.SetString("out_trade_no", out_trade_no)
	p, _ := client.OrderQuery(params)
	fmt.Println(p)
	/*
		map[appid:wxff781548f4e076b1 device_info: mch_id:1489559842 nonce_str:woif0CYwXGupuQJZ out_trade_no:436577857 result_code:SUCCESS return_code:SUCCESS return_msg:OK sign:4315C8F183B3E1BDA8CCB00039EC5EB4 total_fee:1 trade_state:NOTPAY trade_state_desc:订单未支付]
	*/
	// 校验签名
	b := client.ValidSign(params)
	println(b)
}
