package main

import "fmt"

type CalcInput struct {
	total  int
	action string
}

type MathsInput interface {
	Sum(a, b int)
	Res(a, b int)
	Mul(a, b int)
	Div(a, b int)
	GetTotal() int
	GetAction() string
}

func (n *CalcInput) Sum(a, b int) {
	n.action = "Sum"
	n.total = a + b
}

func (n *CalcInput) Res(a, b int) {
	n.action = "Res"
	n.total = a - b
}

func (n *CalcInput) Mul(a, b int) {
	n.action = "Mul"
	n.total = a * b
}

func (n *CalcInput) Div(a, b int) {
	n.action = "Div"
	n.total = a / b
}

func (n *CalcInput) GetTotal() int {
	return n.total
}

func (n *CalcInput) GetAction() string {
	return n.action
}

func main() {
	var calc MathsInput = &CalcInput{}

	calc.Sum(2, 2)
	fmt.Println(calc.GetAction(), calc.GetTotal())
	calc.Mul(2, 5)
	fmt.Println(calc.GetAction(), calc.GetTotal())
	calc.Div(2, 2)
	fmt.Println(calc.GetAction(), calc.GetTotal())
	calc.Res(5, 2)
	fmt.Println(calc.GetAction(), calc.GetTotal())
}
