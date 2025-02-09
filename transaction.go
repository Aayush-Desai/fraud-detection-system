package main

type Transaction struct {
    transactionId string
	user   *User
	ammount   float64
	location string
	device *Device
    timeStamp int64
	account *Account
	fraudulantTransaction bool
}


func GerNewTransaction(transactionId string,user *User, amount float64, location string, device *Device, timestamp int64, account *Account) *Transaction {
	return &Transaction{
		transactionId: transactionId,
		user: user,
		ammount: amount,
		location: location,
		device: device,
		timeStamp: timestamp,
		account: account,
	}
}

func (t *Transaction) ExecuteTransaction()  bool {
	return true
}

func (t *Transaction) SetFraudulantTransaction(isFraudulantTransaction bool)  {
	t.fraudulantTransaction = isFraudulantTransaction
}

func (t *Transaction) GetFraudulantTransaction() bool {
	return t.fraudulantTransaction
}

func (t *Transaction) GetEventType() EventType {
	return EventTypeTransaction
}

func (t *Transaction) GetAmmount() float64 {
	return t.ammount
}

func (t *Transaction) GetLocation() string {
	return t.location
}

func (t *Transaction) GetDevice() *Device {
	return t.device
}

func (t *Transaction) GetAccount() *Account {
	return t.account
}




