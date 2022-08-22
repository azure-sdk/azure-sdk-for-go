//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

import "encoding/json"

func unmarshalInventoryItemPropertiesClassification(rawMsg json.RawMessage) (InventoryItemPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InventoryItemPropertiesClassification
	switch m["inventoryType"] {
	case string(InventoryTypeCluster):
		b = &ClusterInventoryItem{}
	case string(InventoryTypeDatastore):
		b = &DatastoreInventoryItem{}
	case string(InventoryTypeHost):
		b = &HostInventoryItem{}
	case string(InventoryTypeResourcePool):
		b = &ResourcePoolInventoryItem{}
	case string(InventoryTypeVirtualMachine):
		b = &VirtualMachineInventoryItem{}
	case string(InventoryTypeVirtualMachineTemplate):
		b = &VirtualMachineTemplateInventoryItem{}
	case string(InventoryTypeVirtualNetwork):
		b = &VirtualNetworkInventoryItem{}
	default:
		b = &InventoryItemProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
