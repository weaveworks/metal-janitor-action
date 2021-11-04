package action

import (
	"fmt"

	"github.com/packethost/packngo"
)

// cleanProject will clean-up a specific project by name.
func (a *action) cleanProject(proj *packngo.Project, dryRun bool) error {
	Log("cleaning project %s", proj.Name)

	if err := a.cleanupDevices(proj, dryRun); err != nil {
		return fmt.Errorf("cleaning up devices: %w", err)
	}

	Log("deleting project")
	if !dryRun {
		if _, err := a.client.Projects.Delete(proj.ID); err != nil {
			return fmt.Errorf("deleting project %s: %w", proj.Name, err)
		}
	}

	return nil
}

// cleanupDevices will cleanup devices in a project.
func (a *action) cleanupDevices(proj *packngo.Project, dryRun bool) error {
	LogDebug("listing devices in project %s", proj.Name)
	devices, _, err := a.client.Devices.List(proj.ID, nil)
	if err != nil {
		return fmt.Errorf("listing devices for project: %w", err)
	}
	for i := range devices {
		device := devices[i]
		Log("deleting device %s", device.Hostname)
		if !dryRun {
			_, err = a.client.Devices.Delete(device.ID, true)
			if err != nil {
				return fmt.Errorf("deleting device %s: %w", device.Hostname, err)
			}
		}
	}

	return nil
}
