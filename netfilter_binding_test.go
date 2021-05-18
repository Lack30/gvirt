package gvirt

import (
	"fmt"
	"testing"
)

func TestClient_GetAllNWFilterBindings(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	bindings, err := c.GetAllNWFilterBindings()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(bindings)
}
