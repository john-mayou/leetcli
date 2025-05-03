package sandbox

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

type ProblemMeta struct {
	Title      string     `yaml:"title"`
	Number     int        `yaml:"number"`
	Difficulty string     `yaml:"difficulty"`
	Prompt     string     `yaml:"prompt"`
	Tests      []TestCase `yaml:"tests"`
}

type TestCase struct {
	Name     string `yaml:"name"`
	Setup    string `yaml:"setup"`
	Expected string `yaml:"expected"`
}

func LoadProblemsMeta() (map[string]*ProblemMeta, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("could not get problems dir")
	}

	metaDir := filepath.Join(filepath.Dir(filename), "problems")
	metaFiles, err := os.ReadDir(metaDir)
	if err != nil {
		return nil, fmt.Errorf("could not read problems dir: %w", err)
	}

	problemsMeta := make(map[string]*ProblemMeta, len(metaFiles))

	for _, metaFile := range metaFiles {
		data, err := os.ReadFile(filepath.Join(metaDir, metaFile.Name()))
		if err != nil {
			return nil, fmt.Errorf("error reading meta file %q: %w", metaFile.Name(), err)
		}

		var meta ProblemMeta
		if err := yaml.Unmarshal(data, &meta); err != nil {
			return nil, fmt.Errorf("error unmarshaling yaml from %q: %w", metaFile.Name(), err)
		}

		slug := strings.TrimSuffix(metaFile.Name(), filepath.Ext(metaFile.Name()))
		problemsMeta[slug] = &meta
	}

	return problemsMeta, nil
}
