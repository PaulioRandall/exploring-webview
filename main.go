package main

import (
	"github/PaulioRandall/exploring-webview/app"
	"github/PaulioRandall/exploring-webview/window"
)

func main() {

	app := app.New(app.Options{
		Name: "Example Application",
	})

	win := window.New(window.Options{
		Width:  600,
		height: 400,
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

	app.OpenWindow(win)

	if e := app.Run(); e != nil {
		panic(e)
	}
}
