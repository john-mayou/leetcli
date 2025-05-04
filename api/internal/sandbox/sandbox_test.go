package sandbox_test

import (
	"flag"
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

	cases := map[string]string{
		"calculate-total-sales-from-csv":     `awk -F',' '{sum += $3} END {printf "Total: %.2f\n", sum}' sales.csv`,
		"convert-dates-to-iso-format":        `sed -E 's#([0-9]{2})/([0-9]{2})/([0-9]{4})#\3-\1-\2#' dates.txt`,
		"extract-unique-users-from-api-logs": `jq -r 'select(.status == "error") | .user_id' api_logs.jsonl | sort -n | uniq`,
		"find-error-messages-in-logs":        `grep -i "error" logs.txt`,
		"list-all-usernames-from-json":       `jq -r '.[].username' users.json`,
	}

	for slug, script := range cases {
		t.Run(slug+" success", func(t *testing.T) {
			meta, ok := problemsMeta[slug]
			require.True(t, ok)

			result := sandbox.Sandbox(meta, script, defaultSandboxOpts(t))

			assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "problems", slug+"_success.txt"))
		})
	}

	calcTotalMeta, ok := problemsMeta["calculate-total-sales-from-csv"]
	require.True(t, ok)

	t.Run("single test case pass", func(t *testing.T) {
		result := sandbox.Sandbox(calcTotalMeta, "echo 'Total: 12.99'", defaultSandboxOpts(t))
		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "single-test-pass.txt"))
	})
	t.Run("error timeout", func(t *testing.T) {
		result := sandbox.Sandbox(calcTotalMeta, "sleep 0.5", &sandbox.SandboxOpts{Timeout: 50 * time.Millisecond, Timer: &sandbox.FakeTimer{FixedMs: 1}})
		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-timeout.txt"))
	})
	t.Run("error mismatch", func(t *testing.T) {
		result := sandbox.Sandbox(calcTotalMeta, "echo incorrect", defaultSandboxOpts(t))
		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-mismatch.txt"))
	})
	t.Run("error runtime-error", func(t *testing.T) {
		result := sandbox.Sandbox(calcTotalMeta, "!&echo something", defaultSandboxOpts(t))
		assertGoldenResult(t, result, filepath.Join("testdata", "sandbox", "error-runtime-error.txt"))
	})
}

func defaultSandboxOpts(t *testing.T) *sandbox.SandboxOpts {
	t.Helper()

	return &sandbox.SandboxOpts{
		Timeout: time.Second,
		Timer:   &sandbox.FakeTimer{FixedMs: 1},
	}
}

func assertGoldenResult(t *testing.T, result *sandbox.SandboxResult, filepath string) {
	t.Helper()

	actual, err := yaml.Marshal(result)
	require.NoError(t, err)

	if *testutil.Update {
		require.NoError(t, os.WriteFile(filepath, actual, 0644))
	}

	expected, err := os.ReadFile(filepath)
	require.NoError(t, err)
	require.Equal(t, string(actual), string(expected))
}
