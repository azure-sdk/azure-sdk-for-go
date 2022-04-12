package keyvaultapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.2/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	BackupCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.BackupCertificateResult, err error)
	BackupKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.BackupKeyResult, err error)
	BackupSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.BackupSecretResult, err error)
	BackupStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.BackupStorageResult, err error)
	CreateCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateCreateParameters) (result keyvault.CertificateOperation, err error)
	CreateKey(ctx context.Context, vaultBaseURL string, keyName string, parameters keyvault.KeyCreateParameters) (result keyvault.KeyBundle, err error)
	Decrypt(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error)
	DeleteCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.DeletedCertificateBundle, err error)
	DeleteCertificateContacts(ctx context.Context, vaultBaseURL string) (result keyvault.Contacts, err error)
	DeleteCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string) (result keyvault.IssuerBundle, err error)
	DeleteCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateOperation, err error)
	DeleteKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.DeletedKeyBundle, err error)
	DeleteSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.DeletedSasDefinitionBundle, err error)
	DeleteSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.DeletedSecretBundle, err error)
	DeleteStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.DeletedStorageBundle, err error)
	Encrypt(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error)
	FullBackup(ctx context.Context, vaultBaseURL string, azureStorageBlobContainerURI *keyvault.SASTokenParameter) (result keyvault.FullBackupFuture, err error)
	FullBackupStatus(ctx context.Context, vaultBaseURL string, jobID string) (result keyvault.FullBackupOperation, err error)
	FullRestoreOperation(ctx context.Context, vaultBaseURL string, restoreBlobDetails *keyvault.RestoreOperationParameters) (result keyvault.FullRestoreOperationFuture, err error)
	GetCertificate(ctx context.Context, vaultBaseURL string, certificateName string, certificateVersion string) (result keyvault.CertificateBundle, err error)
	GetCertificateContacts(ctx context.Context, vaultBaseURL string) (result keyvault.Contacts, err error)
	GetCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string) (result keyvault.IssuerBundle, err error)
	GetCertificateIssuers(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.CertificateIssuerListResultPage, err error)
	GetCertificateIssuersComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.CertificateIssuerListResultIterator, err error)
	GetCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateOperation, err error)
	GetCertificatePolicy(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificatePolicy, err error)
	GetCertificates(ctx context.Context, vaultBaseURL string, maxresults *int32, includePending *bool) (result keyvault.CertificateListResultPage, err error)
	GetCertificatesComplete(ctx context.Context, vaultBaseURL string, maxresults *int32, includePending *bool) (result keyvault.CertificateListResultIterator, err error)
	GetCertificateVersions(ctx context.Context, vaultBaseURL string, certificateName string, maxresults *int32) (result keyvault.CertificateListResultPage, err error)
	GetCertificateVersionsComplete(ctx context.Context, vaultBaseURL string, certificateName string, maxresults *int32) (result keyvault.CertificateListResultIterator, err error)
	GetDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.DeletedCertificateBundle, err error)
	GetDeletedCertificates(ctx context.Context, vaultBaseURL string, maxresults *int32, includePending *bool) (result keyvault.DeletedCertificateListResultPage, err error)
	GetDeletedCertificatesComplete(ctx context.Context, vaultBaseURL string, maxresults *int32, includePending *bool) (result keyvault.DeletedCertificateListResultIterator, err error)
	GetDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.DeletedKeyBundle, err error)
	GetDeletedKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedKeyListResultPage, err error)
	GetDeletedKeysComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedKeyListResultIterator, err error)
	GetDeletedSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.DeletedSasDefinitionBundle, err error)
	GetDeletedSasDefinitions(ctx context.Context, vaultBaseURL string, storageAccountName string, maxresults *int32) (result keyvault.DeletedSasDefinitionListResultPage, err error)
	GetDeletedSasDefinitionsComplete(ctx context.Context, vaultBaseURL string, storageAccountName string, maxresults *int32) (result keyvault.DeletedSasDefinitionListResultIterator, err error)
	GetDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.DeletedSecretBundle, err error)
	GetDeletedSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedSecretListResultPage, err error)
	GetDeletedSecretsComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedSecretListResultIterator, err error)
	GetDeletedStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.DeletedStorageBundle, err error)
	GetDeletedStorageAccounts(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedStorageListResultPage, err error)
	GetDeletedStorageAccountsComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedStorageListResultIterator, err error)
	GetKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string) (result keyvault.KeyBundle, err error)
	GetKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.KeyListResultPage, err error)
	GetKeysComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.KeyListResultIterator, err error)
	GetKeyVersions(ctx context.Context, vaultBaseURL string, keyName string, maxresults *int32) (result keyvault.KeyListResultPage, err error)
	GetKeyVersionsComplete(ctx context.Context, vaultBaseURL string, keyName string, maxresults *int32) (result keyvault.KeyListResultIterator, err error)
	GetSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.SasDefinitionBundle, err error)
	GetSasDefinitions(ctx context.Context, vaultBaseURL string, storageAccountName string, maxresults *int32) (result keyvault.SasDefinitionListResultPage, err error)
	GetSasDefinitionsComplete(ctx context.Context, vaultBaseURL string, storageAccountName string, maxresults *int32) (result keyvault.SasDefinitionListResultIterator, err error)
	GetSecret(ctx context.Context, vaultBaseURL string, secretName string, secretVersion string) (result keyvault.SecretBundle, err error)
	GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.SecretListResultPage, err error)
	GetSecretsComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.SecretListResultIterator, err error)
	GetSecretVersions(ctx context.Context, vaultBaseURL string, secretName string, maxresults *int32) (result keyvault.SecretListResultPage, err error)
	GetSecretVersionsComplete(ctx context.Context, vaultBaseURL string, secretName string, maxresults *int32) (result keyvault.SecretListResultIterator, err error)
	GetStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.StorageBundle, err error)
	GetStorageAccounts(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.StorageListResultPage, err error)
	GetStorageAccountsComplete(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.StorageListResultIterator, err error)
	ImportCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateImportParameters) (result keyvault.CertificateBundle, err error)
	ImportKey(ctx context.Context, vaultBaseURL string, keyName string, parameters keyvault.KeyImportParameters) (result keyvault.KeyBundle, err error)
	MergeCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateMergeParameters) (result keyvault.CertificateBundle, err error)
	PurgeDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result autorest.Response, err error)
	PurgeDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result autorest.Response, err error)
	PurgeDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result autorest.Response, err error)
	PurgeDeletedStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result autorest.Response, err error)
	RecoverDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateBundle, err error)
	RecoverDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.KeyBundle, err error)
	RecoverDeletedSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.SasDefinitionBundle, err error)
	RecoverDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.SecretBundle, err error)
	RecoverDeletedStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.StorageBundle, err error)
	RegenerateStorageAccountKey(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountRegenerteKeyParameters) (result keyvault.StorageBundle, err error)
	RestoreCertificate(ctx context.Context, vaultBaseURL string, parameters keyvault.CertificateRestoreParameters) (result keyvault.CertificateBundle, err error)
	RestoreKey(ctx context.Context, vaultBaseURL string, parameters keyvault.KeyRestoreParameters) (result keyvault.KeyBundle, err error)
	RestoreSecret(ctx context.Context, vaultBaseURL string, parameters keyvault.SecretRestoreParameters) (result keyvault.SecretBundle, err error)
	RestoreStatus(ctx context.Context, vaultBaseURL string, jobID string) (result keyvault.RestoreOperation, err error)
	RestoreStorageAccount(ctx context.Context, vaultBaseURL string, parameters keyvault.StorageRestoreParameters) (result keyvault.StorageBundle, err error)
	SelectiveKeyRestoreOperationMethod(ctx context.Context, vaultBaseURL string, keyName string, restoreBlobDetails *keyvault.SelectiveKeyRestoreOperationParameters) (result keyvault.SelectiveKeyRestoreOperationMethodFuture, err error)
	SetCertificateContacts(ctx context.Context, vaultBaseURL string, contacts keyvault.Contacts) (result keyvault.Contacts, err error)
	SetCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string, parameter keyvault.CertificateIssuerSetParameters) (result keyvault.IssuerBundle, err error)
	SetSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string, parameters keyvault.SasDefinitionCreateParameters) (result keyvault.SasDefinitionBundle, err error)
	SetSecret(ctx context.Context, vaultBaseURL string, secretName string, parameters keyvault.SecretSetParameters) (result keyvault.SecretBundle, err error)
	SetStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountCreateParameters) (result keyvault.StorageBundle, err error)
	Sign(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeySignParameters) (result keyvault.KeyOperationResult, err error)
	UnwrapKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error)
	UpdateCertificate(ctx context.Context, vaultBaseURL string, certificateName string, certificateVersion string, parameters keyvault.CertificateUpdateParameters) (result keyvault.CertificateBundle, err error)
	UpdateCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string, parameter keyvault.CertificateIssuerUpdateParameters) (result keyvault.IssuerBundle, err error)
	UpdateCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string, certificateOperation keyvault.CertificateOperationUpdateParameter) (result keyvault.CertificateOperation, err error)
	UpdateCertificatePolicy(ctx context.Context, vaultBaseURL string, certificateName string, certificatePolicy keyvault.CertificatePolicy) (result keyvault.CertificatePolicy, err error)
	UpdateKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyUpdateParameters) (result keyvault.KeyBundle, err error)
	UpdateSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string, parameters keyvault.SasDefinitionUpdateParameters) (result keyvault.SasDefinitionBundle, err error)
	UpdateSecret(ctx context.Context, vaultBaseURL string, secretName string, secretVersion string, parameters keyvault.SecretUpdateParameters) (result keyvault.SecretBundle, err error)
	UpdateStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountUpdateParameters) (result keyvault.StorageBundle, err error)
	Verify(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyVerifyParameters) (result keyvault.KeyVerifyResult, err error)
	WrapKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error)
}

var _ BaseClientAPI = (*keyvault.BaseClient)(nil)

// RoleDefinitionsClientAPI contains the set of methods on the RoleDefinitionsClient type.
type RoleDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters keyvault.RoleDefinitionCreateParameters) (result keyvault.RoleDefinition, err error)
	Delete(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (result keyvault.RoleDefinition, err error)
	Get(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (result keyvault.RoleDefinition, err error)
	List(ctx context.Context, vaultBaseURL string, scope string, filter string) (result keyvault.RoleDefinitionListResultPage, err error)
	ListComplete(ctx context.Context, vaultBaseURL string, scope string, filter string) (result keyvault.RoleDefinitionListResultIterator, err error)
}

var _ RoleDefinitionsClientAPI = (*keyvault.RoleDefinitionsClient)(nil)

// RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
type RoleAssignmentsClientAPI interface {
	Create(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string, parameters keyvault.RoleAssignmentCreateParameters) (result keyvault.RoleAssignment, err error)
	Delete(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (result keyvault.RoleAssignment, err error)
	Get(ctx context.Context, vaultBaseURL string, scope string, roleAssignmentName string) (result keyvault.RoleAssignment, err error)
	ListForScope(ctx context.Context, vaultBaseURL string, scope string, filter string) (result keyvault.RoleAssignmentListResultPage, err error)
	ListForScopeComplete(ctx context.Context, vaultBaseURL string, scope string, filter string) (result keyvault.RoleAssignmentListResultIterator, err error)
}

var _ RoleAssignmentsClientAPI = (*keyvault.RoleAssignmentsClient)(nil)

// HSMSecurityDomainClientAPI contains the set of methods on the HSMSecurityDomainClient type.
type HSMSecurityDomainClientAPI interface {
	Download(ctx context.Context, vaultBaseURL string, certificateInfoObject keyvault.CertificateInfoObject) (result keyvault.HSMSecurityDomainDownloadFuture, err error)
	DownloadPending(ctx context.Context, vaultBaseURL string) (result keyvault.SecurityDomainOperationStatus, err error)
	TransferKeyMethod(ctx context.Context, vaultBaseURL string) (result keyvault.TransferKey, err error)
	Upload(ctx context.Context, vaultBaseURL string, securityDomain keyvault.SecurityDomainObject) (result keyvault.HSMSecurityDomainUploadFuture, err error)
	UploadPending(ctx context.Context, vaultBaseURL string) (result keyvault.SecurityDomainOperationStatus, err error)
}

var _ HSMSecurityDomainClientAPI = (*keyvault.HSMSecurityDomainClient)(nil)
