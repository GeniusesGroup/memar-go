# Data raca
A [race condition](https://en.wikipedia.org/wiki/Race_condition) or race hazard is the condition of an electronics, software, or other system where the system's substantive behavior is dependent on the sequence or timing of other uncontrollable events, leading to unexpected or inconsistent results. It becomes a bug when one or more of the possible behaviors is undesirable.

In information technology and computer science, especially in the fields of computer programming, operating systems, multiprocessors, and databases, [concurrency control](https://en.wikipedia.org/wiki/Concurrency_control) ensures that correct results for concurrent operations are generated, while getting those results as quickly as possible.

## Goals 
- Static checker as compile-time or pre-compile in build process or linters not runtime one.

## Other frameworks
[List of tools for static code analysis](https://en.wikipedia.org/wiki/List_of_tools_for_static_code_analysis):::
- Borrow checker >> Rust
- [ThreadSanitizer >> LLVM](https://clang.llvm.org/docs/ThreadSanitizer.html)
- RacerF >> C
- Coderrect Scanner >> C++
- RacerD >> JAVA
- [Chronos >> Go](https://github.com/amit-davidson/Chronos)
- [Go runtime >> Go](https://go.dev/doc/articles/race_detector)
- [Godel Checker](https://bitbucket.org/MobilityReadingGroup/godel-checker)
- [Go Model Checker for MiGo types](https://github.com/JujuYuki/godel2)
- [Go concurrency bugs](http://github.com/timmyyuan/gobench)
- 

## Key Citations:
- [Chronos - A static race detector for the go language - GitHub](https://github.com/amit-davidson/Chronos)
- [Data Race Detector - The Go Programming Language](https://go.dev/doc/articles/race_detector)
- [Static Race Detection and Mutex Safety and Liveness for Go Programs - arXiv](https://arxiv.org/abs/2004.12859)
- [STATIC RACE DETECTION TOOL FOR GOLANG - OAKTrust](https://oaktrust.library.tamu.edu/bitstream/handle/1969.1/194343/SANDERS-FINALTHESIS-2021.pdf?sequence=1&isAllowed=y)
- [Static Race Detection and Mutex Safety and Liveness for Go Programs - Dagstuhl](https://drops.dagstuhl.de/storage/00lipics/lipics-vol166-ecoop2020/LIPIcs.ECOOP.2020.4/LIPIcs.ECOOP.2020.4.pdf)
- [Hybrid Dynamic Data Race Detection - ACM](https://dl.acm.org/doi/10.1145/781498.781528)
- [Dynamic Data Race Detection in Go Code | Uber Blog](https://www.uber.com/en-US/blog/dynamic-data-race-detection-in-go-code/)
- [Concurrency in Go and Rust - Reddit](https://www.reddit.com/r/golang/comments/xwdvl5/concurrency_in_go_and_rust/)
- [Rust vs Go in 2025 - Bitfield Consulting](https://bitfieldconsulting.com/posts/rust-vs-go)
