package main

type FraudDetectionRule interface {
	ValidateTransaction(transaction *Transaction, user *User) bool
}