# GoAdmin登录页面组件

提供多种预设登录页面主题，以及验证码方式。

## 使用

以下代码展示如何使用登录组件，请留意相关注释。

```golang
package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"	    
        _ "github.com/GoAdminGroup/go-admin/adapter/gin"
        _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
        // 引入theme2登录页面主题，如不用，可以不导入
        _ "github.com/GoAdminGroup/components/login/theme2"
	
	"github.com/GoAdminGroup/components/login"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

        // 使用登录页面组件
        login.Init(login.Config{
            Theme: "theme2",
            CaptchaDigits: 5, // 使用图片验证码，这里代表多少个验证码数字
            // 使用腾讯验证码，需提供appID与appSecret
            // TencentWaterProofWallData: login.TencentWaterProofWallData{
            //    AppID:"",
            //    AppSecret: "",
            // }   
        })

	if err := eng.AddConfigFromJson("./config.json").
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}
	
	// 载入对应验证码驱动，如没使用不用载入
	adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})	

	r.Static("/uploads", "./uploads")

	_ = r.Run(":9033")
}
```
