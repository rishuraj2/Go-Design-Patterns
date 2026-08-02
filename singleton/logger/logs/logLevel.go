package logs

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (ll LogLevel) String() string {
	val := []string {"debug", "info", "warn", "error"}

	if int(ll) < len(val) {
		return val[int(ll)]
	}

	return "unknown"
}
