package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/guilhermerodrigues17/sleeper-app/service"
)

func main() {
	fmt.Println("Digite um comando (ex: shutdown 2m, suspend 30s, sair)")
	fmt.Println("Uso: <ação> <em quanto tempo>")
	fmt.Println("Ações disponíveis: shutdown | suspend")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		fullInput := strings.TrimSpace(scanner.Text())

		if fullInput == "sair" || fullInput == "exit" {
			fmt.Println("Encerrando...")
			break
		}

		inputArgs := strings.Fields(fullInput)

		if len(inputArgs) < 2 {
			fmt.Println("Erro ao interpretar o formato do comando...")
			continue
		}

		action := strings.ToLower(inputArgs[0])
		durationStr := inputArgs[1]

		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			fmt.Println("Ocorreu um erro...")
			continue
		}

		if duration < 0 {
			fmt.Println("Duração inválida. Use formatos como 30s, 2m, 1h")
			continue
		}

		if duration > 0 {
			fmt.Printf("A ação '%s' será executada. Duração até executar ação: %s \n", action, duration)
			time.Sleep(duration)
		}

		if action == "shutdown" {
			if err := service.Shutdown(); err != nil {
				fmt.Fprintln(os.Stderr, "Erro: ", err)
				return
			}
		}

		if action == "suspend" {
			if err := service.Suspend(); err != nil {
				fmt.Fprintln(os.Stderr, "Erro: ", err)
				return
			}
		}

		return
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro na entrada: ", err)
	}
}
