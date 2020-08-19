package v7

import (
	"code.cloudfoundry.org/cli/command/flag"
)

type RollbackCommand struct {
	BaseCommand
	RequiredArgs    flag.AppName         `positional-args:"yes"`
	Version         flag.PositiveInteger `long:"revision" required:"true" description:"Revision to which to roll back."`
	usage           interface{}          `usage:"CF_NAME rollback APP_NAME [--revision]"`
	relatedCommands interface{}          `related_commands:"revisions"`
}

func (cmd RollbackCommand) Execute(args []string) error {
	targetRevision := int(cmd.Version.Value)
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	app, warnings, _ := cmd.Actor.GetApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)

	revision, warnings, _ := cmd.Actor.GetRevisionByApplicationAndVersion(app.GUID, targetRevision)
	cmd.UI.DisplayWarnings(warnings)

	cmd.UI.DisplayTextWithFlavor("Rolling back to revision {{.TargetRevision}} for app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":        cmd.RequiredArgs.AppName,
		"TargetRevision": targetRevision,
		"OrgName":        cmd.Config.TargetedOrganization().Name,
		"SpaceName":      cmd.Config.TargetedSpace().Name,
		"Username":       user.Name,
	})
	cmd.UI.DisplayNewline()

	_, warnings, _ = cmd.Actor.CreateDeploymentByApplicationAndRevision(app.GUID, revision.GUID)
	cmd.UI.DisplayWarnings(warnings)

	cmd.UI.DisplayOK()
	return nil
}
