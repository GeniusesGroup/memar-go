---
RFC Number: 495368
Title: "Elimination of Open Generic Type Parameters in memar-go"
Status: Proposed
Start Date: 2026-07-06
Applied to: []
Supersedes: null
Superseded by: null
Related:
  Depends_on: []
  Extends: ["khayyam-containers_without_generics.md"]
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
    contribution: "Drafted initial text from an extended design-review conversation, pressure-tested the reasoning, incorporated revisions"
    task: []
---

## Summary
Open generic type parameters (parameters whose concrete binding is deferred to the caller, e.g. `Encoder[BUF any]`) are removed from memar-go wherever they would need to propagate beyond the package that defines them, and replaced with plain, framework-owned interface types (e.g. `buffer_p.Buffer`). This is a Go-specific implementation correction, not a change to Memar's architectural model: the underlying abstraction (a buffer-agnostic String/Encoder contract) remains valid and is expected to be expressed without this cost in Khayyam, whose contract-based abstraction model does not carry Go's generic-instantiation mechanics.

## Motivation
Several memar-go packages (buffer-backed `String`, JSON `Encoder`, and the `Socket` buffer accessors) were originally written with an open generic parameter for the underlying buffer type (`BUF`), intending to let each call site choose its own buffer implementation without loss of type information. In practice, this benefit never reaches its intended destination:

1. **The path from buffer origin to business logic already collapses to a plain interface at its source.** `Field_Socket_Buffers.SendBuffer()/ReceiveBuffer()` already return `buffer_p.Buffer` as an ordinary interface, not a generic parameter — meaning by the time any encoder or string type could receive a buffer from a socket, the concrete type information the generic parameter was meant to preserve is already erased. Keeping `BUF` generic downstream of this point adds cost for a distinction that has already ceased to exist upstream.
2. **Propagating the parameter through business-facing layers is unacceptable.** Because a socket is a required argument of protocol handler methods (e.g. `Handler.ServeHTTP(sk net_p.Socket, ...)`), making `Socket` generic over its buffer type would force every implementer of that handler — including all business services — to carry the same open type parameter merely to satisfy the compiler, not because business logic has any use for buffer-level detail. Since business-code readability is treated as a higher priority than eliminating this specific, boundable dispatch cost, this path was rejected.
3. **Collapsing the parameter early (type-erasing at the socket boundary, keeping generics only above it) defeats the purpose of having introduced it.** No call site was ever found, across buffer, string, encoder, or socket packages, where a caller actually needs to supply a distinct concrete `BUF` and have it preserved past that call — buffers originate from exactly two places (network sockets, storage), and callers at higher layers have no access to, or use for, the concrete buffer type in either case.

This was not a deliberate, upfront architectural trade-off; it emerged as an implementation path taken without fully tracing where the open parameter would eventually have to surface. This RFC documents the correction and the reasoning, since the original generic-parameterized commits remain visible in version-control history and could otherwise be mistaken for an intentional, still-endorsed design.

## Guide-level explanation
Where a memar-go type previously declared an open type parameter for its underlying storage (e.g. `type Encoder[BUF buffer_p.Buffer] struct { buf BUF }`), it now declares that dependency as a plain field of the framework-owned interface type directly (e.g. `type Encoder struct { buf buffer_p.Buffer }`). Call sites lose no capability they were actually using: `buffer_p.Buffer` was already the only type ever substituted for `BUF` in practice, at every observed call site. Developers writing business services that implement protocol handlers no longer need to reason about, or propagate, any buffer-related type parameter — `net_p.Socket` and everything built on it are ordinary, non-generic interfaces.

## Reference-level explanation
- **Distinguishing closed vs. open parameters:** see the addition proposed to `khayyam-containers_without_generics.md` for the general criterion. This RFC documents its concrete failure mode in Go specifically.
- **Case 1 — `Encoder[BUF buffer_p.Buffer]` (JSON codec):** every method on `Encoder` operates on `buf BUF` purely as an implementer of `buffer_p.Buffer`'s contract (`Push`, `Pop`, `Peek`, `Append`); no method requires knowledge of any concrete buffer type beyond that contract. No call site was found instantiating `Encoder` with anything other than a `buffer_p.Buffer`-satisfying type used generically-in-name-only. Corrected form: `buf buffer_p.Buffer` as a plain field.
- **Case 2 — `STR[BUF buffer_p.Buffer]` (ASCII string builder):** same pattern; `Buffer()` accessor, `Append`/`Prepend`, and comparison methods only ever need the `buffer_p.Buffer` contract. No path exists, across any known or hypothesized memar-go call site, for a concrete, non-erased `BUF` to reach this type directly from either a network socket or a storage backend without first passing through a plain-interface boundary (`Field_Socket_Buffers`, or an equivalent storage-read API). Corrected form: same collapse to `buffer_p.Buffer`.
- **Case 3 — `Socket` / `Field_Socket_Buffers` (network layer):** already correctly implemented as a plain-interface boundary (`SendBuffer() buffer_p.Buffer`, `ReceiveBuffer() buffer_p.Buffer`), predating this RFC. Cited here as existing internal prior art supporting the correction applied to Cases 1 and 2, and as the boundary past which no caller has ever needed — or could obtain — a non-erased buffer type.
- **Scope of this correction:** applies to any memar-go type whose sole use of an open type parameter is a storage/buffer dependency that already terminates, at its origin, in a plain `buffer_p.Buffer` interface. It does **not** apply to closed type parameters bound by the defining package itself to a fixed framework type (e.g. `marker_p.Clone[String]`), which carry no propagation cost and are unaffected by this RFC.

## Drawbacks
- Static dispatch inlining opportunities that Go's generic monomorphization would otherwise provide are given up in favor of interface (dynamic) dispatch at every buffer method call. This is an accepted, deliberate cost given that no call site was found actually depending on the eliminated distinction, and given that business-code readability is prioritized over this specific, boundable runtime cost.
- This correction is scoped to the buffer/string/encoder/socket family described here; other open generic type parameters may exist elsewhere in memar-go that exhibit the same failure mode and have not yet been audited (see Future possibilities).
- Reproduces, in a Go-specific and empirically-motivated way, a form of type erasure that Java adopted for generics from the start — meaning memar-go, for this concern, converges toward Java's trade-off rather than Rust's or C++'s zero-cost-generic default. This is stated plainly rather than minimized.

## Rationale and alternatives
- **Keep `BUF` generic, alias it to a concrete type at a single composition point (considered, rejected):** resolves the runtime-dispatch cost with zero propagation past that point, but only if no two call sites within the same running process ever need genuinely different concrete buffer implementations simultaneously. This was rejected because sockets in memar-go are expected to expose potentially different buffer implementations per connection within a single running process (e.g. a stream-oriented socket vs. a simple request/response socket), which a single process-wide alias cannot express; and because the point where this would need to be aliased is upstream of business code that must remain generic-free regardless.
- **Propagate `BUF` fully through `Socket` and `Handler` (rejected):** examined directly via `net_p.Socket.ServeHTTP(sk net_p.Socket, ...)` — would force every business-logic handler implementation to carry the same open parameter for no benefit to that business logic. Rejected on readability grounds, treated as a hard constraint rather than a cost to be weighed case-by-case.
- **Do nothing / keep generics as originally written (rejected):** the generic parameter provides no benefit that was ever actually exercised at any known call site, while adding real cost to reading and maintaining the affected packages' public APIs.

## Prior art
- Java's generics implementation uses type erasure at compile time by original design choice, not as a later correction — meaning the trade-off accepted here (favoring non-viral, erased types over zero-cost generic propagation) has long-standing precedent as a deliberate language-level decision elsewhere, not merely as a fallback taken under host-language duress.
- Rust and C++ each offer the same two-path choice this RFC navigates in Go: a zero-cost, viral generic path (Rust generics / C++ templates) versus a non-viral, runtime-dispatch path (`dyn Trait` / type-erased wrappers such as `std::function`). The existence of dynamic-dispatch escape hatches in both languages indicates this is a structural trade-off inherent to any statically-typed language offering parametric polymorphism, not a deficiency specific to Go — Go simply lacks Rust/C++'s zero-cost path being reachable here for the reasons given in Reference-level explanation.

## Unresolved questions
- Whether other open generic type parameters exist elsewhere in memar-go with the same failure mode (a parameter that never reaches a call site actually needing it preserved) is not yet audited; this RFC covers only the buffer/string/encoder/socket family examined directly.
- Whether dynamic dispatch is *always* reducible to a compile-time-resolved form given sufficient compiler sophistication (as opposed to being fundamentally required whenever the concrete type is determined by genuinely runtime-only information, e.g. configuration-driven or plugin-loaded types) is an open disagreement, deliberately deferred to a dedicated future discussion of the Khayyam compiler's dispatch-resolution strategy.
- The precise boundary of what belongs in memar-go's `libgo` subfolder (packages needed for compatibility but not considered part of Memar's core architecture) versus what should be corrected in place (as in this RFC) is not yet formalized as a general criterion.

## Future possibilities
- A dedicated audit pass across memar-go for any remaining open generic type parameters exhibiting this same failure mode, once time/resources allow — not undertaken as part of this RFC.
- Once Khayyam's contract-based abstraction model (`ab`) and compiler are further developed, this entire class of problem is expected to be re-expressed there without the open/closed distinction needing to be reasoned about manually at all, per the Khayyam generics-elimination model this RFC extends.
