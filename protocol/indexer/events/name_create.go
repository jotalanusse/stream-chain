package events

// NewNameCreateEvent creates a NameCreateEvent representing creation of a name.
func NewNameCreateEvent(
	id uint32,
	name string,
) *NameCreateEventV1 {
	return &NameCreateEventV1{
		Id:   id,
		Name: name,
	}
}
