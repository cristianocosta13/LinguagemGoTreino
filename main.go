package main

import "fmt"

func main() {

	/*
		SE DECLARAR UMA VARIAVEL E NAO USAR, DA ERRO
		3 FORMAS DE DECLARAR VARIAVEIS
		O ARQUIVO GO.MOD É ONDE SÃO COLOCADAS AS DEPENDENCIAS (DESNECESSARIO POR ENQUANTO)
		O ELSE DEVE ESTAR NA MESMA LINHA DAS CHAVES, NA ESTRUTURA CONDICIONAIS
	*/

	//IMPRIMINDO TEXTO
	fmt.Println("Hello, World!")

	//CHAMADA DE MÉTODO
	AnalisarIdade()
}

func AnalisarIdade() {
	//DECLARANDO E IMPRIMINDO VARIAVEIS
	var idade1 int  //apenas declara, pode ser colocado o valor depois
	var idade2 = 20 //ja define o tipo ao ser colocado o valor, mas precisa utilizar o var
	idade3 := 17    //sem necessidade de colocar o var, ja atribui direto e o valor colocado define o tipo permanentemente
	idade1 = 18

	//PARA CONCATENAR PRECISA DE VIRGULA
	fmt.Println("IDADE 1: ", idade1, "\nIDADE 2:", idade2, "\nIDADE 3:", idade3)

	//ESTRUTURAS CONDICIONAIS COM IF E ELSE E COM OPERADORES LÓGICOS
	if idade1 == 18 {
		fmt.Println("O usuário de idade 1 TEM 18 ANOS!")
	} else if idade1 > 18 {
		fmt.Println("O usuário de idade 2 é maior de idade!")
	} else {
		fmt.Println("O usuário não é maior de idade!")
	}
	if idade1 != 17 {
		fmt.Println("O usuário de idade 1 não tem 17 anos!")
	}
	if idade1 >= 18 && idade2 >= 18 && idade3 >= 18 {
		fmt.Println("Todos os usuários são maiores de idade!")
	} else {
		fmt.Println("Nem todos os usuários são maiores de idade!")
	}
	if idade1 >= 18 || idade2 >= 18 || idade3 >= 18 {
		fmt.Println("Dentre os usuários no mínimo um é maior de idade!")
	} else {
		fmt.Println("Nenhum dos usuários são maior de idade!")
	}

	var i int
	for i = 0; i < 3; i++ { //EM ESTRUTURAS DE REPETIÇÃO E CONDIÇÃO NÃO PRECISA PARENTESE
		fmt.Println("Numero ", i+1, ": ", i)
	}
}
