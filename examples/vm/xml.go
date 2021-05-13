package main

var dev = `<volume type='file'>
  <name>test.img</name>
  <key>/home/kvm/data/test.img</key>
  <source>
  </source>
  <capacity unit='bytes'>21474836480</capacity>
  <allocation unit='bytes'>2044604416</allocation>
  <physical unit='bytes'>21478375424</physical>
  <target>
    <path>/home/kvm/data/test.img</path>
    <format type='qcow2'/>
    <permissions>
      <mode>0600</mode>
      <owner>0</owner>
      <group>0</group>
    </permissions>
    <features>
      <lazy_refcounts/>
    </features>
  </target>
</volume>`

var xml = `<domain type='kvm' id='11'>
  <name>test</name>
  <memory unit='KiB'>1048576</memory>
  <currentMemory unit='KiB'>1048576</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <resource>
    <partition>/machine</partition>
  </resource>
  <os>
    <type arch='x86_64' machine='pc-i440fx-rhel7.0.0'>hvm</type>
	<boot dev='cdrom'/>
    <boot dev='hd'/>
  </os>
  <features>
    <acpi/>
    <apic/>
  </features>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/libexec/qemu-kvm</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source file='/home/kvm/data/test.img'/>
      <backingStore/>
      <target dev='sda' bus='virtio'/>
      <alias name='virtio-disk0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </disk>
    <disk type='file' device='cdrom'>
      <driver name='qemu'/>
	  <source file='/home/kvm/iso/CentOS-7-x86_64-DVD-2003.iso'/>
      <target dev='hda' bus='ide'/>
      <readonly/>
      <alias name='ide0-0-0'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <alias name='usb'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <alias name='usb'/>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <alias name='usb'/>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <alias name='usb'/>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'>
      <alias name='pci.0'/>
    </controller>
    <controller type='ide' index='0'>
      <alias name='ide'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <controller type='virtio-serial' index='0'>
      <alias name='virtio-serial0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </controller>
    <interface type='bridge'>
      <mac address='52:54:00:9c:f3:ec'/>
      <source bridge='br0'/>
      <target dev='vnet0'/>
      <model type='virtio'/>
      <alias name='net0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <source path='/dev/pts/3'/>
      <target type='isa-serial' port='0'>
        <model name='isa-serial'/>
      </target>
      <alias name='serial0'/>
    </serial>
    <console type='pty' tty='/dev/pts/3'>
      <source path='/dev/pts/3'/>
      <target type='serial' port='0'/>
      <alias name='serial0'/>
    </console>
    <input type='tablet' bus='usb'>
      <alias name='input0'/>
      <address type='usb' bus='0' port='1'/>
    </input>
    <input type='mouse' bus='ps2'>
      <alias name='input1'/>
    </input>
    <input type='keyboard' bus='ps2'>
      <alias name='input2'/>
    </input>
    <graphics type='vnc' port='5990' autoport='no' listen='0.0.0.0'>
      <listen type='address' address='0.0.0.0'/>
    </graphics>
    <video>
      <model type='cirrus' vram='16384' heads='1' primary='yes'/>
      <alias name='video0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <alias name='balloon0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </memballoon>
    <rng model='virtio'>
      <backend model='random'>/dev/urandom</backend>
      <alias name='rng0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x08' function='0x0'/>
    </rng>
  </devices>
</domain>`

const DeviceXml = `...
<devices>
  <disk type='file' snapshot='external'>
    <driver name="tap" type="aio" cache="default"/>
    <source file='/var/lib/xen/images/fv0' startupPolicy='optional'>
      <seclabel relabel='no'/>
    </source>
    <target dev='hda' bus='ide'/>
    <iotune>
      <total_bytes_sec>10000000</total_bytes_sec>
      <read_iops_sec>400000</read_iops_sec>
      <write_iops_sec>100000</write_iops_sec>
    </iotune>
    <boot order='2'/>
    <encryption type='...'>
      ...
    </encryption>
    <shareable/>
    <serial>
      ...
    </serial>
  </disk>
    ...
  <disk type='network'>
    <driver name="qemu" type="raw" io="threads" ioeventfd="on" event_idx="off"/>
    <source protocol="sheepdog" name="image_name">
      <host name="hostname" port="7000"/>
    </source>
    <target dev="hdb" bus="ide"/>
    <boot order='1'/>
    <transient/>
    <address type='drive' controller='0' bus='1' unit='0'/>
  </disk>
  <disk type='network'>
    <driver name="qemu" type="raw"/>
    <source protocol="rbd" name="image_name2">
      <host name="hostname" port="7000"/>
      <snapshot name="snapname"/>
      <config file="/path/to/file"/>
      <auth username='myuser'>
        <secret type='ceph' usage='mypassid'/>
      </auth>
    </source>
    <target dev="hdc" bus="ide"/>
  </disk>
  <disk type='block' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <target dev='hdd' bus='ide' tray='open'/>
    <readonly/>
  </disk>
  <disk type='network' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <source protocol="http" name="url_path" query="foo=bar&amp;baz=flurb>
      <host name="hostname" port="80"/>
      <cookies>
        <cookie name="test">somevalue</cookie>
      </cookies>
      <readahead size='65536'/>
      <timeout seconds='6'/>
    </source>
    <target dev='hde' bus='ide' tray='open'/>
    <readonly/>
  </disk>
  <disk type='network' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <source protocol="https" name="url_path">
      <host name="hostname" port="443"/>
      <ssl verify="no"/>
    </source>
    <target dev='hdf' bus='ide' tray='open'/>
    <readonly/>
  </disk>
  <disk type='network' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <source protocol="ftp" name="url_path">
      <host name="hostname" port="21"/>
    </source>
    <target dev='hdg' bus='ide' tray='open'/>
    <readonly/>
  </disk>
  <disk type='network' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <source protocol="ftps" name="url_path">
      <host name="hostname" port="990"/>
    </source>
    <target dev='hdh' bus='ide' tray='open'/>
    <readonly/>
  </disk>
  <disk type='network' device='cdrom'>
    <driver name='qemu' type='raw'/>
    <source protocol="tftp" name="url_path">
      <host name="hostname" port="69"/>
    </source>
    <target dev='hdi' bus='ide' tray='open' rotation_rate='7200'/>
    <readonly/>
  </disk>
  <disk type='block' device='lun'>
    <driver name='qemu' type='raw'/>
    <source dev='/dev/sda'>
      <slices>
        <slice type='storage' offset='12345' size='123'/>
      </slices>
      <reservations managed='no'>
        <source type='unix' path='/path/to/qemu-pr-helper' mode='client'/>
      </reservations>
    </source>
    <target dev='sda' bus='scsi' rotation_rate='1'/>
    <address type='drive' controller='0' bus='0' target='3' unit='0'/>
  </disk>
  <disk type='block' device='disk'>
    <driver name='qemu' type='raw'/>
    <source dev='/dev/sda'/>
    <geometry cyls='16383' heads='16' secs='63' trans='lba'/>
    <blockio logical_block_size='512' physical_block_size='4096'/>
    <target dev='hdj' bus='ide'/>
  </disk>
  <disk type='volume' device='disk'>
    <driver name='qemu' type='raw'/>
    <source pool='blk-pool0' volume='blk-pool0-vol0'/>
    <target dev='hdk' bus='ide'/>
  </disk>
  <disk type='network' device='disk'>
    <driver name='qemu' type='raw'/>
    <source protocol='iscsi' name='iqn.2013-07.com.example:iscsi-nopool/2'>
      <host name='example.com' port='3260'/>
      <auth username='myuser'>
        <secret type='iscsi' usage='libvirtiscsi'/>
      </auth>
    </source>
    <target dev='vda' bus='virtio'/>
  </disk>
  <disk type='network' device='lun'>
    <driver name='qemu' type='raw'/>
    <source protocol='iscsi' name='iqn.2013-07.com.example:iscsi-nopool/1'>
      <host name='example.com' port='3260'/>
      <auth username='myuser'>
        <secret type='iscsi' usage='libvirtiscsi'/>
      </auth>
    </source>
    <target dev='sdb' bus='scsi'/>
  </disk>
  <disk type='network' device='disk'>
    <driver name='qemu' type='raw'/>
    <source protocol='nfs' name='PATH'>
      <host name='example.com'/>
      <identity user='USER' group='GROUP'/>
    </source>
    <target dev='vda' bus='virtio'/>
  </disk>
  <disk type='network' device='lun'>
    <driver name='qemu' type='raw'/>
    <source protocol='iscsi' name='iqn.2013-07.com.example:iscsi-nopool/0'>
      <host name='example.com' port='3260'/>
      <initiator>
        <iqn name='iqn.2013-07.com.example:client'/>
      </initiator>
    </source>
    <target dev='sdb' bus='scsi'/>
  </disk>
  <disk type='volume' device='disk'>
    <driver name='qemu' type='raw'/>
    <source pool='iscsi-pool' volume='unit:0:0:1' mode='host'/>
    <target dev='vdb' bus='virtio'/>
  </disk>
  <disk type='volume' device='disk'>
    <driver name='qemu' type='raw'/>
    <source pool='iscsi-pool' volume='unit:0:0:2' mode='direct'/>
    <target dev='vdc' bus='virtio'/>
  </disk>
  <disk type='file' device='disk'>
    <driver name='qemu' type='qcow2' queues='4'/>
    <source file='/var/lib/libvirt/images/domain.qcow'/>
    <backingStore type='file'>
      <format type='qcow2'/>
      <source file='/var/lib/libvirt/images/snapshot.qcow'/>
      <backingStore type='block'>
        <format type='raw'/>
        <source dev='/dev/mapper/base'/>
        <backingStore/>
      </backingStore>
    </backingStore>
    <target dev='vdd' bus='virtio'/>
  </disk>
  <disk type='nvme' device='disk'>
    <driver name='qemu' type='raw'/>
    <source type='pci' managed='yes' namespace='1'>
      <address domain='0x0000' bus='0x01' slot='0x00' function='0x0'/>
    </source>
    <target dev='vde' bus='virtio'/>
  </disk>
  <disk type='vhostuser' device='disk'>
    <driver name='qemu' type='raw'/>
    <source type='unix' path='/tmp/vhost-blk.sock'>
      <reconnect enabled='yes' timeout='10'/>
    </source>
    <target dev='vdf' bus='virtio'/>
  </disk>
</devices>
...`