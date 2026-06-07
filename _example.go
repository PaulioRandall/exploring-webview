package main

import (
	"github.com/PaulioRandall/exploring-webview/app"
)

// example1 demonstrates minimal code to create an
// application with a single window.
func example1() {
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

// example2 is the same as example1 except statements have
// been split up into their simplest form. Each line of
// code is simpler and easier to understand but more
// verbose overall.
func example2() {
	// Specify options for our new app.
	myAppOptions := app.AppOptions{
		Name:    "Example 2",
		LogMode: app.LogModeSimple,
	}

	// Create an app with the options.
	myApp := app.NewApp(myAppOptions)

	// Specify some HTML to pass to our main window.
	html := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Example 2: Hello world</title>
  </head>
  <body>
    <p>Hello, world!</p>
  </body>
</html>`

	// Specify options for a new window (WebView).
	myWindowOptions := app.WindowOptions{
		ID:     "Main view",
		Width:  600,
		Height: 400,
		HTML:   html,
	}

	// Create the new window using the window options. In
	// this example it is our main and only window.
	myWindow := app.NewWindow(myWindowOptions)

	// Assign the window to our app. This will cause the
	// window to be shown to the user when our app is run.
	myApp.OpenWindow(myWindow)

	// Start the app. This blocks the thread until the app is
	// exited.
	e := myApp.Run()

	// If the app exited with an error, let the user know
	// somehow.
	if e != nil {
		panic(e)
	}
}
