// Para inicializar digit go mod init <module path>
// <module path> Corresponde ao rep/projeto
package main

// Na linguagem GO é necessario que o código faça parte de um pacote
// É necessário um ponto de entrada para que um programa GO se localize
// 1 Programaa só pode ter 1 função "main" por que você só pode ter 1 ponto de entrada

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// Package level variables
var conferenceName = "GO conference"

// var conferenceName = "GO conference"
// Constante não altera durante a execução do programa, sempre manterá o mesmo valor
const conferenceTickets int = 50

var remainingTickets uint = 50

// Exemplo de lista: var bookings = [50]string{}
// var bookings [50]string //outra maneira de declarar lista
// var bookings []string // Isto é um Slice, representa a mesma forma da lista porém de maneira dinâmica
var bookings = []string{} // Criando com a sintaxe compacta

func main() {
	greetUsers()

	// Podemos dar condições de execução ao nosso for
	// Exemplo for remainingTickets > 00 && len(bookings) < 50 {}
	for {
		// Para loops existe apenas o for em GO
		firstName, lastName, email, userTickets := getUserInput()

		fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// Assimilando uma variavel para receber o valor de saída da função getFirtsNames
			firstNames := getFirstNames()

			fmt.Printf("\nO primeiro nome dos compradores: %v\n", firstNames)

			noremainingTickets := remainingTickets == 0
			// if remainingTickets == 0 {
			if noremainingTickets { // end program
				fmt.Println("Os tickets acabaram para esta conferencia. Volte novamente no próximo ano!")
				break
			}

		} else { //else if userTickets == remainingTickets {}
			if !isValidName {
				fmt.Println("O nome informado é muito curto")
			}
			if !isValidEmail {
				fmt.Println("O e-mail não contém o simbolo @")
			}
			if !isValidTicketNumber {
				fmt.Println("Número de tickets que você inseriu é inválido")
			}
			// city := "London"

			// switch city {
			// case "New York":
			// 	// Execute Code for this city
			// case "Singapore":
			// 	// Execute Code for this city
			// case "London", "Mexico":
			// 	// Execute Code for this city
			// case "Berlin":
			// 	// Execute Code for this city
			// case "Hong Kong":
			// 	// Execute Code for this city
			// default: // Caso não seja encontrado nenhum valor, executar o default
			// 	fmt.Println("Cidade inválida")
			// }

			fmt.Printf("O valor informado não está correto, por favor insira novamente.")
			continue
		}

	}
}

// Para executar 1 arquivo, executar no terminal: go run <file name> -> Irá compilar e rodar o código

func greetUsers() {
	fmt.Printf("Bem vindo a %v!\n", conferenceName)
	fmt.Printf("Nós temos o total de %v tickets e %v estão disponíveis atualmente\n", conferenceTickets, remainingTickets)
	fmt.Printf("Garanta seus tickets agora mesmo!\n")

}

func getFirstNames() []string { // Colocamos dentro dos colchetes os parametros de entrada e fora colocamos os parametros que serão retornados / de saida
	firstNames := []string{}
	// Definindo for para iterar o slice Bookings, com a função range conseguimos o index e o valor de cada elemento
	// Tanto para listas ou slices
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your e-mail: ")
	fmt.Scan(&email)

	fmt.Printf("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// Em GO para utilizar lista é necessário definir quandos itens poderám ser armazenados nesta lista
	// E é necessario dizer qual o tipo de dado que terá na lista, ou seja não é possivel utilizar INT e STRING juntos.
	// bookings[0] = firstName + " " + lastName  -> Exemplo de inserção em 1 lista
	bookings = append(bookings, firstName+" "+lastName) // Inserindo em nosso slice

	// fmt.Printf("Slice completo: %v\n", bookings)
	// fmt.Printf("Primeiro valor: %v\n", bookings[0])
	// fmt.Printf("Tipo do Slice: %T\n", bookings)
	// fmt.Printf("Tamanho do Slice: %v\n", len(bookings))

	// Função scanf para pedir informações
	fmt.Printf("\nObrigado %v %v por comprar %v tickets! Você irá receber a confirmação da compra no seguinte e-mail: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Número de tickets em estoque da %v após a compra: %v", conferenceName, remainingTickets)
}
