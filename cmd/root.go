package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/shohi/go-raft-example/pkg/server"
)

var conf server.Config

var rootCmd = &cobra.Command{
	Use:   "goraft",
	Short: "distributed key/value store with HA backed by raft protocol",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	panic("TODO")
}

// Execute is entrance of the service
func Execute() {
	setupFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()
	_ = flagSet

	flagSet.IntVar(&conf.HTTPPort, "http-port", 9090, "HTTP listen port")
	flagSet.IntVar(&conf.RaftPort, "raft-port", 9091, "Raft bind port")
	flagSet.StringVar(&conf.JoinAddr, "join-addr", "", "address for join raft cluster, if any")
	flagSet.StringVar(&conf.NodeID, "node-id", "", "Node ID")

	panic("TODO")
}
