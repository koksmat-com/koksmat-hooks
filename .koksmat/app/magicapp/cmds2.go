package magicapp

import (
	"github.com/365admin/koksmat-hooks/cmds"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic Buttons",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(magicCmd)
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(setupCmd)
	tasksCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Tasks",
		Long:  `Describe the main purpose of this kitchen`,
	}
	TasksStep1PostCmd := &cobra.Command{
		Use:   "step1",
		Short: "Step 1",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TasksStep1Post(ctx, args)
		},
	}
	tasksCmd.AddCommand(TasksStep1PostCmd)

	RootCmd.AddCommand(tasksCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Describe the main purpose of this kitchen`,
	}
	ProvisionWebdeployproductionPostCmd := &cobra.Command{
		Use:   "webdeployproduction",
		Short: "Web deploy to production",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeployproductionPostCmd)

	RootCmd.AddCommand(provisionCmd)
	decommissionCmd := &cobra.Command{
		Use:   "decommission",
		Short: "Decommision",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(decommissionCmd)
}
