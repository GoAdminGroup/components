package theme1

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/login"
	"html/template"
	textTemplate "text/template"
)

type Login struct {
	TencentWaterProofWallID string
	BackgroundColor         string
	LoginBtnColor           string
}

type Config struct {
	TencentWaterProofWallID string
	LoginBtnColor           string
	BackgroundColor         string
}

func Get(cfg ...Config) *Login {
	if len(cfg) > 0 {
		return &Login{
			TencentWaterProofWallID: cfg[0].TencentWaterProofWallID,
			BackgroundColor:         utils.SetDefault(cfg[0].BackgroundColor, "", "#2d3a4b"),
			LoginBtnColor:           utils.SetDefault(cfg[0].LoginBtnColor, "", "#6a83a2"),
		}
	}
	return &Login{BackgroundColor: "#2d3a4b", LoginBtnColor: "#6a83a2"}
}

func (l *Login) GetTemplate() (*template.Template, string) {

	t := textTemplate.New("login_theme1").Delims("{%", "%}")
	t, err := t.Parse(List["login/theme1"])
	if err != nil {
		logger.Error("login component, get template parse error: ", err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, l)
	if err != nil {
		logger.Error("login component, get template execute error: ", err)
	}

	tmpl, err := template.New("login_theme1").
		Funcs(login.DefaultFuncMap).
		Parse(buf.String())

	if err != nil {
		logger.Error("login component, get template error: ", err)
	}

	return tmpl, "login_theme1"
}

func (l *Login) GetAssetList() []string               { return AssetsList }
func (l *Login) GetAsset(name string) ([]byte, error) { return Asset(name[1:]) }
func (l *Login) GetName() string                      { return "login" }
func (l *Login) IsAPage() bool                        { return true }

func (l *Login) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := l.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, l)
	if err != nil {
		logger.Error("login component, compose html error:", err)
	}
	return template.HTML(buffer.String())
}
