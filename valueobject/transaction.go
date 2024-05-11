package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a value object because it has no identifier and it is immutable

type Transaction struct {
	//lower case for no other domain reach and changed the value
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
