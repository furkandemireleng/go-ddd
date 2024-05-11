// package entity holds all the entities that are shared across subdomains

package entity

//Person is an entity that represents a person in all domains

import "github.com/google/uuid"

type Person struct {
	//ID is the identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
