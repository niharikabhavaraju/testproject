package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "currenttime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey! the current time is", prettyTime)
			return nil
		},
	}
}

func main() {
	cmd := &cobra.Command{
		Use:          "hello",
		Short:        "Hello",
		SilenceUsage: true,
	}
	cmd.AddCommand(printTimeCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
