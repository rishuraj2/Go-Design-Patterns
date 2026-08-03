package pizzeria

import "strings"

type pizzaSauce int

const (
	MARINARA pizzaSauce = iota
	BBQ
	WHITE_ALFREDO
	PESTO
	NONE
)

func (p pizzaSauce) String() string {
	val := []string{"marinara", "bbq", "white alfredo", "pesto", "none"}

	if int(p) < len(val) {
		return val[int(p)]
	}

	return "unknown"
}

type pizzaSauces []pizzaSauce

func (p pizzaSauces) String() string {
	var sb strings.Builder

	for _, sauce := range p {
		sb.WriteString(sauce.String() + " ")
	}

	return sb.String()
}
