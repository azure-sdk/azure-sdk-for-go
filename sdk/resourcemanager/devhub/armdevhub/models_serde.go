//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevhub

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ACR.
func (a ACR) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "acrRegistryName", a.AcrRegistryName)
	populate(objectMap, "acrRepositoryName", a.AcrRepositoryName)
	populate(objectMap, "acrResourceGroup", a.AcrResourceGroup)
	populate(objectMap, "acrSubscriptionId", a.AcrSubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ACR.
func (a *ACR) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "acrRegistryName":
			err = unpopulate(val, "AcrRegistryName", &a.AcrRegistryName)
			delete(rawMsg, key)
		case "acrRepositoryName":
			err = unpopulate(val, "AcrRepositoryName", &a.AcrRepositoryName)
			delete(rawMsg, key)
		case "acrResourceGroup":
			err = unpopulate(val, "AcrResourceGroup", &a.AcrResourceGroup)
			delete(rawMsg, key)
		case "acrSubscriptionId":
			err = unpopulate(val, "AcrSubscriptionID", &a.AcrSubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ArtifactGenerationProperties.
func (a ArtifactGenerationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "appName", a.AppName)
	populate(objectMap, "builderVersion", a.BuilderVersion)
	populate(objectMap, "dockerfileGenerationMode", a.DockerfileGenerationMode)
	populate(objectMap, "dockerfileOutputDirectory", a.DockerfileOutputDirectory)
	populate(objectMap, "generationLanguage", a.GenerationLanguage)
	populate(objectMap, "imageName", a.ImageName)
	populate(objectMap, "imageTag", a.ImageTag)
	populate(objectMap, "languageVersion", a.LanguageVersion)
	populate(objectMap, "manifestGenerationMode", a.ManifestGenerationMode)
	populate(objectMap, "manifestOutputDirectory", a.ManifestOutputDirectory)
	populate(objectMap, "manifestType", a.ManifestType)
	populate(objectMap, "namespace", a.Namespace)
	populate(objectMap, "port", a.Port)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactGenerationProperties.
func (a *ArtifactGenerationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "appName":
			err = unpopulate(val, "AppName", &a.AppName)
			delete(rawMsg, key)
		case "builderVersion":
			err = unpopulate(val, "BuilderVersion", &a.BuilderVersion)
			delete(rawMsg, key)
		case "dockerfileGenerationMode":
			err = unpopulate(val, "DockerfileGenerationMode", &a.DockerfileGenerationMode)
			delete(rawMsg, key)
		case "dockerfileOutputDirectory":
			err = unpopulate(val, "DockerfileOutputDirectory", &a.DockerfileOutputDirectory)
			delete(rawMsg, key)
		case "generationLanguage":
			err = unpopulate(val, "GenerationLanguage", &a.GenerationLanguage)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "ImageName", &a.ImageName)
			delete(rawMsg, key)
		case "imageTag":
			err = unpopulate(val, "ImageTag", &a.ImageTag)
			delete(rawMsg, key)
		case "languageVersion":
			err = unpopulate(val, "LanguageVersion", &a.LanguageVersion)
			delete(rawMsg, key)
		case "manifestGenerationMode":
			err = unpopulate(val, "ManifestGenerationMode", &a.ManifestGenerationMode)
			delete(rawMsg, key)
		case "manifestOutputDirectory":
			err = unpopulate(val, "ManifestOutputDirectory", &a.ManifestOutputDirectory)
			delete(rawMsg, key)
		case "manifestType":
			err = unpopulate(val, "ManifestType", &a.ManifestType)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &a.Namespace)
			delete(rawMsg, key)
		case "port":
			err = unpopulate(val, "Port", &a.Port)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeleteWorkflowResponse.
func (d DeleteWorkflowResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "status", d.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeleteWorkflowResponse.
func (d *DeleteWorkflowResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "status":
			err = unpopulate(val, "Status", &d.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeploymentProperties.
func (d DeploymentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "helmChartPath", d.HelmChartPath)
	populate(objectMap, "helmValues", d.HelmValues)
	populate(objectMap, "kubeManifestLocations", d.KubeManifestLocations)
	populate(objectMap, "manifestType", d.ManifestType)
	populate(objectMap, "overrides", d.Overrides)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentProperties.
func (d *DeploymentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helmChartPath":
			err = unpopulate(val, "HelmChartPath", &d.HelmChartPath)
			delete(rawMsg, key)
		case "helmValues":
			err = unpopulate(val, "HelmValues", &d.HelmValues)
			delete(rawMsg, key)
		case "kubeManifestLocations":
			err = unpopulate(val, "KubeManifestLocations", &d.KubeManifestLocations)
			delete(rawMsg, key)
		case "manifestType":
			err = unpopulate(val, "ManifestType", &d.ManifestType)
			delete(rawMsg, key)
		case "overrides":
			err = unpopulate(val, "Overrides", &d.Overrides)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GeneratePreviewArtifactsResponse.
func (g GeneratePreviewArtifactsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "path/to/file1.withExtension", g.PathToFile1WithExtension)
	populate(objectMap, "path/to/file2.withExtension", g.PathToFile2WithExtension)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GeneratePreviewArtifactsResponse.
func (g *GeneratePreviewArtifactsResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "path/to/file1.withExtension":
			err = unpopulate(val, "PathToFile1WithExtension", &g.PathToFile1WithExtension)
			delete(rawMsg, key)
		case "path/to/file2.withExtension":
			err = unpopulate(val, "PathToFile2WithExtension", &g.PathToFile2WithExtension)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubOAuthCallRequest.
func (g GitHubOAuthCallRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "redirectUrl", g.RedirectURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubOAuthCallRequest.
func (g *GitHubOAuthCallRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "redirectUrl":
			err = unpopulate(val, "RedirectURL", &g.RedirectURL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubOAuthInfoResponse.
func (g GitHubOAuthInfoResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authURL", g.AuthURL)
	populate(objectMap, "token", g.Token)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubOAuthInfoResponse.
func (g *GitHubOAuthInfoResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authURL":
			err = unpopulate(val, "AuthURL", &g.AuthURL)
			delete(rawMsg, key)
		case "token":
			err = unpopulate(val, "Token", &g.Token)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubOAuthListResponse.
func (g GitHubOAuthListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", g.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubOAuthListResponse.
func (g *GitHubOAuthListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &g.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubOAuthProperties.
func (g GitHubOAuthProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "username", g.Username)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubOAuthProperties.
func (g *GitHubOAuthProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "username":
			err = unpopulate(val, "Username", &g.Username)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubOAuthResponse.
func (g GitHubOAuthResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", g.ID)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "properties", g.Properties)
	populate(objectMap, "systemData", g.SystemData)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubOAuthResponse.
func (g *GitHubOAuthResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &g.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &g.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &g.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &g.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &g.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubWorkflowProfile.
func (g GitHubWorkflowProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "acr", g.Acr)
	populate(objectMap, "aksResourceId", g.AksResourceID)
	populate(objectMap, "authStatus", g.AuthStatus)
	populate(objectMap, "branchName", g.BranchName)
	populate(objectMap, "deploymentProperties", g.DeploymentProperties)
	populate(objectMap, "dockerBuildContext", g.DockerBuildContext)
	populate(objectMap, "dockerfile", g.Dockerfile)
	populate(objectMap, "lastWorkflowRun", g.LastWorkflowRun)
	populate(objectMap, "namespace", g.Namespace)
	populate(objectMap, "oidcCredentials", g.OidcCredentials)
	populate(objectMap, "prStatus", g.PrStatus)
	populate(objectMap, "prURL", g.PrURL)
	populate(objectMap, "pullNumber", g.PullNumber)
	populate(objectMap, "repositoryName", g.RepositoryName)
	populate(objectMap, "repositoryOwner", g.RepositoryOwner)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubWorkflowProfile.
func (g *GitHubWorkflowProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "acr":
			err = unpopulate(val, "Acr", &g.Acr)
			delete(rawMsg, key)
		case "aksResourceId":
			err = unpopulate(val, "AksResourceID", &g.AksResourceID)
			delete(rawMsg, key)
		case "authStatus":
			err = unpopulate(val, "AuthStatus", &g.AuthStatus)
			delete(rawMsg, key)
		case "branchName":
			err = unpopulate(val, "BranchName", &g.BranchName)
			delete(rawMsg, key)
		case "deploymentProperties":
			err = unpopulate(val, "DeploymentProperties", &g.DeploymentProperties)
			delete(rawMsg, key)
		case "dockerBuildContext":
			err = unpopulate(val, "DockerBuildContext", &g.DockerBuildContext)
			delete(rawMsg, key)
		case "dockerfile":
			err = unpopulate(val, "Dockerfile", &g.Dockerfile)
			delete(rawMsg, key)
		case "lastWorkflowRun":
			err = unpopulate(val, "LastWorkflowRun", &g.LastWorkflowRun)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &g.Namespace)
			delete(rawMsg, key)
		case "oidcCredentials":
			err = unpopulate(val, "OidcCredentials", &g.OidcCredentials)
			delete(rawMsg, key)
		case "prStatus":
			err = unpopulate(val, "PrStatus", &g.PrStatus)
			delete(rawMsg, key)
		case "prURL":
			err = unpopulate(val, "PrURL", &g.PrURL)
			delete(rawMsg, key)
		case "pullNumber":
			err = unpopulate(val, "PullNumber", &g.PullNumber)
			delete(rawMsg, key)
		case "repositoryName":
			err = unpopulate(val, "RepositoryName", &g.RepositoryName)
			delete(rawMsg, key)
		case "repositoryOwner":
			err = unpopulate(val, "RepositoryOwner", &g.RepositoryOwner)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GitHubWorkflowProfileOidcCredentials.
func (g GitHubWorkflowProfileOidcCredentials) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureClientId", g.AzureClientID)
	populate(objectMap, "azureTenantId", g.AzureTenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitHubWorkflowProfileOidcCredentials.
func (g *GitHubWorkflowProfileOidcCredentials) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureClientId":
			err = unpopulate(val, "AzureClientID", &g.AzureClientID)
			delete(rawMsg, key)
		case "azureTenantId":
			err = unpopulate(val, "AzureTenantID", &g.AzureTenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
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

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagsObject.
func (t *TagsObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Workflow.
func (w Workflow) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Workflow.
func (w *Workflow) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type WorkflowListResult.
func (w WorkflowListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkflowListResult.
func (w *WorkflowListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type WorkflowProperties.
func (w WorkflowProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "artifactGenerationProperties", w.ArtifactGenerationProperties)
	populate(objectMap, "githubWorkflowProfile", w.GithubWorkflowProfile)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkflowProperties.
func (w *WorkflowProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactGenerationProperties":
			err = unpopulate(val, "ArtifactGenerationProperties", &w.ArtifactGenerationProperties)
			delete(rawMsg, key)
		case "githubWorkflowProfile":
			err = unpopulate(val, "GithubWorkflowProfile", &w.GithubWorkflowProfile)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkflowRun.
func (w WorkflowRun) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "lastRunAt", w.LastRunAt)
	populate(objectMap, "succeeded", w.Succeeded)
	populate(objectMap, "workflowRunURL", w.WorkflowRunURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkflowRun.
func (w *WorkflowRun) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastRunAt":
			err = unpopulateTimeRFC3339(val, "LastRunAt", &w.LastRunAt)
			delete(rawMsg, key)
		case "succeeded":
			err = unpopulate(val, "Succeeded", &w.Succeeded)
			delete(rawMsg, key)
		case "workflowRunURL":
			err = unpopulate(val, "WorkflowRunURL", &w.WorkflowRunURL)
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
