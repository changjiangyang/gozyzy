package wechatbeans

// Message 微信消息
type Message struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	Event        string `xml:Event`
	EventKey     string `xml:EventKey`
}

// TextMessage 微信文本消息
type TextMessage struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	Content      string `xml:Content`
	MsgId        string `xml:MsgId`
}

type PicMessage struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	PicUrl       string `xml:PicUrl`
	MediaId      string `xml:MediaId`
	MsgId        string `xml:MsgId`
}

type VoiceMessage struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	MediaId      string `xml:MediaId`
	Format       string `xml:Format`
	MsgId        string `xml:MsgId`
}

type VideoMessge struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	MediaId      string `xml:MediaId`
	ThumbMediaId string `xml:ThumbMediaId`
	MsgId        string `xml:MsgId`
}

type ShortVideoMessge struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	MediaId      string `xml:MediaId`
	ThumbMediaId string `xml:ThumbMediaId`
	MsgId        string `xml:MsgId`
}

type LocalMessage struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	Location_X   string `xml:Location_X`
	Location_Y   string `xml:Location_Y`
	Scale        string `xml:Scale`
	Label        string `xml:Label`
	MsgId        string `xml:MsgId`
}

type linkMessage struct {
	ToUserName   string `xml:ToUserName`
	FromUserName string `xml:FromUserName`
	MsgType      string `xml:MsgType`
	CreateTime   string `xml:CreateTime`
	Title        string `xml:Title`
	Description  string `xml:Description`
	Url          string `xml:Url`
	MsgId        string `xml:MsgId`
}
