package redis

import (
	"github.com/nproc/acl-go"
	"gopkg.in/redis.v3"
)

// NewDriver creates a new Driver
func NewDriver(client *redis.Client, prefix string) *Driver {
	return &Driver{
		client: client,
		prefix: prefix,
	}
}

// Driver is a 'in memory' ACL Driver
type Driver struct {
	defaultPolicy acl.Policy
	client        *redis.Client
	prefix        string
}

// Begin - Check github.com/nproc/acl.Driver.Begin
func (d *Driver) Begin() error {
	return d.client.Set(
		d.getDefaltPolicyKey(),
		d.policyToInt(d.defaultPolicy),
		0,
	).Err()
}

// End - Check github.com/nproc/acl.Driver.End
func (d *Driver) End() error {
	return nil
}

// SetDefaultPolicy - Check github.com/nproc/acl.Driver.SetDefaultPolicy
func (d *Driver) SetDefaultPolicy(policy acl.Policy) error {
	d.defaultPolicy = policy
	var value int
	if policy == acl.Allow {
		value = 1
	}
	return d.client.Set(d.getDefaltPolicyKey(), value, 0).Err()
}

// GetActor - Check github.com/nproc/acl.Driver.GetActor
func (d *Driver) GetActor(id string) (acl.Actor, error) {
	return acl.NewSimpleActor(d, id), nil
}

// GetAction - Check github.com/nproc/acl.Driver.GetAction
func (d *Driver) GetAction(id string) (acl.Action, error) {
	return acl.NewSimpleAction(d, id), nil
}

// Set - Check github.com/nproc/acl.Driver.Set
func (d *Driver) Set(actor acl.Actor, action acl.Action, policy acl.Policy) error {
	return d.client.Set(
		d.getRuleKey(actor, action),
		d.policyToInt(policy),
		0,
	).Err()
}

// IsAllowed - Check github.com/nproc/acl.Driver.IsAllowed
func (d *Driver) IsAllowed(actor acl.Actor, action acl.Action) (bool, error) {
	multi := d.client.Multi()
	cmder, err := multi.Exec(func() error {
		multi.SetNX(d.getDefaltPolicyKey(), d.policyToInt(d.defaultPolicy), 0)
		multi.MGet(d.getDefaltPolicyKey(), d.getRuleKey(actor, action))
		return nil
	})
	if err != nil {
		return false, err
	}

	result, err := cmder[1].(*redis.SliceCmd).Result()
	if err != nil {
		return false, err
	}

	defaultValue := false
	if result[0].(string) == "1" {
		defaultValue = true
	}
	d.defaultPolicy = acl.Policy(defaultValue)

	if result[1] == nil {
		return defaultValue, nil
	}

	if result[1].(string) == "1" {
		return true, nil
	}

	return false, nil
}

func (d *Driver) policyToInt(policy acl.Policy) int {
	if policy == acl.Allow {
		return 1
	}
	return 0
}

func (d *Driver) getDefaltPolicyKey() string {
	return d.prefix + ":rule:default"
}

func (d *Driver) getRuleKey(actor acl.Actor, action acl.Action) string {
	return d.prefix + ":rule:" + actor.String() + ":" + action.String()
}
