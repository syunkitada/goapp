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
	Devices       []interface{}  `xml:"devices"`
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
	XMLName  xml.Name `xml:"emulator"`
	Emulator string   `xml:",chardata"`
}

type DeviceDisk struct {
	XMLName xml.Name    `xml:"disk"`
	Type    string      `xml:"type,attr"`
	Device  string      `xml:"device,attr"`
	Driver  DiskDriver  `xml:"driver"`
	Source  DiskSource  `xml:"source"`
	Target  DiskTarget  `xml:"target"`
	Alias   Alias       `xml:"alias"`
	Address interface{} `xml:"address"`
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
	Controller int    `xml:"controller,attr"`
	Bus        int    `xml:"bus,attr"`
	Target     int    `xml:"target,attr"`
	Unit       int    `xml:"unit,attr"`
}

type DeviceController struct {
	XMLName xml.Name    `xml:"controller"`
	Type    string      `xml:"type,attr"`
	Index   int         `xml:"index,attr"`
	Alias   Alias       `xml:"alias"`
	Address interface{} `xml:"address"`
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
	XMLName xml.Name     `xml:"serial"`
	Type    string       `xml:"type,attr"`
	Source  SerialSource `xml:"source"`
	Target  SerialTarget `xml:"target"`
	Alias   Alias        `xml:"alias"`
}

type SerialSource struct {
	Path string `xml:"source,attr"`
}

type SerialTarget struct {
	Port int `xml:"port,attr"`
}

type DeviceConsole struct {
	XMLName xml.Name      `xml:"console"`
	Type    string        `xml:"type,attr"`
	Tty     string        `xml:"tty,attr"`
	Source  ConsoleSource `xml:"source"`
	Target  ConsoleTarget `xml:"target"`
	Alias   Alias         `xml:"alias"`
}

type ConsoleSource struct {
	Path string `xml:"path,attr"`
}

type ConsoleTarget struct {
	Type string `xml:"type,attr"`
	Port int    `xml:"port,attr"`
}

type Membaloon struct {
	XMLName xml.Name   `xml:"membaloon"`
	Alias   Alias      `xml:"alias"`
	Address PciAddress `xml:"address"`
}
