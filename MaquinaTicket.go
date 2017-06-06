package maquina

type Ticket struct {
	precio   int
	balance  int
	total    int
	contador int
}

func (t *Ticket) Init(ticketCost int) {
	t.precio = ticketCost
}
func (t *Ticket) GetPrecio() int {
	return t.precio
}
func (t *Ticket) GetBalance() int {
	return t.balance
}
func (t *Ticket) SetBalance(balanc int) {
	t.balance = balanc
}
func (t *Ticket) GetContador() int {
	return t.contador
}
func (t *Ticket) IncrementarContador() {
	t.contador++
}
func (t *Ticket) GetTotal() int {
	return t.total
}
func (t *Ticket) SetTotal(totaly int) {
	t.total = totaly
}
func (t *Ticket) VaciarMaquina() int {
	TempTotal := t.total
	t.total = 0
	return TempTotal
}
func (t *Ticket) ImprimirTicket() (int, int) {
	return t.contador, t.precio
}
