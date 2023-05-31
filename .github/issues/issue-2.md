---
code: null
id: MDU6SXNzdWU0ODE5NzI3NzQ=
databaseId: 481972774
number: 2
url: https://github.com/octolab/protocol/issues/2
title: "endpoint specification"
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
createdAt: 2019-08-18T09:24:16Z
updatedAt: 2020-03-23T19:31:22Z
lastEditedAt: null
closedAt: 2020-03-11T19:58:26Z
---

# endpoint specification

```go
type Endpoint struct {
  Path string
  Map map[string]string
  Handler http.Handler
  Methods MethodMask
  In Request
  Out Response
}

type Request interface {
  Specified
  Serializable
}

type Response interface {
  Specified
  Serializable
}

var example = Endpoint{
  Path: "/example/{id}",
  Map: {"id" => "item"},
  ...
}

<!-- 2020-03-11T19:58Z https://github.com/octolab/protocol/issues/2#issuecomment-597837759
won't do
-->
