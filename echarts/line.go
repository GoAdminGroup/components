package echarts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
)

type LineChart struct {
	*Chart

	JsContent LineJsContent
}

type LineJsContent struct {
	Options *Options
}

func (l *LineChart) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := l.GetTemplate()

	if l.JsContentOptions != nil {
		l.JsContent.Options = l.JsContentOptions
	}

	jsonByte, _ := json.Marshal(l.JsContent)
	l.Js = template.JS(string(jsonByte))

	err := tmpl.ExecuteTemplate(buffer, defineName, l)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
