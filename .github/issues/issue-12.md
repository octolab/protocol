---
code: null
id: MDU6SXNzdWU3NDMzMDc3ODI=
databaseId: 743307782
number: 12
url: https://github.com/octolab/protocol/issues/12
title: "http: interceptor: request dumper with tracer support"
issueType: null
labels:
  - "difficulty: hard"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2020-11-15T17:49:53Z
updatedAt: 2022-05-28T16:07:17Z
lastEditedAt: 2022-05-28T16:07:17Z
closedAt: null
---

# http: interceptor: request dumper with tracer support

**Motivation:** sometimes is needed to smokes based on real calls.

Steps

- add a possibility to dump requests into `request.http` format
- dump metadata to restore the call graph
- integrate with extended [kamilsk/tracer](https://github.com/kamilsk/tracer)

The metadata allows defining who is the root, the first-level children, etc.

<!-- 2020-11-15T17:53Z https://github.com/octolab/protocol/issues/12#issuecomment-727609527
blocked by kamilsk/tracer#8
-->
