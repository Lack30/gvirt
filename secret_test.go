package gvirt

import (
	"fmt"
	"testing"
)

func TestClient_GetAllSecrets(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	secs, err := c.GetAllSecrets()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(secs)
}
