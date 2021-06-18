package expense

type Wallet struct {
	UserMap         map[string]*User
	UserExpenseList map[string][]string
}

func NewWalletSystem(userMap map[string]*User, expenseList map[string][]string) *Wallet {
	return &Wallet{UserMap: userMap, UserExpenseList: expenseList}
}

func (wallet *Wallet)AddNewUserInMap(name string, user *User){
	if wallet.UserMap !=nil{
		wallet.UserMap[name] = user
	}else{
		wallet.UserMap = make(map[string]*User)
		wallet.UserMap[name] = user
	}
}
