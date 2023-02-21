//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeploymentscripts

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AzureCliScript.
func (a AzureCliScript) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	objectMap["kind"] = ScriptTypeAzureCLI
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureCliScript.
func (a *AzureCliScript) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &a.Identity)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &a.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &a.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &a.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &a.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureCliScriptProperties.
func (a AzureCliScriptProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "arguments", a.Arguments)
	populate(objectMap, "azCliVersion", a.AzCliVersion)
	populate(objectMap, "cleanupPreference", a.CleanupPreference)
	populate(objectMap, "containerSettings", a.ContainerSettings)
	populate(objectMap, "environmentVariables", a.EnvironmentVariables)
	populate(objectMap, "forceUpdateTag", a.ForceUpdateTag)
	populate(objectMap, "outputs", a.Outputs)
	populate(objectMap, "primaryScriptUri", a.PrimaryScriptURI)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "retentionInterval", a.RetentionInterval)
	populate(objectMap, "scriptContent", a.ScriptContent)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "storageAccountSettings", a.StorageAccountSettings)
	populate(objectMap, "supportingScriptUris", a.SupportingScriptUris)
	populate(objectMap, "timeout", a.Timeout)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureCliScriptProperties.
func (a *AzureCliScriptProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "arguments":
			err = unpopulate(val, "Arguments", &a.Arguments)
			delete(rawMsg, key)
		case "azCliVersion":
			err = unpopulate(val, "AzCliVersion", &a.AzCliVersion)
			delete(rawMsg, key)
		case "cleanupPreference":
			err = unpopulate(val, "CleanupPreference", &a.CleanupPreference)
			delete(rawMsg, key)
		case "containerSettings":
			err = unpopulate(val, "ContainerSettings", &a.ContainerSettings)
			delete(rawMsg, key)
		case "environmentVariables":
			err = unpopulate(val, "EnvironmentVariables", &a.EnvironmentVariables)
			delete(rawMsg, key)
		case "forceUpdateTag":
			err = unpopulate(val, "ForceUpdateTag", &a.ForceUpdateTag)
			delete(rawMsg, key)
		case "outputs":
			err = unpopulate(val, "Outputs", &a.Outputs)
			delete(rawMsg, key)
		case "primaryScriptUri":
			err = unpopulate(val, "PrimaryScriptURI", &a.PrimaryScriptURI)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		case "retentionInterval":
			err = unpopulate(val, "RetentionInterval", &a.RetentionInterval)
			delete(rawMsg, key)
		case "scriptContent":
			err = unpopulate(val, "ScriptContent", &a.ScriptContent)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &a.Status)
			delete(rawMsg, key)
		case "storageAccountSettings":
			err = unpopulate(val, "StorageAccountSettings", &a.StorageAccountSettings)
			delete(rawMsg, key)
		case "supportingScriptUris":
			err = unpopulate(val, "SupportingScriptUris", &a.SupportingScriptUris)
			delete(rawMsg, key)
		case "timeout":
			err = unpopulate(val, "Timeout", &a.Timeout)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzurePowerShellScript.
func (a AzurePowerShellScript) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	objectMap["kind"] = ScriptTypeAzurePowerShell
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzurePowerShellScript.
func (a *AzurePowerShellScript) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &a.Identity)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &a.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &a.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &a.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &a.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzurePowerShellScriptProperties.
func (a AzurePowerShellScriptProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "arguments", a.Arguments)
	populate(objectMap, "azPowerShellVersion", a.AzPowerShellVersion)
	populate(objectMap, "cleanupPreference", a.CleanupPreference)
	populate(objectMap, "containerSettings", a.ContainerSettings)
	populate(objectMap, "environmentVariables", a.EnvironmentVariables)
	populate(objectMap, "forceUpdateTag", a.ForceUpdateTag)
	populate(objectMap, "outputs", a.Outputs)
	populate(objectMap, "primaryScriptUri", a.PrimaryScriptURI)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "retentionInterval", a.RetentionInterval)
	populate(objectMap, "scriptContent", a.ScriptContent)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "storageAccountSettings", a.StorageAccountSettings)
	populate(objectMap, "supportingScriptUris", a.SupportingScriptUris)
	populate(objectMap, "timeout", a.Timeout)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzurePowerShellScriptProperties.
func (a *AzurePowerShellScriptProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "arguments":
			err = unpopulate(val, "Arguments", &a.Arguments)
			delete(rawMsg, key)
		case "azPowerShellVersion":
			err = unpopulate(val, "AzPowerShellVersion", &a.AzPowerShellVersion)
			delete(rawMsg, key)
		case "cleanupPreference":
			err = unpopulate(val, "CleanupPreference", &a.CleanupPreference)
			delete(rawMsg, key)
		case "containerSettings":
			err = unpopulate(val, "ContainerSettings", &a.ContainerSettings)
			delete(rawMsg, key)
		case "environmentVariables":
			err = unpopulate(val, "EnvironmentVariables", &a.EnvironmentVariables)
			delete(rawMsg, key)
		case "forceUpdateTag":
			err = unpopulate(val, "ForceUpdateTag", &a.ForceUpdateTag)
			delete(rawMsg, key)
		case "outputs":
			err = unpopulate(val, "Outputs", &a.Outputs)
			delete(rawMsg, key)
		case "primaryScriptUri":
			err = unpopulate(val, "PrimaryScriptURI", &a.PrimaryScriptURI)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		case "retentionInterval":
			err = unpopulate(val, "RetentionInterval", &a.RetentionInterval)
			delete(rawMsg, key)
		case "scriptContent":
			err = unpopulate(val, "ScriptContent", &a.ScriptContent)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &a.Status)
			delete(rawMsg, key)
		case "storageAccountSettings":
			err = unpopulate(val, "StorageAccountSettings", &a.StorageAccountSettings)
			delete(rawMsg, key)
		case "supportingScriptUris":
			err = unpopulate(val, "SupportingScriptUris", &a.SupportingScriptUris)
			delete(rawMsg, key)
		case "timeout":
			err = unpopulate(val, "Timeout", &a.Timeout)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureResourceBase.
func (a AzureResourceBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureResourceBase.
func (a *AzureResourceBase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &a.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &a.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContainerConfiguration.
func (c ContainerConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "containerGroupName", c.ContainerGroupName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContainerConfiguration.
func (c *ContainerConfiguration) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "containerGroupName":
			err = unpopulate(val, "ContainerGroupName", &c.ContainerGroupName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentScript.
func (d DeploymentScript) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "identity", d.Identity)
	objectMap["kind"] = d.Kind
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScript.
func (d *DeploymentScript) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &d.Identity)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &d.Kind)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &d.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &d.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentScriptListResult.
func (d DeploymentScriptListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScriptListResult.
func (d *DeploymentScriptListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &d.NextLink)
			delete(rawMsg, key)
		case "value":
			d.Value, err = unmarshalDeploymentScriptClassificationArray(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentScriptUpdateParameter.
func (d DeploymentScriptUpdateParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScriptUpdateParameter.
func (d *DeploymentScriptUpdateParameter) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EnvironmentVariable.
func (e EnvironmentVariable) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "secureValue", e.SecureValue)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnvironmentVariable.
func (e *EnvironmentVariable) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "secureValue":
			err = unpopulate(val, "SecureValue", &e.SecureValue)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Error.
func (e *Error) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "info", &e.Info)
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

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
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

// MarshalJSON implements the json.Marshaller interface for type LogProperties.
func (l LogProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "log", l.Log)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogProperties.
func (l *LogProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "log":
			err = unpopulate(val, "Log", &l.Log)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedServiceIdentity.
func (m ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
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

// MarshalJSON implements the json.Marshaller interface for type ScriptLog.
func (s ScriptLog) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScriptLog.
func (s *ScriptLog) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type ScriptLogsList.
func (s ScriptLogsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScriptLogsList.
func (s *ScriptLogsList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

// MarshalJSON implements the json.Marshaller interface for type ScriptStatus.
func (s ScriptStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "containerInstanceId", s.ContainerInstanceID)
	populateTimeRFC3339(objectMap, "endTime", s.EndTime)
	populate(objectMap, "error", s.Error)
	populateTimeRFC3339(objectMap, "expirationTime", s.ExpirationTime)
	populateTimeRFC3339(objectMap, "startTime", s.StartTime)
	populate(objectMap, "storageAccountId", s.StorageAccountID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ScriptStatus.
func (s *ScriptStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "containerInstanceId":
			err = unpopulate(val, "ContainerInstanceID", &s.ContainerInstanceID)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &s.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &s.Error)
			delete(rawMsg, key)
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, "ExpirationTime", &s.ExpirationTime)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		case "storageAccountId":
			err = unpopulate(val, "StorageAccountID", &s.StorageAccountID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageAccountConfiguration.
func (s StorageAccountConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "storageAccountKey", s.StorageAccountKey)
	populate(objectMap, "storageAccountName", s.StorageAccountName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageAccountConfiguration.
func (s *StorageAccountConfiguration) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "storageAccountKey":
			err = unpopulate(val, "StorageAccountKey", &s.StorageAccountKey)
			delete(rawMsg, key)
		case "storageAccountName":
			err = unpopulate(val, "StorageAccountName", &s.StorageAccountName)
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
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
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
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
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
