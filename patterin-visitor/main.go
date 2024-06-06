package main

import "fmt"

/*
 Плюсы:
	- Добавление новых операций без изменения классов.
	- Разделение алгоритмов и структур данных.
 Минусы:
	- Нарушение инкапсуляции.
	- Усложнение кода.
*/

// Visitor - интерфейс для посетителей.
type Visitor interface {
	VisitLiteral(literal *Literal)
	VisitAddition(addition *Addition)
}

// ExpressionPrintingVisitor - конкретный посетитель, который печатает выражения.
type ExpressionPrintingVisitor struct{}

func (epv *ExpressionPrintingVisitor) VisitLiteral(literal *Literal) {
	fmt.Println(literal.Value)
}

func (epv *ExpressionPrintingVisitor) VisitAddition(addition *Addition) {
	leftValue := addition.Left.GetValue()
	rightValue := addition.Right.GetValue()
	sum := addition.GetValue()
	fmt.Printf("%v + %v = %v\n", leftValue, rightValue, sum)
}

// Expression - абстрактный класс для выражений.
type Expression interface {
	Accept(visitor Visitor)
	GetValue() float64
}

// Literal - класс для числовых значений.
type Literal struct {
	Value float64
}

func (l *Literal) Accept(visitor Visitor) {
	visitor.VisitLiteral(l)
}

func (l *Literal) GetValue() float64 {
	return l.Value
}

// Addition - класс для выражений сложения.
type Addition struct {
	Left, Right Expression
}

func (a *Addition) Accept(visitor Visitor) {
	a.Left.Accept(visitor)
	a.Right.Accept(visitor)
	visitor.VisitAddition(a)
}

func (a *Addition) GetValue() float64 {
	return a.Left.GetValue() + a.Right.GetValue()
}

func main() {
	// Эмуляция выражения 1 + 2 + 3
	e := &Addition{
		Left: &Addition{
			Left:  &Literal{Value: 1},
			Right: &Literal{Value: 2},
		},
		Right: &Literal{Value: 3},
	}

	printingVisitor := &ExpressionPrintingVisitor{}
	e.Accept(printingVisitor)
}
