package pizzeria

type pizzaSize int

const (
	SMALL pizzaSize = iota
	MEDIUM
	LARGE
)

func (p pizzaSize) String() string {
	val := []string{"small", "medium", "large"}

	if int(p) < len(val) {
		return val[int(p)]
	}

	return "unknown"
}
