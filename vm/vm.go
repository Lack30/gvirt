package main

import (
	"fmt"
	"log"

	"libvirt.org/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu+tcp://192.168.2.189:16509/system")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// TODO: create volume by storage pool
	pool, err := conn.LookupStoragePoolByName("data-1")
	if err != nil {
		log.Fatalln(err)
	}

	vol, err := pool.StorageVolCreateXML(dev, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer vol.Free()

	volName, _ := vol.GetName()
	fmt.Println("create volume: ", volName)

	//dev, err := conn.DeviceCreateXML(dev, 0)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer dev.Free()
	//n, _ := dev.GetName()
	//fmt.Println("create device: ", n)

	d, err := conn.DomainDefineXMLFlags(xml, libvirt.DOMAIN_DEFINE_VALIDATE)
	if err != nil {
		log.Fatalln(err)
	}

	d.SetAutostart(true)

	if err := d.Create(); err != nil {
		log.Fatalln("start domain: ", err)
	}

	//_, err = conn.DomainEventAgentLifecycleRegister(d, func(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventAgentLifecycle) {
	//	fmt.Println(event.State, event.Reason)
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//d, err := conn.DomainCreateXML(xml, libvirt.DOMAIN_NONE)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//libvirt.EventRegisterDefaultImpl()
	//libvirt.EventRegisterImpl(libvirt.EventLoop())

	d.Suspend()
	d.Resume()
	d.Shutdown()
	d.Destroy()
	d.Undefine()

	uuid, err := d.GetUUIDString()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("create domain: ", uuid)
	defer d.Free()
}
