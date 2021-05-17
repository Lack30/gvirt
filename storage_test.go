package gvirt

import (
	"fmt"
	"testing"
)

func TestClient_GetAllStoragePools(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	pools, err := c.GetAllStoragePools()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(pools)
}

func TestClient_GetAllStorageVolumes(t *testing.T) {
	c, err := NewClient(Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	pool, err := c.GetStoragePoolByName("data-1")
	if err != nil {
		t.Fatal(err)
	}

	vols, err := pool.GetAllStorageVolumes()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(vols)
}