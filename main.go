package main

import (
	"awsips/cmd"
	"awsips/internal/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
