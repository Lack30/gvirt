package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lack-io/gvirt"
	"github.com/lack-io/gvirt/spec"
)

func main() {
	cc, err := gvirt.NewClient(gvirt.Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		log.Fatalln(err)
	}

	dom, err := cc.GetDomainByName("oracle")
	if err != nil {
		log.Fatalln(err)
	}

	nw, err := cc.GetNetworkByName("default")
	if err != nil {
		log.Fatalln(err)
	}

	di := &spec.DomainInterface{
		Source: &spec.DomainInterfaceSource{
			Network: &spec.DomainInterfaceSourceNetwork{
				Network: "default",
			},
		},
	}
	fmt.Println(di.MarshalX())

	dom, err = dom.AttachInterface(di)
	if err != nil {
		log.Fatalln(err)
	}

	iface := dom.Devices.Interfaces[len(dom.Devices.Interfaces)-1]
	fmt.Println(iface.MAC.Address)

	var ip string
	for {
		time.Sleep(time.Second * 2)
		leases, _ := nw.Deref().GetDHCPLeases()
		for _, lease := range leases {
			if lease.Mac == iface.MAC.Address {
				ip = lease.IPaddr
			}
			if ip != "" {
				goto LEASE
			}
		}
	}

LEASE:
	fmt.Println(ip)
}
