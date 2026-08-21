package enum

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (this LogLevel) String() string {
	val := []string{"debug", "info", "warn", "error"}
	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
