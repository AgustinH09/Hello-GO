package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if err != nil || !want.MatchString(msg) {
		t.Fatalf(
			`Hello("%q") = %q, %v, want for %#q, nil`,
			name,
			msg,
			err,
			want,
		)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil || msg != "" {
		t.Fatalf(`Hello("") = %q, %v, want ""`, msg, err)
	}
}
