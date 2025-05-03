package sandbox

import (
	"bytes"
	"context"
	"os/exec"
	"testing"
	"time"
)

func Runner(t *testing.T, script string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "nsjail", "--really_quiet", "--config", "sandbox.cfg", "--", "/bin/bash", "-c", script)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
