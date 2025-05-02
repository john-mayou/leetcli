package handler_test

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/john-mayou/leetcli/handler"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("UPDATE", false, "update golden files")

func TestGetProblems(t *testing.T) {
	flag.Parse()

	h := handler.NewTestHandler(&handler.HandlerOpts{Store: &handler.Store{
		Problems:     map[string]*model.Problem{"slug": testutil.FakeProblem()},
		ProblemsMeta: map[string]*sandbox.ProblemMeta{"slug": testutil.FakeProblemMeta()},
	}})

	req := httptest.NewRequest("GET", "/problems", nil)
	r := httptest.NewRecorder()
	h.GetProblems(r, req)

	require.Equal(t, http.StatusOK, r.Code)

	var parsed interface{}
	err := json.Unmarshal(r.Body.Bytes(), &parsed)
	require.NoError(t, err)

	actual, err := json.MarshalIndent(parsed, "", "  ")
	require.NoError(t, err)

	golden := "testdata/GetProblems.txt"
	if *update {
		os.WriteFile(golden, actual, 0644)
	}

	expected, err := os.ReadFile(golden)
	require.NoError(t, err)
	require.Equal(t, actual, expected)
}
