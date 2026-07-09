# Executive Summary
The Go community overwhelmingly endorses the “errors are values” philosophy: explicit `error` returns, values/interfaces, and flexible handling (wrapping, `errors.Is`/`As`) are seen as strengths. Mandating a unique Go `type` for every static error condition (the memar “Static Concept → Distinct Type” RFC) departs from this idiom. Proponents of distinct types note benefits like clear API identity and easy pattern-matching via type assertions. Critics point out significant trade-offs: *increased API surface and coupling* (making future changes hard), and extra boilerplate. No formal academic “proof” exists against the RFC, but many Go experts and proposals emphasize flexibility. The Go 1.13 design doc explicitly lists multiple patterns for signaling identifiable errors — including sentinel variables, custom types, **or** predicates — and advises caution about exposing implementation details. Tools like `golangci-lint` even enforce naming conventions to distinguish sentinels (prefix `Err`) from struct-based errors (suffix `Error`). Performance analyses reveal further nuance: simple boolean or direct comparisons beat wrapped/sentinel checks by 2–6×. In practice, well-known libraries (`pkg/errors`, `xerrors`) and Go 1.13+ wrapping support already allow custom types or sentinels when needed, guided by context. 

**Verdict:** Go’s choice of value-oriented errors is by design, not an oversight. Enforcing “one type per error concept” is *not* idiomatic Go and would meet resistance. It **may improve identity clarity** (good for pattern matching and domain modeling) but at the cost of verbosity, brittleness, and potential performance penalties. We find no community consensus demanding such a rule; on the contrary, thought leaders warn against tightly coupling callers to concrete error types. An RFC should therefore **acknowledge these trade-offs**. We recommend framing the distinct-type requirement as an *option* (when type-specific behavior is truly needed), not an absolute MUST. The proposed RFC text should cite existing guidance (Go 1.13 docs, Cheney’s posts, etc.) and respond to common objections by showing how `errors.Is/As` or new syntax proposals fit into the picture.

```
flowchart LR
    subgraph A["Type-based (distinct error types)"]
        F1[Call function → (T, error)] --> C1{err != nil?}
        C1 -->|no| V1[Use result]
        C1 -->|yes| T1{error type?}
        T1 -->|*MyErrorType| H1[Handle MyErrorType]
        T1 -->|*OtherErrorType| H2[Handle OtherErrorType]
        T1 -->|else| H3[Fallback/return]
    end
    subgraph B["Generic (interface/sentinel)"]
        F2[Call function → (T, error)] --> C2{err != nil?}
        C2 -->|no| V2[Use result]
        C2 -->|yes| I1[Inspect error (errors.Is/As)]
        I1 -->|Condition1| H4[Handle condition1]
        I1 -->|Condition2| H5[Handle condition2]
        I1 -->|else| H6[Fallback/return]
    end
```
*Mermaid flowchart: comparison of error-handling flows under the distinct-type vs generic interface approach.* 

## Pros/Cons: Distinct Type vs Generic Error

| **Dimension**           | **Distinct Type (Type-based)**                                                                                                                                                              | **Generic Error (interface/one struct)**                                                                                                                                                                                                       |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Identity Clarity**    | **High:** Each error kind has its own Go type/name (e.g. `type NotFoundError struct`). Callers can import and compare types or use `errors.As`. The meaning of each error is explicit.      | **Low/Medium:** Errors often distinguished by value or predicate (e.g. sentinel `var ErrNotFound` or error fields). The identity is implicit and may require checking message or interface methods. Opaque errors have *no* public identifier. |
| **Pattern Matching**    | **Direct:** Can use a `type switch` or `errors.As(err, *YourErrorType)` easily. Idiomatic and compile-time safe (unwrapping supported).                                                     | **Indirect:** Must rely on `errors.Is(err, sentinel)` or custom `Is/As` methods, or examine fields on a generic struct. More boilerplate to detect a condition; mistakes (string comparisons) are unsafe.                                      |
| **Performance**         | **Generally Good:** A direct type assertion (`err.(*MyError)`) is fast. Fewer dynamic checks. However, large error type hierarchies add minimal overhead.                                   | **Varies:** Direct equality on a sentinel (if returned unwrapped) is fast, but *any* use of `errors.Is/As` on wrapped errors can be ~5× slower than a boolean result. Complex generic checks (e.g. custom `Is` method) add overhead.           |
| **Ergonomics**          | **Verbosity:** More code needed: one struct type per concept + `Error()`. New types/values to maintain. Easy to forget to wrap vs not.                                                      | **Concise:** Can reuse one struct or shared sentinel with context. Fewer declarations if error kinds are few. But may lead to ambiguous APIs unless documented.                                                                                |
| **Tooling Support**     | **Strong:** Tools expect `Error` suffix on types (golangci-lint’s `errname`). IDEs can autocomplete types. `errors.As` and Go 1.13+ (or 1.20+) unwrapping tools fully support custom types. | **Evolving:** Tools like `errorlint` flag bad wrapping or unchecked errors. Static analysis can enforce sentinel naming. But no tool fully automates pattern-matching by value.                                                                |
| **Codegen/Maintenance** | **Fragmented:** Each new concept adds a file/type. Good modularity but more code. Refactoring error types requires API versioning.                                                          | **Centralized:** A single error type (or few) can cover many cases, easier to extend with new fields. But risk of one god-type with lots of fields, making maintenance hard if conditions multiply.                                            |

## Common Objections & Counterarguments

**“Error types bloat APIs.”** Recall Dave Cheney’s caution: *“Public error types increase surface area, new implementations must only return types specified… error type cannot be changed or deprecated…”*. The RFC should acknowledge that enforcing distinct types **increases coupling**. Counter: Go 1.13 docs explicitly allow *either* returning a sentinel, a type, or a predicate for identifiable errors. Emphasize that “distinct type” is best when callers need to inspect *structured* details (see PathError example), but unnecessary if a simple flag or wrapped error suffices.

**“We want pattern matching like in Rust/Swift.”** While Rust’s `Result<T, E>` uses enums, Go’s design avoids sum types at compile time. The RFC can note cross-language context: Rust/Swift have explicit syntax (`Result`, `try/catch`) to reduce boilerplate. Go has chosen not to add syntax (decades of proposals ultimately declined). The Go team suggests solving verbosity via helper functions (e.g. `cmp.Or` or `cmp.Join` in Go 1.22+) rather than new type systems. Reassure readers that the status quo is a conscious trade-off: explicit error handling yields transparency, not a mistake to be forced into strict typing.

**“What about performance?”** Benchmarks indicate that error handling style matters. Sentinel checks or boolean flags are fastest; any wrapping (`errors.Is`) can multiply cost. The RFC should mention that massive type switches or `errors.As` calls add overhead, but in most I/O/error contexts this cost is negligible. If performance is critical, callers might prefer simpler signals (e.g. boolean returns or tiny structs), a choice orthogonal to type identity.

**“No one stops you from defining error types today.”** The RFC should clarify it’s an *advisory* norm, not a language requirement. It can recommend best practices (e.g. name suffixes per [56†L264-L268]) and point to examples (the Go1.13 guide, common libraries). For instance, the Go blog shows custom errors wrapping sentinel values: e.g., `fmt.Errorf("…: %w", ErrNotFound)`, or defining `type ErrListError struct{…}` when context is needed. Affirm that Go 1.13+ idiom supports both sentinel variables and concrete types, so teams can choose the right abstraction.

**Recommended RFC edits:** Where the draft asserts *“MUST: use distinct types for static concepts”*, we suggest softer language: e.g. “**RECOMMENDED**: for well-defined error conditions that callers must detect, define a distinct error type implementing `error`”. Add a note that Go’s official guidance already uses this pattern sparingly. Include a reference to Cheney’s advice about behavior vs type: “clients should depend on behavior (e.g. interface, `Is`, `As`) rather than concrete types, to avoid coupling”. Mention that `errors.Is` and custom `Is(target error)` methods are available to precisely match error conditions without comparison of error strings. Finally, explicitly address mismatches: if a type-based scheme is used, one can still wrap errors with `%w` to preserve inspectability, or leave them opaque (while noting risks).

## Suggested Citations & Quotes

- Pike: *“Errors are values… values can be programmed. Since errors are values, errors can be programmed.”*. (Emphasizes flexibility of value-based errors.)  
- Cheney: *“Public error types increase surface area… error type cannot be changed or deprecated without breaking compatibility.”*. (Warns against exposing concrete types.)  
- Cheney: *“[Sentinel error] should be immutable… sentinel error values in Go… are not constants.”*. (Critiques naive sentinel use.)  
- Go 1.13 (Newton et al.): *“There are other existing patterns… such as directly returning a sentinel, a specific type, or a value… with a predicate.”*. (Lists accepted idioms.)  
- Go 1.13 (Newton et al.): *“errors.Is behaves like comparison to a sentinel error, and errors.As behaves like a type assertion.”*. (Clarifies how to detect errors.)  
- Go 1.13: *“Whether you wrap or not… wrapping an error makes that error part of your API. If you don’t want to commit to supporting that error… you shouldn’t wrap the error.”*. (Advice on exposing error types.)  
- Russ Cox (Go blog 2024): *“If Go had introduced specific syntactic sugar for error handling early on, few would argue… we are 15 years down the road… Go has a perfectly fine way to handle errors, even if it may seem verbose.”*. (Context on decision to stick with values.)  
- Matt Proud (2024): *“Structured error values… communicate a specific condition **and granular state**… in a manner that can be programmatically inspected (e.g., field access or a method) on the error type.”*. (Pro-type: captures details.)  
- Matt Proud: *“Opaque errors… have no public identifier… no way to discriminate one from another – save from examining their string values, which is unsafe.”*. (Con-generic: opaque strings are fragile.)  
- DoltHub (Musgrave, 2024): *“[errors.Is] strategy has bad performance, at 5.5x slower than the `Bool` strategy when the value is not found, and 6x slower when it is.”*. (Shows cost of generic/sentinel checks.)  
- golangci-lint docs: *“errname – Checks that sentinel errors are prefixed with Err and error types are suffixed with Error… errorlint – Find code that can cause problems with the error wrapping scheme introduced in Go 1.13.”*. (Indicates idioms for error naming.)

These sources provide authoritative viewpoints and concrete advice to integrate into the RFC’s rationale and counter common objections.