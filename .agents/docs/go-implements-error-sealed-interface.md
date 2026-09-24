---
RFC Number: 495414
Title: "implError — Go's Sealed-Interface Realization of abstraction_p.Implements for Error"
Status: Rejected
Start Date: 2026-07-08
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
    contribution: "Identified a package-placement bug in a draft of this pattern (unexported method declared outside the protocol package it needed to satisfy) and corrected it; documented why the fix has no alternative given Go's visibility rules. Later conceded, after pushback, that the core 'sealed interface' guarantee this RFC was built around does not actually hold — see Rejection rationale below."
    task: []
---

## Summary
For `Error` specifically, memar-go realizes the conceptual contract of `abstraction_p.Implements` (RFC 495413) using a Go-specific, package-private mechanism instead of literally implementing that abstraction's
exported `Implements()` method. `ImplementsError`/`implError()`, together with the exported `ErrorMarker`
struct, form a true, compiler-enforced Go "sealed interface": only types that explicitly embed `ErrorMarker`
can satisfy `Error`'s implementation-intent requirement — no other package can independently declare a
matching method, whether by accident or by copy-paste.

## Motivation
RFC 495413 establishes that some abstractions benefit from a real, tooling-facing "implements" declaration
(primarily to support codegen scaffolding of incomplete capsules). Go, uniquely among Memar's current
backends, has package-private visibility available at zero extra cost. Rather than literally implementing
`abstraction_p.Implements`'s exported `Implements()` method — which, being exported, could be redeclared
independently by any package, gaining no additional guarantee beyond the tooling signal itself — memar-go
instead defines its own unexported analog (`implError`), which Go's compiler can enforce as a true sealed
interface: satisfying `Error` requires deliberately embedding `ErrorMarker`, not just happening to write a
method with the same name and signature.

This is presented as a bonus available specifically because Go's type system supports it "for free," not as
a return to the earlier, rejected accidental-satisfaction-prevention framing (RFC 495390) — the primary
motivation remains the same one RFC 495413 establishes: a real, discoverable implementation-intent
declaration for codegen purposes. The compiler-enforced sealing is a secondary property this Go-specific
realization happens to also provide.

## Guide-level explanation
The protocol package declares both the requirement and the only means of satisfying it:

```go
package error_p

type ImplementsError interface {
    implError()
}

type ErrorMarker struct{}

func (ErrorMarker) implError() {}
```

Any implementation package — necessarily a different package from `error_p`, since Memar keeps protocol
definitions separate from concrete implementations — embeds the exported marker to obtain the otherwise
unspellable unexported method:

```go
package service_not_found_p

import error_p "memar/process/error/protocol"

type ErrServiceNotFound struct {
    error_p.ErrorMarker
    // ... other fields specific to this concrete error ...
}
```

## Reference-level explanation
- **Why the marker must live in `error_p` itself — a hard constraint, not a preference.** An earlier draft
  of this pattern placed the marker struct (with its `implError` method) in a separate implementation
  package (named `error`, sibling to `error_p`), reasoning that this would keep the protocol package free of
  concrete types. This does not work: per the Go language specification, an unexported method's identity
  includes the package that declared it. A method literally named `implError`, declared in package `error`,
  is a *different identifier* from `implError` declared in package `error_p`, even with identical spelling
  and signature. A type embedding the `error`-package version would silently fail to satisfy
  `error_p.ImplementsError` — this is a hard compile error the moment such a type is used as an `error_p.Error`,
  not a subtle runtime bug, but it defeats the entire purpose of offering an easy embedding point. There is
  no way to preserve both "protocol package contains only interfaces" and "the sealed-interface mechanism
  works," when Go's unexported-identifier scoping is what it is. The marker must live in `error_p`.
- **What this buys beyond RFC 495413's base guarantee:** RFC 495413's `abstraction_p.Implements` alone gives
  a discoverable, but freely re-implementable, signal — any package could declare its own `Implements()`
  method independent of any other. `ImplementsError`/`ErrorMarker` additionally guarantees that satisfying
  `Error`'s implementation-intent requirement is only possible by deliberately embedding a specific,
  centrally-defined type, giving a single, greppable point of truth for "every capsule that has opted into
  being an Error."
- **Scope of the guarantee:** this prevents independent re-declaration from *outside* `error_p`. It does not
  prevent a type declared *inside* `error_p` itself from defining a conflicting `implError` directly without
  embedding `ErrorMarker` — `error_p` is a small, deliberately curated protocol-only package, so this
  residual risk is negligible and fully within the control of whoever maintains that one package.
- **Cost:** one zero-field embedded struct per concrete Error capsule; the promoted method resolves at
  compile time with no runtime dispatch cost beyond what the interface's other methods already require.

## Drawbacks
- Ties Go's realization of implementation-intent fully to Go's package system — `error_p` cannot be a
  pure-interface-only package if this pattern is used; it must also host `ErrorMarker`.
- Adds one extra embedded field to every concrete Go Error capsule's struct definition.
- Diverges from the literal, generic `abstraction_p.Implements` shape defined at the Memar level (see RFC
  495413's note on per-backend divergence) — any cross-language code generation tooling must know to treat
  Go's `implError` as the realization of the same concept, not attempt a literal, naive translation of
  method names across backends.

## Rationale and alternatives
- **Marker struct and method both unexported (an earlier, rejected draft):** broken — an unexported type
  cannot be embedded from outside its declaring package, making the interface unsatisfiable by any real
  concrete Error capsule, which necessarily lives outside `error_p`.
- **Marker struct in a separate implementation package, method still unexported (the most recent rejected
  draft, corrected by this RFC):** also broken, for the more subtle reason detailed above — unexported
  method identity is scoped to the declaring package, not just to "unexported-ness" in the abstract.
- **Exported marker struct + exported method (i.e. literally implementing `abstraction_p.Implements`):**
  works, and is simpler, but forfeits the compiler-enforced sealing guarantee for no benefit, since Go's
  tools for a stronger guarantee are available at the same implementation cost.
- **Exported marker struct + unexported method, declared in `error_p` (chosen):** the standard Go
  "sealed interface" idiom; gives a true compile-time-enforced guarantee at negligible cost.

## Prior art
This idiom is widely used in the Go ecosystem, most visibly in protobuf/gRPC generated code (patterns like
`mustEmbedUnimplemented...`), specifically to prevent external packages from implementing a given interface
without opting in through a supported, generated embedding point.

## Unresolved questions
- Naming convention for the exported marker type and its unexported method, across future memar-go
  abstractions that might adopt this same hardened pattern, is not yet decided beyond this one example
  (`ErrorMarker`/`implError`). Deliberately left open until a second real abstraction needs it, per the
  project's preference against generalizing from a single data point.
- Whether every backend realization of `abstraction_p.Implements` needs its own dedicated RFC like this one,
  or whether simpler backends (without an equivalent visibility mechanism) can just implement the generic
  method directly without needing a companion document, is open — likely "no companion RFC needed" for the
  simple case, but not yet stated as a rule anywhere.

## Future possibilities
If a second marker-like abstraction in memar-go adopts this pattern, use that occasion to settle the naming
convention question above rather than deciding it speculatively now.

## Rejection rationale (addendum, added after further discussion)
This RFC's central claim — that requiring `ErrorMarker` embedding produces a materially stronger guarantee
than a plain exported method — was examined further and found not to hold. The comparison that matters is
not "embedding vs. not embedding," it is: how hard is it for someone who *wants* to make an unrelated type
satisfy `Error` to do so, under each scheme? Under the scheme this RFC proposed, that person embeds
`error_p.ErrorMarker` — one line, no special knowledge required. Under the plain exported-method
alternative, that same person writes `func (X) ImplError() {}` — also one line. Neither is harder than the
other for anyone acting deliberately, which is the only case where "getting around it" is actually a
concern (the accidental-collision case was already judged implausible on its own, independent of this RFC,
given `Error`'s real, fully-expanded method count — see RFC 495413's Motivation). The sealed-interface
mechanism therefore added real cost (an implementation-package restriction on where the marker may live,
one extra embedded field per concrete type, and a second, harder-to-explain concept alongside RFC 495413's
simpler one) without buying a real guarantee in return.

The one part of this RFC's reasoning that does survive is the observation that a domain-specific method name
(`ImplError`, rather than the fully generic `Implements()`) is useful for disambiguating which abstraction a
capsule is claiming when more than one might apply — but that reasoning has nothing to do with Go
specifically, applies equally in Khayyam, and has been folded into RFC 495413 directly rather than kept here.
Consequently, `error_p.Error` no longer defines `ErrorMarker`; `ImplementsError` requires only a plain,
exported `ImplError()` method, with no required embedding and no stronger guarantee claimed than RFC 495413
already provides on its own.

This file is kept, in full, as the historical record of that reasoning trail, per project convention for
rejected RFCs.
