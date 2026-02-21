package input

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type commandRunner interface {
	Run(ctx context.Context, name string, args ...string) (string, error)
}

type execRunner struct{}

func (r execRunner) Run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func inputCommandError(tool string, err error, output string) error {
	trimmed := strings.TrimSpace(output)
	if trimmed == "" {
		return fmt.Errorf("%s input action failed: %w", tool, err)
	}
	return fmt.Errorf("%s input action failed: %w: %s", tool, err, trimmed)
}
