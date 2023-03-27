package greetings

import (
	"regexp"
	"testing"
)

// TestHelloEmtpy calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, helloErr := Hello("")

	if msg != "" || helloErr == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, helloErr)
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Patrick"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, helloErr := Hello(name)

	if !want.MatchString(msg) || helloErr != nil {
		t.Fatalf(`Hello("Patrick") = %q, %v, want match for %#q, nil`, msg, helloErr, want)
	}
}
