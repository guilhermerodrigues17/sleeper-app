package service

import (
	"fmt"
	"os/exec"
)

func Shutdown() error {
	fmt.Println("Executando shutdown...")
	cmd := exec.Command("shutdown.exe", "/s", "/t", "0")
	return cmd.Run()
}
