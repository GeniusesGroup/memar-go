/* For license and copyright information please see the LEGAL file in the code repository */

package zerocopy

// It means buffer not reuse when need copy more data to it and it can't be resize!
// Use to implement buffer as `zero-copy` feature, But it can cause memory leak.
// Developers must be aware about this, But we introduce `Memar framework` to fix any problems
// by introduce new protocol to replace old protocols like TCP, HTTP, ...
//
// https://en.wikipedia.org/wiki/Zero-copy
