/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/zeeshanahmad0201/ticker_trawler/cmd"
)

func main() {
	// set logs to show the line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cmd.Execute()
}
