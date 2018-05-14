package payment_order

type PaymentOrder struct {
	OrderId       string      `json:"order_id" 		xorm:"'order_id' not null varchar(32) pk"`
	TranscationID string      `json:"transaction_id" xorm:"'transaction_id' not null varchar(32)"`
	PriceTotal    int         `json:"total_fee"      xorm:"'total_fee' bigint "`
	ReturnData    []byte      `json:"return_data" xorm:"'return_data' not null varchar(255)"` // 返回回调数据
	AttachData    interface{} `json:"attach_data" xorm:"'attach_data' not null varchar(500)"` // 用户自定义数据
}

// type PaymentOrder struct {
// 	OrderId       string `json:"order_id" 		xorm:"'order_id' not null varchar(32) pk"`
// 	TranscationID string `json:"transaction_id" xorm:"'transaction_id' not null varchar(32)"`
// 	Appid         string `json:"appid" 		  	xorm:"'appid' not null varchar(32)"`
// 	IsSubscribe   string `json:"is_subscribe"   xorm:"'is_subscribe' varchar(2)"`
// 	OpenId        string `json:"open_id" 	    xorm:"'open_id' varchar(32)"`
// 	MchId         string `json:"mch_id" 	    xorm:"'mch_id'  varchar(32)"`
// 	PriceTotal    int    `json:"total_fee"      xorm:"'total_fee' bigint "`
// 	NonceStr      string `json:"nonce_str"      xorm:"'nonce_str' not null varchar(32)"`
// 	Sign          string `json:"sign" 	        xorm:"'sign' not null varchar(32) "`
// 	TimeEnd       string `json:"time_end" 	    xorm:"'time_end' not null TIMESTAMP"`
// 	BankType      string `json:"bank_type" 	    xorm:"'bank_type' tinyint(10)"`
// 	TradeType     string `json:"trade_type" 	xorm:"'trade_type' varchar(32)"`
// 	FeeType       string `json:"fee_type"  		xorm:"'fee_type' varchar(16)"`
// 	CashFee       string `json:"cash_fee"  		xorm:"'cash_fee' varchar(16)"`
// }

func (p PaymentOrder) TableName() string {
	return "payment_order"
}

func (p *PaymentOrder) Insert() (num int64, err error) {
	num, err = engine.Write.Insert(p)
	return
}
