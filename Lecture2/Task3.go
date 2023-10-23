package main

//непонял как распределение по пакетам/папкам/файлам

import (
	"fmt"
	"time"
)

// User interface for Consumer and Dealer
type User interface {
	getBalance() float64
	getVerification() bool
}

// Consumer uses User interface
type Consumer struct {
	balance float64
	//Usage of composition
	h      []history
	c      credentials
	status bool
}

// Dealer uses User interface
type Dealer struct {
	balance float64
	//Usage of composition
	h      []history
	c      credentials
	status bool
}

// Composition for Consumer and Dealer
type history struct {
	date   string
	time   string
	amount float64
	status string
}

// Composition for Consumer and Dealer
type credentials struct {
	name    string
	email   string
	address string
	bank    string
	bin     int
}

func (c *Consumer) getBalance() float64 {
	return c.balance
}

func (d *Dealer) getBalance() float64 {
	return d.balance
}

// get verified for making transactions for person
func (c *Consumer) getVerification() {
	tempStatus := false
	if sendCredentialsToBank(c.c) && sendEmailVerification(c.c.email) {
		tempStatus = true
	}
	c.status = tempStatus
}

// get verified for making transactions for dealer
func (d *Dealer) getVerification() {
	tempStatus := false
	if sendCredentialsToBank(d.c) && sendEmailVerification(d.c.email) && historyVerification(d.h) {
		tempStatus = true
	}
	d.status = tempStatus
}

// verify bank credentials
func sendCredentialsToBank(credentials credentials) bool {
	//Business logic
	fmt.Println("Checking for bank credentials...")
	return true
}

// verify email
func sendEmailVerification(email string) bool {
	//Business logic
	fmt.Println("Cheking for email")
	return true
}

// verify Dealer for history success rate
func historyVerification(history []history) bool {
	historyCount := len(history)
	failCount := 0
	for _, item := range history {
		if item.status == "fail" {
			failCount += 1
		}
	}
	if failCount/historyCount*100 < 60 {
		return false
	}
	return true
}

// Transaction between Consumer and Dealer
func Transaction(c *Consumer, d *Dealer, amount float64) bool {
	if c.status == false || d.status == false {
		return false
	}
	now := time.Now()
	history := history{date: now.Format("2006-01-02"), time: now.Format("15:04:05"), status: "success"}

	if getMoneyFromDealer(d, amount, history) && sendMoneyToConsumer(c, amount, history) {
		return true
	}

	return false
}

// Get money from Dealer
func getMoneyFromDealer(d *Dealer, amount float64, history history) bool {
	if d.balance < amount {

		d.h = append(d.h, history)
		return false
	}
	d.balance -= amount
	d.h = append(d.h, history)
	return true
}

// Send money to Consumer
func sendMoneyToConsumer(c *Consumer, amount float64, history history) bool {
	temp := c.balance
	c.balance += amount
	if c.balance != temp+amount {
		history.status = "fail"
		c.h = append(c.h, history)
		return false
	}
	c.h = append(c.h, history)
	return true
}

func main() {
	//create Consumer
	consumerCredentials := credentials{name: "Daniar Yermakhan", email: "dy@gmail.com", address: "Abay street 1/1", bank: "Kaspi", bin: 123456789}
	c1 := &Consumer{balance: 0.00, c: consumerCredentials, status: false}
	//verify consumer
	c1.getVerification()

	//create Dealer by default status is true
	dealerCredentials := credentials{name: "John Smith", email: "js@gmail.com", address: "Abay street 10/1", bank: "Kaspi", bin: 987654321}
	d1 := &Dealer{balance: 1000.00, c: dealerCredentials, status: true}

	if Transaction(c1, d1, 500.00) {
		fmt.Println("Transaction is successful")
	} else {
		fmt.Println("Transaction failed")
	}
	//Check if transaction is successful and verify Dealer
	d1.getVerification()

	fmt.Println(c1.balance)
	fmt.Println(d1.balance)

	if Transaction(c1, d1, 600.00) {
		fmt.Println("Transaction is successful")
	} else {
		fmt.Println("Transaction failed")
	}

	fmt.Println(c1.balance)
	fmt.Println(d1.balance)
}
