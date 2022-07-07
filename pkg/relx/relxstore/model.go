package relxstore

type Entities[Entity any] struct {
	Items   []*Entity
	HasMore bool
}
