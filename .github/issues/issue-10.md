---
code: null
id: MDU6SXNzdWU3NDMyOTk2NDA=
databaseId: 743299640
number: 10
url: https://github.com/octolab/protocol/issues/10
title: "support X-Features"
issueType: null
labels:
  - "help wanted"
  - "difficulty: medium"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2020-11-15T17:08:06Z
updatedAt: 2021-01-10T20:18:45Z
lastEditedAt: null
closedAt: 2021-01-10T20:18:45Z
---

# support X-Features

- [ ] add specific header
```go
req.Header.Set(http.XFeatures, config.Feature{Name, false}.String())
```
- [ ] extend config module
```go
features.Find -> alias for FindByName

config.Features.Unmarshal(req.Header.Get(http.XFeatures)).Find(needs)
```
- [ ] context support
```go
ctx = context.WithValue(req.Context, Features, MutatedFeatures)
```
- [ ] timeout/deadline support
```go
context.Value(Features).Find(Deadline).Enabled -> apply middleware
```

<!-- 2020-11-15T17:13Z https://github.com/octolab/protocol/issues/10#issuecomment-727603784
blocked by octolab/config#4
-->

<!-- 2021-01-10T20:18Z https://github.com/octolab/protocol/issues/10#issuecomment-757537263
moved to octolab/propaganda#4
-->
