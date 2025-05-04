package sandbox

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

type ErrReason string

const (
	ErrReasonTimeout      = "timeout"
	ErrReasonMismatch     = "mismatch"
	ErrReasonRuntimeError = "runtime-error"
)

type SandboxResultStatus string

const (
	SandboxStatusAccepted SandboxResultStatus = "accepted"
	SandboxStatusRejected SandboxResultStatus = "rejected"
	SandboxStatusError    SandboxResultStatus = "error"
)

type SandboxResult struct {
	Status      SandboxResultStatus
	ExecTimeMs  int
	TestResults []SandboxTestResult
}

type SandboxTestResult struct {
	Test      ProblemMetaTest
	ErrReason ErrReason
	ExitCode  int
	Stdout    string
	Stderr    string
}

type SandboxOpts struct {
	Timeout time.Duration
	Timer   Timer
}

func Sandbox(meta *ProblemMeta, script string, opts *SandboxOpts) *SandboxResult {
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	status := SandboxStatusAccepted
	results := make([]SandboxTestResult, len(meta.Tests))

	opts.Timer.Start()
	for i := range meta.Tests {
		test := meta.Tests[i]

		var fullScript string
		if test.Setup != "" {
			fullScript = test.Setup + script // setup should have a \n at the end
		} else {
			fullScript = script
		}

		cmd := exec.CommandContext(ctx,
			"nsjail", "--really_quiet", "--config", nsjailCfgPath(),
			"--", "/bin/bash", "-c", fullScript,
		)

		var stdoutBytes, stderrBytes bytes.Buffer
		cmd.Stdout = &stdoutBytes
		cmd.Stderr = &stderrBytes

		err := cmd.Run()

		stdout := stdoutBytes.String()
		stderr := stderrBytes.String()

		exitCode := -1 // default for unknown exit codes
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			}
		} else {
			exitCode = 0
		}

		var errReason ErrReason
		if ctx.Err() == context.DeadlineExceeded {
			errReason = ErrReasonTimeout
			status = setStatus(status, SandboxStatusError)
		} else if err != nil {
			errReason = ErrReasonRuntimeError
			status = setStatus(status, SandboxStatusError)
		} else if test.Expected != stdout {
			errReason = ErrReasonMismatch
			status = setStatus(status, SandboxStatusRejected)
		}

		results[i] = SandboxTestResult{
			Test:      test,
			ErrReason: errReason,
			ExitCode:  exitCode,
			Stdout:    stdout,
			Stderr:    stderr,
		}
	}

	return &SandboxResult{
		Status:      status,
		ExecTimeMs:  opts.Timer.ElapsedMs(),
		TestResults: results,
	}
}

func nsjailCfgPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not get nsjail config dir")
	}
	return filepath.Join(filepath.Dir(filename), "sandbox.cfg")
}

func setStatus(current, next SandboxResultStatus) SandboxResultStatus {
	switch current {
	case SandboxStatusAccepted:
		return next // always promote
	case SandboxStatusRejected:
		if next == SandboxStatusError {
			return next // only promote to error
		} else {
			return current
		}
	case SandboxStatusError:
		return SandboxStatusError // already worst, don't backtrack
	default:
		panic(fmt.Sprintf("unexpected SandboxResultStatus: %s", current))
	}
}
