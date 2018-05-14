package payment_web

import (
	"encoding/json"
	"fmt"

	//"strings"
	"strconv"
)

type AliyPay struct {
	Payment
}

const (
	APAY_URL   = "https://openapi.alipay.com/gateway.do"
	KEY        = "key"
	RETURN_URL = "baidu.com"
)

// 订单单数据包
type OrderJSON struct {
	Appid       string  `json:"app_id"`       // 支付宝分配给开发者的应用ID
	Method      string  `json:"method"`       // 接口名称 alipay.trade.page.pay
	Charset     string  `json:"charset"`      // 请求使用的编码格式 utf-8
	SignType    string  `json:"sign_type"`    // 商户生成签名字符串所使用的签名算法类型
	Sign        string  `json:"sign"`         // 商户请求参数的签名串
	TimeStamp   string  `json:"timestamp"`    // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version     string  `json:"version"`      // 调用的接口版本，固定为：1.0
	Bizcontent  string  `json:"biz_content"`  // 业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	OrderID     string  `json:"out_trade_no"` // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	ProductCode string  `json:"product_code"` // 销售产品码，与支付宝签约的产品码名称
	Totalamount float32 `json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位
	Subject     string  `json:"subject"`      // 商品名称 订单标题
	Body        string  `json:"body"`         // 商品描述
	// ReturnUrl   string  `json:"return_url"`   // 同步返回地址，HTTP/HTTPS开头字符串
	// ClientIP   string `json:"exter_invoke_ip"`
	// SellerID  string  `json:"seller_id"` // 卖家支付宝用户号
	// TotalFee  float32 `json:"total_fee"`
	// NotifyUrl string  `json:"notify_url"` // 支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	// PaymentType int     `json:"payment_type"`
	// ExtraParam  string  `json:"extra_common_param"`
	// Service     string  `json:"service"`
	// Partner     string  `json:"partner"`
	// InputChar   string  `json:"_input_charset"`
}

func NewAlipay(params map[string]interface{}) *AliyPay {
	return &AliyPay{Payment{Config: params}}
}

func (this *AliyPay) GenderPayUrl(order Order) (string, error) {
	var o OrderJSON
	o.Appid = this.getConfig("app_id").(string)
	o.Method = "alipay.trade.page.pay"
	o.Charset = "utf-8"
	o.Subject = order.ProudctName
	o.Body = order.ProudctDescription
	o.OrderID = order.OrderID

	// if sellerId := this.getConfig(SELLER); sellerId == nil {
	// 	o.SellerID = this.getConfig(PARTNER).(string)
	// } else {
	// 	o.SellerID = sellerId.(string)
	// }
	o.Totalamount = float32(order.PriceTotal) / 100
	params := this.struct2map(o)

	o.Sign = this.makeSign(params)
	o.SignType = "MD5"
	urlPrams := this.struct2map(o)
	return fmt.Sprintf("%s?%s", APAY_URL, this.makeUrl(urlPrams)), nil

}
func (this *AliyPay) PayNotify(notify []byte) (*NofiyData, error) {

	if len(notify) == 0 {
		return nil, fmt.Errorf("Notify Data empty")
	}
	var resultData NofiyData
	var notifyData map[string]interface{}
	notifyMap := make(map[string]interface{}, len(notifyData))
	json.Unmarshal(notify, &notifyData)
	for k, v := range notifyData {
		if k == "sign_type" {
			continue
		}
		if item, ok := v.([]interface{}); ok {
			notifyMap[k] = item[0]
		}

	}
	if notifyMap["trade_status"].(string) != "TRADE_SUCCESS" {
		resultData.ReturnData = []byte("Status Error")
		return &resultData, fmt.Errorf("Status Error")
	}
	if !this.checkSign(notifyMap) {

		resultData.ReturnData = []byte("Sign Error")
		return &resultData, fmt.Errorf("Sign Error")
	}
	resultData.OrderID = notifyMap["out_trade_no"].(string)
	tfee, _ := strconv.ParseFloat(notifyMap["total_fee"].(string), 32)
	resultData.TotalFee = int(tfee * 100)
	resultData.TransactionID = notifyMap["trade_no"].(string)
	resultData.ReturnData = []byte("success")
	return &resultData, nil
}

func (this *AliyPay) getConfig(name string) interface{} {
	if v, ok := this.Config[name]; !ok {
		return nil
	} else {
		return v
	}
}

func (this *AliyPay) checkConfig() bool {
	configFiled := []string{KEY, RETURN_URL}
	for _, item := range configFiled {
		if _, ok := this.Payment.Config[item]; !ok {
			// logging.Error("Not found [%v] in the config map", item)
			return false
		}
	}
	return true
}

func (this *AliyPay) makeSign(params map[string]interface{}) string {
	signParams := this.makeUrl(params)
	return this.MD5Sigin(fmt.Sprintf("%s%s", signParams, this.getConfig(KEY)))
}

func (this *AliyPay) checkSign(params map[string]interface{}) bool {
	originSign := params["sign"]
	delete(params, "sign")
	paySign := this.makeSign(params)
	if originSign == paySign {
		return true
	} else {
		return false
	}
}
