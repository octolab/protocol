---
code: null
id: MDU6SXNzdWU1OTcxMjU5ODI=
databaseId: 597125982
number: 8
url: https://github.com/octolab/protocol/issues/8
title: "middleware which wraps http.TimeoutHandler"
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
createdAt: 2020-04-09T08:49:33Z
updatedAt: 2020-04-10T19:45:49Z
lastEditedAt: null
closedAt: 2020-04-10T19:45:49Z
---

# middleware which wraps http.TimeoutHandler

```go
func(handler http.Handler) http.Handler {
		return http.TimeoutHandler(handler, 3*time.Second, "Timeout Error")
	}
```
