package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	shopStr, err := getShopList("vpn_transafer/shop.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	println(len(shopStr))
	isTest := false
	cookie :="sensorsdata2015jssdkcross=%7B%22%24device_id%22%3A%22175549bc3066c0-04fb493b1f6fc5-163a6152-2007040-175549bc307cb1%22%7D; gr_user_id=43067cbc-67d8-4ec2-a8ba-c4c29baacfcb; uid=5f30fae54da090b9c7737ba7; sa_jssdk_2015_dataadmin_xiaoduoai_com=%7B%22distinct_id%22%3A%22%E8%B7%AF%E6%96%8C%22%2C%22first_id%22%3A%221755b61fceba84-007743bcf218f6-32637007-2007040-1755b61fcecf25%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%7D; sso_token=token:4141605d3d8146068e5394606d150ee5; sid=2|1:0|10:1606919806|3:sid|52:XypRbeyWug+5B2v1MzQyNCUlJDRkYTJxM2F3cTIzenhjeHNAIyQ0|6b013e0d284b5d37f25986843b1623b931607eec0b15c617a159e3b406536866"
	err = getNoAddSync(shopStr, isTest,cookie)
	if err != nil {
		fmt.Printf("请求失败，err:%s", err.Error())
	}
}

func getShopList(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileinfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	shopArray := strings.Split(strings.Trim(string(buffer), "\n"), ",")
	return shopArray, nil
}
func getNoAddSync(shopArray []string, isTest bool,cookie string) error {
	yangUrl := "http://wangcai.1yangai.com"
	zhigengUrl := "https://dataadmin.xiaoduoai.com"
	if isTest {
		yangUrl = "http://wangcai-test-mini-src.1yangai.com"
		zhigengUrl = "https://dataadmin-test-mini-src.xiaoduoai.com"
	}
	for _, shop := range shopArray {
		ret,err := findAddSyncStatus(yangUrl,shop)
		if ret.SwitchOn == false {
			err =addSyncReq(shop,zhigengUrl,cookie)
			if err !=nil{
				return err
			}
			time.Sleep(2*time.Second)
			err = insertSyncList(shop,zhigengUrl,cookie)
			if err !=nil{
				return err
			}
			time.Sleep(10*time.Second)
		}
	}
	return nil
}
//查询是否开启增量同步
func findAddSyncStatus(yangUrl,shop string) (ret *GetNoAddSync,err error)  {
	getUrl := fmt.Sprintf("%s/api/internal/mini_data_sync/switch_get?plat_user_id=%s&platform=tb", yangUrl, shop)
	resp, err := http.Get(getUrl)
	if err != nil {
		msg := fmt.Sprintf("%s 查询增量同步接口失败,错误信息：%s", shop, err.Error())
		return ret,errors.New(msg)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("%s 查询结果 %s", shop, string(body))
	fmt.Println(msg)
	ret = &GetNoAddSync{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		msg := fmt.Sprintf("%s 查询增量同步接口失败,错误信息：%s", shop, err.Error())
		return ret, errors.New(msg)
	}
	if ret.Code != 0 {
		msg := fmt.Sprintf("%s 查询增量同步接口失败,错误信息：%s", shop, ret.Message)
		return ret, errors.New(msg)
	}
	return ret,nil
}
//开启增量同步
func addSyncReq(shop ,zhigengUrl,cookie string) error {

	client := &http.Client{}
	getOpenAddSyncUrl := fmt.Sprintf("%s/proxy/admin/config_whitelist",zhigengUrl)
	forData := url.Values{}
	forData.Add("platform","tb")
	forData.Add("nick",shop)
	forData.Add("config","mini_data_sync")
	req ,err := http.NewRequest("POST",getOpenAddSyncUrl,strings.NewReader(forData.Encode()))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	req.Header.Add("Cookie",cookie)
	if err !=nil{
		return errors.New("创建开启增量同步请求接口失败")
	}
	resp, err := client.Do(req)
	if err !=nil{
		msg := fmt.Sprintf("%s 开启增量同步接口失败,错误信息：%s", shop, err.Error())
		return errors.New(msg)
	}

	addRet := &Common{}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("%s 开启增量同步结果 %s", shop, string(body))
	fmt.Println(msg)
	err = json.Unmarshal(body, addRet)
	if err != nil {
		msg := fmt.Sprintf("%s 开启增量同步接口失败,错误信息：%s", shop, err.Error())
		return  errors.New(msg)
	}
	return nil
}
//加入全量同步列表
func insertSyncList(shop ,zhigengUrl,cookie string)  error{
	client := &http.Client{}
	addSyncListUrl := fmt.Sprintf("%s/proxy/admin/data_sync/full_sync/create",zhigengUrl)
	addSyncListFormData := url.Values{}
	addSyncListFormData.Add("platform","tb")
	addSyncListFormData.Add("plat_user_id",shop)
	req ,err := http.NewRequest("POST",addSyncListUrl,strings.NewReader(addSyncListFormData.Encode()))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	req.Header.Add("Cookie",cookie)
	if err !=nil{
		msg := fmt.Sprintf("%s 加入全量同步失败,错误信息：%s", shop, err.Error())
		return errors.New(msg)
	}
	resp, err := client.Do(req)
	if err !=nil{
		msg := fmt.Sprintf("%s 开启增量同步接口失败,错误信息：%s", shop, err.Error())
		return errors.New(msg)
	}
	addSyncListRet := &Common{}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("%s 加入全量同步结果 %s", shop, string(body))
	fmt.Println(msg)
	err = json.Unmarshal(body, addSyncListRet)
	if err != nil {
		msg := fmt.Sprintf("%s 加入全量同步接口失败,错误信息：%s", shop, err.Error())
		return  errors.New(msg)
	}
	if addSyncListRet.Code!=0{
		msg := fmt.Sprintf("%s 加入全量同步接口失败,错误信息：%s", shop, addSyncListRet.Message)
		return errors.New(msg)
	}
	return nil
}

type GetNoAddSync struct {
	Common
	SwitchOn bool   `json:"switch_on"`
	ShowPop  bool   `json:"show_pop"`
}
type Common struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
}
