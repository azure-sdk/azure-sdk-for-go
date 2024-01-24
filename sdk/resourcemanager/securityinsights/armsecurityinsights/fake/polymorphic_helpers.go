//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

func unmarshalAlertRuleClassification(rawMsg json.RawMessage) (armsecurityinsights.AlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.AlertRuleClassification
	switch m["kind"] {
	case string(armsecurityinsights.AlertRuleKindFusion):
		b = &armsecurityinsights.FusionAlertRule{}
	case string(armsecurityinsights.AlertRuleKindMLBehaviorAnalytics):
		b = &armsecurityinsights.MLBehaviorAnalyticsAlertRule{}
	case string(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation):
		b = &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRule{}
	case string(armsecurityinsights.AlertRuleKindNRT):
		b = &armsecurityinsights.NrtAlertRule{}
	case string(armsecurityinsights.AlertRuleKindScheduled):
		b = &armsecurityinsights.ScheduledAlertRule{}
	case string(armsecurityinsights.AlertRuleKindThreatIntelligence):
		b = &armsecurityinsights.ThreatIntelligenceAlertRule{}
	default:
		b = &armsecurityinsights.AlertRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCustomEntityQueryClassification(rawMsg json.RawMessage) (armsecurityinsights.CustomEntityQueryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.CustomEntityQueryClassification
	switch m["kind"] {
	case string(armsecurityinsights.CustomEntityQueryKindActivity):
		b = &armsecurityinsights.ActivityCustomEntityQuery{}
	default:
		b = &armsecurityinsights.CustomEntityQuery{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataConnectorClassification(rawMsg json.RawMessage) (armsecurityinsights.DataConnectorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.DataConnectorClassification
	switch m["kind"] {
	case string(armsecurityinsights.DataConnectorKindAPIPolling):
		b = &armsecurityinsights.CodelessAPIPollingDataConnector{}
	case string(armsecurityinsights.DataConnectorKindAmazonWebServicesCloudTrail):
		b = &armsecurityinsights.AwsCloudTrailDataConnector{}
	case string(armsecurityinsights.DataConnectorKindAmazonWebServicesS3):
		b = &armsecurityinsights.AwsS3DataConnector{}
	case string(armsecurityinsights.DataConnectorKindAzureActiveDirectory):
		b = &armsecurityinsights.AADDataConnector{}
	case string(armsecurityinsights.DataConnectorKindAzureAdvancedThreatProtection):
		b = &armsecurityinsights.AATPDataConnector{}
	case string(armsecurityinsights.DataConnectorKindAzureSecurityCenter):
		b = &armsecurityinsights.ASCDataConnector{}
	case string(armsecurityinsights.DataConnectorKindDynamics365):
		b = &armsecurityinsights.Dynamics365DataConnector{}
	case string(armsecurityinsights.DataConnectorKindGCP):
		b = &armsecurityinsights.GCPDataConnector{}
	case string(armsecurityinsights.DataConnectorKindGenericUI):
		b = &armsecurityinsights.CodelessUIDataConnector{}
	case string(armsecurityinsights.DataConnectorKindIOT):
		b = &armsecurityinsights.IoTDataConnector{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftCloudAppSecurity):
		b = &armsecurityinsights.MCASDataConnector{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection):
		b = &armsecurityinsights.MDATPDataConnector{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftPurviewInformationProtection):
		b = &armsecurityinsights.MicrosoftPurviewInformationProtectionDataConnector{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence):
		b = &armsecurityinsights.MSTIDataConnector{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftThreatProtection):
		b = &armsecurityinsights.MTPDataConnector{}
	case string(armsecurityinsights.DataConnectorKindOffice365):
		b = &armsecurityinsights.OfficeDataConnector{}
	case string(armsecurityinsights.DataConnectorKindOffice365Project):
		b = &armsecurityinsights.Office365ProjectDataConnector{}
	case string(armsecurityinsights.DataConnectorKindOfficeATP):
		b = &armsecurityinsights.OfficeATPDataConnector{}
	case string(armsecurityinsights.DataConnectorKindOfficeIRM):
		b = &armsecurityinsights.OfficeIRMDataConnector{}
	case string(armsecurityinsights.DataConnectorKindOfficePowerBI):
		b = &armsecurityinsights.OfficePowerBIDataConnector{}
	case string(armsecurityinsights.DataConnectorKindThreatIntelligence):
		b = &armsecurityinsights.TIDataConnector{}
	case string(armsecurityinsights.DataConnectorKindThreatIntelligenceTaxii):
		b = &armsecurityinsights.TiTaxiiDataConnector{}
	default:
		b = &armsecurityinsights.DataConnector{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataConnectorDefinitionClassification(rawMsg json.RawMessage) (armsecurityinsights.DataConnectorDefinitionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.DataConnectorDefinitionClassification
	switch m["kind"] {
	case string(armsecurityinsights.DataConnectorDefinitionKindCustomizable):
		b = &armsecurityinsights.CustomizableConnectorDefinition{}
	default:
		b = &armsecurityinsights.DataConnectorDefinition{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataConnectorsCheckRequirementsClassification(rawMsg json.RawMessage) (armsecurityinsights.DataConnectorsCheckRequirementsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.DataConnectorsCheckRequirementsClassification
	switch m["kind"] {
	case string(armsecurityinsights.DataConnectorKindAmazonWebServicesCloudTrail):
		b = &armsecurityinsights.AwsCloudTrailCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindAmazonWebServicesS3):
		b = &armsecurityinsights.AwsS3CheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindAzureActiveDirectory):
		b = &armsecurityinsights.AADCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindAzureAdvancedThreatProtection):
		b = &armsecurityinsights.AATPCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindAzureSecurityCenter):
		b = &armsecurityinsights.ASCCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindDynamics365):
		b = &armsecurityinsights.Dynamics365CheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindIOT):
		b = &armsecurityinsights.IoTCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftCloudAppSecurity):
		b = &armsecurityinsights.MCASCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection):
		b = &armsecurityinsights.MDATPCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftPurviewInformationProtection):
		b = &armsecurityinsights.MicrosoftPurviewInformationProtectionCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence):
		b = &armsecurityinsights.MSTICheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindMicrosoftThreatProtection):
		b = &armsecurityinsights.MtpCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindOffice365Project):
		b = &armsecurityinsights.Office365ProjectCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindOfficeATP):
		b = &armsecurityinsights.OfficeATPCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindOfficeIRM):
		b = &armsecurityinsights.OfficeIRMCheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindOfficePowerBI):
		b = &armsecurityinsights.OfficePowerBICheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindThreatIntelligence):
		b = &armsecurityinsights.TICheckRequirements{}
	case string(armsecurityinsights.DataConnectorKindThreatIntelligenceTaxii):
		b = &armsecurityinsights.TiTaxiiCheckRequirements{}
	default:
		b = &armsecurityinsights.DataConnectorsCheckRequirements{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSecurityMLAnalyticsSettingClassification(rawMsg json.RawMessage) (armsecurityinsights.SecurityMLAnalyticsSettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.SecurityMLAnalyticsSettingClassification
	switch m["kind"] {
	case string(armsecurityinsights.SecurityMLAnalyticsSettingsKindAnomaly):
		b = &armsecurityinsights.AnomalySecurityMLAnalyticsSettings{}
	default:
		b = &armsecurityinsights.SecurityMLAnalyticsSetting{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSettingsClassification(rawMsg json.RawMessage) (armsecurityinsights.SettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armsecurityinsights.SettingsClassification
	switch m["kind"] {
	case string(armsecurityinsights.SettingKindAnomalies):
		b = &armsecurityinsights.Anomalies{}
	case string(armsecurityinsights.SettingKindEntityAnalytics):
		b = &armsecurityinsights.EntityAnalytics{}
	case string(armsecurityinsights.SettingKindEyesOn):
		b = &armsecurityinsights.EyesOn{}
	case string(armsecurityinsights.SettingKindUeba):
		b = &armsecurityinsights.Ueba{}
	default:
		b = &armsecurityinsights.Settings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
