// Package oa https://developer.work.weixin.qq.com/document/path/91853
package oa

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 发送应用消息的接口地址
	sendURL = "https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent?access_token=%s"
)

type T struct {
}
type (
	Contents struct {
		Control string `json:"control"`
		Id      string `json:"id"`
		Value   struct {
			Text string `json:"text"`
		} `json:"value"`
	}
	ApplyData struct {
		Contents []Contents `json:"contents"`
	}
	// SendRequestCommon 发送应用消息请求公共参数
	SendRequestCommon struct {
		CreatorUserid       string `json:"creator_userid"` //申请人userid，此审批申请将以此员工身份提交，申请人需在应用可见范围内
		TemplateId          string `json:"template_id"`    //模板id。可在“获取审批申请详情”、“审批状态变化回调通知”中获得，也可在审批模板的模板编辑页面链接中获得。暂不支持通过接口提交[打卡补卡][调班]模板审批单。
		UseTemplateApprover int    `json:"use_template_approver"`
		Approver            []struct {
			Attr   int      `json:"attr"`
			Userid []string `json:"userid"`
		} `json:"approver"`
		ApplyData ApplyData `json:"apply_data"`

		SummaryList []struct {
			SummaryInfo []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"summary_info"`
		} `json:"summary_list"`
	}
	// SendResponse 发送应用消息响应参数
	SendResponse struct {
		util.CommonError
		SpNo string `json:"sp_no"` // 审批单编号
	}
)

// Send 发送应用消息
// @desc 实现企业微信发送应用消息接口：https://developer.work.weixin.qq.com/document/path/90236
func (r *Client) Send(apiName string, request interface{}) (*SendResponse, error) {
	// 获取accessToken
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 请求参数转 JSON 格式
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	// 发起http请求
	response, err := util.HTTPPost(fmt.Sprintf(sendURL, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	// 按照结构体解析返回值
	result := &SendResponse{}
	err = json.Unmarshal(response, result)
	// 返回数据
	return result, err
}
