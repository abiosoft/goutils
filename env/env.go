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
