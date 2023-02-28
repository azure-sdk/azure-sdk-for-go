# Release History

## 0.3.0 (2023-02-28)
### Breaking Changes

- Function `NewDeveloperHubServiceClient` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, string, string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*DeveloperHubServiceClient.GitHubOAuthCallback` parameter(s) have been changed from `(context.Context, string, string, string, *DeveloperHubServiceClientGitHubOAuthCallbackOptions)` to `(context.Context, string, *DeveloperHubServiceClientGitHubOAuthCallbackOptions)`
- Function `NewWorkflowClient` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, *string, azcore.TokenCredential, *arm.ClientOptions)`
- Type of `GitHubWorkflowProfile.AuthStatus` has been changed from `*ManifestType` to `*AuthorizationStatus`
- Field `ManagedClusterResource` of struct `WorkflowClientListByResourceGroupOptions` has been removed

### Features Added

- New type alias `AuthorizationStatus` with values `AuthorizationStatusAuthorized`, `AuthorizationStatusError`, `AuthorizationStatusNotFound`
- New type alias `DockerfileGenerationMode` with values `DockerfileGenerationModeDisabled`, `DockerfileGenerationModeEnabled`
- New type alias `GenerationLanguage` with values `GenerationLanguageClojure`, `GenerationLanguageCsharp`, `GenerationLanguageErlang`, `GenerationLanguageGo`, `GenerationLanguageGomodule`, `GenerationLanguageGradle`, `GenerationLanguageJava`, `GenerationLanguageJavascript`, `GenerationLanguagePhp`, `GenerationLanguagePython`, `GenerationLanguageRuby`, `GenerationLanguageRust`, `GenerationLanguageSwift`
- New type alias `GenerationManifestType` with values `GenerationManifestTypeHelm`, `GenerationManifestTypeKube`
- New type alias `ManifestGenerationMode` with values `ManifestGenerationModeDisabled`, `ManifestGenerationModeEnabled`
- New function `*DeveloperHubServiceClient.GeneratePreviewArtifacts(context.Context, string, *DeveloperHubServiceClientGeneratePreviewArtifactsOptions) (DeveloperHubServiceClientGeneratePreviewArtifactsResponse, error)`
- New struct `ArtifactGenerationProperties`
- New struct `GeneratePreviewArtifactsResponse`
- New field `ArtifactGenerationProperties` in struct `WorkflowProperties`


## 0.2.0 (2022-10-13)
### Breaking Changes

- Function `NewWorkflowClient` parameter(s) have been changed from `(string, *string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `NewDeveloperHubServiceClient` parameter(s) have been changed from `(string, string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*DeveloperHubServiceClient.GitHubOAuthCallback` parameter(s) have been changed from `(context.Context, string, *DeveloperHubServiceClientGitHubOAuthCallbackOptions)` to `(context.Context, string, string, string, *DeveloperHubServiceClientGitHubOAuthCallbackOptions)`

### Features Added

- New field `ManagedClusterResource` in struct `WorkflowClientListByResourceGroupOptions`


## 0.1.1 (2022-10-12)
### Other Changes
- Loosen Go version requirement.

## 0.1.0 (2022-09-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).