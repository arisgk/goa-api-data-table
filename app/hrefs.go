// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Data Table": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/arisgk/goa-api-data-table/design
// --out=$(GOPATH)/src/github.com/arisgk/goa-api-data-table
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/users/%v", paramuserID)
}