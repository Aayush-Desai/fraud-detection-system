package main

type VelocityCheck struct {
	id string
	highValueThreshold float64
}

func GetVelocityCheckRule(id string,highValueThreshold float64) *VelocityCheck {
	return &VelocityCheck{
		id:                 id,
		highValueThreshold: highValueThreshold,
	}
}

func (c *VelocityCheck) ValidateTransaction(transaction *Transaction, user *User) bool {
	if transaction.GetAmmount() > c.GetHighValueThreshold() {
		return true
	}
	return false
}

func (c *VelocityCheck) GetHighValueThreshold() float64 {
	return c.highValueThreshold
}