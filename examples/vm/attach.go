package main

import (
	"fmt"
	"log"

	"github.com/lack-io/gvirt"
	"github.com/lack-io/gvirt/spec"
)

func main() {
	cc, err := gvirt.NewClient(gvirt.Addr("tcp", "192.168.2.189:16509"))
	if err != nil {
		log.Fatalln(err)
	}

	dom, err := cc.GetDomainByName("centos")
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

	//time.Sleep(time.Second * 2)
	//
	//for _, i := range dom.Devices.Interfaces {
	//	fmt.Println("mac:", i.MAC.Address)
	//	if i.MAC.Address != "52:54:00:bb:f3:ac" {
	//		x, _ := i.MarshalX()
	//		_, err = dom.DetachDevice(x)
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//	}
	//}
}
