package main

type AnomalyDetection struct {
	id string
}

func GetAnomalyDetectionRule(id string) *AnomalyDetection {
	return &AnomalyDetection{
		id: id,
	}
}

func (c *AnomalyDetection) ValidateTransaction(transaction *Transaction) bool {
	if transaction.GetAmmount()==0 {
		return true
	}
	return false
}