package msg

import (
	"encoding/json"

	"github.com/samhuangszu/wechat/mp/media"
)

// 消息类型
const (
	MsgTypeWechat = 1  //微信
	MsgTypeAlipay = 2  //阿里
	MsgTypeSms    = 4  //短信
	MsgTypeMail   = 8  //邮件
	MsgTypeApp    = 16 //App
	MsgTypeWeiMin = 32 //微信小程序
	MsgTypeAliMin = 64 //支付宝小程序
)

// 短信厂商
const (
	SenderTypeTX = 0 //腾讯
	SenderTypeAL = 1 //阿里
	SenderTypeYP = 2 //云片
	SenderTypeFG = 3 //飞鸽
	SenderTypeYX = 4 //羽信
)

// 返回Code
const (
	ResultCodeNormal    = 0
	ResultCodeError     = 1
	ResultCodeNotFollow = 10000
)

// 消息模板
const (
	WxTextTmpID         = "OPENTM206919910" //微信文本消息模板ID
	WxRefundTmpID       = "TM00004"         //微信退款消息模板ID
	WxOrderTmpID        = "OPENTM206305613" //微信设备状态消息模板ID
	AliTextTmpID        = "TM000001374"     //阿里服务提醒消息模板ID
	AliRefundTmpID      = "TM000003980"     //阿里退款消息模板ID
	AliOrderTmpID       = "TM000002289"     //阿里设备状态消息模板ID
	WxMpRefundTmpID     = "4860"            //微信小程序退款模板ID
	WxMpOrderStartTmpID = "764"             //微信小程序开始充电通知模板ID
	WxMpOrderStopTmpID  = "646"             //微信小程序开始结束通知模板ID
	WxMpOrderOddTmpID   = "1429"            //微信小程序充电异常通知模板ID
)

// Alert 提醒
type Alert struct {
	AppID   string `json:"appId"`
	Title   string `json:"title"`
	Remark  string `json:"remark"`
	Content string `json:"content"`
	Time    string `json:"Time"`
	MsgType int    `json:"msgType"`
	UserID  string `json:"userId"`
	URL     string `json:"url,omitempty"`   // 对应的H5页面，没有可以留空
	TplID   string `json:"tplId,omitempty"` //如果是发短信，则必填
}

// Message 请求
type Message struct {
	AppID string `json:"app_id"` //服务商ID
	Msg   *Msg   `json:"msg"`    //消息内容
}

// Msg 消息
type Msg struct {
	Type  int    `json:"type"`             //类型
	Extra string `json:"extra,omitempty"`  //透传字段
	Ext   string `json:"ext"`              //类型字符串
	MsgID string `json:"msg_id,omitempty"` //消息ID
}

// ToString 转成str
func (c Message) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// WeappMessage 小程序消息
type WeappMessage struct {
	ToUser        string         `json:"touser"`          //必须, 接受者OpenID
	MpTemplateMsg *MpTemplateMsg `json:"mp_template_msg"` //公众号消息模版
}

// MpTemplateMsg 微信小程序消息
type MpTemplateMsg struct {
	Appid       string               `json:"appid"`
	TemplateID  string               `json:"template_id"`
	URL         string               `json:"url"`
	Data        map[string]*DataItem `json:"data"`        // 必须, 模板数据
	MiniProgram *MiniProgram         `json:"miniprogram"` // 可选,跳转至小程序地址
}

//MiniProgram 小程序跳转配置
type MiniProgram struct {
	AppID    string `json:"appid"`    //所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
	PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
}

//WxMessage 发送的模板消息内容
type WxMessage struct {
	ToUser      string               `json:"touser"`          // 必须, 接受者OpenID
	TemplateID  string               `json:"template_id"`     // 必须, 模版ID
	URL         string               `json:"url,omitempty"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color       string               `json:"color,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Extra       string               `json:"extra,omitempty"` // 可选，预留字段
	Data        map[string]*DataItem `json:"data"`            // 必须, 模板数据
	MiniProgram *MiniProgram         `json:"miniprogram"`     // 可选, 跳转至小程序地址
}

// ToString 转成str
func (c WxMessage) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// WeappMessage 转化成WeappMessage
func (c WxMessage) WeappMessage() *WeappMessage {
	msg := &WeappMessage{
		ToUser: c.ToUser,
		MpTemplateMsg: &MpTemplateMsg{
			Appid:       c.Extra,
			TemplateID:  c.TemplateID,
			Data:        c.Data,
			MiniProgram: c.MiniProgram,
		},
	}
	return msg
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// AliMessage alipay 消息模板
type AliMessage struct {
	ToUserID string      `json:"to_user_id"` // 必须, ali_user_id
	Template AliTemplate `json:"template"`   //必填，消息内容
}

// ToString 转成str
func (c AliMessage) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// AliTemplate 消息内容
type AliTemplate struct {
	TemplateID string     `json:"template_id"`
	Context    AliContext `json:"context"`
}

// AliContext 消息上下文
type AliContext struct {
	HeadColor  string   `json:"head_color"`  //顶部色条的色值
	URL        string   `json:"url"`         // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	ActionName string   `json:"action_name"` //底部链接描述文字，如“查看详情”
	Keyword1   DataItem `json:"keyword1,omitempty"`
	Keyword2   DataItem `json:"keyword2,omitempty"`
	Keyword3   DataItem `json:"keyword3,omitempty"`
	Keyword4   DataItem `json:"keyword4,omitempty"`
	Keyword5   DataItem `json:"keyword5,omitempty"`
	Keyword6   DataItem `json:"keyword6,omitempty"`
	Keyword7   DataItem `json:"keyword7,omitempty"`
	First      DataItem `json:"first,omitempty"`
	Remark     DataItem `json:"remark,omitempty"`
}

// AppMessage App消息
type AppMessage struct {
	ToUserID  string                 `json:"toUserId"`
	Content   string                 `json:"content"`
	Title     string                 `json:"title,omitempty"`
	DetailURL string                 `json:"detailUrl,omitempty"`
	Extras    map[string]interface{} `json:"extras,omitempty"`
}

// ToString 转成str
func (c AppMessage) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// 短信业务
const (
	SmsBizTypeNone   = 0 //未知业务
	SmsBizTypeLogin  = 1 //登录验证
	SmsBizTypeNormal = 2 //通用验证
)

// BizTypDesc 业务模板
var BizTypDesc = map[int]string{
	SmsBizTypeNone:   "#0为您的登录验证码，请于两分钟内填写。如非本人操作，请忽略本短信",
	SmsBizTypeLogin:  "#0为您的登录验证码，请于两分钟内填写。如非本人操作，请忽略本短信",
	SmsBizTypeNormal: "#0为您的#1验证码，请于两分钟内填写。如非本人操作，请忽略本短信",
}

// 通用验证码类型
const (
	BizSubTypeDraw       = 1 //提现验证
	BizSubTypeBindCard   = 2 //绑定银行卡
	BizSubTypeUnBindCard = 3 //解绑银行卡
	BizSubTypeLogin      = 4 //登录
)

// SubTypeDesc 映射
var SubTypeDesc = map[int]string{
	BizSubTypeDraw:       "提现",
	BizSubTypeBindCard:   "绑定银行卡",
	BizSubTypeUnBindCard: "解绑银行卡",
	BizSubTypeLogin:      "登录",
}

// SmsMessage 短信消息
type SmsMessage struct {
	NationCode string   `json:"nation_code"` //国家码
	Mobile     string   `json:"mobile"`      //手机号码
	BizType    int      `json:"bizType"`     //业务类型
	TplID      string   `json:"tpl_id"`      //模板ID
	Params     []string `json:"params"`      //对应tpl中的参数
}

// ToString 转成str
func (c SmsMessage) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// MailerMessage 邮件消息
type MailerMessage struct {
	FromName string `json:"from_name"` //发件人名称
	To       string `json:"to"`        //收件人邮箱地址，以,号分开
	Bcc      string `json:"bcc"`       //cc接收人，以,号分开
	Subject  string `json:"subject"`   //标题
	Content  string `json:"content"`   //内容：目前支持html
}

// ToString 转成str
func (c MailerMessage) ToString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// NewsArticle 媒体新闻结构体
type NewsArticle struct {
	ServiceID        string `json:"serviceId"`        //商户ID
	MediaID          string `json:"mediaId"`          //媒体图片ID
	Title            string `json:"title"`            //标题
	Digest           string `json:"digest"`           //新闻摘要
	Content          string `json:"content"`          //新闻内容，支持html标签
	ContentSourceURL string `json:"contentSourceUrl"` //点击查看原文链接
}

// ToWxArticle 转换成微信结构
func (c NewsArticle) ToWxArticle(author ...string) *media.Article {
	a := &media.Article{
		ThumbMediaId:     c.MediaID,
		Title:            c.Title,
		Digest:           c.Digest,
		Content:          c.Content,
		ContentSourceURL: c.ContentSourceURL,
		ShowCoverPic:     1,
	}
	if len(author) > 0 {
		a.Author = author[0]
	}
	return a
}
