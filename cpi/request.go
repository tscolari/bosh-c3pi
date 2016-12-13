package cpi

import (
	"fmt"

	"github.com/tscolari/bosh-c3pi/cloud"
)

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
	properties := new(cloud.CloudProperties)

	convertToType(r[index], properties)
	return *properties
}

func (r RequestArguments) ToResourcePool(index int) (resourcePool cloud.ResourcePool) {
	convertToType(r[index], &resourcePool)
	return resourcePool
}

func (r RequestArguments) ToNewtworks(index int) (networks cloud.Networks) {
	convertToType(r[index], &networks)
	return networks
}

func (r RequestArguments) ToEnvironment(index int) (environment cloud.Environment) {
	convertToType(r[index], &environment)
	return environment
}

func (r RequestArguments) ToInt(index int) int {
	return r[index].(int)
}

func (r RequestArguments) ToMetadata(index int) (metadata cloud.Metadata) {
	return r[index].(cloud.Metadata)
}

func convertToType(input interface{}, result interface{}) {
	data := input.(map[string]interface{})
	typeMapString := map[string]string{}
	typeMapNetwork := map[string]cloud.Network{}

	for key, value := range data {
		switch value := value.(type) {
		case string:
			typeMapString[key] = value
		case cloud.Network:
			typeMapNetwork[key] = value
		}
	}

	fmt.Printf("============ %#v\n", typeMapString)
	fmt.Printf("============ %#v\n", typeMapNetwork)
	switch result.(type) {
	case *cloud.CloudProperties, cloud.ResourcePool, cloud.Metadata, cloud.Environment:
		fmt.Printf("BLA")
		result = typeMapString
	case *cloud.Networks:
		fmt.Printf("BLE")
		result = typeMapNetwork
	}
}
