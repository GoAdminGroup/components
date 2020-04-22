# GoAdmin Login Theme Components

Offers a variety of default login page theme, and the ways of verification code.

## How To

The following code shows how to use the login module, please pay attention to the relevant comments.

```golang
package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"	    
        _ "github.com/GoAdminGroup/go-admin/adapter/gin"
        _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
        // import the theme2 login theme, if you don`t use, don`t import
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

        // use the login theme component
        login.Init(login.Config{
            Theme: "theme2", // theme name
            CaptchaDigits: 5, // Use captcha images, here on behalf of how many authentication code Numbers
            // Use tencent verification code, need to offer appID and appSecret
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
	
	// load the CAPTCHA driver if you use it
	adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})	

	r.Static("/uploads", "./uploads")

	_ = r.Run(":9033")
}
```
