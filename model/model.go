package model

/*
 *日期/開市價/最高價/最低價/收市價/成交量
 */

type ExchangeData struct {

	//交易日期,index:0
	ExchangeDate string
	//開市價,index:1
	PriceFirst float64
	//最高價,index:2
	PriceHigh float64
	//最低價,index:3
	PriceLow float64
	//收市價,index:4
	PriceLast float64
	//成交量,index:5
	ExchangeAmount int64
	//10日均线,index:6
	MA10 float64
	//30日均线,index:7
	MA30 float64
	//60日均线,index:8
	MA60 float64
	//年均线,index:9
	MA365 float64
}
