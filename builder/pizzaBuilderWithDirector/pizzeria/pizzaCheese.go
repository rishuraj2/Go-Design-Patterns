package pizzeria

import "strings"

type pizzaCheese int

const (
	MOZZARELLA pizzaCheese = iota
	PARMASAN
	PROVOLONE
	GORGONZOLA
	CHEDDAR
	VEGAN
)

func (p pizzaCheese) String() string {
	val := []string{"mozzarella", "parmasan", "provolone", "gorgonzola", "cheddar", "vegan"}

	if int(p) < len(val) {
		return val[int(p)]
	}

	return "unknown"
}

type pizzaCheeses []pizzaCheese

func (p pizzaCheeses) String() string {
	var sb strings.Builder

	for _, sauce := range p {
		sb.WriteString(sauce.String() + " ")
	}

	return sb.String()
}
