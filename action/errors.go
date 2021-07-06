package action

import "errors"

var (
	ErrAPIKeyRequired      = errors.New("auth token is required")
	ErrProjectNameRequired = errors.New("project name is required")
	ErrAPIUrlRequired      = errors.New("api url is required")
)
