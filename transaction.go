package transaction

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"time"
)

// AmountFigure represents an amount of money. It is stored as an int64
// equaling (amount x 100)
type AmountFigure int64

// NewAmountFigure creates an AmountFigure from units and fractions of
// currencies and the debit flag. If debit is true, then AmountFigure is
// considered to be a debit, and is positive. Otherwise, AmountFigure is
// considered to be a credit, and is negative.
func NewAmountFigure(units int64, fractions int64, debit bool) (af AmountFigure, err error) {
	if units < 0 {
		return af, fmt.Errorf("units %d is less than 0", units)
	}
	if fractions < 0 {
		return af, fmt.Errorf("fractions %d is less than 0", fractions)
	}
	if fractions >= 100 {
		return af, fmt.Errorf("fractions %d is not less than 100", fractions)
	}
	sign := int64(-1)
	if debit {
		sign = 1
	}
	return AmountFigure(sign * (units*100 + fractions)), nil
}

// Float64 returns an AmountFigure in float64
func (a AmountFigure) Float64() float64 {
	return float64(a) / 100
}

// Add takes another AmountFigure and returns their sum.
func (a AmountFigure) Add(b AmountFigure) AmountFigure {
	return AmountFigure(int64(a) + int64(b))
}

// Debit represents the amount a person owes. Amount is negative when the
// person instead of owing money, is owed money.
type Debit struct {
	Timestamp int64
	Person    Person
	Amount    AmountFigure
	Resolved  bool
}

func NewDebit(p Person, a AmountFigure) *Debit {
	return &Debit{
		Timestamp: time.Now().UnixNano(),
		Person:    p,
		Amount:    a,
	}
}

// Record represents the record of a single payment. It holds a slice of
// Debits, one for each person that must to a payment.
type Record []Debit

// Valid returns if a record is valid. A record in only valid if all Debit
// amounts in it add up to 0 ie. all contributors to the payment are
// accounted for.
func (r *Record) Valid() bool {
	var total AmountFigure
	for _, d := range *r {
		total = total.Add(d.Amount)
	}
	return total.Float64() == 0
}
