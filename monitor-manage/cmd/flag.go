package main

import "flag"


var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./config.yml", "config path")
}