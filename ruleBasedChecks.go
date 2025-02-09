package main

type RuleBasedCheck struct {
	id string
}

func GetRuleBasedCheck(id string) *RuleBasedCheck {
	return &RuleBasedCheck{
		id: id,
	}
}

func (* RuleBasedCheck) ValidateTransaction(transaction *Transaction, user *User) bool {
	if(transaction.account.IsAccountBlacklisted()) {
		return false
	}else if(transaction.GetAmmount() > user.GetTransactionLimit(transaction.GetAccount())) {
		return false
	}
	return true
}