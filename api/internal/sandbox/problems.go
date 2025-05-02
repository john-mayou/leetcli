package sandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ProblemMeta struct {
	Number     int    `json:"number"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Difficulty string `json:"difficulty"`
	Prompt     string `json:"prompt"`
	Input      string `json:"input"`
	Expected   string `json:"expected"`
}

// for the metadata.json file
type metadata struct {
	Number       int    `json:"number"`
	Title        string `json:"title"`
	Difficulty   string `json:"difficulty"`
	InputFile    string `json:"input_file"`
	ExpectedFile string `json:"expected_file"`
}

func LoadProblemsMeta() (map[string]*ProblemMeta, error) {
	problemsDir, err := getProblemsDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		return nil, fmt.Errorf("could not read problems dir: %w", err)
	}

	problems := make(map[string]*ProblemMeta, len(entries))

	for _, entry := range entries {
		problemDir := filepath.Join(problemsDir, entry.Name())

		metaBytes, err := os.ReadFile(filepath.Join(problemDir, "metadata.json"))
		if err != nil {
			return nil, fmt.Errorf("error reading metadata file from %s: %w", problemDir, err)
		}

		var meta metadata
		err = json.Unmarshal(metaBytes, &meta)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling metadata from %s: %w", problemDir, err)
		}

		input, err := os.ReadFile(filepath.Join(problemDir, meta.InputFile))
		if err != nil {
			return nil, fmt.Errorf("error reading input file data from %s: %w", problemDir, err)
		}

		expected, err := os.ReadFile(filepath.Join(problemDir, meta.ExpectedFile))
		if err != nil {
			return nil, fmt.Errorf("error reading expected file data from %s: %w", problemDir, err)
		}

		prompt, err := os.ReadFile(filepath.Join(problemDir, "prompt.md"))
		if err != nil {
			return nil, fmt.Errorf("error reading prompt file from %s: %w", problemDir, err)
		}

		slug := entry.Name()

		problems[slug] = &ProblemMeta{
			Number:     meta.Number,
			Title:      meta.Title,
			Slug:       slug,
			Difficulty: meta.Difficulty,
			Prompt:     strings.TrimSpace(string(prompt)),
			Input:      strings.TrimSpace(string(input)),
			Expected:   strings.TrimSpace(string(expected)),
		}
	}

	return problems, nil
}

func getProblemsDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("could not get problems dir")
	}
	return filepath.Join(filepath.Dir(filename), "problems"), nil
}
