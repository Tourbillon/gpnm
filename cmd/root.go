// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd will create a new instance of the root command.
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gpnm",
		Short: "Golang package manager web",
	}

	cmd.AddCommand(newServerStartCmd())

	return cmd
}
