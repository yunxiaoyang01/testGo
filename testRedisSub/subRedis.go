package main

type RDSSwitcher struct {
	IsOpenTradeRdsWhiteList bool     `json:"is_open_trade_rds_white_list"`
	TradeRdsWhiteList       []string `json:"trade_rds_white_list"`
}

func main() {
	// v :="is_open_trade_rds_white_list=true&trade_rds_white_list=wangjq_1990"
	// paramArray := strings.Split(v,"&")
	// tradeRdsWhite :=RDSSwitcher{}
	// for _,param := range paramArray{
	// 	subParamArray := strings.Split(param,"=")

	// 	tradeRdsWhite[subParamArray[0]]=subParamArray[1]
	// }

}
