package listnic

import (
	"net"
	"testing"
)

func TestIsLoopback(t *testing.T) {
	i0 := net.Interface{
		Flags: net.FlagLoopback,
	}
	switch IsLoopback(i0) {
	case true:
		break
	default:
		t.Errorf("got False; want: True")
	}

	i1 := net.Interface{
		Flags: net.FlagUp,
	}
	switch IsLoopback(i1) {
	case false:
		break
	default:
		t.Errorf("got True; want: False")
	}
}

func TestGetNicsF(t *testing.T) {
	nics := []net.Interface{
		{Flags: net.FlagUp},
		{Flags: net.FlagLoopback},
	}
	filtered := GetNicsF(nics, IsLoopback)

	switch len(filtered) {
	case 1:
		break
	default:
		t.Errorf("got: %v; want: 1", len(filtered))
	}

	switch filtered[0].Flags {
	case net.FlagLoopback:
		break
	default:
		t.Errorf("got: %v; want: FlagLoopback", filtered[0].Flags)
	}
}

func TestGetNicsFF(t *testing.T) {
	nics := []net.Interface{
		{Flags: net.FlagUp},
		{Flags: net.FlagLoopback},
	}
	filtered := GetNicsFF(nics, IsLoopback)

	switch len(filtered) {
	case 1:
		break
	default:
		t.Errorf("got: %v; want: 1", len(filtered))
	}

	switch filtered[0].Flags {
	case net.FlagUp:
		break
	default:
		t.Errorf("got: %v; want: FlagUp", filtered[0].Flags)
	}
}
