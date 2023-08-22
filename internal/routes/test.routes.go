// routes/test.routes.go
package routes

import (
	"lixIQ/backend/internal/controllers"
	"lixIQ/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// TestRouteController,
type TestRouteController struct {
	testController controllers.TestController
}

// NewTestRouteController
func NewTestRouteController(testController controllers.TestController) TestRouteController {
	return TestRouteController{testController}
}

// TestRoute, route defination
func (trc *TestRouteController) TestRoute(rg *gin.RouterGroup, emailService services.EmailService) {
	router := rg.Group("/test")

	router.POST("/mailservice", trc.testController.TestSendEmail)
}
