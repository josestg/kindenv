package kindenv

import (
	"fmt"
	"strings"
)

// Env is the type for the environment modes.
type Env uint8

// List of environment modes.
const (
	Develop Env = iota
	Testing
	Staging
	Sandbox
	Release // or Production.
)

// Unknown represents an unknown environment.
const Unknown Env = 255

var _texts = [Release + 1]string{
	Develop: "develop",
	Testing: "testing",
	Staging: "staging",
	Sandbox: "sandbox",
	Release: "release",
}

// Parse parses the string into a Env.
func Parse(s string) (Env, error) {
	for i, v := range _texts {
		if strings.EqualFold(s, v) {
			return Env(i), nil
		}
	}
	return Unknown, fmt.Errorf("kindenv: invalid kind %q", s)
}

func (e Env) String() string {
	if e > Release {
		return "unknown"
	}
	return _texts[e]
}
