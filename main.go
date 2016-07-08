package main

import "github.com/mlambrichs/graphite-tools/commands"

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
