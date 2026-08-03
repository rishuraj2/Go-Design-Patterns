package email

type EmailPriority int

const (
	LOW EmailPriority = iota
	NORMAL
	URGENT
)

func (e EmailPriority) String() string {
	val := []string{"low", "normal", "urgent"}

	if int(e) < len(val) {
		return val[int(e)]
	}

	return "unknown"
}
