package model

/*
 *日期/開市價/最高價/最低價/收市價/成交量
 */

type ExchangeData struct {

	//交易日期
	ExchageDate string

	//交易量
	ExchangeAmount int64

	//最高成交价
	PriceHigh float64

	//最低价格
	PriceLow float64

	//开盘价
	PriceFirst float64

	//收盘价格
	PriceLast float64
}
