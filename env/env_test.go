package env

import (
	"strings"
	"testing"
)

func TestEnvVar(t *testing.T) {
	var vars EnvVar

	vars.Set("IDE", "vim")
	if ide := vars.Get("IDE"); ide != "vim" {
		t.Errorf("Expected vim found %s", ide)
	}
	if str := vars.String(); strings.TrimSpace(str) != "IDE=vim" {
		t.Errorf("Expected IDE=vim found %s", str)
	}

	// invalid strings
	vars = append(vars, "SOMEVALS")
	if str := vars.String(); strings.TrimSpace(str) != "IDE=vim" {
		t.Errorf("Expected IDE=vim found %s", str)
	}
	vars.SetStr("SOME STRING")
	if str := vars.String(); strings.TrimSpace(str) != "IDE=vim" {
		t.Errorf("Expected IDE=vim found %s", str)
	}

	if e := vars.Get("UNKNOWN_KEY"); e != "" {
		t.Errorf("Expected '' found %s", e)
	}

	vars.Set("KEY", "val")
	if key := vars.Get("KEY"); key != "val" {
		t.Errorf("Expected vim found %s", key)
	}

	vars.Set("KEY", "val")
	if key := vars.Get("KEY"); key != "val" {
		t.Errorf("Expected vim found %s", key)
	}

	vars.SetStr("KEY=val1")
	if key := vars.Get("KEY"); key != "val1" {
		t.Errorf("Expected val1 found %s", key)
	}

	if l := len(vars); l != 3 {
		t.Errorf("Expected 3 found %d", l)
	}
	if str := vars.String(); !strings.Contains(str, "IDE=vim") || !strings.Contains(str, "KEY=val1") {
		t.Errorf("Must contain IDE=vim and KEY=val1")
	}

}
