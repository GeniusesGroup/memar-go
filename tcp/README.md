# TCP on UserSpace
In services as API era, almost all services call transfer under two TCP packets, So many logics that consume both memory (e.g. Go just use [2KB for each goroutine](https://go.dev/doc/go1.4#runtime)) and computing resources (e.g. many [context switch](https://en.wikipedia.org/wiki/Context_switch) between kernel and userspace) to handle these few packets are not acceptable anymore.

Some suggestion such as using [epool](https://github.com/eranyanay/1m-go-websockets) has many other problems 
- OS depend on implementation can very tricky tasks
- OS depend optimization need to change such as the number of the active and open file in UNIX based OSs known as ulimit
- Balance between the number of events and timeout(milliseconds) in high and low app load isn't easy.
- Need some runtime adaptors, because other packages not being ready for this architecture even the Go net library.

## Why (OS Kernel level disadvantages)
- The Linux or other OSs networking stack has a limit on how many packets per second they can handle. When the limit is reached all CPUs become busy just receiving and routing packets.
- 

## Goals
- Improve performance by reducing resource usage. e.g.
    - No context switch need (L3 as IP, ... need context switch but not as much as kernel-based logics)
    - No need for separate files for each TCP stream (L3 as IP, ... need some mechanism to take packets from OS)
    - Just have one buffer and no huge memory copy need more.
    - Just have one lock mechanism for each stream
    - Just have one timeout mechanism for a stream of any connections, not many in kernel and user-space to handle the same requirements.
    - Mix congestion control with rate limiting
    - Keep-alive a stream for almost free. Just store some bytes in RAM for a stream without impacting other parts of the application
- Track connections and streams metrics for any purpose like security, ...
- Easily add or changed logic whereas upgrading the host kernel is quite challenging. e.g. add machine learning algorithms, ...
- Have protocol implementation in user space to build applications as unikernel image without need huge os kernel.

## Non-Goals (Non Considering that can be treated as disadvantages)
- Don't want to know how TCP packets come from. So we don't consider or think about how other layers work.

## Still considering
- Support some protocols like [PLPMTUD - Packetization Layer Path MTU Discovery](https://www.ietf.org/rfc/rfc4821.txt) for bad networks that don't serve L3 IP/ICMP services?
- Why [tcp checksum computation](https://en.wikipedia.org/wiki/Transmission_Control_Protocol#Checksum_computation) must change depending on the below layer!!??

## Little Endian vs Big Endian
Go don't support [automatic endian detection](https://groups.google.com/g/golang-nuts/c/3GEzwKfRRQw), So we must have some mechanism to detect CPU architecture and do logics correctly. So by some [help](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) we must indicate 

## Similar Projects
- https://github.com/Xilinx-CNS/onload
- https://github.com/mtcp-stack/mtcp
- https://github.com/golang/go/issues/15735
- https://github.com/xtaci/gaio
- https://github.com/lesismal/nbio
- https://github.com/saminiir/level-ip
- https://github.com/google/gopacket/blob/master/layers/tcp.go

## Resources
- https://en.wikipedia.org/wiki/Transmission_Control_Protocol
- https://github.com/torvalds/linux/blob/master/net/ipv4/tcp.c
- https://github.com/torvalds/linux/blob/master/net/ipv6/tcp_ipv6.c

## Articles
- https://ieeexplore.ieee.org/document/8672289
- https://engineering.salesforce.com/performance-analysis-of-linux-kernel-library-user-space-tcp-stack-be75fb198730
- https://tempesta-tech.com/blog/user-space-tcp
- https://blog.cloudflare.com/kernel-bypass/
- https://blog.cloudflare.com/why-we-use-the-linux-kernels-tcp-stack/
- https://blog.cloudflare.com/path-mtu-discovery-in-practice/
- https://www.fastly.com/blog/measuring-quic-vs-tcp-computational-efficiency
- https://stackoverflow.com/questions/8509152/max-number-of-goroutines

# Abbreviations
- L3    >> layer 3 OSI
- IP    >> Internet Protocol
