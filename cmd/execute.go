package cmd

import "log"

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}