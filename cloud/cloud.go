package cloud

type Cloud interface {
	CreateStemcell(imagePath string, cloudProperties CloudProperties) (string, error)
	DeleteStemcell(stemcellID string) error

	CurrentVmID() string
	CreateVm(agentID, stemcellID string, resourcePool ResourcePool, networks Networks, diskLocality string, env Environment) (string, error)
	DeleteVm(vmID string) error
	HasVm(vmID string) (bool, error)
	RebootVm(vmID string) error
	SetVmMetadata(vm string, metadata Metadata) error

	CreateDisk(size int, cloudProperties CloudProperties, vmLocality string) (string, error)
	GetDisks(vmID string) ([]string, error)
	HasDisk(diskID string) (bool, error)
	DeleteDisk(diskID string) error
	AttachDisk(vmID, diskID string) error
	DetachDisk(vmID, diskID string) error

	SnapshotDisk(diskID string, metadata Metadata) (string, error)
	DeleteSnapshot(snapshotID string) error
}
