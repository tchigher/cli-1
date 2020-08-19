package v7action

import "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"

// Revisions represents all revisions for application
type Revisions []ccv3.Revision

// GetRevisionsByApplicationNameAndSpace returns revisions for application.
func (actor *Actor) GetRevisionsByApplicationNameAndSpace(appName string, spaceGUID string) (Revisions, Warnings, error) {
	app, warnings, appErr := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if appErr != nil {
		return Revisions{}, warnings, appErr
	}

	revisions, v3Warnings, apiErr := actor.CloudControllerClient.GetApplicationRevisions(app.GUID)
	warnings = append(warnings, v3Warnings...)

	return revisions, warnings, apiErr
}

func (actor Actor) GetRevisionByApplicationAndVersion(appGUID string, revisionVersion int) (ccv3.Revision, Warnings, error) {
	revisions, warnings, _ := actor.CloudControllerClient.GetApplicationRevisions(appGUID)

	for _, revision := range revisions {
		if revision.Version == revisionVersion {
			return revision, Warnings(warnings), nil
		}
	}
	return ccv3.Revision{}, Warnings(warnings), nil
}
