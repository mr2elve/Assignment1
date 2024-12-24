package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited: %.2f. New Balance: %.2f\n", amount, b.Balance)
	} else {
		fmt.Println("Invalid deposit amount. Must be greater than 0.")
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 {
		if b.Balance >= amount {
			b.Balance -= amount
			fmt.Printf("Withdrew: %.2f. New Balance: %.2f\n", amount, b.Balance)
		} else {
			fmt.Println("Insufficient balance.")
		}
	} else {
		fmt.Println("Invalid withdraw amount. Must be greater than 0.")
	}
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, t := range transactions {
		if t > 0 {
			account.Deposit(t)
		} else if t < 0 {
			account.Withdraw(-t)
		}
	}
}

func main() {
	account := &BankAccount{
		AccountNumber: "1234567890",
		HolderName:    "Alice",
		Balance:       1000.00,
	}

	for {
		fmt.Println("\nBank Account System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
