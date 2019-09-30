package theme1

import (
	"github.com/chenhg5/go-admin/modules/logger"
	"html/template"
)

type Login struct {
}

func Get() *Login {
	return new(Login)
}

func (*Login) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("login_theme1").Parse(List["login/theme1"])

	if err != nil {
		logger.Error("Login GetTemplate Error: ", err)
	}

	return tmpl, "login_theme1"
}

func (*Login) GetAssetList() []string {
	return asserts
}

func (*Login) GetAsset(name string) ([]byte, error) {
	return Asset(name[1:])
}
