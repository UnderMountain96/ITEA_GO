package dotenv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(envFilePath string) error {
	f, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf("loadEnv: cannot open file: %q: %w", envFilePath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		name, value, err := parseEnvVar(line)
		if err != nil {
			return fmt.Errorf("loadEnv: invalid line: %q: %w", line, err)
		}

		if err := setEnvVar(name, value); err != nil {
			return fmt.Errorf("loadEnv: cannot set env var: %q: %w", name, err)
		}
	}

	return nil
}

func parseEnvVar(envVar string) (string, string, error) {
	parts := strings.Split(envVar, "=")
	if len(parts) < 2 {
		return "", "", errors.New("invalid value")
	}

	name, value := parts[0], parts[1]

	return name, value, nil
}

func setEnvVar(name, value string) error {
	return os.Setenv(name, value)
}
