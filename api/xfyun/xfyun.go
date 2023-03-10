// 科大讯飞api客户端
// 如果不使用 NewXfyun，必须先 Init 初始化
// 不同服务的密钥不同
// https://www.xfyun.cn/doc/
package xfyun

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/crypto/sign"
	"github.com/violetpupil/components/lib/resty"
	"github.com/violetpupil/components/std/time"
)

type Client struct {
	Appid       string `json:"appid"`
	LfAsrSecret string `json:"lfAsrSecret"` // 语音转写密钥
}

// NewXfyun 创建科大讯飞api客户端，cfg是配置json字符串
func NewXfyun(cfg string) (*Client, error) {
	c := new(Client)
	err := json.Unmarshal([]byte(cfg), c)
	return c, err
}

var Xfyun *Client

// Init 初始化api客户端
func Init(appid string) {
	Xfyun = &Client{Appid: appid}
}

// InitEnv 用环境变量初始化api客户端
func InitEnv() {
	Init(os.Getenv("XfyunAppid"))
	SetLfAsrSecret(os.Getenv("XfyunLfAsrSecret"))
}

// SetLfAsrSecret 设置语音转写密钥
func SetLfAsrSecret(s string) {
	Xfyun.LfAsrSecret = s
}

// SignA 生成签名并构造请求参数，secret为服务密钥
func (a *Client) SignA(secret string) map[string]string {
	ts := time.Ts()
	signA := sign.Sign(a.Appid, ts, secret)
	m := map[string]string{
		"signa": signA,
		"appId": a.Appid,
		"ts":    ts,
	}
	return m
}

// 响应成功时，才能获取到响应内容
const ResCodeSuss = "000000"

// ResBody 通用响应体
type ResBody struct {
	Code     string `json:"code"`     // 响应代码
	DescInfo string `json:"descInfo"` // 响应信息
}

// Error 处理失败时，获取错误
func (r ResBody) Error() error {
	return fmt.Errorf("xfyun api: %s %s", r.Code, r.DescInfo)
}

// CodeI 获取响应代码，实现ResBodyI
func (r ResBody) CodeI() string {
	return r.Code
}

// ResBodyI 通用响应体接口
// 嵌入了 ResBody 的结构体可以统一使用该接口
type ResBodyI interface {
	Error() error
	CodeI() string
}

// Unmarshal 响应体json解码，响应代码不是成功时返回错误
// body 必须传不为nil的指针
func Unmarshal(res *resty.Response, body ResBodyI) error {
	err := json.Unmarshal(res.Body, body)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return res.ToError()
	}
	// 解码成功，但没有获取到响应代码
	if body.CodeI() == "" {
		return res.ToError()
	} else if body.CodeI() != ResCodeSuss {
		return body.Error()
	}
	return nil
}
