// := operador curto de declaração
// variáveis dentro e fora do bloco/escopo
// função variádica e _

package main

import "fmt"

var y string = "hey, hi"
// não é possível atribuir um valor fora de um code-block


func main() {
	var x int = 10
	z := true
	z = false
	fmt.Println(x, y, z)

	// declaração, inicialização e atribuição
	var w int
	w = 10
	w = 15
	fmt.Printf("%v, %T", w, w)

	//ou
	a := "hello"
	fmt.Println(a)

	pacotefmt()
	propriavariavel()
}

func pacotefmt(){
	//literal "" ou interpretado ``
	x := "oi bom dia"
	fmt.Println(x)

	y := "tudo bem?"

	z := fmt.Sprint(x," ",y)

	fmt.Println(z)

}

type hotdog int
var b hotdog

func propriavariavel() {
	x := 10
	fmt.Printf("%T", x)
	fmt.Printf("%T", b)
}

