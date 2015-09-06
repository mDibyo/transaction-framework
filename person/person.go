package person

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go-uuid/uuid"
)


type Person struct {
	id uuid.UUID
	name string
	pCard PaymentCard

}


type PaymentCard struct {
	number [16]byte //stub
}

