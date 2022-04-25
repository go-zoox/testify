package testify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ttacon/chalk"
)

// Type logs error if type is not equal with the reflect.Type of the given value.
// typ:
//	string
//	bool
//  int | int8 | int16 | int32 | int64
//  uint | uint8 | uint16 | uint32 | uint64
//  float32 | float64
//  complex64 | complex128
//  byte
//  rune
//  error
func Type(t *testing.T, typ string, value any, message ...string) {
	messageX := fmt.Sprintf("value if not type(%s)", typ)
	if len(message) > 0 {
		messageX = message[0]
	}

	if typ != reflect.TypeOf(value).String() {
		expectedX := fmt.Sprintf("%#v", typ)
		receivedX := fmt.Sprintf("%#v", reflect.TypeOf(value).String())
		t.Errorf(
			"%s\nexpected: %s\nreceived: %s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
			chalk.Green.Color(expectedX),
			chalk.Red.Color(receivedX),
		)
	}
}
