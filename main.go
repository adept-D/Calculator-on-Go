package main

import (
	"errors"
	"fmt"
	"log"
)

type Calculator struct {
	firstNumber  float32
	secondNumber float32
}

func (c *Calculator) sum() float32 {
	return c.firstNumber + c.secondNumber
}

func (c *Calculator) multiply() float32 {
	return c.firstNumber * c.secondNumber
}

func (c *Calculator) divide() (float32, error) {
	if c.secondNumber == 0 {
		return 0, errors.New("divide by zero")
	}
	return c.firstNumber / c.secondNumber, nil

}

func main() {
	calc := &Calculator{5, 20}
	fmt.Println("Sum ", calc.sum())
	fmt.Println("Multiply ", calc.multiply())
	n, err := calc.divide()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Divide ", n)
}
