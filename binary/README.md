# Binary

## Little Endian vs Big Endian
Go added library to compiler don't support [automatic endian detection](https://groups.google.com/g/golang-nuts/c/3GEzwKfRRQw), So we must have some mechanism to detect CPU architecture and do logics correctly. So by some [help](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) we can implement by compile time or runtime.

Runtime mechanism suggestion like [x/net](https://github.com/golang/net/tree/master/internal/socket/sys.go) is not efficient enough for very low level tasks, So we implement this compile time logic package.

## [Conversions - Cast](https://go.dev/ref/spec#Conversions)
Converting uint64 to int64 costs nothing: the bits don't change. Only the interpretation of them does. Just be careful about how large your uint64 is because it could end up becoming negative once it's converted to int64.

## Resources
- https://groups.google.com/g/golang-nuts/c/Y95NNox15Ns
