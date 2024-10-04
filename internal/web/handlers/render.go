package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/lafetz/htmx-starter/internal/web/views/layout"
)
func renderComponent(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}
func renderWithLayout(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return layout.Base("Sign in", template).Render(ctx.Request.Context(), ctx.Writer)
	
}
