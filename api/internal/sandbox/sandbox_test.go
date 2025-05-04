package sandbox_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestSandbox(t *testing.T) {
	flag.Parse()

	problemsMeta, err := sandbox.LoadProblemsMeta()
	require.NoError(t, err, "error loading problems meta")

	getMeta := func(slug string) *sandbox.ProblemMeta {
		meta, ok := problemsMeta[slug]
		require.True(t, ok, fmt.Sprintf("failed to find slug key in problemsMeta: %q", slug))
		return meta
	}

	cases := map[string]string{
		"calculate-total-sales-from-csv":     `awk -F',' '{sum += $3} END {printf "Total: %.2f\n", sum}' sales.csv`,
		"convert-dates-to-iso-format":        `sed -E 's#([0-9]{2})/([0-9]{2})/([0-9]{4})#\3-\1-\2#' dates.txt`,
		"extract-unique-users-from-api-logs": `jq -r 'select(.status == "error") | .user_id' api_logs.jsonl | sort -n | uniq`,
		"find-error-messages-in-logs":        `grep -i "error" logs.txt`,
		"list-all-usernames-from-json":       `jq -r '.[].username' users.json`,
	}

	for slug, script := range cases {
		t.Run(slug+" success", func(t *testing.T) {
			meta := getMeta(slug)

			result := sandbox.Sandbox(meta, script, &sandbox.SandboxOpts{Timeout: time.Second})

			assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "problems", slug+"_success.txt"))
		})
	}

	t.Run("single test case pass", func(t *testing.T) {
		meta := getMeta("calculate-total-sales-from-csv")

		result := sandbox.Sandbox(meta, "echo 'Total: 12.99'", &sandbox.SandboxOpts{Timeout: time.Second})

		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "single-test-pass.txt"))
	})
	t.Run("error timeout", func(t *testing.T) {
		meta := getMeta("calculate-total-sales-from-csv")

		result := sandbox.Sandbox(meta, "sleep 0.5", &sandbox.SandboxOpts{Timeout: 50 * time.Millisecond})

		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-timeout.txt"))
	})
	t.Run("error mismatch", func(t *testing.T) {
		meta := getMeta("calculate-total-sales-from-csv")

		result := sandbox.Sandbox(meta, "echo incorrect", &sandbox.SandboxOpts{Timeout: time.Second})

		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-mismatch.txt"))
	})
	t.Run("error runtime-error", func(t *testing.T) {
		meta := getMeta("calculate-total-sales-from-csv")

		result := sandbox.Sandbox(meta, "!&echo something", &sandbox.SandboxOpts{Timeout: time.Second})

		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-runtime-error.txt"))
	})
}

func assertGoldenResult(t *testing.T, result *sandbox.SandboxResult, filepath string) {
	actual, err := yaml.Marshal(result)
	require.NoError(t, err)

	if *testutil.Update {
		os.WriteFile(filepath, actual, 0644)
	}

	expected, err := os.ReadFile(filepath)
	require.NoError(t, err)
	require.Equal(t, string(actual), string(expected))
}
