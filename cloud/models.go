package cloud

type CloudProperties map[string]string

type ResourcePool map[string]string

type Networks map[string]Network

type Metadata map[string]string

type Environment map[string]string

type Network struct {
	Netmask         string
	IP              string
	Gateway         string
	Dns             []string
	CloudProperties CloudProperties
}
