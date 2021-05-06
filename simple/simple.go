package main

import (
	"fmt"
	"log"
	"net/url"

	libvirt "libvirt.org/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu+tcp://192.168.2.189:16509/system")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	u, _ := conn.GetURI()
	uu, _ := url.Parse(u)
	fmt.Println(uu.Scheme)
	fmt.Println(conn.GetVersion())
	fmt.Println(conn.GetLibVersion())

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d running domains:\n", len(doms))
	fmt.Printf("Id\tUuid\tName\tState\n")
	for _, dom := range doms {
		id, _ := dom.GetID()
		uuid, _ := dom.GetUUIDString()
		name, _ := dom.GetName()
		info, _ := dom.GetInfo()
		hostname, _ := dom.GetHostname(libvirt.DOMAIN_GET_HOSTNAME_AGENT)
		fmt.Println(hostname)
		fmt.Printf("%d \t %s \t %s\t %d \n", id, string(uuid), name, info.State)
		_ = dom.Free()
	}

	dom, err := conn.LookupDomainByName("centos")
	if err != nil {
		log.Fatalln(err)
	}
	defer dom.Free()

	//out, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_INACTIVE | libvirt.DOMAIN_XML_SECURE | libvirt.DOMAIN_XML_UPDATE_CPU)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(out)

	//info, err := dom.GetGuestInfo(libvirt.DOMAIN_GUEST_INFO_HOSTNAME, 0)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(info)

	//fs, err := dom.GetFSInfo(0)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(fs[0])

	//pools, err := conn.ListAllStoragePools(libvirt.CONNECT_LIST_STORAGE_POOLS_ACTIVE)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, pool := range pools {
	//	x, _ := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	//	fmt.Println(x)
	//	//info, _ := pool.GetInfo()
	//}
}
