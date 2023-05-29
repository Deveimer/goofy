package log

type level int

const (
	Fatal level = iota + 1
	Error
	Warn
	Info
	Debug
)

const (
	redColor    = 31
	yellowColor = 33
	blueColor   = 36
	normalColor = 37
)

func (l level) String() string {
	logLevel := "INFO"

	switch l {
	case Fatal:
		logLevel = "FATAL"
	case Error:
		logLevel = "ERROR"
	case Warn:
		logLevel = "WARN"
	case Debug:
		logLevel = "DEBUG"
	case Info:
		logLevel = "INFO"
	}

	return logLevel
}

func (l level) colorCode() int {
	colorCode := normalColor

	switch l {
	case Error, Fatal:
		colorCode = redColor
	case Warn:
		colorCode = yellowColor
	case Info:
		colorCode = blueColor
	case Debug:
		colorCode = normalColor
	}

	return colorCode
}
