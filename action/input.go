package action

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Input represents the input parameters.
type Input struct {
	APIKey   string `env:"INPUT_METAL_AUTH_TOKEN"`
	Projects string `env:"INPUT_PROJECT_NAMES"`
	DryRun   bool   `env:"INPUT_DRY_RUN"`
}

// NewInput creates a new input from the environment variables.
func NewInput() (*Input, error) {
	input := &Input{}
	if err := env.Parse(input); err != nil {
		return nil, fmt.Errorf("parsing environment variables: %w", err)
	}

	return input, nil
}
