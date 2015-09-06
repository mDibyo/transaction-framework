package transaction

import (
	"code.google.com/p/go-uuid/uuid"
)

type PaymentCard struct {
	number [16]byte //stub
}

type Person struct {
	Id   uuid.UUID
	Name string
	// Debits represents the amounts that a person is owed by each of
	// his debitors. These amounts may be negative, which means that
	// the amount is owed by this person. There might be multiple
	// debits for a single debitor.
	Debits []Debit
	PCard  PaymentCard
}

func NewPerson(name string) *Person {
	return &Person{
		Id:   uuid.NewRandom(),
		Name: name,
	}
}

// Owed returns the amount a person is owed by his debitors.
func (p *Person) Owed() AmountFigure {
	var total AmountFigure
	for _, d := range p.Debits {
		total = total.Add(d.Amount)
	}
	return total
}
