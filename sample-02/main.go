package main

import (
	"fmt"

	"./accounts"
)

func main() {
	fmt.Println("main start")

	accounts := accounts.NewAccount("mango")
	fmt.Println("accounts.Balance(): ", accounts.Balance())

	accounts.Deposit(10)
	fmt.Println(*accounts)

	fmt.Println("accounts.Balance(): ", accounts.Balance())

	err := accounts.WithDraw(20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("accounts.Balance(): ", accounts.Balance())
}
