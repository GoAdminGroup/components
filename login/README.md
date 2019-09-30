# GoAdmin login themes

how to use: 

```golang
package main

import (
	"github.com/GoAdminGroup/components/login/theme1"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
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

	examplePlugin := example.NewExample()

    // here to add login theme
	template.AddLoginComp(theme1.Get())

	if err := eng.AddConfigFromJson("./config.json").
		AddPlugins(adminPlugin, examplePlugin).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	r.GET("/admin/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	_ = r.Run(":9033")
}
```
