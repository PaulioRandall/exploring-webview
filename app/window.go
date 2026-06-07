package app

type WindowOptions struct {
	ID     string
	Width  int
	Height int
	HTML   string
}

type Window struct {
	id     string
	width  int
	height int
	html   string
}

func NewWindow(options WindowOptions) *Window {
	w := &Window{
		id:     options.ID,
		width:  options.Width,
		height: options.Height,
		html:   options.HTML,
	}

	return w
}
