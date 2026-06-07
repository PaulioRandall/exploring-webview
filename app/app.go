package app

type LogMode int

// TODO: Should we use more meaningful string names instead?
const (
	// LogModeError logs only errors being logged to output.
	LogModeError LogMode = iota

	// LogModeWarning extends LogModeError by logging
	// warnings.
	LogModeWarning

	// LogModeSimple extends LogModeWarning by concisely
	// logging major events such as startup, shutdown, and
	// opening and closing of windows.
	LogModeSimple

	// LogModeConcise extends LogModeSimple by concisely
	// logging minor events such as interactions between
	// Go and JS (e.g. bound Go function calls, eval, etc).
	LogModeConcise

	// LogModeVerbose extends LogModeConcise by logging
	// everything in detail.
	LogModeVerbose
)

type AppOptions struct {
	Name    string
	LogMode LogMode
}

type App struct {
	name    string
	logMode LogMode
}

func NewApp(options AppOptions) *App {
	a := &App{
		name: orElseString(options.Name, "Unnamed App"),
		// TODO: panic if bad log mode.
		logMode: options.LogMode,
	}

	a.logSimple("'%v' application started", a.name)
	return a
}

func (a *App) OpenWindow(w *Window) {
	// TODO

	a.logSimple("'%v' window opened in '%v' application", w.id, a.name)
}

func (a *App) Run() error {
	// TODO
	return nil
}

func (a *App) logSimple(format string, args ...any) {
	if a.logMode < LogModeSimple {
		return
	}

	// TODO: Should we really just panic?
	if _, e := log(format, args...); e != nil {
		panic(e)
	}
}
