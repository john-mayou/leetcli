package sandbox_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestLoadProblemsMeta(t *testing.T) {
	flag.Parse()

	meta, err := sandbox.LoadProblemsMeta()
	require.NoError(t, err)

	actual, err := yaml.Marshal(meta)
	require.NoError(t, err)

	golden := filepath.Join("testdata", "LoadProblemsMeta.txt")
	if *testutil.Update {
		require.NoError(t, os.WriteFile(golden, actual, 0644))
	}

	expected, err := os.ReadFile(golden)
	require.NoError(t, err)
	require.Equal(t, string(actual), string(expected))
}
