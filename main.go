package main

import (
	"fmt"
	"os"

	"github.com/9to6/aws-autoscaling-tester/cmd"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	cobra *cobra.Command
}

var rootCmd = &RootCmd{
	cobra: &cobra.Command{
		Use:     "AWS metrics generator",
		Short:   "To generate metrics on the CloudWatch, Server, Client",
		Version: fmt.Sprintf("%v git=%v build=%v", "0.1.0", 0, 0),
	},
}

func init() {
	rootCmd.cobra.AddCommand(cmd.NewCmdServer())
	rootCmd.cobra.AddCommand(cmd.NewCmdClient())
}

func main() {
	if err := rootCmd.cobra.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
