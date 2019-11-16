package echarts

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"html/template"
)

type Chart struct {
	ID     string
	Title  template.HTML
	Js     template.JS
	Height int
	Width  int

	JsContentOptions *Options
}

func (c *Chart) SetID(id string) *Chart {
	c.ID = id
	return c
}

func (c *Chart) SetTitle(title template.HTML) *Chart {
	c.Title = title
	return c
}

func (c *Chart) SetHeight(height int) *Chart {
	c.Height = height
	return c
}

func (c *Chart) SetWidth(width int) *Chart {
	c.Width = width
	return c
}

type Options struct {
}

type Color string

func NewChart() *Chart {
	return new(Chart)
}

func (c *Chart) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("echarts").
		Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
			"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
				if cdnUrl == "" {
					return prefixUrl + assetsUrl
				}
				return cdnUrl + assetsUrl
			},
			"isLinkUrl": func(s string) bool {
				return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
			},
		}).
		Parse(List["echarts"])

	if err != nil {
		logger.Error("Chart GetTemplate Error: ", err)
	}

	return tmpl, "echarts"
}

func (c *Chart) GetAssetList() []string {
	return AssetsList
}

func (c *Chart) GetAsset(name string) ([]byte, error) {
	return Asset(name[1:])
}

func (c *Chart) IsAPage() bool {
	return false
}

func (c *Chart) GetName() string {
	return "echarts"
}

func (c *Chart) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := c.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, c)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
