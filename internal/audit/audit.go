package audit

import "github.com/google/uuid"

type Entry struct {
	ID     string
	Action string
	Amount float64
}

func NewEntry(action string, amount float64) Entry {
	return Entry{ID: uuid.NewString(), Action: action, Amount: amount}
}
