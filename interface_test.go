package gvirt

import (
	"fmt"
	"testing"
)

func TestClient_GetAllInterfaces(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	secs, err := c.GetAllInterfaces()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(secs)
}
