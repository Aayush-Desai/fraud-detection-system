package main

type FraudDetectionSystem struct {
	rules []FraudDetectionRule
}

func (s *FraudDetectionSystem) ProcessTransaction(transaction *Transaction, user *User) bool {
	for _, rule := range s.rules {
		if rule.ValidateTransaction(transaction, user) {
			return true
		}
	}
	transaction.ExecuteTransaction()
	return false
}

func (s *FraudDetectionSystem) AddNewFalutDetectionRule(rule FraudDetectionRule) {
	s.rules = append(s.rules, rule)
}