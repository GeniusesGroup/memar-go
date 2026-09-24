---
RFC Number: 495390
Title: "implError — Intentional Abstraction Satisfaction"
Status: Draft
Start Date: 2026-07-07
Applied to: ["error.go"]
Supersedes: null
Superseded by: null
Related:
  Depends_on: []
  Extends: ["495413"]
  Conflicts with: []
Contributor(s):
  - name: "Omid Hekayati"
    uri: "mailto:omid@geniuses.group"
    contribution: ""
    task: []
  - name: "Claude"
    uri: "https://claude.ai"
    model: "claude-sonnet-5"
    effort: "Medium - extended thinking enabled"
    contribution: ""
    task: []
---

# Intentional Abstraction Satisfaction — Go Backend Pattern

## Problem

Go uses structural typing: any concrete type whose method set matches an interface satisfies that interface automatically, regardless of whether the author of that type ever intended to implement it. For most interfaces with several distinct, domain-specific methods, an accidental match is vanishingly unlikely and this is a non-issue. But for small, "marker-like" abstractions — few methods, often themselves generic (e.g. `Error`, which composes `DataType`, `Field_MediaType`, and `ADT`) — the risk of an unrelated type accidentally satisfying the abstraction is real, and there is no compiler error to catch it: the mistaken type simply gets treated as a valid instance of the abstraction wherever it's used.

This is not a Khayyam or Memar framework problem — it is specific to how Go's type system works. Whether Khayyam itself needs an equivalent discipline at the language level is tracked separately (see RFC 495390, "Intentional Abstraction Satisfaction", in the memar RFC set) and is not assumed to have the same answer.

## Pattern
For any Memar abstraction ported to Go as an `interface` where accidental structural satisfaction is a meaningful risk (small method count, identity-bearing/marker-like role), add an unexported marker method to the interface, named `impl<AbstractionName>()`, and provide a small unexported marker struct implementing it:

```go
type <name>Marker struct{}

func (<name>Marker) impl<AbstractionName>() {}
```

Every concrete capsule that implements the abstraction embeds this marker struct (typically via code
generation, since it is boilerplate with zero variation across capsules). This makes satisfaction explicit
and intentional — a type cannot accidentally satisfy the abstraction just by having matching method
signatures; it must deliberately embed the marker.

## When to apply this pattern

Only where the risk is real. Do not apply this reflexively to every abstraction in memar-go — most
abstractions have enough distinct, domain-specific methods that accidental structural satisfaction is not a
practical concern, and adding a marker method to all of them would be unnecessary ceremony working against
Memar's minimalism principle. Apply it specifically to small, marker-like, identity-bearing abstractions.
`Error` (see `error.go`) is the first applied example.

## Cost

Negligible. The marker method is resolvable at compile time and adds no runtime dispatch cost beyond what
the interface's other methods already require; it exists purely to change what the Go compiler will accept
as a valid implementation, not to be called at runtime by any business logic.

## Applied example

See `error.go`: `Error` embeds `implError()`, satisfied via the exported `Error` struct.
