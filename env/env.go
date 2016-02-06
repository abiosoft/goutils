// Package env implements a Environment Variables using []string.
package env

import (
	"bytes"
	"fmt"
	"strings"
)

// Env represents environment variables
type EnvVar []string

// Set sets the environment key to value.
func (e *EnvVar) Set(key, value string) {
	keyVal := key + "=" + value
	for i, v := range *e {
		env := strings.SplitN(v, "=", 2)
		if len(env) < 2 {
			continue
		}
		if env[0] == key {
			(*e)[i] = keyVal
			return
		}
	}
	*e = append(*e, keyVal)
}

// SetStr sets environment variable from key=val string format.
func (e EnvVar) SetStr(keyVal string) {
	s := strings.SplitN(keyVal, "=", 2)
	if len(s) == 2 {
		e.Set(s[0], s[1])
	}
}

// Get retrieves the environment variable key
func (e EnvVar) Get(key string) string {
	for _, v := range e {
		env := strings.SplitN(v, "=", 2)
		if len(env) < 2 {
			continue
		}
		if env[0] == key {
			return env[1]
		}
	}
	return ""
}

// String returns key=val format of the environment variables.
// Each on a line.
func (e EnvVar) String() string {
	b := bytes.NewBuffer(nil)
	for _, env := range e {
		// don't include invalid strings
		if len(strings.SplitN(env, "=", 2)) < 2 {
			continue
		}
		fmt.Fprintln(b, env)
	}
	return b.String()
}
