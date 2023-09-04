// controllers/test.controller.go

package controllers

import (
	"lixIQ/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestController, test.
type TestController struct {
	emailService services.EmailService
}

// NewTestController,
func NewTestController(emailService services.EmailService) TestController {
	return TestController{emailService}
}

// TestSendEmail, test amaçlı e-posta gönderimini gerçekleştirir.
func (tc *TestController) TestSendEmail(ctx *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	data := make(map[string]interface{})
	data["Email"] = request.Email
	// Test e-mail sending
	success, err := tc.emailService.SendEmail(request.Email, "Test Email", "test_mail_template.html", data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error sending test email"})
		return
	}

	if success {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Test email sent successfully"})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Test email could not be sent"})
	}
}
