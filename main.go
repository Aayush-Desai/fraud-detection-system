package main

import "fmt"

func main()  {
	fraudDetectionSystem := &FraudDetectionSystem{}

	device1 :=GetNewDevice("1", "Mobile")

	user1 := RegesterUser("1", "Aayush", "Delhi", device1)

	account1 := GetNewAccount("1", "SBI", user1, 1000)

	transaction1 := GerNewTransaction("1", user1, 100, "Delhi", device1, 0, account1)

	ruleBasedCheck := GetRuleBasedCheck("1")

	fraudDetectionSystem.AddNewFalutDetectionRule(ruleBasedCheck)
	
	x := fraudDetectionSystem.ProcessTransaction(transaction1, user1)
	fmt.Println(x)
}