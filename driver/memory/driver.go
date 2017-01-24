package memory

import (
	"sync"

	"github.com/txgruppi/acl-go"
)

// NewDriver creates a new Driver
func NewDriver() *Driver {
	return &Driver{
		rules: map[string]map[string]bool{},
		mutex: sync.RWMutex{},
	}
}

// Driver is a 'in memory' ACL Driver
type Driver struct {
	policy bool
	rules  map[string]map[string]bool
	mutex  sync.RWMutex
}

// Begin - Check github.com/txgruppi/acl.Driver.Begin
func (d *Driver) Begin() error {
	return nil
}

// End - Check github.com/txgruppi/acl.Driver.End
func (d *Driver) End() error {
	return nil
}

// SetDefaultPolicy - Check github.com/txgruppi/acl.Driver.SetDefaultPolicy
func (d *Driver) SetDefaultPolicy(policy acl.Policy) error {
	d.policy = bool(policy)
	return nil
}

// GetActor - Check github.com/txgruppi/acl.Driver.GetActor
func (d *Driver) GetActor(id string) (acl.Actor, error) {
	return acl.NewSimpleActor(d, id), nil
}

// GetAction - Check github.com/txgruppi/acl.Driver.GetAction
func (d *Driver) GetAction(id string) (acl.Action, error) {
	return acl.NewSimpleAction(d, id), nil
}

// Set - Check github.com/txgruppi/acl.Driver.Set
func (d *Driver) Set(actor acl.Actor, action acl.Action, policy acl.Policy) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.rules[actor.String()] == nil {
		d.rules[actor.String()] = map[string]bool{}
	}
	d.rules[actor.String()][action.String()] = bool(policy)
	return nil
}

// IsAllowed - Check github.com/txgruppi/acl.Driver.IsAllowed
func (d *Driver) IsAllowed(actor acl.Actor, action acl.Action) (bool, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if d.rules[actor.String()] == nil {
		return d.policy, nil
	}
	if v, ok := d.rules[actor.String()][action.String()]; ok {
		return v, nil
	}
	return d.policy, nil
}
