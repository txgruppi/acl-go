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

// Actor is an user which has access or not to one or more actions
type Actor interface {
	// IsAllowed checks if this user has access to an action
	IsAllowed(Action) (bool, error)

	// String returns the string ID of this actor
	String() string
}

// Action is an action which can be allowed or denied to one or more actors
type Action interface {
	// Allows checks if this action is allowed to an actor
	Allows(Actor) (bool, error)

	// String returns the string ID of this action
	String() string
}

// Driver defines the basic methods for an ACL manager driver
type Driver interface {
	// Begin the communication with the Driver's backend
	//
	// This method should be used to do any initialization before the driver can
	// be used
	Begin() error

	// End the communication with the Driver's backend
	//
	// This method should be used to do any cleanup after the driver is no more
	// needed
	End() error

	// SetDefaultPolicy defines the default access policy, whether to deny or allow
	//
	// The default policy is returned when a rule is not defined.
	// Its value should be false by default.
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
