package helper

import "net/textproto"

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

type Attachment struct {
	Filename    string
	ContentType string
	Header      textproto.MIMEHeader
	Content     []byte
	HTMLRelated bool
}

func NewEmail() *Email {
	return &Email{Headers: textproto.MIMEHeader{}}
}
