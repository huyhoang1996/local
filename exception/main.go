package main

import (
	"fmt"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func main() {
	fmt.Println("We started")
	Block{
		Try: func() {
			fmt.Println("I tried")
			// Create Error
			panic("Create panic")
		},
		Catch: func(e Exception) {
			// Catch error
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
	fmt.Println("We went on")
}
