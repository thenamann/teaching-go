// package main

// import "fmt"

// /////////////////////
// // 1. ABSTRACTION  //
// /////////////////////

// // Payment defines WHAT can be done
// type Payment interface {
//     Pay(amount int)
// }

// /////////////////////
// // 2. ENCAPSULATION//
// /////////////////////

// // UPI struct (data hidden)
// type UPI struct {
//     upiId string // private field
// }

// // Constructor
// func NewUPI(id string) *UPI {
//     return &UPI{upiId: id}
// }

// // Method (implements Payment)
// func (u *UPI) Pay(amount int) {
//     fmt.Println("Paid", amount, "using UPI:", u.upiId)
// }

// /////////////////////
// // 3. COMPOSITION //
// /////////////////////

// // CardDetails reused by multiple cards
// type CardDetails struct {
//     number string
// }

// // CreditCard HAS-A CardDetails
// type CreditCard struct {
//     CardDetails
// }

// // Constructor
// func NewCreditCard(number string) *CreditCard {
//     return &CreditCard{
//         CardDetails: CardDetails{number: number},
//     }
// }

// // Method (implements Payment)
// func (c *CreditCard) Pay(amount int) {
//     fmt.Println("Paid", amount, "using Credit Card:", c.number)
// }

// /////////////////////
// // 4. POLYMORPHISM //
// /////////////////////

// // This function works for ALL payments
// func ProcessPayment(p Payment, amount int) {
//     p.Pay(amount)
// }

// /////////////////////
// // MAIN            //
// /////////////////////

// func main() {
//     upi := NewUPI("naman@upi")
//     card := NewCreditCard("1234-XXXX-XXXX-5678")

//     ProcessPayment(upi, 500)
//     ProcessPayment(card, 1000)
// }