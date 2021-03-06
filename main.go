package main

import (
	"github.com/Sho372/grc/commands"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app, err := commands.New()
	if err != nil {
		log.Fatal(err)
	}
	cmdRoot := buildRootCommand(app)
	cmdZadd := buildZaddCommand(app)
	cmdRoot.AddCommand(cmdZadd)
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCommand(app *commands.App) *cobra.Command {
	cmdRoot := &cobra.Command{
		Use:   "grc",
		Short: "grc is golang redis client.",
	}
	return cmdRoot
}

func buildZaddCommand(app *commands.App) *cobra.Command {
	var period int
	var repeat int

	cmdZadd := &cobra.Command{
		Use: "zadd",
		Run: func(cmd *cobra.Command, args []string) {
			key, score, value := args[0], args[1], args[2]
			app.Zadd(key, score, value, period, repeat)
		},
		Args: cobra.ExactArgs(3),
	}

	cmdZadd.Flags().IntVar(&period, "period", 0, "How long does it keep executing?")
	cmdZadd.Flags().IntVar(&repeat, "repeat", 1, "How many times does it repeat?")

	return cmdZadd
}
