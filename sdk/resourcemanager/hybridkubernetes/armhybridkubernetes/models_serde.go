//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AADProfile.
func (a AADProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "adminGroupObjectIDs", a.AdminGroupObjectIDs)
	populate(objectMap, "enableAzureRBAC", a.EnableAzureRBAC)
	populate(objectMap, "tenantID", a.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AADProfile.
func (a *AADProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "adminGroupObjectIDs":
			err = unpopulate(val, "AdminGroupObjectIDs", &a.AdminGroupObjectIDs)
			delete(rawMsg, key)
		case "enableAzureRBAC":
			err = unpopulate(val, "EnableAzureRBAC", &a.EnableAzureRBAC)
			delete(rawMsg, key)
		case "tenantID":
			err = unpopulate(val, "TenantID", &a.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ArcAgentProfile.
func (a ArcAgentProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "agentAutoUpgrade", a.AgentAutoUpgrade)
	populate(objectMap, "desiredAgentVersion", a.DesiredAgentVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArcAgentProfile.
func (a *ArcAgentProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "agentAutoUpgrade":
			err = unpopulate(val, "AgentAutoUpgrade", &a.AgentAutoUpgrade)
			delete(rawMsg, key)
		case "desiredAgentVersion":
			err = unpopulate(val, "DesiredAgentVersion", &a.DesiredAgentVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedCluster.
func (c ConnectedCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "kind", c.Kind)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedCluster.
func (c *ConnectedCluster) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &c.Identity)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &c.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &c.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &c.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &c.Tags)
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

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterIdentity.
func (c ConnectedClusterIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", c.PrincipalID)
	populate(objectMap, "tenantId", c.TenantID)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterIdentity.
func (c *ConnectedClusterIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &c.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &c.TenantID)
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

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterList.
func (c ConnectedClusterList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterList.
func (c *ConnectedClusterList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &c.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &c.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterPatch.
func (c ConnectedClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterPatch.
func (c *ConnectedClusterPatch) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &c.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterPatchProperties.
func (c ConnectedClusterPatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureHybridBenefit", c.AzureHybridBenefit)
	populate(objectMap, "distribution", c.Distribution)
	populate(objectMap, "distributionVersion", c.DistributionVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterPatchProperties.
func (c *ConnectedClusterPatchProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureHybridBenefit":
			err = unpopulate(val, "AzureHybridBenefit", &c.AzureHybridBenefit)
			delete(rawMsg, key)
		case "distribution":
			err = unpopulate(val, "Distribution", &c.Distribution)
			delete(rawMsg, key)
		case "distributionVersion":
			err = unpopulate(val, "DistributionVersion", &c.DistributionVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterProperties.
func (c ConnectedClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aadProfile", c.AADProfile)
	populate(objectMap, "agentPublicKeyCertificate", c.AgentPublicKeyCertificate)
	populate(objectMap, "agentVersion", c.AgentVersion)
	populate(objectMap, "arcAgentProfile", c.ArcAgentProfile)
	populate(objectMap, "azureHybridBenefit", c.AzureHybridBenefit)
	populate(objectMap, "connectivityStatus", c.ConnectivityStatus)
	populate(objectMap, "distribution", c.Distribution)
	populate(objectMap, "distributionVersion", c.DistributionVersion)
	populate(objectMap, "infrastructure", c.Infrastructure)
	populate(objectMap, "kubernetesVersion", c.KubernetesVersion)
	populateDateTimeRFC3339(objectMap, "lastConnectivityTime", c.LastConnectivityTime)
	populateDateTimeRFC3339(objectMap, "managedIdentityCertificateExpirationTime", c.ManagedIdentityCertificateExpirationTime)
	populate(objectMap, "miscellaneousProperties", c.MiscellaneousProperties)
	populate(objectMap, "offering", c.Offering)
	populate(objectMap, "privateLinkScopeResourceId", c.PrivateLinkScopeResourceID)
	populate(objectMap, "privateLinkState", c.PrivateLinkState)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "totalCoreCount", c.TotalCoreCount)
	populate(objectMap, "totalNodeCount", c.TotalNodeCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterProperties.
func (c *ConnectedClusterProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aadProfile":
			err = unpopulate(val, "AADProfile", &c.AADProfile)
			delete(rawMsg, key)
		case "agentPublicKeyCertificate":
			err = unpopulate(val, "AgentPublicKeyCertificate", &c.AgentPublicKeyCertificate)
			delete(rawMsg, key)
		case "agentVersion":
			err = unpopulate(val, "AgentVersion", &c.AgentVersion)
			delete(rawMsg, key)
		case "arcAgentProfile":
			err = unpopulate(val, "ArcAgentProfile", &c.ArcAgentProfile)
			delete(rawMsg, key)
		case "azureHybridBenefit":
			err = unpopulate(val, "AzureHybridBenefit", &c.AzureHybridBenefit)
			delete(rawMsg, key)
		case "connectivityStatus":
			err = unpopulate(val, "ConnectivityStatus", &c.ConnectivityStatus)
			delete(rawMsg, key)
		case "distribution":
			err = unpopulate(val, "Distribution", &c.Distribution)
			delete(rawMsg, key)
		case "distributionVersion":
			err = unpopulate(val, "DistributionVersion", &c.DistributionVersion)
			delete(rawMsg, key)
		case "infrastructure":
			err = unpopulate(val, "Infrastructure", &c.Infrastructure)
			delete(rawMsg, key)
		case "kubernetesVersion":
			err = unpopulate(val, "KubernetesVersion", &c.KubernetesVersion)
			delete(rawMsg, key)
		case "lastConnectivityTime":
			err = unpopulateDateTimeRFC3339(val, "LastConnectivityTime", &c.LastConnectivityTime)
			delete(rawMsg, key)
		case "managedIdentityCertificateExpirationTime":
			err = unpopulateDateTimeRFC3339(val, "ManagedIdentityCertificateExpirationTime", &c.ManagedIdentityCertificateExpirationTime)
			delete(rawMsg, key)
		case "miscellaneousProperties":
			err = unpopulate(val, "MiscellaneousProperties", &c.MiscellaneousProperties)
			delete(rawMsg, key)
		case "offering":
			err = unpopulate(val, "Offering", &c.Offering)
			delete(rawMsg, key)
		case "privateLinkScopeResourceId":
			err = unpopulate(val, "PrivateLinkScopeResourceID", &c.PrivateLinkScopeResourceID)
			delete(rawMsg, key)
		case "privateLinkState":
			err = unpopulate(val, "PrivateLinkState", &c.PrivateLinkState)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &c.ProvisioningState)
			delete(rawMsg, key)
		case "totalCoreCount":
			err = unpopulate(val, "TotalCoreCount", &c.TotalCoreCount)
			delete(rawMsg, key)
		case "totalNodeCount":
			err = unpopulate(val, "TotalNodeCount", &c.TotalNodeCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResult.
func (c CredentialResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populateByteArray(objectMap, "value", c.Value, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CredentialResult.
func (c *CredentialResult) UnmarshalJSON(data []byte) error {
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
		case "value":
			err = runtime.DecodeByteArray(string(val), &c.Value, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResults.
func (c CredentialResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "hybridConnectionConfig", c.HybridConnectionConfig)
	populate(objectMap, "kubeconfigs", c.Kubeconfigs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CredentialResults.
func (c *CredentialResults) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "hybridConnectionConfig":
			err = unpopulate(val, "HybridConnectionConfig", &c.HybridConnectionConfig)
			delete(rawMsg, key)
		case "kubeconfigs":
			err = unpopulate(val, "Kubeconfigs", &c.Kubeconfigs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
			err = unpopulate(val, "Info", &e.Info)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
			err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HybridConnectionConfig.
func (h HybridConnectionConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "expirationTime", h.ExpirationTime)
	populate(objectMap, "hybridConnectionName", h.HybridConnectionName)
	populate(objectMap, "relay", h.Relay)
	populate(objectMap, "token", h.Token)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HybridConnectionConfig.
func (h *HybridConnectionConfig) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTime":
			err = unpopulate(val, "ExpirationTime", &h.ExpirationTime)
			delete(rawMsg, key)
		case "hybridConnectionName":
			err = unpopulate(val, "HybridConnectionName", &h.HybridConnectionName)
			delete(rawMsg, key)
		case "relay":
			err = unpopulate(val, "Relay", &h.Relay)
			delete(rawMsg, key)
		case "token":
			err = unpopulate(val, "Token", &h.Token)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListClusterUserCredentialProperties.
func (l ListClusterUserCredentialProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authenticationMethod", l.AuthenticationMethod)
	populate(objectMap, "clientProxy", l.ClientProxy)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListClusterUserCredentialProperties.
func (l *ListClusterUserCredentialProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticationMethod":
			err = unpopulate(val, "AuthenticationMethod", &l.AuthenticationMethod)
			delete(rawMsg, key)
		case "clientProxy":
			err = unpopulate(val, "ClientProxy", &l.ClientProxy)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
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
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
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

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationList.
func (o *OperationList) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
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

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrackedResource.
func (t *TrackedResource) UnmarshalJSON(data []byte) error {
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
		case "location":
			err = unpopulate(val, "Location", &t.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
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

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func populateByteArray(m map[string]any, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
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
