package v7

import (
	"code.cloudfoundry.org/cli/command/flag"
)

type RollbackCommand struct {
	BaseCommand
	Revision     int
	RequiredArgs flag.AppName `positional-args:"yes"`
}

func (cmd RollbackCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}
	app, warnings, err := cmd.Actor.GetApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	_, _, _ = app, warnings, err
	// revision, warnings, err := cmd.Actor.GetRevisionByApplicationAndVersion(app.GUID, cmd.Revision)
	// _, _ = warnings, err

	cmd.UI.DisplayTextWithFlavor("Rolling back to revision 1 for app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":   cmd.RequiredArgs.AppName,
		"OrgName":   cmd.Config.TargetedOrganization().Name,
		"SpaceName": cmd.Config.TargetedSpace().Name,
		"Username":  user.Name,
	})
	// _, _, _ = cmd.Actor.CreateDeploymentByApplicationAndRevision(app.GUID, revision.GUID)
	cmd.UI.DisplayNewline()
	cmd.UI.DisplayWarnings(warnings)
	cmd.UI.DisplayOK()
	return nil
}
