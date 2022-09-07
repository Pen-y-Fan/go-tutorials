package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosNames calls greetings.Hellos with names, checking
// for a valid return value.
func TestHellosNames(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Unexpected error: Hellos(%v) = %q, %v`, names, messages, err)
	}

	if len(names) != len(messages) {
		t.Fatalf(`Names and messages are not the same length: %v = %q`, names, messages)
	}
	
	for _, name := range names {
		_, ok := messages[name]
		if !ok {
			t.Fatalf(`Name (%v) not found: Hellos(%v) = %q`, name, names, messages)
		}
	}
}

// TestHellosEmpty calls greetings.Hellos with an empty slice,
// checking for an error.
func TestHellosEmpty(t *testing.T) {
	names := []string{""}
	messages, err := Hellos(names)
	if messages != nil || err == nil {
		t.Fatalf(`Hellos("") = %q, %v, want "", error`, messages, err)
	}
}
