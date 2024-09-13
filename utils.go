package zlogwrap

import (
	"fmt"
	"strings"
)

func toString(anything ...interface{}) string {
	var list []string
	for _, val := range anything {
		list = append(list, fmt.Sprintf("%v", val))
	}
	content := strings.Join(list, " ")
	return content
}
