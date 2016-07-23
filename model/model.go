package model

/*
 *日期/開市價/最高價/最低價/收市價/成交量
 */

type ExchangeData struct {
	ExchageDate    string
	ExchangeAmount int64
	PriceHigh      float64
	PriceLow       float64
	PriceFirst     float64
	PriceLast      float64
}
