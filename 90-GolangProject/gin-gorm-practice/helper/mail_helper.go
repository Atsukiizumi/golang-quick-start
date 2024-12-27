package helper

import "net/textproto"

// 邮件
type Email struct {
	ReplyTo     []string
	From        string
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	Text        []byte
	HTML        []byte
	Sender      string
	Headers     textproto.MIMEHeader
	Attachments []*Attachment
	ReadReceipt []string
}

// 附件
type Attachment struct {
	Filename    string
	ContentType string
	Header      textproto.MIMEHeader
	Content     []byte
	HTMLRelated bool
}

// NewEmail 生成email对象，返回指针
func NewEmail() *Email {
	return &Email{Headers: textproto.MIMEHeader{}}
}
