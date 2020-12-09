package main

import (
	"strconv"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TextContent struct {
	TemplateId   string `json:"template_id"`
	TemplateData struct {
		Text string `json:"text"`
	} `json:"template_data"`
}

type ContentType int

const (
	ContentTypeText ContentType = 1 //文本
	ContentTypePic  ContentType = 6 //图片
)

type ChatMessage struct {
	CreateTime  int64             `json:"create_time"`
	MsgId       string            `json:"msg_id"`
	BizUniqueId string            `json:"biz_unique_id"`
	Sender      User              `json:"sender"`
	Receivers   []User            `json:"receivers"`
	ContentType ContentType       `json:"content_type"`
	Content     string            `json:"content"`
	ChannelType string            `json:"channel_type"`
	Ext         map[string]string `json:"ext"`
}

const (
	transfer_answer_for_report = "转接"
	separate_send              = "{分行发送↓}"
)

//var reg *regexp.Regexp = regexp.MustCompile(separate_send)
//var reg1 *regexp.Regexp = regexp.MustCompile()
func main() {
	numIdStr:=""
	tid, err := strconv.ParseInt(numIdStr, 10, 64)
	if err != nil {
		print(err.Error())
	}
	print(tid)
}

func removeEmptyString(textArray []string) []string {
	j := 0
	for _, val := range textArray {
		val = strings.TrimPrefix(val,"\n")
		if val != "" {
			textArray[j]=val
			j++
		}
	}
	return textArray[:j]

}