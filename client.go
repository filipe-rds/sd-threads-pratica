package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/hprose/hprose-go/hprose"
)

type RemoteData struct {
	Update func(string, int)
	Get    func(string) (int, error)
	Remove func(string) error
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") 
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// Conectando ao servidor Hprose
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var rm *RemoteData
	client.UseService(&rm)

	for {
		// Limpar a tela antes de exibir o menu
		clearScreen()

		// Exibir o menu de opções
		fmt.Println("Escolha uma operação:")
		fmt.Println("1. Adicionar/atualizar par no dicionário (Update)")
		fmt.Println("2. Obter valor de uma chave do dicionário (Get)")
		fmt.Println("3. Remover par no dicionário através da chave (Remove)")
		fmt.Println("4. Sair\n")
		fmt.Print("Digite sua opção: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Erro ao ler a entrada. Tente novamente.")
			continue
		}

		switch choice {
		case 1:
			// Operação de Update
			var key string
			var value int
			fmt.Print("\nDigite a chave: ")
			_, err := fmt.Scanf("%s", &key)
			if err != nil {
				fmt.Println("Erro ao ler a chave.")
				continue
			}
			fmt.Print("Digite o valor: ")
			_, err = fmt.Scanf("%d", &value)
			if err != nil {
				fmt.Println("Erro ao ler o valor.")
				continue
			}
			rm.Update(key, value)
			fmt.Println("Par inserido com sucesso!")

		case 2:
			// Operação de Get
			var key string
			fmt.Print("\nDigite a chave: ")
			_, err := fmt.Scanf("%s", &key)
			if err != nil {
				fmt.Println("Erro ao ler a chave.")
				continue
			}
			value, err := rm.Get(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Valor da chave '%s': %d\n", key, value)
			}

		case 3:
			// Operação de Remove
			var key string
			fmt.Print("\nDigite a chave para remover: ")
			_, err := fmt.Scanf("%s", &key)
			if err != nil {
				fmt.Println("Erro ao ler a chave.")
				continue
			}
			err = rm.Remove(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Chave removida com sucesso!")
			}

		case 4:
			// Sair
			fmt.Println("\nSaindo do programa...")
			return

		default:
			// Operação inválida
			fmt.Println("\nOpção inválida, tente novamente.")
		}

		fmt.Println("\nPressione qualquer tecla para continuar...")
		fmt.Scanln()
	}
}
