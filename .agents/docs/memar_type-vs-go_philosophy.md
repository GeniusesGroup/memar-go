---
RFC Number: 495422
Title: "Static Concepts Must Be Types — vs Go Philosophy"
Status: Proposed
Start Date: 2026-07-08
Applied to: []
Supersedes: null
Superseded by: null
Related:
  Depends_on: ["495421", "000001"]
  Extends: ["495421"]
  Conflicts with: []
Contributor(s):
  - name: "Omid Hekayati"
    uri: "mailto:omid@geniuses.group"
    contribution: "Go philosophy analysis, Memar vs Go optimization target comparison, insight that Go ecosystem rebuilds type identity around errors without admitting it"
    task: []
  - name: "Super Z"
    uri: ""
    model: "GLM"
    effort: ""
    contribution: "Compiled community research, documented trade-offs, drafted analysis"
    task: []
  - name: "ChatGPT"
    uri: "https://openai.com"
    model: "GPT-5.5"
    effort: "RFC review and architectural refinement"
    contribution: "Internet research on Go community critiques; compiled citations from Rob Pike, Dave Cheney, Go 1.13 design docs, DoltHub, Matt Proud, golangci-lint, CedarDB, yegor256"
    task: []
  - name: "Claude"
    uri: "https://claude.ai"
    model: "claude-sonnet-5"
    effort: "Medium - extended thinking enabled"
    contribution: "Verified Pike, Cheney, and DoltHub citations against primary sources; flagged Proud, Cox, and the Go 1.13 design doc citations as not independently re-verified; noted the exhaustiveness gap applies to the Go implementation as well and has no Go-native closed-type mechanism to resolve it"
    task: []
---

## Summary

This document is a companion to RFC 495421 ("Static Concepts Must Be Types"). That RFC establishes a universal Memar principle: identity belongs to the type system, and static capsules MUST be distinct types. This document analyzes what happens when that principle is implemented in Go — a language whose community has deliberately chosen a value-oriented error model where identity is carried by runtime data.

The Go ecosystem and the Memar framework optimize for different targets. Neither is objectively "correct" in isolation — each is correct for its own goals. Go values flexibility, simplicity, and ergonomics. Memar values type-level contract identity, architectural clarity, and compile-time verifiability. This document documents the specific tensions, trade-offs, and community critiques that arise at this intersection, and explains why memar-go — which compiles with the Go compiler but follows Memar's architectural authority (RFC 000001) — deliberately diverges from Go idioms.

## Why This Document Exists

RFC 495421 is a universal principle document. It applies to Memar regardless of implementation language. It intentionally does not contain language-specific analysis, because the principle is not about Go — it is about what the type system is responsible for in the Memar framework.

However, memar-go is currently the primary implementation. Developers working in memar-go will encounter friction between this RFC's principle and Go's established idioms. This friction is real, expected, and worth documenting in detail — not to argue that Go is wrong, but to explain why Memar makes a different choice and what the practical consequences are.

This document also serves as a record of the research and analysis that informed RFC 495421, preserving the community critiques, performance data, and philosophical comparisons that would be out of place in the principle RFC itself but are valuable as context for future development decisions.

## The Go Ecosystem's Implicit Question

The Go community's response to its own value-oriented error model is itself instructive. Over more than a decade, the Go ecosystem has progressively built type-level identity mechanisms around errors — without explicitly acknowledging that this is what it is doing:

- **`errors.New`** (Go 1.0): pure string-based identity. The type system provides no information about which error is represented.

- **Custom error types** (community practice from early Go): developers who needed structured, identifiable errors began defining their own types — `type NotFoundError struct{}` — because sentinel strings were insufficient for their needs. This is the community implicitly rediscovering the principle that RFC 495421 makes explicit.

- **`errors.Is` / `errors.As`** (Go 1.13): introduced to solve the problem that wrapped errors could no longer be compared by pointer equality. These functions restore the ability to determine "which error is this" — but they do so through runtime chain-walking, not through the type system. The type system still only knows `error`.

- **`golangci-lint` `errname` rule**: enforces that sentinel errors use the `Err` prefix (e.g., `ErrNotFound`) and error types use the `Error` suffix (e.g., `NotFoundError`). This lint rule encodes a naming convention that distinguishes identity-as-data (`ErrNotFound`, a variable) from identity-as-type (`NotFoundError`, a type definition) — a distinction the language itself does not enforce but the community has found necessary enough to automate.

- **Interface-based predicates** (custom `Is(target error) bool` methods): allow error types to define their own matching logic, effectively implementing type-level identity checking through a runtime method dispatch. This is the most sophisticated form of identity simulation that Go's value-oriented model permits — and it is still a simulation, because the method dispatch happens at runtime, not at compile time.

Each of these innovations addresses a real pain point caused by the absence of type-level identity for errors. Each is, in its own way, an attempt to rebuild some of the guarantees that type-level identity would have provided natively. The Memar framework's position is that these guarantees should be provided by the type system directly, not approximated through runtime mechanisms layered on top of a value-based model.

## The Memar–Go Design Target Comparison

The Go community's critiques of distinct error types are not wrong — they are correctly analyzing the trade-off from within Go's design frame. The disagreement is about the design frame itself, not about the analysis within it.

### What Go Optimizes For

- **Flexibility**: a function can return any error without committing to a specific type in its signature. This makes API evolution easier: adding a new error case does not require changing the return type.

- **Simplicity of the error model**: `error` is a single interface. There is no error type hierarchy to learn, no inheritance chain to navigate.

- **Ergonomics for common cases**: `errors.New("not found")` is three words. Defining a custom type requires a struct definition, an `Error() string` method, and usually a constructor.

- **Composability**: `fmt.Errorf("context: %w", err)` wraps any error without needing to know its type. Wrapping chains are transparent to `errors.Is` and `errors.As`.

These are genuine strengths, and they explain why the Go community has not converged on "one type per error concept" as an idiom despite more than a decade of discussion.

### What Memar Optimizes For

- **Type-level contract identity**: the answer to "which concept is this?" is known at compile time, verified by the type checker, and expressible in method signatures.

- **Architectural clarity**: identity, behavior, and data have distinct, non-overlapping homes. No mechanism simulates another's job.

- **Framework-as-authority**: the framework (RFC 000001) decides what the type system is responsible for, and RFC 495421 decides that identity is one of those responsibilities. The framework's goals take precedence over language-community conventions.

- **Cross-language fidelity**: when capsule types are compiled to other languages, identity maps 1:1 through the type system. No string-based dispatch, no runtime chain-walking.

### The Verdict

These are different optimization targets. Neither is objectively "correct" in isolation — each is correct for its framework's goals. Go's value-oriented error model serves Go's goals well. Memar's type-oriented model serves Memar's goals well. The disagreement is architectural, not technical: it is about what the type system should be responsible for, and that question has no universal answer.

The Go ecosystem's own evolution, however, is instructive. The progressive introduction of custom error types, `errors.Is/As`, the `errname` lint rule, and interface-based predicates reveals a pattern: the Go community keeps building type-level identity mechanisms around errors without explicitly acknowledging that this is what it is doing. Each innovation addresses a real pain point caused by the absence of type-level identity. Each is, in its own way, an attempt to approximate the guarantees that type-level identity would have provided natively. Memar's position is that these guarantees should be provided by the type system directly, not approximated through runtime mechanisms.

## How memar-go Implements the Principle

In the current Go implementation (`error.go`), the `Error` interface embeds `datatype_p.DataType`, `mediatype_p.Field_MediaType`, and `adt_p.ADT`, plus the `ImplementsError` marker and `Field_Error` accessor. Each concrete error capsule type implements the full method set:

```go
func (datatype_p.Quiddity) Abbreviation() string
func (datatype_p.Quiddity) Aliases() []string
func (datatype_p.Field_ID) DataTypeID() datatype_p.ID
func (datatype_p.Field_ID_Base64) DataTypeID_Base64() string
func (datatype_p.Detail) DevActionNote() string
func (datatype_p.Detail) Domain() string
func (datatype_p.Details) ExpireInFavorOf() datatype_p.DataType
func (datatype_p.Details) ExpiryDate() string
func (adt_p.Nil) IsNil() bool
func (datatype_p.Details) IssueDate() string
func (datatype_p.Field_LifeCycle) LifeCycle() datatype_p.LifeCycle
func (mediatype_p.Field_MediaType) MediaType() mediatype_p.MediaType
func (datatype_p.Quiddity) Name() string
func (datatype_p.Detail) Overview() string
func (datatype_p.Details) ReferenceURI() string
func (datatype_p.Detail) Summary() string
func (datatype_p.Detail) TAGS() []string
func (datatype_p.Detail) UserActionNote() string
```

For a static Error capsule, every one of these methods returns a compile-time constant. Two variables of the same error type are always indistinguishable — there is no per-instance state to compare.

### What You Lose Without the Principle (Go Code)

Consider a system that uses a single `GenericError` type with an `Init` method. A developer writes:

```go
func (s Service) Find(id ID) (Result, error) {
    // ...
    return Result{}, GenericError{}.Init("ErrNotFound", "memar/storage", "...")
}
```

The caller receives an `error` interface. To determine which error occurred, it must:
1. Type-assert to `GenericError` (or use `errors.As`).
2. Inspect the `Name()` field's return value.
3. Compare that string to the expected value.

This is a three-step runtime process to answer a question that the type system could have answered in one step: "is this a `ErrNotFound`?" Furthermore, the string comparison is fragile — a typo in either the producer or the consumer silently breaks the check. The compiler cannot help because it does not know that `"ErrNotFound"` is a concept identity — it sees a string like any other.

Now consider the same system under RFC 495421's principle:

```go
func (s Service) Find(id ID) (Result, ErrNotFound) {
    // ...
    return Result{}, ErrNotFound{}
}
```

The caller receives a value whose type *is* the concept. No string comparison, no type assertion chain, no runtime inspection. The type system carries the full identity. The compiler verifies that the caller handles the specific error type declared in the signature. A typo in the error type name is a compile-time error, not a silent runtime mismatch.

### Multi-Cause Returns in Go

When a method can fail in multiple ways, the return type is the `Error` interface:

```go
func (s Service) Find(id ID) (Result, error_p.Error)
```

The caller uses a type assertion to determine the specific error. A critical distinction must be drawn between this approach and the three-step process above: under RFC 495421's principle, there is no generic container and no string field. The caller performs a *single type assertion* against the specific error type. It does not need to dereference the Error abstraction's vtable or inspect any method return values to determine identity — the type itself is the identity. The Error interface's rich method set (`Name()`, `Domain()`, `Summary()`, etc.) is available for logging, display, and diagnostic purposes, but it is never needed to answer "which error occurred?"

This resolves *identification* cost and mechanism, but not *exhaustiveness*: `Service.Find`'s signature does not tell the compiler (or `go vet`, or `golangci-lint`) the closed set of concrete error types it can actually return, so nothing today flags a type-assertion chain that forgot one. Go has no sealed-interface construct to close that gap natively — unlike Kotlin's `sealed class` or Swift's closed protocol hierarchies. This is the same open need documented in RFC 495421 (Drawback 4: No syntax-level exhaustiveness) and deferred to the same future linter/compiler tooling discussion. It is not a Go-specific problem, though Go's lack of a sealed-type construct means the gap cannot be closed experimentally at the language level the way it might be in a language that already has one.

## Community Critiques and Responses

### Dave Cheney on Public Error Types

Dave Cheney, one of the most respected voices on Go error handling, has cautioned against exposing concrete error types in public APIs:

> "Public error types increase surface area, new implementations must only return types specified… error type cannot be changed or deprecated without breaking compatibility."

This is a valid concern. Cheney's advice — that clients should depend on behavior (interfaces, `Is`, `As`) rather than concrete types — is sound *within Go's value-oriented error model*, where the type system does not carry identity. In Memar's model, where the type system *does* carry identity, depending on the concrete type is not coupling to an implementation detail — it is depending on the concept itself. The type *is* the concept. Cheney's concern about "error type cannot be changed or deprecated" applies with full force: removing a static concept type is a breaking change because it removes a concept from the system's identity model. This is accepted as a deliberate consequence of the principle.

Cheney has also critiqued sentinel errors:

> "[Sentinel error] should be immutable… sentinel error values in Go… are not constants."

This observation — that Go's sentinel errors are variables, not constants, and can therefore be reassigned — is another example of the identity-as-data model's weakness. When identity is a variable, it can be accidentally or maliciously changed. When identity is a type, it cannot.

### Matt Proud on Structured vs. Opaque Errors

Matt Proud (2024) provided a clear analysis of the trade-off between structured error types and opaque errors:

**In favor of types:**

> "Structured error values… communicate a specific condition **and granular state**… in a manner that can be programmatically inspected (e.g., field access or a method) on the error type."

**Against opaque errors:**

> "Opaque errors… have no public identifier… no way to discriminate one from another — save from examining their string values, which is unsafe."

Proud's analysis aligns with RFC 495421's position: when identity matters (when the caller needs to discriminate between concepts), opaque errors are unsafe and structured types are necessary. The RFC extends Proud's observation from "when discrimination is needed" to "always, for static concepts," on the grounds that if a concept has no per-instance data, there is no reason *not* to make it a type.

### DoltHub Performance Analysis

DoltHub (Musgrave, 2024) published benchmark results showing the performance cost of Go's runtime error-inspection mechanisms:

> "[`errors.Is`] strategy has bad performance, at 5.5x slower than the `Bool` strategy when the value is not found, and 6x slower when it is."

These numbers demonstrate that the identity-as-data model carries a measurable runtime cost. In most I/O-bound contexts, this cost is negligible — but in tight inner loops or performance-critical paths, it is real. Type-level identity (a single type assertion) is faster because it does not require chain-walking. This performance difference is a secondary benefit of type-level identity, not a primary argument for it — but it is worth documenting as a concrete consequence.

### CedarDB Boilerplate Critique

CedarDB, a database system written in Rust, has publicly critiqued the boilerplate cost of defining many error types. Their argument is that for systems with large numbers of error conditions (hundreds in a database's case), the mechanical overhead of defining a distinct type for each — even with Rust's relatively concise enum syntax — becomes a real maintenance burden.

This critique is acknowledged. In Memar, the cost is mitigated by the code generator: error capsule types are produced mechanically from a YAML/DSL definition, and the developer authoring a new error concept writes a declarative entry, not a full type definition. The boilerplate exists in the generated output, not in the developer's input. This does not eliminate the output verbosity, but it eliminates the human effort of producing it.

For the rare case where a developer must define a static capsule by hand (e.g., before the code generator exists for a given target language), the verbosity is substantial (~20 method implementations for an Error capsule in Go). This is accepted as a deliberate signal: the framework expects static capsules to be generated, and the manual path's weight reinforces that expectation. CedarDB's critique applies to manually-authored types in a system without code generation; it does not apply — or applies with greatly reduced force — to a system where the generator absorbs the mechanical work.

### yegor256 and Checked Exceptions Advocacy

yegor256 (Yegor Bugayenko), a prominent critic of Go's error handling, has advocated for checked exceptions as a mechanism for enforcing error contracts at the call site. While Memar does not use checked exceptions (it uses explicit error returns), yegor256's underlying argument — that the compiler should enforce error handling contracts — is aligned with RFC 495421's position. The disagreement is about the mechanism (checked exceptions vs. explicit return types) rather than the principle (the compiler should know about error contracts).

yegor256's advocacy is relevant here because it represents a strand of opinion within the broader programming community that agrees with Memar's premise: that error handling benefits from compiler-enforced contracts, and that the Go ecosystem's value-oriented model leaves too much to runtime convention. RFC 495421 does not adopt yegor256's specific mechanism (checked exceptions), but it shares the conviction that the type system should carry error identity.

## Go-Specific Prior Art

### Go Sentinel Errors and `errors.Is` / `errors.As`

Go uses `errors.New("not found")` — string-based values inside a shared `error` interface. This is the pattern RFC 495421 explicitly rejects for static capsules: identity as data, not as type.

The Go 1.13 design document (Newton et al.) explicitly lists multiple acceptable patterns for signaling identifiable errors:

> "There are other existing patterns… such as directly returning a sentinel, a specific type, or a value… with a predicate."

The document further clarifies the relationship between the two mechanisms:

> "`errors.Is` behaves like comparison to a sentinel error, and `errors.As` behaves like a type assertion."

This is a fair description of the status quo: Go provides *both* value-based and type-based identity mechanisms, treats them as equally valid, and lets the developer choose. Memar's RFC removes that choice for static capsules: the type-based mechanism is mandatory. This is a constraint, not a freedom — but it is a constraint that the framework is entitled to impose (RFC 000001) and that serves its architectural goals.

**Rob Pike's framing** of the Go position is the clearest statement of the underlying philosophy:

> "Errors are values… values can be programmed. Since errors are values, errors can be programmed."

This is a statement about *what errors are* in Go: they are values, and values are programmable. The flexibility this provides is real and deliberate. Memar's position is that for *static concepts* — concepts with no per-instance variation — this flexibility is unnecessary and the cost (loss of compile-time identity) is unacceptable.

**Russ Cox's 2024 assessment** of Go's error handling trajectory acknowledges the verbosity but defends the choice:

> "If Go had introduced specific syntactic sugar for error handling early on, few would argue… we are 15 years down the road… Go has a perfectly fine way to handle errors, even if it may seem verbose."

This is a pragmatic defense of Go's value-oriented model: it works, it is stable, and the cost of changing it now outweighs the benefit. Memar, as a new framework designing from first principles rather than evolving an existing ecosystem, does not face this constraint.

### Go Naming Conventions and Tooling

The `golangci-lint` `errname` rule enforces a naming convention that implicitly distinguishes identity-as-data from identity-as-type:

- Sentinel errors (identity-as-data): prefixed with `Err` — `ErrNotFound`, `ErrPermissionDenied`.
- Error types (identity-as-type): suffixed with `Error` — `NotFoundError`, `PermissionDeniedError`.

The `errorlint` rule flags code that can cause problems with Go 1.13's error wrapping scheme. Together, these lint rules encode community knowledge that identity-as-data and identity-as-type are *different things that require different treatment* — a distinction that the language itself does not enforce but the community has found necessary enough to automate.

Memar's position is that this distinction should be enforced by the type system, not by a lint rule that can be disabled. If identity belongs to the type system, then identity-as-data for static concepts is not a style violation to be caught by a linter — it is an architectural violation that should be structurally impossible.

## Go-Specific Drawbacks of Following RFC 495421

These drawbacks apply specifically to the Go implementation of the principle. They are not drawbacks of the principle itself (which is universal) but of applying it in a language whose community conventions diverge from it.

1. **Not idiomatic Go.** This RFC's principle is not how the Go community writes error-handling code. Go's established idiom uses sentinel error variables (`var ErrNotFound = errors.New("not found")`), wrapped errors inspected via `errors.Is`, and occasional custom types when structured data is needed. Requiring a distinct type for every static concept — including concepts that would be sentinel variables in idiomatic Go — is a deliberate departure from community norms. Developers working in memar-go must understand that they are writing Memar-framework code that happens to be compiled by the Go compiler, not idiomatic Go code. The framework's architectural authority (RFC 000001) gives it the right to make this choice, but the friction is real and acknowledged.

2. **Go tooling friction.** Go tooling (IDEs, debuggers, error formatters) is designed around the `error` interface. Error types that carry Memar's full capsule method set may not display cleanly in Go debugging tools, stack trace formatters, or error reporting libraries. This is a practical friction point for developers using standard Go tooling alongside memar-go code.

## Rejected Go-Specific Alternatives

### Rejected: Go-Style Sentinel Errors with `errors.Is`

Using `errors.New("not found")` (or `var ErrNotFound = errors.New(...)`) as sentinel values, distinguished at runtime via `errors.Is(err, ErrNotFound)`.

**Rejected because:** identity is carried by a variable's value (a pointer to an internal struct containing a string), not by the type system. The type of both `ErrNotFound` and `ErrPermissionDenied` is the same: `*errors.errorString` (or whatever internal type `errors.New` returns). The type system cannot distinguish them. `errors.Is` provides structured comparison, but it is runtime chain-walking, not compile-time type checking. As DoltHub's benchmark analysis demonstrated, this chain-walking carries a performance cost: `errors.Is` is approximately 5.5x slower than a boolean return when the value is not found, and 6x slower when it is found. This performance cost is a secondary concern — the primary rejection reason remains the identity demotion.

This rejection is specific to Memar's architectural goals. Within the Go ecosystem, sentinel errors with `errors.Is` are a well-designed solution for Go's optimization targets. The Go 1.13 design doc (Newton et al.) explicitly lists sentinel variables as one of the accepted patterns for identifiable errors, alongside custom types and predicate functions. The Go team's position is that flexibility in error representation is more valuable than type-level identity. Memar's position is the reverse. These are different optimization targets, both valid for their respective frameworks.

### Rejected: Interface-Based Predicates

Defining custom `Is(target error) bool` methods on error types, allowing each error to define its own matching logic.

**Rejected because:** while more sophisticated than sentinel comparison, this is still runtime behavior. The method dispatch happens at runtime; the type system does not know which `Is` method will match until the code executes. This approach can express rich identity logic (e.g., "this error matches if it carries a specific code *and* a specific domain"), but it cannot express identity at compile time. For static capsules — where identity is the entire point and there is no instance data to factor into the comparison — the added runtime complexity is unnecessary overhead for a weaker guarantee than the type system provides natively.

## Citation Verification Status

The following citations were verified against primary sources during the research phase. The remaining citations (Matt Proud, Russ Cox, Go 1.13 design doc) were not independently re-verified in this session and should be confirmed before final commit.

- ✅ Rob Pike: *"Errors are values… values can be programmed. Since errors are values, errors can be programmed."* — Verified.
- ✅ Dave Cheney: *"Public error types increase surface area, new implementations must only return types specified… error type cannot be changed or deprecated without breaking compatibility."* — Verified.
- ✅ Dave Cheney: *"[Sentinel error] should be immutable… sentinel error values in Go… are not constants."* — Verified.
- ✅ DoltHub (Musgrave, 2024): *"[errors.Is] strategy has bad performance, at 5.5x slower than the Bool strategy when the value is not found, and 6x slower when it is."* — Verified.
- ⚠️ Matt Proud (2024): Structured vs. opaque error analysis — Not independently re-verified.
- ⚠️ Russ Cox (2024): Error handling trajectory assessment — Not independently re-verified.
- ⚠️ Go 1.13 design doc (Newton et al.): Multi-pattern guidance — Not independently re-verified.

---

## Change Rationale

| Revision | Section | Change | Reason |
|----------|---------|--------|--------|
| 2 (this) | Contributor(s) | Added missing Claude entry | Change Rationale already credited "Claude's suggestion" in revision 1, but the front-matter had no corresponding entry |
| 2 | "Multi-Cause Returns in Go" | Added paragraph distinguishing identification cost (resolved) from signature-level exhaustiveness (still open, deferred to tooling per RFC 495421) | Cross-reference RFC 495421's revision-7 clarification so the two documents stay consistent |
| 1 (this) | Entire document | Created as companion to RFC 495421; extracted all Go-specific content from the principle RFC | User direction: split into two files, Go analysis belongs in memar-go |
| 1 | — | Added citation verification status table | Claude's suggestion: verify quotes before committing to permanent RFC |
