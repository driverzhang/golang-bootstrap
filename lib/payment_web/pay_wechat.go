package payment_web

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"git.zhuzi.me/zzjz/zhuzi-payment/api/payments"
	"git.zhuzi.me/zzjz/zhuzi-payment/lib/log"
)

type Wechat struct {
	Payment
}

// 这里参数填入 可以接入后端脚手架的config.yaml进行统一配置获取
const PAY_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" // 统一下单 微信支付 接口
var TotalPrice int                                               // 用于取 商品价格

// 统一下单数据包
type OrderXML struct {
	XMLName    xml.Name `xml:"xml"`
	APPID      string   `xml:"appid"`            // 公众账号ID String(32) 企业号corpid即为此appId
	MchID      string   `xml:"mch_id"`           // 商户号 String(32) 微信支付分配的商户号
	Body       string   `xml:"body"`             // 商品描述 String(128) 商品简单描述 浏览器打开的网站主页title名
	Detail     string   `xml:"detail"`           // 商品详情 String(6000)
	NonceStr   string   `xml:"nonce_str"`        // 随机字符串 String(32) 长度要求在32位以内 调用随机数函数生成，将得到的值转换为字符串。
	NotifyUrl  string   `xml:"notify_url"`       // 通知地址 String(256) 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	OrderID    string   `xml:"out_trade_no"`     // 商户订单号 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。
	ClientIP   string   `xml:"spbill_create_ip"` // 终端IP APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
	PriceTotal int      `xml:"total_fee"`        // 标价金额 订单总金额，单位为分 100分=1元
	TradeType  string   `xml:"trade_type"`       // 交易类型 JSAPI 公众号支付; NATIVE 扫码支付; APP APP支付
	ProductId  string   `xml:"product_id"`       // 商品ID trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义
	Sign       string   `xml:"sign"`             // 签名 通过签名算法计算得出的签名值
}

// 微信统一下单返回数据包
// 以下字段在 return_code result_code 为 SUCCESS 的时候返回
type WeChatPayXML struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"` // return_code 为 SUCCESS
	ReturnMsg  string   `xml:"return_msg"`
	MchID      string   `xml:"mch_id"`
	APPID      string   `xml:"appid"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	ResultCode string   `xml:"result_code"` // 业务结果 SUCCESS/FAIL 交易标识，交易是否成功用 result_code 来判断
	PrepayID   string   `xml:"prepay_id"`   // 预支付交易会话标识 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	TradeType  string   `xml:"trade_type"`
	CodeUrl    string   `xml:"code_url"` // 二维码链接 trade_type为NATIVE时有返回，用于生成二维码，展示给用户进行扫码支付
	ErrorMSG   string   `xml:"err_code_des"`
}

// 微信异步通知数据包
type WechatNotify struct {
	XMLName       xml.Name `xml:"xml"`
	APPID         string   `xml:"appid"`
	MchID         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`      // 随机字符串
	Sign          string   `xml:"sign"`           // 签名
	ResultCode    string   `xml:"result_code"`    // 业务结果
	OpenID        string   `xml:"openid"`         // 用户标识 用户在商户appid下的唯一标识
	TradeType     string   `xml:"trade_type"`     // 交易类型
	BankType      string   `xml:"bank_type"`      // 付款银行 银行类型，采用字符串类型的银行标识
	TotalFee      int      `xml:"total_fee"`      // 订单金额
	FeeType       string   `xml:"fee_type"`       // 货币类型
	CashFee       string   `xml:"cash_fee"`       // 现金支付金额
	TranscationID string   `xml:"transaction_id"` // 微信支付订单号
	OrderID       string   `xml:"out_trade_no"`   // 商户订单号
	IsSubscribe   string   `xml:"is_subscribe"`   // 是否关注公众账号 Y-关注，N-未关注，仅在公众账号类型支付有效
	TimeEnd       string   `xml:"time_end"`       // 支付完成时间	格式为yyyyMMddHHmmss
	ReturnCode    string   `xml:"return_code"`    // 返回状态码
}

// 微信支付结果通知 返回给微信 表示收到通知
type returnMSG struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"` // 返回状态码 SUCCESS/FAIL SUCCESS表示商户接收通知成功并校验成功
	ReturnMSG  string   `xml:"return_msg"`  // 返回信息 如非空，为错误原因; 签名失败; 参数格式校验错误
}

func (w WechatNotify) NewWechatNotify() *WechatNotify {
	return &WechatNotify{}
}

func NewWechat(params map[string]interface{}) *Wechat {
	return &Wechat{Payment{Config: params}}
}

/*
 *	 调用 统一下单API 获取预交易链接
 */

func (this *Wechat) GenderPayUrl(order Order) (string, error) {
	var o OrderXML
	o.APPID = this.getConfig("appid").(string)
	o.MchID = this.getConfig("mch_id").(string)
	o.Body = order.ProudctName
	o.NonceStr = this.GenerateString(16) // 生成16位的随机字符串 => "tGc1WTnu7ao6mG5R"
	o.NotifyUrl = order.NotifyUrl        // 通知地址 微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	o.OrderID = order.OrderID            // 订单ID
	o.ClientIP = order.IP                // 订单用户IP
	o.PriceTotal = order.PriceTotal      // 价格
	o.TradeType = "NATIVE"               // NATIVE 扫码支付;
	o.ProductId = order.ProductID        // 商品ID
	params := this.struct2map(o)         // 将 OrderXML 转化成 map[string]interface{} 格式
	o.Sign = this.makeSign(params)       // 生成签名算法
	TotalPrice = o.PriceTotal
	log.Printf("【原始下单传入数据包】%+v\n", o)
	data, err := xml.Marshal(o)               // 返回 o 的XML编码
	r, err := this.httpPostXML(PAY_URL, data) // 模拟客户端 开始调用 统一下单接口
	if err != nil {
		return "", fmt.Errorf("response Error | %v\n", err)
	}
	res, err := ioutil.ReadAll(r.Body)
	log.Printf("【下单接口返回XML数据包】 %+v\n", string(res))
	if err != nil {
		return "", fmt.Errorf("get body Error | %v\n", err)
	}
	var wechatPay WeChatPayXML
	xml.Unmarshal(res, &wechatPay)      // 转化格式 获得返回数据包
	if wechatPay.ResultCode == "FAIL" { // 交易结果 未成功
		return "", fmt.Errorf("%v", wechatPay.ErrorMSG)
	}
	payMap := this.struct2map(wechatPay)
	if !this.checkSign(payMap) {
		return "", fmt.Errorf("Sign Error")

	} else {
		return wechatPay.CodeUrl, nil // 最终成功就返回预交易code_url取此链接生成二维码,直接扔给前端生成即可
	}
}

/*
 *	微信支付通结果通知 商户端 返回 收到支付结果通知信息
 *	这里的 notify 是 XML 形式的数据格式
 */

func (this *Wechat) PayNotify(notify []byte) (*NofiyData, error) {
	if len(notify) == 0 {
		return nil, fmt.Errorf("Notify Data empty")
	}
	var notifyPay WechatNotify
	var msg returnMSG     // 接受的数据包
	var paydata NofiyData // 返回的数据包
	xml.Unmarshal(notify, &notifyPay)
	log.Printf("【支付后返回的订单数据】 %+v\n", notifyPay)

	// TODO: 线上应该放开下面注释代码！

	// notifyMap := this.struct2map(notifyPay)
	// if !this.checkSign(notifyMap) { // 检验 签名 是否一致
	// 	msg.ReturnCode = "FAIL"
	// 	msg.ReturnMSG = "Sign Error"

	// 	paydata.ReturnData, _ = xml.Marshal(msg)
	// 	return &paydata, fmt.Errorf("Sign Error")
	// }

	if notifyPay.TotalFee != TotalPrice { // 检验 商品订单价格 是否一致
		msg.ReturnCode = "FAIL"
		msg.ReturnMSG = "total_fee Error"

		paydata.ReturnData, _ = xml.Marshal(msg)
		return &paydata, fmt.Errorf("total_fee Error")
	}

	paydata.OrderID = notifyPay.OrderID
	paydata.TotalFee = notifyPay.TotalFee
	paydata.TransactionID = notifyPay.TranscationID
	msg.ReturnMSG = "OK"
	msg.ReturnCode = "SUCCESS"
	paydata.ReturnData, _ = xml.Marshal(msg)
	paydata.AttachData = fmt.Sprintf("%+v,%+v,%+v,%+v,%+v,%+v,%+v,%+v",
		notifyPay.OpenID, notifyPay.Sign, notifyPay.NonceStr,
		notifyPay.BankType, notifyPay.CashFee, notifyPay.FeeType,
		notifyPay.IsSubscribe, notifyPay.TimeEnd)

	log.Print(paydata.AttachData)

	if notifyPay.ResultCode == "SUCCESS" && notifyPay.ReturnCode == "SUCCESS" {
		// 生成系统调用者订单
		wx := &payments.Wx{
			OrderID:       paydata.OrderID,
			TransactionID: paydata.TransactionID,
			TotalFee:      paydata.TotalFee,
			ReturnData:    paydata.ReturnData,
			AttachData:    paydata.AttachData,
		}

		payments.InsertOrder(wx)
	}
	return &paydata, nil
}

// 模拟客户端 开始调用 统一下单接口
func (this *Wechat) httpPostXML(url string, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	resp, err := client.Do(req)
	defer req.Body.Close()
	return resp, err

}

func (this *Wechat) getConfig(name string) interface{} {

	if v, ok := this.Config[name]; !ok {
		return nil
	} else {
		return v
	}
}

// 检验输入参数是否存在
func (this *Wechat) checkConfig() bool {
	configFiled := []string{"mch_id", "appid", "key"}
	for _, item := range configFiled {
		if _, ok := this.Payment.Config[item]; !ok {
			log.Error("Not found [%v] in the config map", item)
			return false
		}
	}
	return true
}

// 生成 签名算法 参考微信扫码支付官网文档说明
func (this *Wechat) makeSign(params map[string]interface{}) string {
	signParams := this.makeUrl(params)
	str := strings.ToUpper(this.MD5Sigin(fmt.Sprintf("%s&key=%s", signParams, this.getConfig("key"))))
	return str
}

func (this *Wechat) CheckSign(params map[string]interface{}) bool {
	return this.checkSign(params)
}

// 验证返回 sign , 此返回sign 就是返回数据生成好的，只需取返回数据以同样签名方法再次签名即可验证。
func (this *Wechat) checkSign(params map[string]interface{}) bool {
	originSign := params["sign"]
	delete(params, "sign")
	paySign := this.makeSign(params)
	if originSign == paySign {
		return true
	} else {
		return false
	}
}
