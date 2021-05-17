package gvirt

import (
	"fmt"
	"testing"

	"github.com/lack-io/gvirt/spec"
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

	dom, err := c.GetDomainByName("centos")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(dom)
}

func TestDomain_SetVCPUs(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	dom, err := c.GetDomainByName("centos")
	if err != nil {
		t.Log(err)
	}
	err = dom.SetVCPUs(&spec.DomainVCPU{
		Placement: "auto",
		Current:   1,
		Value:     4,
	})
	if err != nil {
		t.Fatal(err)
	}
}