package models

import (
	"strconv"
	"strings"
)

func GenerateOutput(output Output) string {
	s := []string{strconv.Itoa(len(output.Pizza.Ingredients))}
	for name := range output.Pizza.Ingredients {
		s = append(s, name)
	}
	return strings.Join(s, " ")
}
