package cpi

import (
	"errors"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/tscolari/bosh-c3pi/cloud"
)

func NewDispatcher(cloud cloud.Cloud, logger boshlog.Logger) *Dispatcher {
	return &Dispatcher{
		cloud:  cloud,
		logger: logger,
	}
}

type Dispatcher struct {
	cloud  cloud.Cloud
	logger boshlog.Logger
}

func (r *Dispatcher) Dispatch(method string, arguments RequestArguments) (result interface{}, err error) {
	switch method {
	case "create_stemcell":
		result, err = r.cloud.CreateStemcell(
			arguments.ToString(0),
			arguments.ToCloudProperties(1),
		)
	case "delete_stemcell":
		err = r.cloud.DeleteStemcell(arguments.ToString(0))

	case "create_vm":
		result, err = r.cloud.CreateVm(
			arguments.ToString(0),
			arguments.ToString(1),
			arguments.ToResourcePool(2),
			arguments.ToNewtworks(3),
			arguments.ToString(4),
			arguments.ToEnvironment(5),
		)
	case "has_vm":
		result, err = r.cloud.HasVm(arguments.ToString(0))
	case "delete_vm":
		err = r.cloud.DeleteVm(arguments.ToString(0))
	case "reboot_vm":
		err = r.cloud.RebootVm(arguments.ToString(0))
	case "set_vm_metadata":
		err = r.cloud.SetVmMetadata(
			arguments.ToString(0),
			arguments.ToMetadata(1),
		)

	case "create_disk":
		result, err = r.cloud.CreateDisk(
			arguments.ToInt(0),
			arguments.ToCloudProperties(1),
			arguments.ToString(2),
		)
	case "has_disk":
		result, err = r.cloud.HasDisk(arguments.ToString(0))
	case "delete_disk":
		err = r.cloud.DeleteDisk(arguments.ToString(0))
	case "attach_disk":
		err = r.cloud.AttachDisk(
			arguments.ToString(0),
			arguments.ToString(1),
		)
	case "get_disks":
		result, err = r.cloud.GetDisks(arguments.ToString(0))
	case "detach_disk":
		err = r.cloud.DetachDisk(
			arguments.ToString(0),
			arguments.ToString(1),
		)

	case "snapshot_disk":
		result, err = r.cloud.SnapshotDisk(
			arguments.ToString(0),
			arguments.ToMetadata(1),
		)
	case "delete_snapshot":
		err = r.cloud.DeleteSnapshot(arguments.ToString(0))
	default:
		err = errors.New("Invalid cpi method: '" + method + "'")
	}

	return result, err
}
