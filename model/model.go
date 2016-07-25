package model

/*
 *日期/開市價/最高價/最低價/收市價/成交量
 */

type ExchangeData struct {

	//交易日期
	ExchangeDate string

	//開市價
	PriceFirst float64

	//最高價
	PriceHigh float64

	//最低價
	PriceLow float64

	//收市價
	PriceLast float64

	//成交量
	ExchangeAmount int64
}
