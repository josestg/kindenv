package kindenv

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		s string
		e Env
	}{
		{"develop", Develop},
		{"testing", Testing},
		{"staging", Staging},
		{"sandbox", Sandbox},
		{"release", Release},
		{"unknown", Unknown},
		{"abc_def", Unknown},
	}

	for _, tt := range tests {
		k, _ := Parse(tt.s)
		if k != tt.e {
			t.Errorf("got %s, want %s", k, tt.e)
		}
	}
}

func TestEnv_String(t *testing.T) {
	tests := []struct {
		e Env
		s string
	}{
		{Develop, "develop"},
		{Testing, "testing"},
		{Staging, "staging"},
		{Sandbox, "sandbox"},
		{Release, "release"},
		{Unknown, "unknown"},
		{Release + 1, "unknown"},
	}

	for _, tt := range tests {
		s := tt.e.String()
		if s != tt.s {
			t.Errorf("got %s, want %s", s, tt.s)
		}
	}
}
