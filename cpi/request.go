package cpi

import "github.com/tscolari/bosh-c3pi/cloud"

type Request struct {
	Method    string           `json:"method"`
	Arguments RequestArguments `json:"arguments"`
	Context   RequestContext   `json:"context"`
}

type RequestContext struct {
	DirectorUUID string `json:"director_uuid"`
}

type RequestArguments []interface{}

func (r RequestArguments) ToString(index int) string {
	return r[index].(string)
}

func (r RequestArguments) ToCloudProperties(index int) cloud.CloudProperties {
	return r[index].(cloud.CloudProperties)
}

func (r RequestArguments) ToResourcePool(index int) cloud.ResourcePool {
	return r[index].(cloud.ResourcePool)
}

func (r RequestArguments) ToNewtworks(index int) cloud.Networks {
	return r[index].(cloud.Networks)
}

func (r RequestArguments) ToEnvironment(index int) cloud.Environment {
	return r[index].(cloud.Environment)
}

func (r RequestArguments) ToInt(index int) int {
	return r[index].(int)
}

func (r RequestArguments) ToMetadata(index int) cloud.Metadata {
	return r[index].(cloud.Metadata)
}
