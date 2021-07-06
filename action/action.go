package action

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/packethost/packngo"
	"go.uber.org/zap"
)

const (
	deleteAllprojects = "DELETEALL"
)

// MetalJanitorAction represents the interface for the cleanup action
type MetalJanitorAction interface {
	// Cleanup is used to cleanup the Equinix Metal projects.
	Cleanup(projectNames string, dryRun bool) error
}

// New will create a new instance of the metal janitor action
func New(authToken string, logger *zap.Logger, httpClient *http.Client) (MetalJanitorAction, error) {
	if authToken == "" {
		return nil, ErrAPIKeyRequired
	}

	return &action{
		client: packngo.NewClientWithAuth("metal-janitor-action", authToken, httpClient),
		logger: logger.Sugar(),
	}, nil
}

// NewWithURL will create a new instance of the metal janitor action using a specific URL for the Equinix api
func NewWithURL(authToken string, logger *zap.Logger, httpClient *http.Client, apiURL string) (MetalJanitorAction, error) {
	if authToken == "" {
		return nil, ErrAPIKeyRequired
	}
	if apiURL == "" {
		return nil, ErrAPIUrlRequired
	}

	client, err := packngo.NewClientWithBaseURL("metal-janitor-action", authToken, httpClient, apiURL)
	if err != nil {
		return nil, fmt.Errorf("creating equinix metal client: %w", err)
	}

	return &action{
		client: client,
		logger: logger.Sugar(),
	}, nil
}

// action represents the main action implementation
type action struct {
	client *packngo.Client
	logger *zap.SugaredLogger
}

// Cleanup will cleanup the supplied projects
func (a *action) Cleanup(projectNames string, dryRun bool) error {
	a.logger.Infow("cleaning up projects", "names", projectNames, "dryrun", dryRun)

	a.logger.Debug("listing projects")
	projects, _, err := a.client.Projects.List(&packngo.ListOptions{})
	if err != nil {
		return fmt.Errorf("getting projects list: %w", err)
	}

	projectsToDelete := strings.Split(projectNames, ",")

	for i := range projects {
		proj := projects[i]

		if !shouldDelete(projectsToDelete, proj.Name) {
			continue
		}

		if err := a.cleanProject(&proj, dryRun); err != nil {
			return fmt.Errorf("cleaning up project %s: %w", proj.Name, err)
		}
	}

	return nil
}

// shouldDelete determines if a project should be deleted
func shouldDelete(projectsToDelete []string, project string) bool {
	if len(projectsToDelete) == 0 {
		return false
	}

	if projectsToDelete[0] == deleteAllprojects {
		return true
	}

	for _, projToDelete := range projectsToDelete {
		if projToDelete == project {
			return true
		}
	}

	return false
}
