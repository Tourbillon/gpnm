// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package main

import (
	"log"

	"anbillon.com/gpnm/cmd"
)

func main() {
	c := cmd.NewRootCmd()
	if err := c.Execute(); err != nil {
		log.Printf("%v", err)
	}
}
