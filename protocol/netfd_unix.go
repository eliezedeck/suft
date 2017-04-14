// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package suft

import (
	"net"
	"syscall"
)

// Network file descriptor.
type netFD struct {
	// locking/lifetime of sysfd + serialize access to Read and Write methods
	fdmu fdMutex

	// immutable until Close
	sysfd       int
	family      int
	sotype      int
	isStream    bool
	isConnected bool
	net         string
	laddr       net.Addr
	raddr       net.Addr

	// writev cache.
	iovecs *[]syscall.Iovec

	// wait server
	pd pollDesc
}
