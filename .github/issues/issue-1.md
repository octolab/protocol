---
code: null
id: MDU6SXNzdWU0ODE5NzE2MDI=
databaseId: 481971602
number: 1
url: https://github.com/octolab/protocol/issues/1
title: "middleware with tracer injection"
issueType: null
labels:
  - "help wanted"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2019-08-18T09:08:54Z
updatedAt: 2020-03-23T19:31:21Z
lastEditedAt: 2020-03-12T05:05:05Z
closedAt: 2020-03-12T05:07:58Z
---

# middleware with tracer injection

- [x] implementation
```go
func Middleware(http.Handler) http.Handler {
  return http.HandlerFunc
    req = req.WithContext(context.WithValue)
    handler.ServerHTTP
    if ctx.Err log the tracer
}
```
- [x] example to the tracer project
