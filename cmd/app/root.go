package app

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "transmitter",
	Short: "federatorai agent",
	Long:  "",
}

var (
	configurationFilePath string
)

func init() {
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(ProbeCmd)

	RootCmd.PersistentFlags().StringVar(
		&configurationFilePath,
		"config",
		"/etc/alameda/federatorai-agent/transmitter.toml",
		"The path to federatorai-agent configuration file.")
}

