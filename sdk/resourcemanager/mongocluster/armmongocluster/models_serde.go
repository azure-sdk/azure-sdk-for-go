// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CheckNameAvailabilityRequest.
func (c CheckNameAvailabilityRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameAvailabilityRequest.
func (c *CheckNameAvailabilityRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CheckNameAvailabilityResponse.
func (c CheckNameAvailabilityResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "nameAvailable", c.NameAvailable)
	populate(objectMap, "reason", c.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameAvailabilityResponse.
func (c *CheckNameAvailabilityResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &c.Message)
			delete(rawMsg, key)
		case "nameAvailable":
			err = unpopulate(val, "NameAvailable", &c.NameAvailable)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &c.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectionString.
func (c ConnectionString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionString", c.ConnectionString)
	populate(objectMap, "description", c.Description)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectionString.
func (c *ConnectionString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionString":
			err = unpopulate(val, "ConnectionString", &c.ConnectionString)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &c.Description)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FirewallRule.
func (f FirewallRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", f.ID)
	populate(objectMap, "name", f.Name)
	populate(objectMap, "properties", f.Properties)
	populate(objectMap, "systemData", f.SystemData)
	populate(objectMap, "type", f.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FirewallRule.
func (f *FirewallRule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &f.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &f.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &f.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &f.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &f.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FirewallRuleListResult.
func (f FirewallRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FirewallRuleListResult.
func (f *FirewallRuleListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &f.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &f.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FirewallRuleProperties.
func (f FirewallRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "endIpAddress", f.EndIPAddress)
	populate(objectMap, "provisioningState", f.ProvisioningState)
	populate(objectMap, "startIpAddress", f.StartIPAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FirewallRuleProperties.
func (f *FirewallRuleProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endIpAddress":
			err = unpopulate(val, "EndIPAddress", &f.EndIPAddress)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &f.ProvisioningState)
			delete(rawMsg, key)
		case "startIpAddress":
			err = unpopulate(val, "StartIPAddress", &f.StartIPAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListConnectionStringsResult.
func (l ListConnectionStringsResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionStrings", l.ConnectionStrings)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListConnectionStringsResult.
func (l *ListConnectionStringsResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionStrings":
			err = unpopulate(val, "ConnectionStrings", &l.ConnectionStrings)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListResult.
func (l ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListResult.
func (l *ListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MongoCluster.
func (m MongoCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MongoCluster.
func (m *MongoCluster) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &m.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &m.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &m.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NodeGroupSpec.
func (n NodeGroupSpec) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "diskSizeGB", n.DiskSizeGB)
	populate(objectMap, "enableHa", n.EnableHa)
	populate(objectMap, "kind", n.Kind)
	populate(objectMap, "nodeCount", n.NodeCount)
	populate(objectMap, "sku", n.SKU)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NodeGroupSpec.
func (n *NodeGroupSpec) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "diskSizeGB":
			err = unpopulate(val, "DiskSizeGB", &n.DiskSizeGB)
			delete(rawMsg, key)
		case "enableHa":
			err = unpopulate(val, "EnableHa", &n.EnableHa)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &n.Kind)
			delete(rawMsg, key)
		case "nodeCount":
			err = unpopulate(val, "NodeCount", &n.NodeCount)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &n.SKU)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actionType", o.ActionType)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionType":
			err = unpopulate(val, "ActionType", &o.ActionType)
			delete(rawMsg, key)
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpoint.
func (p PrivateEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateEndpoint.
func (p *PrivateEndpoint) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnection.
func (p PrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateEndpointConnection.
func (p *PrivateEndpointConnection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &p.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionProperties.
func (p PrivateEndpointConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "groupIds", p.GroupIDs)
	populate(objectMap, "privateEndpoint", p.PrivateEndpoint)
	populate(objectMap, "privateLinkServiceConnectionState", p.PrivateLinkServiceConnectionState)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateEndpointConnectionProperties.
func (p *PrivateEndpointConnectionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "groupIds":
			err = unpopulate(val, "GroupIDs", &p.GroupIDs)
			delete(rawMsg, key)
		case "privateEndpoint":
			err = unpopulate(val, "PrivateEndpoint", &p.PrivateEndpoint)
			delete(rawMsg, key)
		case "privateLinkServiceConnectionState":
			err = unpopulate(val, "PrivateLinkServiceConnectionState", &p.PrivateLinkServiceConnectionState)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionResource.
func (p PrivateEndpointConnectionResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateEndpointConnectionResource.
func (p *PrivateEndpointConnectionResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &p.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionResourceListResult.
func (p PrivateEndpointConnectionResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateEndpointConnectionResourceListResult.
func (p *PrivateEndpointConnectionResourceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &p.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &p.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResource.
func (p PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateLinkResource.
func (p *PrivateLinkResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &p.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateLinkResourceListResult.
func (p *PrivateLinkResourceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &p.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &p.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateLinkResourceProperties.
func (p *PrivateLinkResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "groupId":
			err = unpopulate(val, "GroupID", &p.GroupID)
			delete(rawMsg, key)
		case "requiredMembers":
			err = unpopulate(val, "RequiredMembers", &p.RequiredMembers)
			delete(rawMsg, key)
		case "requiredZoneNames":
			err = unpopulate(val, "RequiredZoneNames", &p.RequiredZoneNames)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkServiceConnectionState.
func (p PrivateLinkServiceConnectionState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actionsRequired", p.ActionsRequired)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "status", p.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateLinkServiceConnectionState.
func (p *PrivateLinkServiceConnectionState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionsRequired":
			err = unpopulate(val, "ActionsRequired", &p.ActionsRequired)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "administratorLogin", p.AdministratorLogin)
	populate(objectMap, "administratorLoginPassword", p.AdministratorLoginPassword)
	populate(objectMap, "clusterStatus", p.ClusterStatus)
	populate(objectMap, "connectionString", p.ConnectionString)
	populate(objectMap, "createMode", p.CreateMode)
	populate(objectMap, "earliestRestoreTime", p.EarliestRestoreTime)
	populate(objectMap, "nodeGroupSpecs", p.NodeGroupSpecs)
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	populate(objectMap, "restoreParameters", p.RestoreParameters)
	populate(objectMap, "serverVersion", p.ServerVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "administratorLogin":
			err = unpopulate(val, "AdministratorLogin", &p.AdministratorLogin)
			delete(rawMsg, key)
		case "administratorLoginPassword":
			err = unpopulate(val, "AdministratorLoginPassword", &p.AdministratorLoginPassword)
			delete(rawMsg, key)
		case "clusterStatus":
			err = unpopulate(val, "ClusterStatus", &p.ClusterStatus)
			delete(rawMsg, key)
		case "connectionString":
			err = unpopulate(val, "ConnectionString", &p.ConnectionString)
			delete(rawMsg, key)
		case "createMode":
			err = unpopulate(val, "CreateMode", &p.CreateMode)
			delete(rawMsg, key)
		case "earliestRestoreTime":
			err = unpopulate(val, "EarliestRestoreTime", &p.EarliestRestoreTime)
			delete(rawMsg, key)
		case "nodeGroupSpecs":
			err = unpopulate(val, "NodeGroupSpecs", &p.NodeGroupSpecs)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, "PrivateEndpointConnections", &p.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, "PublicNetworkAccess", &p.PublicNetworkAccess)
			delete(rawMsg, key)
		case "restoreParameters":
			err = unpopulate(val, "RestoreParameters", &p.RestoreParameters)
			delete(rawMsg, key)
		case "serverVersion":
			err = unpopulate(val, "ServerVersion", &p.ServerVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreParameters.
func (r RestoreParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "pointInTimeUTC", r.PointInTimeUTC)
	populate(objectMap, "sourceResourceId", r.SourceResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreParameters.
func (r *RestoreParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "pointInTimeUTC":
			err = unpopulateDateTimeRFC3339(val, "PointInTimeUTC", &r.PointInTimeUTC)
			delete(rawMsg, key)
		case "sourceResourceId":
			err = unpopulate(val, "SourceResourceID", &r.SourceResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateDateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
