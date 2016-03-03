package acl

// NewSimpleAction creates a new SimpleAction with the given ID and Driver
func NewSimpleAction(driver Driver, id string) *SimpleAction {
	return &SimpleAction{
		id:     id,
		driver: driver,
	}
}

// SimpleAction represents an ACL action, it implements the Action interface
type SimpleAction struct {
	id     string
	driver Driver
}

// Allows - Check Action's Allows
func (s *SimpleAction) Allows(actor Actor) (bool, error) {
	return s.driver.IsAllowed(actor, s)
}

// String - Check Action's String
func (s *SimpleAction) String() string {
	return s.id
}
