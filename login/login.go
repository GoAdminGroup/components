package login

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	textTemplate "text/template"
	"time"

	"github.com/GoAdminGroup/components/login/theme1"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	captcha2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/captcha"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/dchest/captcha"
)

var themes = map[string]Theme{
	"theme1": new(theme1.Theme1),
}

func Register(key string, theme Theme) {
	if _, ok := themes[key]; ok {
		panic("duplicate login theme")
	}
	themes[key] = theme
}

type Login struct {
	TencentWaterProofWallData TencentWaterProofWallData `json:"tencent_water_proof_wall_data"`
	CaptchaDigits             int                       `json:"captcha_digits"`
	CaptchaID                 string                    `json:"captcha_id"`
	CaptchaImgSrc             string                    `json:"captcha_img_src"`
	Theme                     string                    `json:"theme"`
}

type TencentWaterProofWallData struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Config struct {
	TencentWaterProofWallData TencentWaterProofWallData `json:"tencent_water_proof_wall_data"`
	CaptchaDigits             int                       `json:"captcha_digits"`
	Theme                     string                    `json:"theme"`
}

func Init(cfg ...Config) {
	template2.AddLoginComp(Get(cfg...))
}

func Get(cfg ...Config) *Login {
	if len(cfg) > 0 {

		if cfg[0].CaptchaDigits != 0 && cfg[0].TencentWaterProofWallData.AppID == "" {
			captchaData.Clean()
			captcha2.Add(CaptchaDriverKeyDefault, new(DigitsCaptcha))
		}

		if cfg[0].TencentWaterProofWallData.AppID != "" {
			captcha2.Add(CaptchaDriverKeyTencent, &TencentCaptcha{
				AppID:     cfg[0].TencentWaterProofWallData.AppID,
				AppSecret: cfg[0].TencentWaterProofWallData.AppSecret,
			})
		}

		if cfg[0].Theme == "" {
			cfg[0].Theme = "theme1"
		}

		return &Login{
			TencentWaterProofWallData: cfg[0].TencentWaterProofWallData,
			CaptchaDigits:             cfg[0].CaptchaDigits,
			Theme:                     cfg[0].Theme,
		}
	}
	return &Login{Theme: "theme1"}
}

func byteToStr(b []byte) string {
	s := ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%v", b[i])
	}
	return s
}

func (l *Login) GetTemplate() (*template.Template, string) {

	if l.CaptchaDigits != 0 {
		id := utils.Uuid(10)
		digitByte := captcha.RandomDigits(l.CaptchaDigits)
		captchaData[id] = CaptchaDataItem{
			Data: byteToStr(digitByte),
			Time: time.Now(),
			Num:  0,
		}
		img := captcha.NewImage(id, digitByte, 110, 34)
		buf := new(bytes.Buffer)
		_, _ = img.WriteTo(buf)
		l.CaptchaID = id
		l.CaptchaImgSrc = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
	}

	t := textTemplate.New("login").Delims("{%", "%}")
	t, err := t.Parse(themes[l.Theme].GetHTML())
	if err != nil {
		logger.Error("login component, get template parse error: ", err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, l)
	if err != nil {
		logger.Error("login component, get template execute error: ", err)
	}

	tmpl, err := template.New("login").
		Funcs(login.DefaultFuncMap).
		Parse(buf.String())

	if err != nil {
		logger.Error("login component, get template error: ", err)
	}

	return tmpl, "login"
}

func (l *Login) GetAssetList() []string               { return themes[l.Theme].GetAssetList() }
func (l *Login) GetAsset(name string) ([]byte, error) { return themes[l.Theme].GetAsset(name[1:]) }
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

type Theme interface {
	GetAssetList() []string
	GetAsset(name string) ([]byte, error)
	GetHTML() string
}
