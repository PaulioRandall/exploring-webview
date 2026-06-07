package main

import (
	"github.com/PaulioRandall/exploring-webview/app"
)

func main() {
	myApp := app.NewApp(app.AppOptions{
		Name:    "Example 1",
		LogMode: app.LogModeSimple,
	})

	window := app.NewWindow(app.WindowOptions{
		ID:     "Main view",
		Width:  600,
		Height: 400,
		HTML: `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Example 1: Hello world</title>
  </head>
  <body>
    <p>Hello, world!</p>
  </body>
</html>`,
	})

	myApp.OpenWindow(window)

	if e := myApp.Run(); e != nil {
		panic(e)
	}
}
