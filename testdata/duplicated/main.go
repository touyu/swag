package composition

import (
	"github.com/gin-gonic/gin"

	"github.com/touyu/swag/testdata/duplicated/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	r := gin.New()
	r.GET("/testapi/get-foo", api.GetFoo)
	r.POST("/testapi/post-bar", api.PostBar)
	r.Run()
}
