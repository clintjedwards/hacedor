package cmd

import (
	"log"
	"strings"

	"github.com/clintjedwards/cursor/api"
	"github.com/clintjedwards/cursor/client"
	"github.com/clintjedwards/cursor/config"
	"github.com/spf13/cobra"
)

var cmdPipelinesList = &cobra.Command{
	Use:   "list",
	Short: "list all available pipelines",
	Long:  ``,
	Run:   runPipelinesListCmd,
}

func runPipelinesListCmd(cmd *cobra.Command, args []string) {

	config, err := config.FromEnv()
	if err != nil {
		log.Fatalf("failed to read configuration")
	}

	hostPortTuple := strings.Split(config.Master.URL, ":")

	cursorClient := client.CursorClient{}
	err = cursorClient.Connect(hostPortTuple[0], hostPortTuple[1])
	if err != nil {
		log.Fatalf("could not connect to host: %v", err)
	}
	defer cursorClient.Close()

	_, err = cursorClient.ListPipelines(&api.ListPipelinesRequest{})
	if err != nil {
		log.Fatalf("could not list pipelines: %v", err)
	}
}

func init() {
	var description string
	cmdPipelinesCreate.Flags().StringVarP(&description, "description", "d", "", "long form description of pipeline")

	var gitBranch string
	cmdPipelinesCreate.Flags().StringVarP(&gitBranch, "git_branch", "b", "master", "git branch name")

	cmdPipeline.AddCommand(cmdPipelinesCreate)
}