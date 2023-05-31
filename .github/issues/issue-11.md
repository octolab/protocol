---
code: null
id: MDU6SXNzdWU3NDMzMDY5OTY=
databaseId: 743306996
number: 11
url: https://github.com/octolab/protocol/issues/11
title: "RoundTripper middlewares"
issueType: null
labels:
  - "help wanted"
  - "difficulty: easy"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2020-11-15T17:45:53Z
updatedAt: 2021-01-10T20:15:06Z
lastEditedAt: null
closedAt: 2021-01-10T20:15:06Z
---

# RoundTripper middlewares

- [ ] define type
```go
type Middleware func (http.RoundTripper) http.RoundTripper
```
- [ ] prefixed RoundTripper
```go
func RoundTripper(prefix) Middleware
```
- [ ] timeout RoundTripper
```go
func RoundTripperMiddleware(http.RoundTripper) http.RoundTripper

inject X-Deadline header
```

<!-- 2020-11-15T17:58Z https://github.com/octolab/protocol/issues/11#issuecomment-727610291
see kamilsk/tracer#9
-->

<!-- 2021-01-10T20:14Z https://github.com/octolab/protocol/issues/11#issuecomment-757536666
moved to the separated project: octolab/propaganda#1
-->
