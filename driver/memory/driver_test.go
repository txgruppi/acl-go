package memory_test

import (
	"testing"

	"github.com/nproc/acl-go"
	"github.com/nproc/acl-go/driver/memory"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemoryDriver(t *testing.T) {
	Convey("memory.Driver", t, func() {
		driver := memory.NewDriver()
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

			can, err = driver.IsAllowed(actor, action)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.Set(actor, action, acl.Deny)

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
