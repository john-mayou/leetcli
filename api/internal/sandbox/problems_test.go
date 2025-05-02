package sandbox_test

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("UPDATE", false, "update golden files")

func TestLoadProblemsMeta(t *testing.T) {
	flag.Parse()

	meta, err := sandbox.LoadProblemsMeta()
	require.NoError(t, err)

	actual, err := json.MarshalIndent(meta, "", "  ")
	require.NoError(t, err)

	golden := filepath.Join("testdata", "loadProblemsMeta.txt")
	if *update {
		os.WriteFile(golden, actual, 0644)
	}

	expected, err := os.ReadFile(golden)
	require.NoError(t, err)
	require.Equal(t, actual, expected)
}
