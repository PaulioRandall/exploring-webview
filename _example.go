package main

import (
	"github/PaulioRandall/exploring-webview/app"
	"github/PaulioRandall/exploring-webview/window"
)

// example1 demonstrates minimal code to create an
// application with a single window.
func example1() {
	app := app.New(app.Options{
		Name: "Example Application",
	})

	win := window.New(window.Options{
		Width: 600,
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

// example2 is the same as example1 except statements have
// been split up into their simplest form. Each line of
// code is simpler and easier to understand but more
// verbose overall.
func example2() {
	// Specify options for our new app.
	appOptions := app.Options{
		Name: "Example Application"
	}

	// Create an app with the options.
	app := app.New(appOptions)

	// Specify options for a new window (WebView).
	winOptions := window.Options{
		Width: 600,
		height: 400,
		HTML: `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Example 2: Hello world</title>
  </head>
  <body>
    <p>Hello, world!</p>
  </body>
</html>`,
	}

	// Create the new window using the window options. In
	// this example it is our main and only window.
	win := window.New(winOptions)
	
	// Assign the window to our app. This will cause the
	// window to be shown to the user when our app is run.
	app.OpenWindow(win)

	// Start the app. This blocks the thread until the app is
	// exited.
	e := app.Run();

	// If the app exited with an error, let the user know
	// somehow.
	if e != nil {
		panic(e)
	}
}
