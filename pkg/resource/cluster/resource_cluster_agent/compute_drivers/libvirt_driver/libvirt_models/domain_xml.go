package libvirt_models

import "encoding/xml"

type Domain struct {
	XMLName       xml.Name       `xml:"domain"`
	ID            uint           `xml:"id,attr"`
	Type          string         `xml:"type,attr"`
	Name          string         `xml:"name"`
	Memory        Memory         `xml:"memory"`
	CurrentMemory Memory         `xml:"currentMemory"`
	Vcpu          Vcpu           `xml:"vcpu"`
	Resource      Resource       `xml:"resource"`
	Os            Os             `xml:"os"`
	Features      Features       `xml:"features"`
	Cpu           Cpu            `xml:"cpu"`
	MemoryBacking *MemoryBacking `xml:"memoryBacking"`
	Clock         Clock          `xml:"clock"`
	OnPoweroff    string         `xml:"on_poweroff"`
	OnReboot      string         `xml:"on_reboot"`
	OnCrash       string         `xml:"on_crash"`
	Devices       Devices        `xml:"devices"`
}

type Devices struct {
	Emulators   []DeviceEmulator   `xml:"emulator"`
	Interfaces  []DeviceInterface  `xml:"interface"`
	Disks       []DeviceDisk       `xml:"disk"`
	Controllers []DeviceController `xml:"controller"`
	Serials     []DeviceSerial     `xml:"serial"`
	Consoles    []DeviceConsole    `xml:"console"`
	Membaloons  []DeviceMembaloon  `xml:"membaloon"`
}

type Memory struct {
	Unit   string `xml:"unit,attr"`
	Memory uint   `xml:",chardata"`
}

type Vcpu struct {
	Placement string `xml:"placement,attr"`
	Vcpu      uint   `xml:",chardata"`
}

type Resource struct {
	Partition string `xml:"partition"`
}

type OsType struct {
	Arch string `xml:"arch,attr"`
	Type string `xml:",chardata"`
}

type OsBoot struct {
	Dev string `xml:"dev,attr"`
}

type Os struct {
	Type OsType `xml:"type"`
	Boot OsBoot `xml:"boot"`
}

type FeaturesAcpi struct{}
type FeaturesApic struct{}
type FeaturesPae struct{}

type Features struct {
	Acpi FeaturesAcpi `xml:"acpi"`
	Apic FeaturesApic `xml:"apic"`
	Pae  FeaturesPae  `xml:"pae"`
}

type CpuModel struct {
	Fallback string `xml:"fallback,attr"`
}

type CpuTopology struct {
	Sockets int `xml:"sockets,attr"`
	Cores   int `xml:"cores,attr"`
	Threads int `xml:"threads,attr"`
}

type Cpu struct {
	Mode     string      `xml:"mode,attr"`
	Model    CpuModel    `xml:"model"`
	Topology CpuTopology `xml:"topology"`
}

type Hugepages struct {
}

type MemoryBacking struct {
	Hogepages Hugepages `xml:"hugepages"`
}

type Clock struct {
	Offset string `xml:"offset,attr"`
}

type DeviceEmulator struct {
	Emulator string `xml:",chardata"`
}

type DeviceDisk struct {
	Type    string       `xml:"type,attr"`
	Device  string       `xml:"device,attr"`
	Driver  DiskDriver   `xml:"driver"`
	Source  DiskSource   `xml:"source"`
	Target  DiskTarget   `xml:"target"`
	Alias   Alias        `xml:"alias"`
	Address DriveAddress `xml:"address"`
}

type DiskDriver struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Cache string `xml:"cache,attr"`
}

type DiskSource struct {
	File string `xml:"file,attr"`
}

type DiskTarget struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type Alias struct {
	Name string `xml:"name,attr"`
}

type DriveAddress struct {
	Type       string `xml:"type,attr"`
	Controller uint   `xml:"controller,attr"`
	Bus        uint   `xml:"bus,attr"`
	Target     uint   `xml:"target,attr"`
	Unit       uint   `xml:"unit,attr"`
}

type DeviceController struct {
	Type    string     `xml:"type,attr"`
	Index   int        `xml:"index,attr"`
	Alias   Alias      `xml:"alias"`
	Address PciAddress `xml:"address"`
}

type PciAddress struct {
	Type     string `xml:"type,attr"`
	Domain   string `xml:"domain,attr"`
	Bus      string `xml:"bus,attr"`
	Slot     string `xml:"slot,attr"`
	Function string `xml:"function,attr"`
}

type DeviceInterface struct {
	Type    string          `xml:"type,attr"`
	Driver  InterfaceDriver `xml:"driver"`
	Mac     InterfaceMac    `xml:"mac"`
	Source  InterfaceSource `xml:"source"`
	Target  InterfaceTarget `xml:"target"`
	Model   InterfaceModel  `xml:"model"`
	Alias   Alias           `xml:"alias"`
	Address PciAddress      `xml:"address"`
}

type InterfaceDriver struct {
	Name   string `xml:"name,attr"`
	Queues int    `xml:"queues,attr"`
}
type InterfaceMac struct {
	Address string `xml:"address,attr"`
}

type InterfaceSource struct {
	Bridge string `xml:"bridge,attr"`
}

type InterfaceTarget struct {
	Dev string `xml:"dev,attr"`
}

type InterfaceModel struct {
	Type string `xml:"type,attr"`
}

type DeviceSerial struct {
	Type   string       `xml:"type,attr"`
	Source SerialSource `xml:"source"`
	Target SerialTarget `xml:"target"`
	Alias  Alias        `xml:"alias"`
}

type SerialSource struct {
	Path string `xml:"source,attr"`
}

type SerialTarget struct {
	Port uint `xml:"port,attr"`
}

type DeviceConsole struct {
	Type   string        `xml:"type,attr"`
	Tty    string        `xml:"tty,attr"`
	Source ConsoleSource `xml:"source"`
	Target ConsoleTarget `xml:"target"`
	Alias  Alias         `xml:"alias"`
}

type ConsoleSource struct {
	Path string `xml:"path,attr"`
}

type ConsoleTarget struct {
	Type string `xml:"type,attr"`
	Port uint   `xml:"port,attr"`
}

type DeviceMembaloon struct {
	Model   string     `xml:"model,attr"`
	Alias   Alias      `xml:"alias"`
	Address PciAddress `xml:"address"`
}
