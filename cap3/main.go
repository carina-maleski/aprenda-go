package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

// - Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
//     1. 42
//     2. "James Bond"
//     3. true
// - Agora demonstre os valores nestas variáveis, com:
//     1. Uma única declaração print
//     2. Múltiplas declarações print

func ex1() {
	fmt.Println("Exercício 1")

	x1 := 42
	y1 := "James Bond"
	z1 := true

	fmt.Println(x1, y1, z1)
	fmt.Printf("%v %T\n", x1, x1)
	fmt.Printf("%v %T\n", y1, y1)
	fmt.Printf("%v %T\n", z1, z1)
}

// - Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
//     1. Identificador "x" deverá ter tipo int
//     2. Identificador "y" deverá ter tipo string
//     3. Identificador "z" deverá ter tipo bool
// - Na função main:
//     1. Demonstre os valores de cada identificador
//     2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

var x2 int
var y2 string
var z2 bool

func ex2() {
	fmt.Println("\nExercício 2")

	fmt.Println(x2, y2, z2)
	fmt.Printf("%v\n%v\n%v\n", x2, y2, z2)
	fmt.Print("Esses valores se chamam zero-values. \n")
}


// - Utilizando a solução do exercício anterior:
//     1. Em package-level scope, atribua os seguintes valores às variáveis:
//         1. para "x" atribua 42
//         2. para "y" atribua "James Bond"
//         3. para "z" atribua true
//     2. Na função main:
//         1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
//         2. Demonstre a variável "s".

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

func ex3() {
	fmt.Println("\nExercício 3")

	s := fmt.Sprintf("As variáveis são %v, %v e %v.\n", x3, y3, z3)
	fmt.Println(s)
}

// - Crie um tipo. O tipo subjacente deve ser int.
// - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// - Na função main:
//     1. Demonstre o valor da variável "x"
//     2. Demonstre o tipo da variável "x"
//     3. Atribua 42 à variável "x" utilizando o operador "="
//     4. Demonstre o valor da variável "x"

type Oi int
var x4 Oi 

func ex4() {
	fmt.Println("Exercício 4")
	fmt.Println(x4)
	x4 = 4837
	fmt.Println(x4)
	fmt.Printf("O valor de x4 é %v e o seu tipo é %T", x4, x4)
}

// - Utilizando a solução do exercício anterior:
//     1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//     2. Na função main:
//         1. Isto já deve estar feito:
//             1. Demonstre o valor da variável "x"
//             2. Demonstre o tipo da variável "x"
//             3. Atribua 42 à variável "x" utilizando o operador "="
//             4. Demonstre o valor da variável "x"
//         2. Agora faça tambem:
//             1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//             2. Demonstre o valor de "y"
//             3. Demonstre o tipo de "y"

var y5 int
func ex5() {
	fmt.Println("\n\nExercício 5")
	fmt.Printf("O valor de x4 é %v e o seu tipo é %T\n", x4, x4)

	x4 = 42

	fmt.Println(x4)

	y5 = int(x4)

	fmt.Printf("O valor de y5 é %v e o seu tipo é %T\n",y5,y5)
}