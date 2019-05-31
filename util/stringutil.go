package util
import (
	"strings"
)

func TabToSpace(input string) string {
	replacer := strings.NewReplacer("\t", "    ")
	return replacer.Replace(input)
}
