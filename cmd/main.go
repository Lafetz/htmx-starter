package main

import (
	"log/slog"

	"github.com/lafetz/htmx-starter/internal/web"
)
func main(){
	
	app:=web.NewApp(8080,slog.Default())
	app.Run()
}