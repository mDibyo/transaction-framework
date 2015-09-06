package transaction

import (
	"code.google.com/p/go-uuid/uuid"
)


type Person struct {
	id uuid.UUID
	name string
	debitors []Debit
	pCard PaymentCard
}


type PaymentCard struct {
	number [16]byte //stub
}

