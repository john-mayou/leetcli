package handler_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/john-mayou/leetcli/handler"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitProblem(t *testing.T) {
	flag.Parse()

	runCases := map[string]struct {
		Code   string
		Golden string
	}{
		"run accepted": {
			Code:   "echo 'Hello world'",
			Golden: "run-accepted.txt",
		},
		"run rejected": {
			Code:   "echo 'Hello'",
			Golden: "run-rejected.txt",
		},
		"run error": {
			Code:   "!&echo 'Hello world'",
			Golden: "run-error.txt",
		},
	}

	for tcName, tc := range runCases {
		t.Run(tcName, func(*testing.T) {
			req := buildSubmitProblemBody(t, &handler.SubmitProblemBody{Slug: "hello-world", Type: "run", Code: tc.Code})
			r := httptest.NewRecorder()

			handler := buildSubmitProblemHandler(t)
			handler.SubmitProblem(r, req)

			assertGoldenResult(t, r, filepath.Join("testdata", "SubmitProblem", tc.Golden))
		})
	}

	submitCases := map[string]struct {
		Code           string
		Golden         string
		ExpectedStatus model.ProblemSubmissionStatus
	}{
		"submit accepted": {
			Code:           "echo 'Hello world'",
			Golden:         "submit-accepted.txt",
			ExpectedStatus: model.ProblemSubmissionStatusAccepted,
		},
		"submit rejected": {
			Code:           "echo 'Hello'",
			Golden:         "submit-rejected.txt",
			ExpectedStatus: model.ProblemSubmissionStatusRejected,
		},
		"submit error": {
			Code:           "!&echo 'Hello world'",
			Golden:         "submit-error.txt",
			ExpectedStatus: model.ProblemSubmissionStatusError,
		},
	}

	for tcName, tc := range submitCases {
		t.Run(tcName, func(*testing.T) {
			req := buildSubmitProblemBody(t, &handler.SubmitProblemBody{Slug: "hello-world", Type: "submit", Code: tc.Code})
			r := httptest.NewRecorder()

			handler := buildSubmitProblemHandler(t)
			handler.SubmitProblem(r, req)

			// assert golden
			assertGoldenResult(t, r, filepath.Join("testdata", "SubmitProblem", tc.Golden))

			// assert submission
			subs, err := handler.DBClient.ListProblemSubmissions()
			require.NoError(t, err)
			require.Len(t, subs, 1)
			assert.Equal(t, tc.ExpectedStatus, subs[0].Status)
			assert.Equal(t, tc.Code, subs[0].Code)
			assert.Equal(t, 0, subs[0].ExecTimeMs)
		})
	}
}

func buildSubmitProblemHandler(t *testing.T) *handler.Handler {
	dbClient := testutil.SetupTestClient(t)

	problem := testutil.FakeProblem()
	problem.Slug = "hello-world"
	problem, err := dbClient.CreateProblem(problem)
	require.NoError(t, err)

	return handler.NewTestHandler(&handler.HandlerOpts{
		Now:      func() time.Time { return time.Unix(0, 0) },
		DBClient: dbClient,
		Store: &handler.Store{
			Problems: map[string]*model.Problem{
				"hello-world": problem,
			},
			ProblemsMeta: map[string]*sandbox.ProblemMeta{
				"hello-world": {
					Title:      "Hello World",
					Number:     1,
					Difficulty: "easy",
					Prompt:     "Echo 'Hello world' using `echo`",
					Tests: []sandbox.ProblemMetaTest{{
						Name:     "Test 1",
						Setup:    "",
						Expected: "Hello world",
					}},
				},
			},
		},
	})
}

func buildSubmitProblemBody(t *testing.T, body *handler.SubmitProblemBody) *http.Request {
	jsonBytes, err := json.Marshal(body)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/submit", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func assertGoldenResult(t *testing.T, r *httptest.ResponseRecorder, filepath string) {
	t.Helper()

	// unmarshal
	var result sandbox.SandboxResult
	err := json.Unmarshal(r.Body.Bytes(), &result)
	require.NoError(t, err)

	// marshal with indent so its readable
	actual, err := json.MarshalIndent(result, "", "  ")
	require.NoError(t, err)

	// update if wanted
	if *testutil.Update {
		require.NoError(t, os.WriteFile(filepath, actual, 0644))
	}

	// assert
	expected, err := os.ReadFile(filepath)
	require.NoError(t, err)
	require.Equal(t, string(actual), string(expected))
}
