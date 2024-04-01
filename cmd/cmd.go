package cmd

import (
	"github.com/hsmtkk/aukabucomgo/cmd/selloptiondelta"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use: "aukabocom",
}

var test bool

func init() {
	RootCmd.AddCommand(selloptiondelta.SellOptionDeltaCmd)
	RootCmd.PersistentFlags().BoolVar(&test, "test", false, "test")
	viper.BindPFlag("test", RootCmd.PersistentFlags().Lookup("test"))
}
