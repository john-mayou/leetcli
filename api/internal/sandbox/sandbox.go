package sandbox

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

type ErrReason string

const (
	ErrReasonTimeout      = "timeout"
	ErrReasonMismatch     = "mismatch"
	ErrReasonRuntimeError = "runtime-error"
)

type SandboxResult struct {
	Success     bool
	TestResults []SandboxTestResult
}

type SandboxTestResult struct {
	Test      ProblemMetaTest
	ErrReason string
	ExitCode  int
	Stdout    string
	Stderr    string
}

type SandboxOpts struct {
	Timeout time.Duration
}

func Sandbox(meta *ProblemMeta, script string, opts *SandboxOpts) *SandboxResult {
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	success := true
	results := make([]SandboxTestResult, len(meta.Tests))
	for i := range meta.Tests {
		test := meta.Tests[i]

		var fullScript string
		if test.Setup != "" {
			fullScript = test.Setup + script // setup should have a \n at the end
		} else {
			fullScript = script
		}

		cmd := exec.CommandContext(ctx,
			"nsjail", "--really_quiet", "--config", "sandbox.cfg",
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

		errReason := ""
		if ctx.Err() == context.DeadlineExceeded {
			errReason = ErrReasonTimeout
		} else if err != nil {
			errReason = ErrReasonRuntimeError
		} else if test.Expected != stdout {
			errReason = ErrReasonMismatch
		}

		if exitCode != 0 || errReason != "" {
			success = false
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
		Success:     success,
		TestResults: results,
	}
}
