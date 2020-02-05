package service_account

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	orgv1 "github.com/confluentinc/ccloudapis/org/v1"
	"github.com/confluentinc/go-printer"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	v2 "github.com/confluentinc/cli/internal/pkg/config/v2"
	"github.com/confluentinc/cli/internal/pkg/errors"
)

type command struct {
	*pcmd.AuthenticatedCLICommand
}

var (
	listFields      = []string{"Id", "ServiceName", "ServiceDescription"}
	listLabels      = []string{"Id", "Name", "Description"}
	describeFields  = []string{"Id", "ServiceName", "ServiceDescription"}
	describeRenames = map[string]string{"ServiceName": "Name", "ServiceDescription": "Description"}
)

const nameLength = 32
const descriptionLength = 128

// New returns the Cobra command for service accounts.
func New(prerunner pcmd.PreRunner, config *v2.Config) *cobra.Command {
	cliCmd := pcmd.NewAuthenticatedCLICommand(
		&cobra.Command{
			Use:   "service-account",
			Short: `Manage service accounts.`,
		},
		config, prerunner)
	cmd := &command{
		AuthenticatedCLICommand: cliCmd,
	}
	cmd.init()
	return cmd.Command
}

func (c *command) init() {
	c.AddCommand(&cobra.Command{
		Use:   "list",
		Short: `List service accounts.`,
		RunE:  c.list,
		Args:  cobra.NoArgs,
	})

	createCmd := &cobra.Command{
		Use:   "create <name>",
		Short: `Create a service account.`,
		Example: `
Create a service account named ` + "``DemoServiceAccount``" + `.

::

  ccloud service-account create "DemoServiceAccount" \
  --description "This is a demo service account."

`,
		RunE: c.create,
		Args: cobra.ExactArgs(1),
	}
	createCmd.Flags().String("description", "", "Description of the service account.")
	_ = createCmd.MarkFlagRequired("description")
	createCmd.Flags().SortFlags = false
	c.AddCommand(createCmd)

	updateCmd := &cobra.Command{
		Use:   "update <id>",
		Short: `Update a service account.`,
		Example: `
Update the description of a service account with the ID ` + "``2786``" + `.

::

    ccloud service-account update service-account-id 2786 \
    --description "Update demo service account information."

`,
		RunE: c.update,
		Args: cobra.ExactArgs(1),
	}
	updateCmd.Flags().String("description", "", "Description of the service account.")
	_ = updateCmd.MarkFlagRequired("description")
	updateCmd.Flags().SortFlags = false
	c.AddCommand(updateCmd)

	c.AddCommand(&cobra.Command{
		Use:   "delete <id>",
		Short: `Delete a service account.`,
		Example: `
Delete a service account with the ID ` + "``2786``" + `.

::

    ccloud service-account delete 2786

`,
		RunE: c.delete,
		Args: cobra.ExactArgs(1),
	})
}

func requireLen(val string, maxLen int, field string) error {
	if len(val) > maxLen {
		return fmt.Errorf(field+" length should be less then %d characters.", maxLen)
	}

	return nil
}

func (c *command) create(cmd *cobra.Command, args []string) error {
	name := args[0]

	if err := requireLen(name, nameLength, "service name"); err != nil {
		return errors.HandleCommon(err, cmd)
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}

	if err := requireLen(description, descriptionLength, "description"); err != nil {
		return errors.HandleCommon(err, cmd)
	}

	user := &orgv1.User{
		ServiceName:        name,
		ServiceDescription: description,
		OrganizationId:     c.State.Auth.User.OrganizationId,
		ServiceAccount:     true,
	}
	user, err = c.Client.User.CreateServiceAccount(context.Background(), user)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}

	return printer.RenderTableOut(user, describeFields, describeRenames, os.Stdout)
}

func (c *command) update(cmd *cobra.Command, args []string) error {
	idp, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	id := int32(idp)

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}

	if err := requireLen(description, descriptionLength, "description"); err != nil {
		return errors.HandleCommon(err, cmd)
	}

	user := &orgv1.User{
		Id:                 id,
		ServiceDescription: description,
	}
	err = c.Client.User.UpdateServiceAccount(context.Background(), user)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	return nil
}

func (c *command) delete(cmd *cobra.Command, args []string) error {
	idp, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	id := int32(idp)

	user := &orgv1.User{
		Id: id,
	}
	err = c.Client.User.DeleteServiceAccount(context.Background(), user)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	return nil
}

func (c *command) list(cmd *cobra.Command, args []string) error {
	users, err := c.Client.User.GetServiceAccounts(context.Background())
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}

	var data [][]string
	for _, u := range users {
		data = append(data, printer.ToRow(u, listFields))
	}

	printer.RenderCollectionTable(data, listLabels)
	return nil
}
