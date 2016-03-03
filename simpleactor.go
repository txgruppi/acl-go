package acl

// NewSimpleActor creates a new SimpleActor with the given ID and Driver
func NewSimpleActor(driver Driver, id string) *SimpleActor {
	return &SimpleActor{
		id:     id,
		driver: driver,
	}
}

// SimpleActor represents an ACL actor, it implements the Actor interface
type SimpleActor struct {
	id     string
	driver Driver
}

// IsAllowed - Check Actor's IsAllowed
func (s *SimpleActor) IsAllowed(action Action) (bool, error) {
	return s.driver.IsAllowed(s, action)
}

// String - Check Actor's String
func (s *SimpleActor) String() string {
	return s.id
}
