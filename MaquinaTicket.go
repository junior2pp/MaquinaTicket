package maquina

import "fmt"

type Ticket struct {
	precio   int
	balance  int
	total    int
	contador int
}

func (t *Ticket) Init(ticketCost int) {
	t.precio = ticketCost
	t.balance = 0
	t.total = 0
	t.contador = 0
}
func (t *Ticket) GetPrecio() int {
	return t.precio
}
func (t *Ticket) GetBalance() int {
	return t.balance
}
func (t *Ticket) SetBalance(balanc int) {
	if balanc > 0 {
		t.balance += balanc
	}

}
func (t *Ticket) GetContador() int {
	return t.contador
}
func (t *Ticket) GetTotal() int {
	return t.total
}
func (t *Ticket) SetTotal(totaly int) {
	t.total = totaly
}
func (t *Ticket) VaciarMaquina() int {
	tempTotal := t.total
	t.total = 0
	t.contador = 0
	return tempTotal
}
func (t *Ticket) ImprimirTicket() {
	if t.balance >= t.precio {
		t.contador++
		t.balance = t.balance - t.precio
		t.total = t.total + t.precio
		fmt.Println("----------------------")
		fmt.Println("TICKET #", t.contador)
		fmt.Println("PRECIO", t.precio)
		fmt.Println("----------------------")
	} else {
		fmt.Printf("Necesita %v centavos mas", (t.precio - t.balance))
	}
}
