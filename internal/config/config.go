package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	Backend    string
	JSONPath   string
	SQLitePath string
}

func Load() Config {
	cfg := Config{
		Backend:    "json",
		JSONPath:   "todos.json",
		SQLitePath: "todos.db",
	}

	values := readEnvFile(".env")

	if v, ok := values["BACKEND"]; ok && v != "" {
		cfg.Backend = v
	}
	if v, ok := values["JSON_PATH"]; ok && v != "" {
		cfg.JSONPath = v
	}
	if v, ok := values["SQLITE_PATH"]; ok && v != "" {
		cfg.SQLitePath = v
	}

	return cfg
}

func readEnvFile(path string) map[string]string {
	values := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return values
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		value = strings.Trim(value, `"`)
		values[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	if err := scanner.Err(); err != nil {
		return values
	}
	return values
}
