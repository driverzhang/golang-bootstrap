package payment_web

import (
	"fmt"
	"testing"
)

var Oreder Order

func TestMakeUrl(t *testing.T) {
	wxconfig := map[string]interface{}{
		"appid":  "wx60b42cdbf7055559",
		"mch_id": "1491810802",
		"key":    "9c32adcbfc89e064ec68c61e969e7b81",
	}
	pay, err := NewPayment().Init(1, wxconfig)
	if err != nil {
		fmt.Println("Happend err", err)
		return
	}

	Oreder = Order{
		OrderID:            "2018052s918221234", // 每次调用不能重复
		ProudctName:        "1233",
		PriceTotal:         1,
		ProductID:          "12322",
		IP:                 "127.0.0.1",
		ProudctDescription: "好的很东西",
		NotifyUrl:          "baidu.com",
	}
	url, err := pay.GenderPayUrl(Oreder)
	if err != nil {
		t.Log("死掉了啦：", err)
	}
	t.Log(url)
	body := `<xml>
	<appid><![CDATA[wx60b42cdbf7055559]]></appid>
	<attach><![CDATA[支付测试]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[1491810802]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[2018052s918221234]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
	<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>1</total_fee>
	<trade_type><![CDATA[NATIVE]]></trade_type>
	<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
  </xml>`
	// 注意，测试一下代码时 要将 149-155 注释
	stru, err := pay.PayNotify([]byte(body))
	if err != nil {
		t.Log("死掉了啦：", err)
	}

	t.Log(stru)
}
