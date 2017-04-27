// 앱 실행되도록 함

package main

import (
	"github.com/letsget23/go-playground/projects/memo-hub/app"
	"github.com/letsget23/go-playground/projects/memo-hub/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
