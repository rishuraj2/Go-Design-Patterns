package model

type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

func (this LogLevel) String() string {
	val := []string{"info", "warn", "error"}

	if int(this) < len(val) {
		return val[int(this)]
	}
	return "unknown"
}
