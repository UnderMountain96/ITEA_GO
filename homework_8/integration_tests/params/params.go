package params

import (
	"fmt"
	"strings"
)

type MapParams map[string]string

func (m *MapParams) String() string {
	return fmt.Sprintf("%v", *m)
}

func (m *MapParams) Set(value string) error {
	parts := strings.Split(value, "=")
	if len(parts) != 2 {
		return fmt.Errorf("invalid map format: %s", value)
	}

	if *m == nil {
		*m = make(map[string]string)
	}

	(*m)[parts[0]] = parts[1]
	return nil
}
