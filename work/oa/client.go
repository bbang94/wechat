// Package oa
package oa

import (
	"github.com/silenceper/wechat/v2/work/context"
)

// Client 消息推送接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{ctx}
}
