package action

import (
	"fmt"

	"github.com/packethost/packngo"
	"go.uber.org/zap"
)

// cleanProject will clean-up a specific project by name
func (a *action) cleanProject(proj *packngo.Project, dryRun bool) error {
	logger := a.logger.With("project", proj.Name)
	logger.Info("cleaning project")

	if err := a.cleanupDevices(logger, proj, dryRun); err != nil {
		return fmt.Errorf("cleaning up devices: %w", err)
	}

	if err := a.cleanupVolumes(logger, proj, dryRun); err != nil {
		return fmt.Errorf("cleaning up volumes: %w", err)
	}

	logger.Info("deleting project")
	if !dryRun {
		if _, err := a.client.Projects.Delete(proj.ID); err != nil {
			return fmt.Errorf("deleting project %s: %w", proj.Name, err)
		}
	}

	return nil
}

// cleanupDevices will cleanup devices in a project
func (a *action) cleanupDevices(logger *zap.SugaredLogger, proj *packngo.Project, dryRun bool) error {
	logger.Debug("listing devices in project")
	devices, _, err := a.client.Devices.List(proj.ID, nil)
	if err != nil {
		return fmt.Errorf("listing devices for project: %w", err)
	}
	for i := range devices {
		device := devices[i]
		logger.Infow("deleting device", "name", device.Hostname, "id", device.ID)
		if !dryRun {
			_, err = a.client.Devices.Delete(device.ID, true)
			if err != nil {
				return fmt.Errorf("deleting device %s: %w", device.Hostname, err)
			}
		}
	}

	return nil
}

// cleanupVolumes will cleanup volumes in a project
func (a *action) cleanupVolumes(logger *zap.SugaredLogger, proj *packngo.Project, dryRun bool) error {
	logger.Debug("listing volumes in project")
	volumes, _, err := a.client.Volumes.List(proj.ID, nil)
	if err != nil {
		return fmt.Errorf("listing volumes for project: %w", err)
	}
	for i := range volumes {
		volume := volumes[i]
		logger.Infow("deleting volume", "volume", volume.Name)
		if !dryRun {
			_, err = a.client.Volumes.Delete(volume.ID)
			if err != nil {
				return fmt.Errorf("deleting volume %s: %w", volume.Name, err)
			}
		}
	}

	return nil
}
