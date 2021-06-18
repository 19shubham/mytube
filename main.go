package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"expense"
)

func main() {

	eWallet := expense.NewWalletSystem(make(map[string]*expense.User), make(map[string][]string))

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	for _, eachLine := range fileTextLines {
		commands := strings.Split(eachLine," ")

		switch commands[0] {

		case "CreateWallet":
			name := commands[1]
			balance,_ := strconv.ParseFloat(commands[2],64)
			user,err := expense.CreateWallet(name,balance)
			if err!=nil{
				fmt.Println(err.Error())
				continue
			}
			eWallet.AddNewUserInMap(name,user)

		case "Overview":
			userMap := eWallet.UserMap
			for _, user := range userMap{
				fmt.Println(user.Name(), " ",user.Balance())
			}

		case "TransferMoney":
			creditUser := commands[2]
			debitUser := commands[1]
			amt, _ := strconv.ParseFloat(commands[3],64)
			expense.TransferMoney(eWallet,creditUser, debitUser,amt)
		case "Statement":
			name := commands[1]
			if userExpList, ok := eWallet.UserExpenseList[name];ok{
				for _,exp := range userExpList{
					fmt.Println(exp)
				}
			}else{
				fmt.Println("No expenses occurred for user")
			}
		case "Offer2":
			expense.AddOffer2Expense(eWallet)
		}
	}
}