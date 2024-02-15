//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabasewatcher

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Datastore.
func (d Datastore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "adxClusterResourceId", d.AdxClusterResourceID)
	populate(objectMap, "kustoClusterDisplayName", d.KustoClusterDisplayName)
	populate(objectMap, "kustoClusterUri", d.KustoClusterURI)
	populate(objectMap, "kustoDataIngestionUri", d.KustoDataIngestionURI)
	populate(objectMap, "kustoDatabaseName", d.KustoDatabaseName)
	populate(objectMap, "kustoManagementUrl", d.KustoManagementURL)
	populate(objectMap, "kustoOfferingType", d.KustoOfferingType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datastore.
func (d *Datastore) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "adxClusterResourceId":
			err = unpopulate(val, "AdxClusterResourceID", &d.AdxClusterResourceID)
			delete(rawMsg, key)
		case "kustoClusterDisplayName":
			err = unpopulate(val, "KustoClusterDisplayName", &d.KustoClusterDisplayName)
			delete(rawMsg, key)
		case "kustoClusterUri":
			err = unpopulate(val, "KustoClusterURI", &d.KustoClusterURI)
			delete(rawMsg, key)
		case "kustoDataIngestionUri":
			err = unpopulate(val, "KustoDataIngestionURI", &d.KustoDataIngestionURI)
			delete(rawMsg, key)
		case "kustoDatabaseName":
			err = unpopulate(val, "KustoDatabaseName", &d.KustoDatabaseName)
			delete(rawMsg, key)
		case "kustoManagementUrl":
			err = unpopulate(val, "KustoManagementURL", &d.KustoManagementURL)
			delete(rawMsg, key)
		case "kustoOfferingType":
			err = unpopulate(val, "KustoOfferingType", &d.KustoOfferingType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DatastoreUpdate.
func (d DatastoreUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "adxClusterResourceId", d.AdxClusterResourceID)
	populate(objectMap, "kustoClusterDisplayName", d.KustoClusterDisplayName)
	populate(objectMap, "kustoClusterUri", d.KustoClusterURI)
	populate(objectMap, "kustoDataIngestionUri", d.KustoDataIngestionURI)
	populate(objectMap, "kustoDatabaseName", d.KustoDatabaseName)
	populate(objectMap, "kustoManagementUrl", d.KustoManagementURL)
	populate(objectMap, "kustoOfferingType", d.KustoOfferingType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatastoreUpdate.
func (d *DatastoreUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "adxClusterResourceId":
			err = unpopulate(val, "AdxClusterResourceID", &d.AdxClusterResourceID)
			delete(rawMsg, key)
		case "kustoClusterDisplayName":
			err = unpopulate(val, "KustoClusterDisplayName", &d.KustoClusterDisplayName)
			delete(rawMsg, key)
		case "kustoClusterUri":
			err = unpopulate(val, "KustoClusterURI", &d.KustoClusterURI)
			delete(rawMsg, key)
		case "kustoDataIngestionUri":
			err = unpopulate(val, "KustoDataIngestionURI", &d.KustoDataIngestionURI)
			delete(rawMsg, key)
		case "kustoDatabaseName":
			err = unpopulate(val, "KustoDatabaseName", &d.KustoDatabaseName)
			delete(rawMsg, key)
		case "kustoManagementUrl":
			err = unpopulate(val, "KustoManagementURL", &d.KustoManagementURL)
			delete(rawMsg, key)
		case "kustoOfferingType":
			err = unpopulate(val, "KustoOfferingType", &d.KustoOfferingType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedServiceIdentity.
func (m ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagedServiceIdentity.
func (m *ManagedServiceIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &m.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &m.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		case "userAssignedIdentities":
			err = unpopulate(val, "UserAssignedIdentities", &m.UserAssignedIdentities)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
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

// MarshalJSON implements the json.Marshaller interface for type SQLDbElasticPoolTargetProperties.
func (s SQLDbElasticPoolTargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "anchorDatabaseResourceId", s.AnchorDatabaseResourceID)
	populate(objectMap, "connectionServerName", s.ConnectionServerName)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "readIntent", s.ReadIntent)
	populate(objectMap, "sqlEpResourceId", s.SQLEpResourceID)
	populate(objectMap, "targetAuthenticationType", s.TargetAuthenticationType)
	populate(objectMap, "targetCollectionStatus", s.TargetCollectionStatus)
	objectMap["targetType"] = "SqlEp"
	populate(objectMap, "targetVault", s.TargetVault)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLDbElasticPoolTargetProperties.
func (s *SQLDbElasticPoolTargetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "anchorDatabaseResourceId":
			err = unpopulate(val, "AnchorDatabaseResourceID", &s.AnchorDatabaseResourceID)
			delete(rawMsg, key)
		case "connectionServerName":
			err = unpopulate(val, "ConnectionServerName", &s.ConnectionServerName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "readIntent":
			err = unpopulate(val, "ReadIntent", &s.ReadIntent)
			delete(rawMsg, key)
		case "sqlEpResourceId":
			err = unpopulate(val, "SQLEpResourceID", &s.SQLEpResourceID)
			delete(rawMsg, key)
		case "targetAuthenticationType":
			err = unpopulate(val, "TargetAuthenticationType", &s.TargetAuthenticationType)
			delete(rawMsg, key)
		case "targetCollectionStatus":
			err = unpopulate(val, "TargetCollectionStatus", &s.TargetCollectionStatus)
			delete(rawMsg, key)
		case "targetType":
			err = unpopulate(val, "TargetType", &s.TargetType)
			delete(rawMsg, key)
		case "targetVault":
			err = unpopulate(val, "TargetVault", &s.TargetVault)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLDbSingleDatabaseTargetProperties.
func (s SQLDbSingleDatabaseTargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionServerName", s.ConnectionServerName)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "readIntent", s.ReadIntent)
	populate(objectMap, "sqlDbResourceId", s.SQLDbResourceID)
	populate(objectMap, "targetAuthenticationType", s.TargetAuthenticationType)
	populate(objectMap, "targetCollectionStatus", s.TargetCollectionStatus)
	objectMap["targetType"] = "SqlDb"
	populate(objectMap, "targetVault", s.TargetVault)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLDbSingleDatabaseTargetProperties.
func (s *SQLDbSingleDatabaseTargetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionServerName":
			err = unpopulate(val, "ConnectionServerName", &s.ConnectionServerName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "readIntent":
			err = unpopulate(val, "ReadIntent", &s.ReadIntent)
			delete(rawMsg, key)
		case "sqlDbResourceId":
			err = unpopulate(val, "SQLDbResourceID", &s.SQLDbResourceID)
			delete(rawMsg, key)
		case "targetAuthenticationType":
			err = unpopulate(val, "TargetAuthenticationType", &s.TargetAuthenticationType)
			delete(rawMsg, key)
		case "targetCollectionStatus":
			err = unpopulate(val, "TargetCollectionStatus", &s.TargetCollectionStatus)
			delete(rawMsg, key)
		case "targetType":
			err = unpopulate(val, "TargetType", &s.TargetType)
			delete(rawMsg, key)
		case "targetVault":
			err = unpopulate(val, "TargetVault", &s.TargetVault)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLMiTargetProperties.
func (s SQLMiTargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionServerName", s.ConnectionServerName)
	populate(objectMap, "connectionTcpPort", s.ConnectionTCPPort)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "readIntent", s.ReadIntent)
	populate(objectMap, "sqlMiResourceId", s.SQLMiResourceID)
	populate(objectMap, "targetAuthenticationType", s.TargetAuthenticationType)
	populate(objectMap, "targetCollectionStatus", s.TargetCollectionStatus)
	objectMap["targetType"] = "SqlMi"
	populate(objectMap, "targetVault", s.TargetVault)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLMiTargetProperties.
func (s *SQLMiTargetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionServerName":
			err = unpopulate(val, "ConnectionServerName", &s.ConnectionServerName)
			delete(rawMsg, key)
		case "connectionTcpPort":
			err = unpopulate(val, "ConnectionTCPPort", &s.ConnectionTCPPort)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "readIntent":
			err = unpopulate(val, "ReadIntent", &s.ReadIntent)
			delete(rawMsg, key)
		case "sqlMiResourceId":
			err = unpopulate(val, "SQLMiResourceID", &s.SQLMiResourceID)
			delete(rawMsg, key)
		case "targetAuthenticationType":
			err = unpopulate(val, "TargetAuthenticationType", &s.TargetAuthenticationType)
			delete(rawMsg, key)
		case "targetCollectionStatus":
			err = unpopulate(val, "TargetCollectionStatus", &s.TargetCollectionStatus)
			delete(rawMsg, key)
		case "targetType":
			err = unpopulate(val, "TargetType", &s.TargetType)
			delete(rawMsg, key)
		case "targetVault":
			err = unpopulate(val, "TargetVault", &s.TargetVault)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLVMTargetProperties.
func (s SQLVMTargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionServerName", s.ConnectionServerName)
	populate(objectMap, "connectionTcpPort", s.ConnectionTCPPort)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "sqlNamedInstanceName", s.SQLNamedInstanceName)
	populate(objectMap, "sqlVmResourceId", s.SQLVMResourceID)
	populate(objectMap, "targetAuthenticationType", s.TargetAuthenticationType)
	populate(objectMap, "targetCollectionStatus", s.TargetCollectionStatus)
	objectMap["targetType"] = "SqlVm"
	populate(objectMap, "targetVault", s.TargetVault)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLVMTargetProperties.
func (s *SQLVMTargetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionServerName":
			err = unpopulate(val, "ConnectionServerName", &s.ConnectionServerName)
			delete(rawMsg, key)
		case "connectionTcpPort":
			err = unpopulate(val, "ConnectionTCPPort", &s.ConnectionTCPPort)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "sqlNamedInstanceName":
			err = unpopulate(val, "SQLNamedInstanceName", &s.SQLNamedInstanceName)
			delete(rawMsg, key)
		case "sqlVmResourceId":
			err = unpopulate(val, "SQLVMResourceID", &s.SQLVMResourceID)
			delete(rawMsg, key)
		case "targetAuthenticationType":
			err = unpopulate(val, "TargetAuthenticationType", &s.TargetAuthenticationType)
			delete(rawMsg, key)
		case "targetCollectionStatus":
			err = unpopulate(val, "TargetCollectionStatus", &s.TargetCollectionStatus)
			delete(rawMsg, key)
		case "targetType":
			err = unpopulate(val, "TargetType", &s.TargetType)
			delete(rawMsg, key)
		case "targetVault":
			err = unpopulate(val, "TargetVault", &s.TargetVault)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SharedPrivateLinkResource.
func (s SharedPrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SharedPrivateLinkResource.
func (s *SharedPrivateLinkResource) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type SharedPrivateLinkResourceListResult.
func (s SharedPrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SharedPrivateLinkResourceListResult.
func (s *SharedPrivateLinkResourceListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type SharedPrivateLinkResourceProperties.
func (s SharedPrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "dnsZone", s.DNSZone)
	populate(objectMap, "groupId", s.GroupID)
	populate(objectMap, "privateLinkResourceId", s.PrivateLinkResourceID)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "requestMessage", s.RequestMessage)
	populate(objectMap, "status", s.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SharedPrivateLinkResourceProperties.
func (s *SharedPrivateLinkResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dnsZone":
			err = unpopulate(val, "DNSZone", &s.DNSZone)
			delete(rawMsg, key)
		case "groupId":
			err = unpopulate(val, "GroupID", &s.GroupID)
			delete(rawMsg, key)
		case "privateLinkResourceId":
			err = unpopulate(val, "PrivateLinkResourceID", &s.PrivateLinkResourceID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "requestMessage":
			err = unpopulate(val, "RequestMessage", &s.RequestMessage)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
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

// MarshalJSON implements the json.Marshaller interface for type Target.
func (t Target) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Target.
func (t *Target) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "properties":
			t.Properties, err = unmarshalTargetPropertiesClassification(val)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &t.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TargetListResult.
func (t TargetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", t.NextLink)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TargetListResult.
func (t *TargetListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &t.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &t.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TargetProperties.
func (t TargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionServerName", t.ConnectionServerName)
	populate(objectMap, "provisioningState", t.ProvisioningState)
	populate(objectMap, "targetAuthenticationType", t.TargetAuthenticationType)
	populate(objectMap, "targetCollectionStatus", t.TargetCollectionStatus)
	objectMap["targetType"] = t.TargetType
	populate(objectMap, "targetVault", t.TargetVault)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TargetProperties.
func (t *TargetProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionServerName":
			err = unpopulate(val, "ConnectionServerName", &t.ConnectionServerName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &t.ProvisioningState)
			delete(rawMsg, key)
		case "targetAuthenticationType":
			err = unpopulate(val, "TargetAuthenticationType", &t.TargetAuthenticationType)
			delete(rawMsg, key)
		case "targetCollectionStatus":
			err = unpopulate(val, "TargetCollectionStatus", &t.TargetCollectionStatus)
			delete(rawMsg, key)
		case "targetType":
			err = unpopulate(val, "TargetType", &t.TargetType)
			delete(rawMsg, key)
		case "targetVault":
			err = unpopulate(val, "TargetVault", &t.TargetVault)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserAssignedIdentity.
func (u UserAssignedIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "clientId", u.ClientID)
	populate(objectMap, "principalId", u.PrincipalID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserAssignedIdentity.
func (u *UserAssignedIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clientId":
			err = unpopulate(val, "ClientID", &u.ClientID)
			delete(rawMsg, key)
		case "principalId":
			err = unpopulate(val, "PrincipalID", &u.PrincipalID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VaultSecret.
func (v VaultSecret) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "akvResourceId", v.AkvResourceID)
	populate(objectMap, "akvTargetPassword", v.AkvTargetPassword)
	populate(objectMap, "akvTargetUser", v.AkvTargetUser)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VaultSecret.
func (v *VaultSecret) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "akvResourceId":
			err = unpopulate(val, "AkvResourceID", &v.AkvResourceID)
			delete(rawMsg, key)
		case "akvTargetPassword":
			err = unpopulate(val, "AkvTargetPassword", &v.AkvTargetPassword)
			delete(rawMsg, key)
		case "akvTargetUser":
			err = unpopulate(val, "AkvTargetUser", &v.AkvTargetUser)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Watcher.
func (w Watcher) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Watcher.
func (w *Watcher) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &w.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &w.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &w.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &w.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &w.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &w.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &w.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WatcherListResult.
func (w WatcherListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WatcherListResult.
func (w *WatcherListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &w.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &w.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WatcherProperties.
func (w WatcherProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "datastore", w.Datastore)
	populate(objectMap, "provisioningState", w.ProvisioningState)
	populate(objectMap, "status", w.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WatcherProperties.
func (w *WatcherProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "datastore":
			err = unpopulate(val, "Datastore", &w.Datastore)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &w.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &w.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WatcherUpdate.
func (w WatcherUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WatcherUpdate.
func (w *WatcherUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "identity":
			err = unpopulate(val, "Identity", &w.Identity)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &w.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WatcherUpdateProperties.
func (w WatcherUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "datastore", w.Datastore)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WatcherUpdateProperties.
func (w *WatcherUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "datastore":
			err = unpopulate(val, "Datastore", &w.Datastore)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
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
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
