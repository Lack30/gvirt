package gkvm

import (
	"fmt"
	"testing"
)

func TestClient_GetDomains(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	domains, err := c.GetDomains()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(domains)
}
