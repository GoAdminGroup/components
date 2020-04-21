package login

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/GoAdminGroup/components/login/theme1"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	captcha2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/captcha"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/dchest/captcha"
	"html/template"
	"strings"
	textTemplate "text/template"
	"time"
)

type Login struct {
	TencentWaterProofWallData TencentWaterProofWallData
	BackgroundColor           string
	LoginBtnColor             string
	CaptchaDigits             int
	CaptchaID                 string
	CaptchaImgSrc             string
	Theme                     int
}

type TencentWaterProofWallData struct {
	ID           string
	AppID        string
	AppSecretKey string
}

type Config struct {
	TencentWaterProofWallData TencentWaterProofWallData
	LoginBtnColor             string
	BackgroundColor           string
	CaptchaDigits             int
	Theme                     int
}

const (
	CaptchaDriverKeyTencent = "tencent"
	CaptchaDriverKeyDefault = "digits"

	CaptchaDisableDuration = time.Minute * 2
)

type CaptchaDataItem struct {
	Time time.Time
	Data string
	Num  int
}

type CaptchaData map[string]CaptchaDataItem

func (c *CaptchaData) Clean() {
	for key, value := range *c {
		if value.Time.Add(CaptchaDisableDuration).Before(time.Now()) {
			delete(*c, key)
		}
	}
}

var captchaData = make(CaptchaData)

type DigitsCaptcha struct{}

func (c *DigitsCaptcha) Validate(token string) bool {
	tokenArr := strings.Split(token, ",")
	if len(tokenArr) < 2 {
		return false
	}
	if v, ok := captchaData[tokenArr[1]]; ok {
		if v.Data == tokenArr[0] {
			delete(captchaData, tokenArr[1])
			return true
		} else {
			v.Num++
			captchaData[tokenArr[1]] = v
			return false
		}
	}
	return false
}

type TencentCaptcha struct{}

func (c *TencentCaptcha) Validate(token string) bool {
	return true
}

func Get(cfg ...Config) *Login {
	if len(cfg) > 0 {

		if cfg[0].CaptchaDigits != 0 && cfg[0].TencentWaterProofWallData.ID == "" {
			captchaData.Clean()
			captcha2.Add(CaptchaDriverKeyDefault, new(DigitsCaptcha))
		}

		if cfg[0].TencentWaterProofWallData.ID != "" {
			captcha2.Add(CaptchaDriverKeyTencent, new(TencentCaptcha))
		}

		return &Login{
			TencentWaterProofWallData: cfg[0].TencentWaterProofWallData,
			BackgroundColor:           utils.SetDefault(cfg[0].BackgroundColor, "", theme1.DefaultBackgroundColor),
			LoginBtnColor:             utils.SetDefault(cfg[0].LoginBtnColor, "", theme1.DefaultLoginBtnColor),
			CaptchaDigits:             cfg[0].CaptchaDigits,
			Theme:                     cfg[0].Theme,
		}
	}
	return &Login{BackgroundColor: theme1.DefaultBackgroundColor, LoginBtnColor: theme1.DefaultLoginBtnColor, Theme: 0}
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

	t := textTemplate.New("login_theme1").Delims("{%", "%}")
	t, err := t.Parse(theme1.List["login/theme1"])
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

func (l *Login) GetAssetList() []string               { return theme1.AssetsList }
func (l *Login) GetAsset(name string) ([]byte, error) { return theme1.Asset(name[1:]) }
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
