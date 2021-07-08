package main

import (
	"net/http"

	"github.com/weaveworks/metal-janitor-action/action"
)

func main() {
	action.Log("running metal janitor action")

	input, err := action.NewInput()
	if err != nil {
		action.LogErrorAndExit("failed parsing action input: %s", err.Error())
	}
	a, err := action.New(input.APIKey, http.DefaultClient)
	if err != nil {
		action.LogErrorAndExit("failed to created action: %s", err.Error())
	}

	if err := a.Cleanup(input.Projects, input.DryRun); err != nil {
		action.LogErrorAndExit("failed to cleanup projects: %s", err.Error())
	}
}
