package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lafetz/htmx-starter/internal/web/form"
	pages "github.com/lafetz/htmx-starter/internal/web/views/pages/signin"
)
func SigninGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	
		renderWithLayout(ctx,http.StatusOK,pages.Signin(form.SigninUser{}, "", ""))
		// p := pages.Signin(form.SigninUser{}, "", "")
		// err := layout.Base("Sign in", p).Render(c, c)
		// if err != nil {
		// 	// ServerError(w, r, err, logger)
		// 	return
		// }

	}
}
// func SigninPost() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	}
// }