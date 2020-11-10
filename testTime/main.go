package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var tmp_str = `{"tid":1347456243421196949,"tid_str":"1347456243421196949","status":"WAIT_BUYER_PAY","type":"fixed","seller_nick":"瓷之舞旗舰店","buyer_nick":"糖糖熊9966","created":"2020-11-04 14:59:33","modified":"2020-11-04 14:59:33","trade_attr":"{\"erpHold\":\"0\"}","adjust_fee":"0.00","alipay_no":"2020110422001119725704565719","alipay_point":0,"available_confirm_fee":"1.80","buyer_alipay_no":"151********","buyer_area":"未知","buyer_cod_fee":"0.00","buyer_email":"","buyer_obtain_point_fee":0,"buyer_rate":false,"cod_fee":"0.00","cod_status":"NEW_CREATED","coupon_fee":0,"commission_fee":"0.00","discount_fee":"3.00","has_post_fee":true,"is_3D":false,"is_brand_sale":false,"is_daixiao":false,"is_force_wlb":false,"is_sh_ship":false,"is_lgtype":false,"is_part_consign":false,"is_wt":false,"is_gift":false,"num":1,"num_iid":623451801629,"new_presell":false,"nr_shop_guide_id":"","nr_shop_guide_name":"","orders":{"order":[{"adjust_fee":"0.00","buyer_rate":false,"cid":50003449,"discount_fee":"0.10","is_daixiao":false,"is_oversold":false,"num":1,"num_iid":623451801629,"oid":1347456243421196949,"order_from":"WAP,WAP","part_mjz_discount":"3.00","payment":"1.80","price":"4.90","refund_status":"NO_REFUND","seller_rate":false,"seller_type":"B","sku_id":"4648863642366","sku_properties_name":"颜色分类:【特价款】小麦pp32.5*21.5","snapshot_url":"s:1347456243421196949_1","status":"WAIT_BUYER_PAY","store_code":"cizhiwu001","timeout_action_time":"2020-11-05 14:59:33","title":"德国316不锈钢菜板家用抗菌防霉砧板厨房双面切菜板粘板塑料案板","total_fee":"4.80"}]},"payment":"1.80","pcc_af":0,"platform_subsidy_fee":"0.00","point_fee":0,"post_fee":"0.00","price":"4.90","real_point_fee":0,"received_payment":"0.00","receiver_address":"民*街道**街**号菜鸟驿**城国**园三期店","receiver_city":"沈阳市","receiver_country":"","receiver_district":"苏家屯区","receiver_mobile":"151********","receiver_name":"软*","receiver_state":"辽宁省","receiver_town":"民主街道","receiver_zip":"000000","seller_alipay_no":"***uyao@163.com","seller_can_rate":false,"seller_cod_fee":"0.00","seller_email":"3539755978@qq.com","seller_flag":0,"seller_mobile":"15228918881","seller_name":"苏州沐遥贸易有限公司","seller_rate":false,"service_type":"","shipping_type":"express","sid":"1347456243421196949","snapshot_url":"s:1347456243421196949_1","timeout_action_time":"2020-11-05 14:59:33","title":"瓷之舞旗舰店","total_fee":"4.90","trade_from":"WAP,WAP"}`

func transOidIntoString(info *map[string]interface{}) {
	fmt.Printf("info: %v", info)
	orders, ok := (*info)["orders"]
	fmt.Println(orders, ok)
	if !ok {
		fmt.Println("1")
		return
	}
	o, ok := orders.(map[string]interface{})
	if !ok {
		fmt.Println("2")
		return
	}
	o1, ok := o["order"]
	if !ok {
		return
	}
	o2, ok := o1.([]interface{})
	if !ok {
		return
	}

	for i, item := range o2 {
		order, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		fmt.Println("i")

		if v, ok := order["oid"]; ok {
			fmt.Println("xx", v)
			if v1, ok := v.(float64); ok {
				order["oid"] = strconv.FormatFloat(v1, 'f', -1, 64)
			} else {
				order["oid"] = fmt.Sprint(v1)
			}
		}
		o2[i] = order
	}

	o["order"] = o2
	(*info)["orders"] = o
}

func main() {
	decoder := json.NewDecoder(strings.NewReader(tmp_str))
	decoder.UseNumber()

	mp := map[string]interface{}{}

	decoder.Decode(&mp)

	transOidIntoString(&mp)
	fmt.Println("end")
}
