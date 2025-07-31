package actions

import (
	"fmt"
	"os/exec"
)

func Suspend() error {
	fmt.Println("Executando suspend...")
	cmd := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0")
	return cmd.Run()
}
