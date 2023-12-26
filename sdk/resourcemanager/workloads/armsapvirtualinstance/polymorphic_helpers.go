//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsapvirtualinstance

import "encoding/json"

func unmarshalFileShareConfigurationClassification(rawMsg json.RawMessage) (FileShareConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FileShareConfigurationClassification
	switch m["configurationType"] {
	case string(ConfigurationTypeCreateAndMount):
		b = &CreateAndMountFileShareConfiguration{}
	case string(ConfigurationTypeMount):
		b = &MountFileShareConfiguration{}
	case string(ConfigurationTypeSkip):
		b = &SkipFileShareConfiguration{}
	default:
		b = &FileShareConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalInfrastructureConfigurationClassification(rawMsg json.RawMessage) (InfrastructureConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InfrastructureConfigurationClassification
	switch m["deploymentType"] {
	case string(SAPDeploymentTypeSingleServer):
		b = &SingleServerConfiguration{}
	case string(SAPDeploymentTypeThreeTier):
		b = &ThreeTierConfiguration{}
	default:
		b = &InfrastructureConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOSConfigurationClassification(rawMsg json.RawMessage) (OSConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OSConfigurationClassification
	switch m["osType"] {
	case string(OSTypeLinux):
		b = &LinuxConfiguration{}
	case string(OSTypeWindows):
		b = &WindowsConfiguration{}
	default:
		b = &OSConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSAPConfigurationClassification(rawMsg json.RawMessage) (SAPConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SAPConfigurationClassification
	switch m["configurationType"] {
	case string(SAPConfigurationTypeDeployment):
		b = &DeploymentConfiguration{}
	case string(SAPConfigurationTypeDeploymentWithOSConfig):
		b = &DeploymentWithOSConfiguration{}
	case string(SAPConfigurationTypeDiscovery):
		b = &DiscoveryConfiguration{}
	default:
		b = &SAPConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSAPSizingRecommendationResultClassification(rawMsg json.RawMessage) (SAPSizingRecommendationResultClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SAPSizingRecommendationResultClassification
	switch m["deploymentType"] {
	case string(SAPDeploymentTypeSingleServer):
		b = &SingleServerRecommendationResult{}
	case string(SAPDeploymentTypeThreeTier):
		b = &ThreeTierRecommendationResult{}
	default:
		b = &SAPSizingRecommendationResult{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSingleServerCustomResourceNamesClassification(rawMsg json.RawMessage) (SingleServerCustomResourceNamesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SingleServerCustomResourceNamesClassification
	switch m["namingPatternType"] {
	case string(NamingPatternTypeFullResourceName):
		b = &SingleServerFullResourceNames{}
	default:
		b = &SingleServerCustomResourceNames{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSoftwareConfigurationClassification(rawMsg json.RawMessage) (SoftwareConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SoftwareConfigurationClassification
	switch m["softwareInstallationType"] {
	case string(SAPSoftwareInstallationTypeExternal):
		b = &ExternalInstallationSoftwareConfiguration{}
	case string(SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig):
		b = &SAPInstallWithoutOSConfigSoftwareConfiguration{}
	case string(SAPSoftwareInstallationTypeServiceInitiated):
		b = &ServiceInitiatedSoftwareConfiguration{}
	default:
		b = &SoftwareConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalThreeTierCustomResourceNamesClassification(rawMsg json.RawMessage) (ThreeTierCustomResourceNamesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ThreeTierCustomResourceNamesClassification
	switch m["namingPatternType"] {
	case string(NamingPatternTypeFullResourceName):
		b = &ThreeTierFullResourceNames{}
	default:
		b = &ThreeTierCustomResourceNames{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
