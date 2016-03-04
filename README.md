[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/nproc/acl-go)
![Codeship](https://img.shields.io/codeship/6149a4b0-c485-0133-222a-265f477b0567.svg?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/nproc/acl-go.svg?style=flat-square)](https://codecov.io/github/nproc/acl-go)
[![Go Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat-square)](https://goreportcard.com/report/github.com/nproc/acl-go)

# ACL - Access Control List

ACL is a simple but powerful Access Control List manager

## Installation

```
go get -u github.com/nproc/acl-go
```

## Example

*You should not ignored the errors returned by the methods*

```go
package main

import (
  "fmt"

  "github.com/nproc/acl-go"
  "github.com/nproc/acl-go/driver/memory"
)

func main() {
  driver := memory.NewDriver()

  // Driver can be directly used as ACL managers
  var manager acl.ACL = driver

  // Set the default policy as Deny
  acl.SetDefaultPolicy(acl.Deny)

  // Get some users
  userCEO, _ := acl.GetActor("userCEO_UUID")
  userDeveloper, _ := acl.GetActor("userDeveloper_UUID")

  // Get some actions
  accessBackAccount, _ := acl.GetAction("accessBackAccount")
  accessProductionServer, _ := acl.GetAction("accessProductionServer")

  // Set rules
  acl.Set(userCEO, accessBackAccount, acl.Allow)
  acl.Set(userDeveloper, accessProductionServer, acl.Allow)

  // Check using the ACL manager
  allowed, _ := acl.IsAllowed(userCEO, accessBackAccount)
  fmt.Println(allowed) // true
  allowed, _ = acl.IsAllowed(userDeveloper, accessBackAccount)
  fmt.Println(allowed) // false

  // Check using the Actor or Action struct
  allowed, _ := userCEO.IsAllowed(accessProductionServer)
  fmt.Println(allowed) // false
  allowed, _ = accessProductionServer.Allows(userDeveloper)
  fmt.Println(allowed) // true
}
```

## License

MIT
