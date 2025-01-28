//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

// BenefitRecommendationPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetBenefitRecommendationProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BenefitRecommendationProperties, *SharedScopeBenefitRecommendationProperties, *SingleScopeBenefitRecommendationProperties
type BenefitRecommendationPropertiesClassification interface {
	// GetBenefitRecommendationProperties returns the BenefitRecommendationProperties content of the underlying type.
	GetBenefitRecommendationProperties() *BenefitRecommendationProperties
}

// BenefitUtilizationSummaryClassification provides polymorphic access to related types.
// Call the interface's GetBenefitUtilizationSummary() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BenefitUtilizationSummary, *IncludedQuantityUtilizationSummary, *SavingsPlanUtilizationSummary
type BenefitUtilizationSummaryClassification interface {
	// GetBenefitUtilizationSummary returns the BenefitUtilizationSummary content of the underlying type.
	GetBenefitUtilizationSummary() *BenefitUtilizationSummary
}

// SettingClassification provides polymorphic access to related types.
// Call the interface's GetSetting() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Setting, *TagInheritanceSetting
type SettingClassification interface {
	// GetSetting returns the Setting content of the underlying type.
	GetSetting() *Setting
}
