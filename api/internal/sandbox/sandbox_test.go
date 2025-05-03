package sandbox_test

import (
	"strings"
	"testing"

	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	stdout, stderr, err := sandbox.Runner(t, "echo hello")
	assert.NoError(t, err)
	assert.Equal(t, "", stderr)
	assert.Equal(t, "hello", strings.TrimSpace(stdout))
}

// 1: grep -i "error" logs.txt
// 2: sed -E 's#([0-9]{2})/([0-9]{2})/([0-9]{4})#\3-\1-\2#' dates.txt
// 3: awk -F',' '{sum += $3} END {printf "Total: %.2f\n", sum}' sales.csv
// 4: jq -r '.[].username' users.json
// 5: jq -r 'select(.status == "error") | .user_id' api_logs.jsonl | sort -n | uniq
