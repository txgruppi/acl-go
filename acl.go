package acl

// Policy defines the Allow and Deny policies
type Policy bool

const (
	// Deny should be used when the actor has no access to an action
	Deny Policy = false

	// Allow should be used when the actor has access to an action
	Allow = true
)

// ACL defines the basic methods for an ACL manager
type ACL interface {
	// SetDefaultPolicy - Check Driver's .SetDefaultPolicy
	SetDefaultPolicy(Policy) error

	// GetActor - Check Driver's .GetActor
	GetActor(string) (Actor, error)

	// GetAction - Check Driver's .GetAction
	GetAction(string) (Action, error)

	// Set - Check Driver's .Set
	Set(Actor, Action, Policy) error
}

// Actor is a user which has access or not to one or more actions
type Actor interface {
	// IsAllowed checks if this user has access to an action
	IsAllowed(Action) (bool, error)

	// String returns the string ID of this actor
	String() string
}

// Action is a action which can be allowed or denied to one or more actors
type Action interface {
	// Allows checks if this action is allowed to an actor
	Allows(Actor) (bool, error)

	// String returns the string ID of this action
	String() string
}

// Driver defines the basic methods for an ACL manager driver
type Driver interface {
	// Open the communication with the Driver's backend
	Open() error

	// Close the communication with the Driver's backend
	Close() error

	// SetDefaultPolicy defines what is the default policy
	//
	// The default policy is return when a rule is not defined
	SetDefaultPolicy(Policy) error

	// GetActor returns an Actor with the given ID
	GetActor(string) (Actor, error)

	// GetAction returns an Action with the given ID
	GetAction(string) (Action, error)

	// Set defines a rule for a Actor and Action, which can be Allow or Deny
	Set(Actor, Action, Policy) error

	// IsAllowed checks if an Actor has access to an Action
	IsAllowed(Actor, Action) (bool, error)
}
