//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import "encoding/json"

func unmarshalCustomDomainHTTPSParametersClassification(rawMsg json.RawMessage) (CustomDomainHTTPSParametersClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CustomDomainHTTPSParametersClassification
	switch m["certificateSource"] {
	case string(CertificateSourceAzureKeyVault):
		b = &UserManagedHTTPSParameters{}
	case string(CertificateSourceCdn):
		b = &ManagedHTTPSParameters{}
	default:
		b = &CustomDomainHTTPSParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDeliveryRuleActionClassification(rawMsg json.RawMessage) (DeliveryRuleActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DeliveryRuleActionClassification
	switch m["name"] {
	case string(DeliveryRuleActionNameCacheExpiration):
		b = &DeliveryRuleCacheExpirationAction{}
	case string(DeliveryRuleActionNameCacheKeyQueryString):
		b = &DeliveryRuleCacheKeyQueryStringAction{}
	case string(DeliveryRuleActionNameModifyRequestHeader):
		b = &DeliveryRuleRequestHeaderAction{}
	case string(DeliveryRuleActionNameModifyResponseHeader):
		b = &DeliveryRuleResponseHeaderAction{}
	case string(DeliveryRuleActionNameOriginGroupOverride):
		b = &OriginGroupOverrideAction{}
	case string(DeliveryRuleActionNameRouteConfigurationOverride):
		b = &DeliveryRuleRouteConfigurationOverrideAction{}
	case string(DeliveryRuleActionNameURLRedirect):
		b = &URLRedirectAction{}
	case string(DeliveryRuleActionNameURLRewrite):
		b = &URLRewriteAction{}
	case string(DeliveryRuleActionNameURLSigning):
		b = &URLSigningAction{}
	default:
		b = &DeliveryRuleAction{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDeliveryRuleActionClassificationArray(rawMsg json.RawMessage) ([]DeliveryRuleActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DeliveryRuleActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDeliveryRuleActionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDeliveryRuleConditionClassification(rawMsg json.RawMessage) (DeliveryRuleConditionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DeliveryRuleConditionClassification
	switch m["name"] {
	case string(MatchVariableClientPort):
		b = &DeliveryRuleClientPortCondition{}
	case string(MatchVariableCookies):
		b = &DeliveryRuleCookiesCondition{}
	case string(MatchVariableHostName):
		b = &DeliveryRuleHostNameCondition{}
	case string(MatchVariableHTTPVersion):
		b = &DeliveryRuleHTTPVersionCondition{}
	case string(MatchVariableIsDevice):
		b = &DeliveryRuleIsDeviceCondition{}
	case string(MatchVariablePostArgs):
		b = &DeliveryRulePostArgsCondition{}
	case string(MatchVariableQueryString):
		b = &DeliveryRuleQueryStringCondition{}
	case string(MatchVariableRemoteAddress):
		b = &DeliveryRuleRemoteAddressCondition{}
	case string(MatchVariableRequestBody):
		b = &DeliveryRuleRequestBodyCondition{}
	case string(MatchVariableRequestHeader):
		b = &DeliveryRuleRequestHeaderCondition{}
	case string(MatchVariableRequestMethod):
		b = &DeliveryRuleRequestMethodCondition{}
	case string(MatchVariableRequestScheme):
		b = &DeliveryRuleRequestSchemeCondition{}
	case string(MatchVariableRequestURI):
		b = &DeliveryRuleRequestURICondition{}
	case string(MatchVariableServerPort):
		b = &DeliveryRuleServerPortCondition{}
	case string(MatchVariableSocketAddr):
		b = &DeliveryRuleSocketAddrCondition{}
	case string(MatchVariableSSLProtocol):
		b = &DeliveryRuleSSLProtocolCondition{}
	case string(MatchVariableURLFileExtension):
		b = &DeliveryRuleURLFileExtensionCondition{}
	case string(MatchVariableURLFileName):
		b = &DeliveryRuleURLFileNameCondition{}
	case string(MatchVariableURLPath):
		b = &DeliveryRuleURLPathCondition{}
	default:
		b = &DeliveryRuleCondition{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDeliveryRuleConditionClassificationArray(rawMsg json.RawMessage) ([]DeliveryRuleConditionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DeliveryRuleConditionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDeliveryRuleConditionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSecretParametersClassification(rawMsg json.RawMessage) (SecretParametersClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SecretParametersClassification
	switch m["type"] {
	case string(SecretTypeAzureFirstPartyManagedCertificate):
		b = &AzureFirstPartyManagedCertificateParameters{}
	case string(SecretTypeCustomerCertificate):
		b = &CustomerCertificateParameters{}
	case string(SecretTypeManagedCertificate):
		b = &ManagedCertificateParameters{}
	case string(SecretTypeURLSigningKey):
		b = &URLSigningKeyParameters{}
	default:
		b = &SecretParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSecurityPolicyPropertiesParametersClassification(rawMsg json.RawMessage) (SecurityPolicyPropertiesParametersClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SecurityPolicyPropertiesParametersClassification
	switch m["type"] {
	case string(SecurityPolicyTypeWebApplicationFirewall):
		b = &SecurityPolicyWebApplicationFirewallParameters{}
	default:
		b = &SecurityPolicyPropertiesParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
