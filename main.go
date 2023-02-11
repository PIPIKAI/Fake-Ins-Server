package main

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/core"
)

// @title          Swagger Fake Ins API
// @version        1.0
// @description    Fake Ins的API文档
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:1016
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Cookie

func main() {
	core.Run()
}
