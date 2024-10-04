package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lafetz/htmx-starter/internal/web/handlers"
)

//go:embed static/*
var staticFiles embed.FS

func (a *App) initAppRoutes() {
	staticFs, err := fs.Sub(staticFiles, "static")
	if err != nil {
		println(err)
	}


    a.gin.GET("/static/*filepath",func(ctx *gin.Context) {
		http.StripPrefix("/static/", http.FileServer(http.FS(staticFs))).ServeHTTP(ctx.Writer,ctx.Request)
	})

	a.gin.GET("/signin",handlers.SigninGet())

}
