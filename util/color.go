package util

import "fmt"

var ColorMap = map[string]string{
	"Black":   "40",
	"Red":     "41",
	"Green":   "42",
	"Yellow":  "43",
	"Blue":    "44",
	"Magenta": "45",
	"Cyan":    "46",
	"White":   "47",
}

func PrintColorBlock(color, message string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", ColorMap[color], message)
}
