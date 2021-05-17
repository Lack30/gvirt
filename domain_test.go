package gkvm

import (
	"fmt"
	"testing"
)

func TestClient_GetAllDomains(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	domains, err := c.GetAllDomains()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(domains)
}

func TestClient_GetDomainById(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	dom, err := c.GetDomainById(5)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(dom)
}

func TestClient_GetDomainByName(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	dom, err := c.GetDomainByName("windows")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(dom)
}