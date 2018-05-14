package payments

import (
	"fmt"
	"reflect"

	"git.zhuzi.me/zzjz/zhuzi-payment/lib/log"
	"git.zhuzi.me/zzjz/zhuzi-payment/model/payment_order"
)

type Wx struct {
	OrderID       string      // 订单号
	TransactionID string      // 支付交易号
	TotalFee      int         // 订单金额(单位分)
	ReturnData    []byte      // 返回回调数据
	AttachData    interface{} // 用户自定义数据
}

func InsertOrder(wx *Wx) {
	log.Print("wx: ", wx)
	var err error
	log.Print(fmt.Sprintf("%s %+v\n", wx.AttachData, wx.AttachData))
	payment := &payment_order.PaymentOrder{
		OrderId:       wx.OrderID,
		TranscationID: wx.TransactionID,
		PriceTotal:    wx.TotalFee,
		ReturnData:    wx.ReturnData,
		AttachData:    wx.AttachData,
	}

	_, err = payment.Insert()

	if err != nil {
		log.Error("生成订单失败！,err", err.Error())
		return
	}
}

func struct2map(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ { // 返回v持有的结构体类型值的字段数
		var filedName string
		if filedName = t.Field(i).Tag.Get("json"); len(filedName) == 0 {
			if filedName = t.Field(i).Tag.Get("xml"); len(filedName) == 0 {
				filedName = t.Field(i).Name
			}
		}
		if filedName == "xml" {
			continue
		}
		data[filedName] = v.Field(i).Interface()
	}
	return data
}
