//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import "encoding/json"

func unmarshalAcceleratorAuthSettingClassification(rawMsg json.RawMessage) (AcceleratorAuthSettingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AcceleratorAuthSettingClassification
	switch m["authType"] {
	case "BasicAuth":
		b = &AcceleratorBasicAuthSetting{}
	case "Public":
		b = &AcceleratorPublicSetting{}
	case "SSH":
		b = &AcceleratorSSHSetting{}
	default:
		b = &AcceleratorAuthSetting{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCertificatePropertiesClassification(rawMsg json.RawMessage) (CertificatePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CertificatePropertiesClassification
	switch m["type"] {
	case "ContentCertificate":
		b = &ContentCertificateProperties{}
	case "KeyVaultCertificate":
		b = &KeyVaultCertificateProperties{}
	default:
		b = &CertificateProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalContainerRegistryCredentialsClassification(rawMsg json.RawMessage) (ContainerRegistryCredentialsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ContainerRegistryCredentialsClassification
	switch m["type"] {
	case "BasicAuth":
		b = &ContainerRegistryBasicCredentials{}
	default:
		b = &ContainerRegistryCredentials{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCustomPersistentDiskPropertiesClassification(rawMsg json.RawMessage) (CustomPersistentDiskPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CustomPersistentDiskPropertiesClassification
	switch m["type"] {
	case string(TypeAzureFileVolume):
		b = &AzureFileVolume{}
	default:
		b = &CustomPersistentDiskProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalGatewayResponseCachePropertiesClassification(rawMsg json.RawMessage) (GatewayResponseCachePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GatewayResponseCachePropertiesClassification
	switch m["responseCacheType"] {
	case "LocalCachePerInstance":
		b = &GatewayLocalResponseCachePerInstanceProperties{}
	case "LocalCachePerRoute":
		b = &GatewayLocalResponseCachePerRouteProperties{}
	default:
		b = &GatewayResponseCacheProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMaintenanceScheduleConfigurationClassification(rawMsg json.RawMessage) (MaintenanceScheduleConfigurationClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MaintenanceScheduleConfigurationClassification
	switch m["frequency"] {
	case string(FrequencyWeekly):
		b = &WeeklyMaintenanceScheduleConfiguration{}
	default:
		b = &MaintenanceScheduleConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProbeActionClassification(rawMsg json.RawMessage) (ProbeActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProbeActionClassification
	switch m["type"] {
	case string(ProbeActionTypeExecAction):
		b = &ExecAction{}
	case string(ProbeActionTypeHTTPGetAction):
		b = &HTTPGetAction{}
	case string(ProbeActionTypeTCPSocketAction):
		b = &TCPSocketAction{}
	default:
		b = &ProbeAction{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalStoragePropertiesClassification(rawMsg json.RawMessage) (StoragePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StoragePropertiesClassification
	switch m["storageType"] {
	case string(StorageTypeStorageAccount):
		b = &StorageAccount{}
	default:
		b = &StorageProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalUserSourceInfoClassification(rawMsg json.RawMessage) (UserSourceInfoClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UserSourceInfoClassification
	switch m["type"] {
	case "BuildResult":
		b = &BuildResultUserSourceInfo{}
	case "Container":
		b = &CustomContainerUserSourceInfo{}
	case "Jar":
		b = &JarUploadedUserSourceInfo{}
	case "NetCoreZip":
		b = &NetCoreZipUploadedUserSourceInfo{}
	case "Source":
		b = &SourceUploadedUserSourceInfo{}
	case "UploadedUserSourceInfo":
		b = &UploadedUserSourceInfo{}
	case "War":
		b = &WarUploadedUserSourceInfo{}
	default:
		b = &UserSourceInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
