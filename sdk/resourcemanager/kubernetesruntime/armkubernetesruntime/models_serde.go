// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armkubernetesruntime

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BgpPeer.
func (b BgpPeer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "properties", b.Properties)
	populate(objectMap, "systemData", b.SystemData)
	populate(objectMap, "type", b.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BgpPeer.
func (b *BgpPeer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &b.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &b.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &b.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &b.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &b.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BgpPeerListResult.
func (b BgpPeerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", b.NextLink)
	populate(objectMap, "value", b.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BgpPeerListResult.
func (b *BgpPeerListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &b.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &b.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BgpPeerProperties.
func (b BgpPeerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "myAsn", b.MyAsn)
	populate(objectMap, "peerAddress", b.PeerAddress)
	populate(objectMap, "peerAsn", b.PeerAsn)
	populate(objectMap, "provisioningState", b.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BgpPeerProperties.
func (b *BgpPeerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "myAsn":
			err = unpopulate(val, "MyAsn", &b.MyAsn)
			delete(rawMsg, key)
		case "peerAddress":
			err = unpopulate(val, "PeerAddress", &b.PeerAddress)
			delete(rawMsg, key)
		case "peerAsn":
			err = unpopulate(val, "PeerAsn", &b.PeerAsn)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &b.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BlobStorageClassTypeProperties.
func (b BlobStorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureStorageAccountKey", b.AzureStorageAccountKey)
	populate(objectMap, "azureStorageAccountName", b.AzureStorageAccountName)
	objectMap["type"] = SCTypeBlob
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BlobStorageClassTypeProperties.
func (b *BlobStorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureStorageAccountKey":
			err = unpopulate(val, "AzureStorageAccountKey", &b.AzureStorageAccountKey)
			delete(rawMsg, key)
		case "azureStorageAccountName":
			err = unpopulate(val, "AzureStorageAccountName", &b.AzureStorageAccountName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &b.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LoadBalancer.
func (l LoadBalancer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LoadBalancer.
func (l *LoadBalancer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &l.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &l.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &l.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &l.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &l.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LoadBalancerListResult.
func (l LoadBalancerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LoadBalancerListResult.
func (l *LoadBalancerListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type LoadBalancerProperties.
func (l LoadBalancerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "addresses", l.Addresses)
	populate(objectMap, "advertiseMode", l.AdvertiseMode)
	populate(objectMap, "bgpPeers", l.BgpPeers)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "serviceSelector", l.ServiceSelector)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LoadBalancerProperties.
func (l *LoadBalancerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "addresses":
			err = unpopulate(val, "Addresses", &l.Addresses)
			delete(rawMsg, key)
		case "advertiseMode":
			err = unpopulate(val, "AdvertiseMode", &l.AdvertiseMode)
			delete(rawMsg, key)
		case "bgpPeers":
			err = unpopulate(val, "BgpPeers", &l.BgpPeers)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &l.ProvisioningState)
			delete(rawMsg, key)
		case "serviceSelector":
			err = unpopulate(val, "ServiceSelector", &l.ServiceSelector)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NativeStorageClassTypeProperties.
func (n NativeStorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["type"] = SCTypeNative
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NativeStorageClassTypeProperties.
func (n *NativeStorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "type":
			err = unpopulate(val, "Type", &n.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NfsStorageClassTypeProperties.
func (n NfsStorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "mountPermissions", n.MountPermissions)
	populate(objectMap, "onDelete", n.OnDelete)
	populate(objectMap, "server", n.Server)
	populate(objectMap, "share", n.Share)
	populate(objectMap, "subDir", n.SubDir)
	objectMap["type"] = SCTypeNFS
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NfsStorageClassTypeProperties.
func (n *NfsStorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "mountPermissions":
			err = unpopulate(val, "MountPermissions", &n.MountPermissions)
			delete(rawMsg, key)
		case "onDelete":
			err = unpopulate(val, "OnDelete", &n.OnDelete)
			delete(rawMsg, key)
		case "server":
			err = unpopulate(val, "Server", &n.Server)
			delete(rawMsg, key)
		case "share":
			err = unpopulate(val, "Share", &n.Share)
			delete(rawMsg, key)
		case "subDir":
			err = unpopulate(val, "SubDir", &n.SubDir)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &n.Type)
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

// MarshalJSON implements the json.Marshaller interface for type RwxStorageClassTypeProperties.
func (r RwxStorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "backingStorageClassName", r.BackingStorageClassName)
	objectMap["type"] = SCTypeRWX
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RwxStorageClassTypeProperties.
func (r *RwxStorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backingStorageClassName":
			err = unpopulate(val, "BackingStorageClassName", &r.BackingStorageClassName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceProperties.
func (s ServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "rpObjectId", s.RpObjectID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceProperties.
func (s *ServiceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "rpObjectId":
			err = unpopulate(val, "RpObjectID", &s.RpObjectID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceResource.
func (s ServiceResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceResource.
func (s *ServiceResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &s.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceResourceListResult.
func (s ServiceResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceResourceListResult.
func (s *ServiceResourceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SmbStorageClassTypeProperties.
func (s SmbStorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "domain", s.Domain)
	populate(objectMap, "password", s.Password)
	populate(objectMap, "source", s.Source)
	populate(objectMap, "subDir", s.SubDir)
	objectMap["type"] = SCTypeSMB
	populate(objectMap, "username", s.Username)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmbStorageClassTypeProperties.
func (s *SmbStorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "domain":
			err = unpopulate(val, "Domain", &s.Domain)
			delete(rawMsg, key)
		case "password":
			err = unpopulate(val, "Password", &s.Password)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &s.Source)
			delete(rawMsg, key)
		case "subDir":
			err = unpopulate(val, "SubDir", &s.SubDir)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		case "username":
			err = unpopulate(val, "Username", &s.Username)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassProperties.
func (s StorageClassProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "accessModes", s.AccessModes)
	populate(objectMap, "allowVolumeExpansion", s.AllowVolumeExpansion)
	populate(objectMap, "dataResilience", s.DataResilience)
	populate(objectMap, "failoverSpeed", s.FailoverSpeed)
	populate(objectMap, "limitations", s.Limitations)
	populate(objectMap, "mountOptions", s.MountOptions)
	populate(objectMap, "performance", s.Performance)
	populate(objectMap, "priority", s.Priority)
	populate(objectMap, "provisioner", s.Provisioner)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "typeProperties", s.TypeProperties)
	populate(objectMap, "volumeBindingMode", s.VolumeBindingMode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassProperties.
func (s *StorageClassProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessModes":
			err = unpopulate(val, "AccessModes", &s.AccessModes)
			delete(rawMsg, key)
		case "allowVolumeExpansion":
			err = unpopulate(val, "AllowVolumeExpansion", &s.AllowVolumeExpansion)
			delete(rawMsg, key)
		case "dataResilience":
			err = unpopulate(val, "DataResilience", &s.DataResilience)
			delete(rawMsg, key)
		case "failoverSpeed":
			err = unpopulate(val, "FailoverSpeed", &s.FailoverSpeed)
			delete(rawMsg, key)
		case "limitations":
			err = unpopulate(val, "Limitations", &s.Limitations)
			delete(rawMsg, key)
		case "mountOptions":
			err = unpopulate(val, "MountOptions", &s.MountOptions)
			delete(rawMsg, key)
		case "performance":
			err = unpopulate(val, "Performance", &s.Performance)
			delete(rawMsg, key)
		case "priority":
			err = unpopulate(val, "Priority", &s.Priority)
			delete(rawMsg, key)
		case "provisioner":
			err = unpopulate(val, "Provisioner", &s.Provisioner)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "typeProperties":
			s.TypeProperties, err = unmarshalStorageClassTypePropertiesClassification(val)
			delete(rawMsg, key)
		case "volumeBindingMode":
			err = unpopulate(val, "VolumeBindingMode", &s.VolumeBindingMode)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassPropertiesUpdate.
func (s StorageClassPropertiesUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "accessModes", s.AccessModes)
	populate(objectMap, "allowVolumeExpansion", s.AllowVolumeExpansion)
	populate(objectMap, "dataResilience", s.DataResilience)
	populate(objectMap, "failoverSpeed", s.FailoverSpeed)
	populate(objectMap, "limitations", s.Limitations)
	populate(objectMap, "mountOptions", s.MountOptions)
	populate(objectMap, "performance", s.Performance)
	populate(objectMap, "priority", s.Priority)
	populate(objectMap, "typeProperties", s.TypeProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassPropertiesUpdate.
func (s *StorageClassPropertiesUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessModes":
			err = unpopulate(val, "AccessModes", &s.AccessModes)
			delete(rawMsg, key)
		case "allowVolumeExpansion":
			err = unpopulate(val, "AllowVolumeExpansion", &s.AllowVolumeExpansion)
			delete(rawMsg, key)
		case "dataResilience":
			err = unpopulate(val, "DataResilience", &s.DataResilience)
			delete(rawMsg, key)
		case "failoverSpeed":
			err = unpopulate(val, "FailoverSpeed", &s.FailoverSpeed)
			delete(rawMsg, key)
		case "limitations":
			err = unpopulate(val, "Limitations", &s.Limitations)
			delete(rawMsg, key)
		case "mountOptions":
			err = unpopulate(val, "MountOptions", &s.MountOptions)
			delete(rawMsg, key)
		case "performance":
			err = unpopulate(val, "Performance", &s.Performance)
			delete(rawMsg, key)
		case "priority":
			err = unpopulate(val, "Priority", &s.Priority)
			delete(rawMsg, key)
		case "typeProperties":
			err = unpopulate(val, "TypeProperties", &s.TypeProperties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassResource.
func (s StorageClassResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassResource.
func (s *StorageClassResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &s.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassResourceListResult.
func (s StorageClassResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassResourceListResult.
func (s *StorageClassResourceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassResourceUpdate.
func (s StorageClassResourceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassResourceUpdate.
func (s *StorageClassResourceUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassTypeProperties.
func (s StorageClassTypeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassTypeProperties.
func (s *StorageClassTypeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageClassTypePropertiesUpdate.
func (s StorageClassTypePropertiesUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureStorageAccountKey", s.AzureStorageAccountKey)
	populate(objectMap, "azureStorageAccountName", s.AzureStorageAccountName)
	populate(objectMap, "backingStorageClassName", s.BackingStorageClassName)
	populate(objectMap, "domain", s.Domain)
	populate(objectMap, "mountPermissions", s.MountPermissions)
	populate(objectMap, "onDelete", s.OnDelete)
	populate(objectMap, "password", s.Password)
	populate(objectMap, "server", s.Server)
	populate(objectMap, "share", s.Share)
	populate(objectMap, "source", s.Source)
	populate(objectMap, "subDir", s.SubDir)
	populate(objectMap, "username", s.Username)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageClassTypePropertiesUpdate.
func (s *StorageClassTypePropertiesUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureStorageAccountKey":
			err = unpopulate(val, "AzureStorageAccountKey", &s.AzureStorageAccountKey)
			delete(rawMsg, key)
		case "azureStorageAccountName":
			err = unpopulate(val, "AzureStorageAccountName", &s.AzureStorageAccountName)
			delete(rawMsg, key)
		case "backingStorageClassName":
			err = unpopulate(val, "BackingStorageClassName", &s.BackingStorageClassName)
			delete(rawMsg, key)
		case "domain":
			err = unpopulate(val, "Domain", &s.Domain)
			delete(rawMsg, key)
		case "mountPermissions":
			err = unpopulate(val, "MountPermissions", &s.MountPermissions)
			delete(rawMsg, key)
		case "onDelete":
			err = unpopulate(val, "OnDelete", &s.OnDelete)
			delete(rawMsg, key)
		case "password":
			err = unpopulate(val, "Password", &s.Password)
			delete(rawMsg, key)
		case "server":
			err = unpopulate(val, "Server", &s.Server)
			delete(rawMsg, key)
		case "share":
			err = unpopulate(val, "Share", &s.Share)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &s.Source)
			delete(rawMsg, key)
		case "subDir":
			err = unpopulate(val, "SubDir", &s.SubDir)
			delete(rawMsg, key)
		case "username":
			err = unpopulate(val, "Username", &s.Username)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
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
