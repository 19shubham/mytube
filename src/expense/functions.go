package expense

import (
	"fmt"
	"utils"
)

func CheckOffer1Applicability ( user1 *User, user2 *User) bool{
	if user1.Balance() == user2.Balance(){
		return true
	}
	return false
}

func TransferMoney(eWallet *Wallet,creditUser string, debitUser string,amt float64){
	if eWallet!=nil{
         cUser := eWallet.UserMap[creditUser]
         cUser.CreditBalance(amt)
         dUser := eWallet.UserMap[debitUser]
         dUser.DebitBalance(amt)

         AddExpenseinList(eWallet,cUser,dUser,amt)

         if CheckOffer1Applicability(cUser,dUser){
         	cUser.CreditBalance(10)
         	dUser.CreditBalance(10)
         	AddOffer1ExpenseInList(eWallet,cUser,dUser)
		 }
	}
}

func AddExpenseinList(eWallet *Wallet,cUser *User , dUser *User, amt float64 ){
	if eWallet!=nil {
		creditString := dUser.Name() + " credit " + fmt.Sprintf("%f",amt)
		if expList,ok:=eWallet.UserExpenseList[cUser.Name()];ok{
			expList = append(expList,creditString)
			eWallet.UserExpenseList[cUser.Name()]= expList
		}else{
			expList := make([]string,0)
			expList = append(expList,creditString)
			eWallet.UserExpenseList[cUser.Name()]= expList
		}

		debitString := cUser.Name() + " debit " + fmt.Sprintf("%0.1f",amt)
		if expList,ok:=eWallet.UserExpenseList[dUser.Name()];ok{
			expList = append(expList,debitString)
			eWallet.UserExpenseList[dUser.Name()]= expList
		}else{
			expList := make([]string,0)
			expList = append(expList,debitString)
			eWallet.UserExpenseList[dUser.Name()]= expList
		}
	}
}

func AddOffer1ExpenseInList(eWallet *Wallet,cUser *User , dUser *User ){
	if eWallet!=nil {
		offerString := "Offer1 credit 10"
		if expList,ok:=eWallet.UserExpenseList[cUser.Name()];ok{
			expList = append(expList, offerString)
			eWallet.UserExpenseList[cUser.Name()]= expList
		}else{
			expList := make([]string,0)
			expList = append(expList, offerString)
			eWallet.UserExpenseList[cUser.Name()]= expList
		}

		if expList,ok:=eWallet.UserExpenseList[dUser.Name()];ok{
			expList = append(expList,offerString)
			eWallet.UserExpenseList[dUser.Name()]= expList
		}else{
			expList := make([]string,0)
			expList = append(expList,offerString)
			eWallet.UserExpenseList[dUser.Name()]= expList
		}
	}
}

func AddOffer2Expense(eWallet *Wallet){
	if eWallet!=nil{
       userTransactionMap := make(map[string]int)
       for name, userExpList:= range eWallet.UserExpenseList{
       	    userTransactionMap[name]= len(userExpList)
	   }
		pairList := utils.SortMapByIntValue(userTransactionMap)
		userList := pairList.DescendingOrder()

		offerString1 := "Offer2 credit 10"
		eWallet.UserMap[userList[0]].CreditBalance(10)
		 if userExpList,ok := eWallet.UserExpenseList[userList[0]];ok{
		 	userExpList = append(userExpList,offerString1)
		 	eWallet.UserExpenseList[userList[0]]= userExpList
		 }

		eWallet.UserMap[userList[1]].CreditBalance(5)
		offerString2 := "Offer2 credit 5"
		if userExpList,ok := eWallet.UserExpenseList[userList[1]];ok{
			userExpList = append(userExpList,offerString2)
			eWallet.UserExpenseList[userList[1]]= userExpList
		}

		eWallet.UserMap[userList[2]].CreditBalance(2)
		offerString3 := "Offer2 credit 2"
		if userExpList,ok := eWallet.UserExpenseList[userList[2]];ok{
			userExpList = append(userExpList,offerString3)
			eWallet.UserExpenseList[userList[2]]= userExpList
		}
	}
}