package main

import (
	"fmt"
	"github.com/caos/zitadel/cmd/zitadelctl/cmds"
	"os"
)

var (
	version = "none"
)

func main() {
	rootCmd, rootValues := cmds.RootCommand(version)
	rootCmd.Version = fmt.Sprintf("%s\n", version)

	startCmd := cmds.StartOperator(rootValues)
	takeoffCmd := cmds.TakeoffCommand(rootValues)
	backuplistCmd := cmds.BackupListCommand(rootValues)
	restoreCmd := cmds.RestoreCommand(rootValues)
	readsecretCmd := cmds.ReadSecretCommand(rootValues)
	writesecretCmd := cmds.WriteSecretCommand(rootValues)

	rootCmd.AddCommand(
		startCmd,
		takeoffCmd,
		backuplistCmd,
		restoreCmd,
		readsecretCmd,
		writesecretCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
