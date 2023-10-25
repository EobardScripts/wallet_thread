package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"threads/pkg/bankClient"
	"time"
)

// Случайное вещественное число
func RandFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Случайное целочисленное число
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	wallet := bankClient.NewWallet()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				wallet.Deposit(RandInt(1, 10))
				time.Sleep(time.Duration(RandFloat(0.5, 1.0)) * time.Second)
			}
		}()

		if i < 5 {
			go func() {
				for {
					err := wallet.Withdrawal(RandInt(1, 5))
					if err != nil {
						fmt.Println(err)
					}
					time.Sleep(time.Duration(RandFloat(0.5, 1.0)) * time.Second)
				}
			}()
		}
	}

	operation := ""
	fmt.Println("Команды: balance, deposit, withdrawal, exit")
mark:
	_, err := fmt.Scanln(&operation)
	if err != nil {
		log.Fatalln(err)
	}

	switch strings.ToLower(operation) {
	case "balance":
		fmt.Println(wallet.Balance())
	case "deposit":
		{
			fmt.Println("Введите количество денег для зачисления на счет: ")
			var amount int
			_, err = fmt.Scanln(&amount)
			if err != nil {
				fmt.Println(err)
			}
			wallet.Deposit(amount)
		}
	case "withdrawal":
		{
			fmt.Println("Введите количество денег для вывода со счета: ")
			var amount int
			_, err = fmt.Scanln(&amount)
			if err != nil {
				fmt.Println(err)
			}
			err = wallet.Withdrawal(amount)
			if err != nil {
				fmt.Println(err)
			}
		}

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")

	}
	goto mark

}
