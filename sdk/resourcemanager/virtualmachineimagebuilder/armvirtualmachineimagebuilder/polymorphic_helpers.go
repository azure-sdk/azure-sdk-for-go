// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import "encoding/json"

func unmarshalDistributeVersionerClassification(rawMsg json.RawMessage) (DistributeVersionerClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DistributeVersionerClassification
	switch m["scheme"] {
	case "Latest":
		b = &DistributeVersionerLatest{}
	case "Source":
		b = &DistributeVersionerSource{}
	default:
		b = &DistributeVersioner{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalImageTemplateCustomizerClassification(rawMsg json.RawMessage) (ImageTemplateCustomizerClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateCustomizerClassification
	switch m["type"] {
	case "File":
		b = &ImageTemplateFileCustomizer{}
	case "PowerShell":
		b = &ImageTemplatePowerShellCustomizer{}
	case "Shell":
		b = &ImageTemplateShellCustomizer{}
	case "WindowsRestart":
		b = &ImageTemplateRestartCustomizer{}
	case "WindowsUpdate":
		b = &ImageTemplateWindowsUpdateCustomizer{}
	default:
		b = &ImageTemplateCustomizer{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalImageTemplateCustomizerClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateCustomizerClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateCustomizerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateCustomizerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateDistributorClassification(rawMsg json.RawMessage) (ImageTemplateDistributorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateDistributorClassification
	switch m["type"] {
	case "ManagedImage":
		b = &ImageTemplateManagedImageDistributor{}
	case "SharedImage":
		b = &ImageTemplateSharedImageDistributor{}
	case "VHD":
		b = &ImageTemplateVhdDistributor{}
	default:
		b = &ImageTemplateDistributor{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalImageTemplateDistributorClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateDistributorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateDistributorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateDistributorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateInVMValidatorClassification(rawMsg json.RawMessage) (ImageTemplateInVMValidatorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateInVMValidatorClassification
	switch m["type"] {
	case "File":
		b = &ImageTemplateFileValidator{}
	case "PowerShell":
		b = &ImageTemplatePowerShellValidator{}
	case "Shell":
		b = &ImageTemplateShellValidator{}
	default:
		b = &ImageTemplateInVMValidator{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalImageTemplateInVMValidatorClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateInVMValidatorClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateInVMValidatorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateInVMValidatorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateSourceClassification(rawMsg json.RawMessage) (ImageTemplateSourceClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateSourceClassification
	switch m["type"] {
	case "ManagedImage":
		b = &ImageTemplateManagedImageSource{}
	case "PlatformImage":
		b = &ImageTemplatePlatformImageSource{}
	case "SharedImageVersion":
		b = &ImageTemplateSharedImageVersionSource{}
	default:
		b = &ImageTemplateSource{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTriggerPropertiesClassification(rawMsg json.RawMessage) (TriggerPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerPropertiesClassification
	switch m["kind"] {
	case "SourceImage":
		b = &SourceImageTriggerProperties{}
	default:
		b = &TriggerProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
