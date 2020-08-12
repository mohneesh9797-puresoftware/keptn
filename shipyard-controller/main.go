package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keptn/keptn/shipyard-controller/api"
	docs "github.com/keptn/keptn/shipyard-controller/docs" // docs is generated by Swag CLI, you have to import it.
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Shipyard Controller API
// @version 1.0
// @description This is the API documentation of the Shipyard Controller.

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name x-token

// @contact.name Keptn Team
// @contact.url http://www.keptn.sh

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	router := gin.Default()

	if os.Getenv("GIN_MODE") == "release" {
		docs.SwaggerInfo.Version = os.Getenv("version")
		docs.SwaggerInfo.BasePath = "/api/shipyard-controller/v1"
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	apiV1 := router.Group("/v1")
	apiV1.GET("/event/triggered/:eventType", api.GetTriggeredEvents)

	apiV1.POST("/event", api.HandleEvent)

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
