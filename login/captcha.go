package login

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	CaptchaDriverKeyTencent = "tencent"
	CaptchaDriverKeyDefault = "digits"

	CaptchaDisableDuration = time.Minute * 2
)

type CaptchaDataItem struct {
	Time time.Time `json:"time"`
	Data string    `json:"data"`
	Num  int       `json:"num"`
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

type TencentCaptcha struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type TencentCaptchaRes struct {
	Response  string `json:"response"`
	EvilLevel string `json:"evil_level"`
	ErrMsg    string `json:"err_msg"`
}

func (c *TencentCaptcha) Validate(token string) bool {

	u := "https://ssl.captcha.qq.com/ticket/verify?"

	tokenArr := strings.Split(token, ",")
	if len(tokenArr) < 2 {
		return false
	}

	v := url.Values{
		"aid":          {c.AppID},
		"AppSecretKey": {c.AppSecret},
		"Ticket":       {tokenArr[0]},
		"Randstr":      {tokenArr[1]},
		"UserIP":       {"127.0.0.1"},
	}

	req, err := http.NewRequest("GET", u+v.Encode(), nil)
	if err != nil {
		return false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	defer func() {
		_ = res.Body.Close()
	}()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return false
	}

	var captchaRes TencentCaptchaRes
	err = json.Unmarshal(body, &captchaRes)

	if err != nil {
		logger.Debug("tencent captch validate response: ", captchaRes)
		return false
	}

	return captchaRes.Response == "1"
}
