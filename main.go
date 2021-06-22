package main

import (
	"fmt"

	"github.com/windo-developer/tutorial/mydict"
)

func main() {
	// account := accounts.NewAccount("windo")
	// account.Deposit(10)
	// err := account.Withdraw(50)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account.Balance())

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err2 := dictionary.Update("3adfdsf", "Greeting")
	if err2 != nil {
		fmt.Println((err2))
	}
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
