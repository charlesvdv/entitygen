package gogen

import "strings"

func goifyIdentifier(name string) string {
	return strings.Title(name)
}
