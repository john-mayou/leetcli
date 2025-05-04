package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/john-mayou/leetcli/internal/httpx"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/model"
)

type SubmitProblemBody struct {
	Slug string `json:"slug"`
	Type string `json:"type"`
	Code string `json:"code"`
}

func (h *Handler) SubmitProblem(w http.ResponseWriter, r *http.Request) {
	userID, ok := CtxUserID(r.Context())
	if !ok {
		http.Error(w, "unauthorized: missing user id", http.StatusUnauthorized)
		return
	}

	var body SubmitProblemBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON body %v", err), http.StatusBadRequest)
		return
	}

	problem, ok := h.Store.Problems[body.Slug]
	if !ok {
		http.Error(w, fmt.Sprintf("problem not found with slug: %q", body.Slug), http.StatusBadRequest)
		return
	}

	problemMeta, ok := h.Store.ProblemsMeta[body.Slug]
	if !ok {
		http.Error(w, fmt.Sprintf("problem meta not found with slug: %q", body.Slug), http.StatusBadRequest)
		return
	}

	result := sandbox.Sandbox(problemMeta, body.Code, &sandbox.SandboxOpts{
		Timeout: time.Second,
		Timer:   &sandbox.RealTimer{Now: h.Now},
	})

	switch body.Type {
	case "run":
	case "submit":
		status, err := mapSandboxStatusToSubmissionStatus(result.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = h.DBClient.CreateProblemSubmission(&model.ProblemSubmission{
			ID:         uuid.NewString(),
			ProblemID:  problem.ID,
			UserID:     userID,
			Status:     status,
			Code:       body.Code,
			ExecTimeMs: result.ExecTimeMs,
		})
		if err != nil {
			http.Error(w, fmt.Errorf("error creating problem submission: %w", err).Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("invalid type parameter: %q", body.Type), http.StatusBadRequest)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, result)
}

func mapSandboxStatusToSubmissionStatus(status sandbox.SandboxResultStatus) (model.ProblemSubmissionStatus, error) {
	switch status {
	case sandbox.SandboxStatusAccepted:
		return model.ProblemSubmissionStatusAccepted, nil
	case sandbox.SandboxStatusRejected:
		return model.ProblemSubmissionStatusRejected, nil
	case sandbox.SandboxStatusError:
		return model.ProblemSubmissionStatusError, nil
	default:
		return "", fmt.Errorf("unknown sandbox status: %q", status)
	}
}
