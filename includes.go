package testify

import (
	"fmt"
	"testing"

	"github.com/go-zoox/core-utils/array"
	"github.com/ttacon/chalk"
)

// Includes logs error if expected does not includes received.
func Includes[T comparable](t *testing.T, list []T, item T, message ...string) {
	messageX := "expected not includes received"
	if len(message) > 0 {
		messageX = message[0]
	}

	if !array.Includes(list, item) {
		listX := fmt.Sprintf("%#v", list)
		itemX := fmt.Sprintf("%#v", item)
		t.Errorf(
			"%s\nlist: %s\nitem: %s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
			chalk.Green.Color(listX),
			chalk.Red.Color(itemX),
		)
	}
}
