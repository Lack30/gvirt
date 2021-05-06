package gkvm

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	v1, v2 := c.Version()
	t.Log(v1, v2)
}
