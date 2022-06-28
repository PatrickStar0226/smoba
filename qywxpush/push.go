package qywxpush

import (
	"bytes"
	"github.com/superggfun/smoba/wxpush"
	"text/template"

	"github.com/superggfun/smoba/config"
)

func Push(markdown wxpush.Markdown) error {
	tmpl, err := template.ParseFiles("static/template.md")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, markdown)
	err = qywxpush(config.ReadFile().Webhookurl, buffer.String())
	if err != nil {
		return err
	}
	return nil
}

func PushE(markdown wxpush.Markdown) error {
	tmpl, err := template.ParseFiles("static/error.md")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, markdown)
	err = qywxpush(config.ReadFile().Webhookurl, buffer.String())
	if err != nil {
		return err
	}
	return nil
}
