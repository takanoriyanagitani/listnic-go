package listnic

import (
	"net"
)

func GetNics() ([]net.Interface, error) { return net.Interfaces() }

func IsLoopback(i net.Interface) bool { return 0 != (i.Flags & net.FlagLoopback) }

func GetNicsF(nics []net.Interface, filter func(net.Interface) bool) (r []net.Interface) {
	for _, n := range nics {
		switch filter(n) {
		case true:
			r = append(r, n)
		default:
			continue
		}
	}
	return
}

func GetNicsFF(nics []net.Interface, filterfalse func(net.Interface) bool) (r []net.Interface) {
	for _, n := range nics {
		switch filterfalse(n) {
		case false:
			r = append(r, n)
		default:
			continue
		}
	}
	return
}
