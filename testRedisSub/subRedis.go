package main

import (
	"encoding/json"
	"fmt"
)

type RDSSwitcher struct {
	IsOpenTradeRdsWhiteList bool     `json:"is_open_trade_rds_white_list"`
	TradeRdsWhiteList       []string `json:"trade_rds_white_list"`
}

type GetTradeFullInfoRes map[string]interface{}
func main() {
	// v :="is_open_trade_rds_white_list=true&trade_rds_white_list=wangjq_1990"
	// paramArray := strings.Split(v,"&")
	// tradeRdsWhite :=RDSSwitcher{}
	// for _,param := range paramArray{
	// 	subParamArray := strings.Split(param,"=")

	// 	tradeRdsWhite[subParamArray[0]]=subParamArray[1]
	// }
	reply := struct {
		Data GetTradeFullInfoRes `json:"trade"`
	}{}
	retMsg :="{\"trade\":{\"buyer_nick\":\"lubin1808368487\",\"buyer_rate\":false,\"created\":\"2020-11-13 16:22:48\",\"modified\":\"2020-11-13 16:22:48\",\"new_presell\":false,\"num\":1,\"orders\":{\"order\":[{\"adjust_fee\":\"0.00\",\"buyer_rate\":false,\"cid\":50012725,\"discount_fee\":\"0.00\",\"is_daixiao\":false,\"is_oversold\":false,\"num\":1,\"num_iid\":626628054416,\"oid\":\"1367811901342024015\",\"order_from\":\"WAP,WAP\",\"payment\":\"0.01\",\"pic_path\":\"https:\\/\\/img.alicdn.com\\/bao\\/uploaded\\/i3\\/4291096516\\/O1CN01ifqBOd1y0NKzhyw6B_!!4291096516.jpg\",\"price\":\"0.01\",\"refund_status\":\"NO_REFUND\",\"seller_rate\":false,\"seller_type\":\"C\",\"snapshot_url\":\"s:1367811901342024015_1\",\"status\":\"WAIT_BUYER_PAY\",\"title\":\"二手荧光笔01\"}]},\"payment\":\"0.01\",\"post_fee\":\"0.00\",\"receiver_city\":\"成都市\",\"receiver_state\":\"四川省\",\"receiver_town\":\"华阳镇街道\",\"receiver_zip\":\"000000\",\"seller_nick\":\"tb425734443\",\"status\":\"WAIT_BUYER_PAY\",\"tid\":\"1367811901342024015\",\"total_fee\":\"0.01\",\"trade_from\":\"WAP,WAP\",\"type\":\"fixed\"},\"request_id\":\"f1zw1rtwpqv0\"}"
	json.Unmarshal([]byte(retMsg),&reply)
	ori := dealTradeFullInfoReply(reply.Data)
	print(ori["tid_str"].(string))
}



func  dealTradeFullInfoReply( ori map[string]interface{}) map[string]interface{} {
	// 小程序无法通过api获取到tid_str，需要填充

	if _, ok := ori["tid_str"]; !ok {
		ori["tid_str"] = fmt.Sprint(ori["tid"])
	}
	return ori
}