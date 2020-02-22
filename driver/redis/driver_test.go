package redis_test

import (
	"testing"

	"github.com/alicebob/miniredis"
	rds "github.com/go-redis/redis/v7"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/txgruppi/acl-go"
	"github.com/txgruppi/acl-go/driver/redis"
)

func TestRedisDriver(t *testing.T) {
	server, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	client := rds.NewClient(&rds.Options{
		Addr: server.Addr(),
		DB:   0,
	})

	Convey("redis.Driver", t, func() {
		Reset(func() {
			client.FlushDB()
		})

		driver := redis.NewDriver(client, "acl")
		actor, err := driver.GetActor("testActor")
		So(err, ShouldBeNil)
		action, err := driver.GetAction("testAction")
		So(err, ShouldBeNil)

		Convey(".Begin should return nil", func() {
			So(driver.Begin(), ShouldBeNil)
		})

		Convey(".End should return nil", func() {
			So(driver.End(), ShouldBeNil)
		})

		Convey("it should start with default policy as Deny", func() {
			can, err := driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})

		Convey("it should set the default policy", func() {
			err := driver.SetDefaultPolicy(acl.Allow)
			So(err, ShouldBeNil)
			can, err := driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.SetDefaultPolicy(acl.Deny)
			So(err, ShouldBeNil)
			can, err = driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})

		Convey("it should set a access rule", func() {
			err := driver.SetDefaultPolicy(acl.Deny)
			So(err, ShouldBeNil)

			can, err := driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)

			err = driver.Set(actor, action, acl.Allow)
			So(err, ShouldBeNil)

			can, err = driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.Set(actor, action, acl.Deny)
			So(err, ShouldBeNil)

			can, err = driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})

		Convey("it should return the default policy if actor is defined but action is not", func() {
			anotherAction, err := driver.GetAction("someAction")
			So(err, ShouldBeNil)

			err = driver.SetDefaultPolicy(acl.Allow)
			So(err, ShouldBeNil)

			err = driver.Set(actor, action, acl.Deny)
			So(err, ShouldBeNil)

			can, err := driver.IsAllowed(actor, anotherAction)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.SetDefaultPolicy(acl.Deny)
			So(err, ShouldBeNil)

			err = driver.Set(actor, action, acl.Allow)
			So(err, ShouldBeNil)

			can, err = driver.IsAllowed(actor, anotherAction)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})

		Convey("it should return the default policy if action is defined but actor is not", func() {
			anotherActor, err := driver.GetActor("someActor")
			So(err, ShouldBeNil)

			err = driver.SetDefaultPolicy(acl.Allow)
			So(err, ShouldBeNil)

			err = driver.Set(actor, action, acl.Deny)
			So(err, ShouldBeNil)

			can, err := driver.IsAllowed(anotherActor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.SetDefaultPolicy(acl.Deny)
			So(err, ShouldBeNil)

			err = driver.Set(actor, action, acl.Allow)
			So(err, ShouldBeNil)

			can, err = driver.IsAllowed(anotherActor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})
	})
}
