---
code: null
id: MDU6SXNzdWU2NDAxOTU0MDE=
databaseId: 640195401
number: 9
url: https://github.com/octolab/protocol/issues/9
title: "reduce boilerplate"
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
createdAt: 2020-06-17T07:06:06Z
updatedAt: 2020-06-17T11:06:12Z
lastEditedAt: null
closedAt: 2020-06-17T11:06:12Z
---

# reduce boilerplate

```go
	{
		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			_, _, spanName := runtime.Caller().Meta()
			child := span.Tracer().StartSpan(
				spanName,
				opentracing.ChildOf(span.Context()),
			)
			defer child.Finish()
			ctx = opentracing.ContextWithSpan(ctx, child)
		}
	}
```
instead of that, I need `defer tracer.Child(ctx).Finish()`

<!-- 2020-06-17T07:07Z https://github.com/octolab/protocol/issues/9#issuecomment-645193852
Forget about the context:

```go
defer tracer.Child(ctx).RewriteContext(func(newCtx) { ctx = newCtx }).Finish()
```
-->

<!-- 2020-06-17T11:06Z https://github.com/octolab/protocol/issues/9#issuecomment-645307982
won't do, cause .StartFromContext exists
-->
