package pizzeria

type pizzaCrust int

const (
	THIN pizzaCrust = iota
	THICK
	STUFFED
	GLUTEN_FREE
)

func (p pizzaCrust) String() string {
	val := []string{"thin", "thick", "stuffed", "gluten free"}

	if int(p) < len(val) {
		return val[int(p)]
	}

	return "unknown"
}
