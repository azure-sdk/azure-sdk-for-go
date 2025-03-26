// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package azcertificates

import "time"

// AdministratorContact - Details of the organization administrator of the certificate issuer.
type AdministratorContact struct {
	// Email address.
	Email *string

	// First name.
	FirstName *string

	// Last name.
	LastName *string

	// Phone number.
	Phone *string
}

// BackupCertificateResult - The backup certificate result, containing the backup blob.
type BackupCertificateResult struct {
	// READ-ONLY; The backup blob containing the backed up certificate.
	Value []byte
}

// Certificate - A certificate bundle consists of a certificate (X509) plus its attributes.
type Certificate struct {
	// The certificate attributes.
	Attributes *CertificateAttributes

	// CER contents of x509 certificate.
	CER []byte

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string

	// Specifies whether the certificate chain preserves its original order. The default value is false, which sets the leaf certificate
	// at index 0.
	PreserveCertOrder *bool

	// Application specific metadata in the form of key-value pairs
	Tags map[string]*string

	// READ-ONLY; The certificate id.
	ID *string

	// READ-ONLY; The key id.
	KID *string

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy

	// READ-ONLY; The secret id.
	SID *string

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte
}

// CertificateAttributes - The certificate management attributes.
type CertificateAttributes struct {
	// Determines whether the object is enabled.
	Enabled *bool

	// Expiry date in UTC.
	Expires *time.Time

	// Not before date in UTC.
	NotBefore *time.Time

	// READ-ONLY; Creation time in UTC.
	Created *time.Time

	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32

	// READ-ONLY; Reflects the deletion recovery level currently in effect for certificates in the current vault. If it contains
	// 'Purgeable', the certificate can be permanently deleted by a privileged user; otherwise, only the system can purge the
	// certificate, at the end of the retention interval.
	RecoveryLevel *DeletionRecoveryLevel

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time
}

// CertificateOperation - A certificate operation is returned in case of asynchronous requests.
type CertificateOperation struct {
	// The certificate signing request (CSR) that is being used in the certificate operation.
	CSR []byte

	// Indicates if cancellation was requested on the certificate operation.
	CancellationRequested *bool

	// Error encountered, if any, during the certificate operation.
	Error *KeyVaultErrorError

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters

	// Specifies whether the certificate chain preserves its original order. The default value is false, which sets the leaf certificate
	// at index 0.
	PreserveCertOrder *bool

	// Identifier for the certificate operation.
	RequestID *string

	// Status of the certificate operation.
	Status *string

	// The status details of the certificate operation.
	StatusDetails *string

	// Location which contains the result of the certificate operation.
	Target *string

	// READ-ONLY; The certificate id.
	ID *string
}

// CertificatePolicy - Management policy for a certificate.
type CertificatePolicy struct {
	// The certificate attributes.
	Attributes *CertificateAttributes

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters

	// Properties of the key backing a certificate.
	KeyProperties *KeyProperties

	// Actions that will be performed by Key Vault over the lifetime of a certificate.
	LifetimeActions []*LifetimeAction

	// Properties of the secret backing a certificate.
	SecretProperties *SecretProperties

	// Properties of the X509 component of a certificate.
	X509CertificateProperties *X509CertificateProperties

	// READ-ONLY; The certificate id.
	ID *string
}

// CertificateProperties - The certificate item containing certificate metadata.
type CertificateProperties struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes

	// Certificate identifier.
	ID *string

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// Thumbprint of the certificate.
	X509Thumbprint []byte
}

// CertificatePropertiesListResult - The certificate list result.
type CertificatePropertiesListResult struct {
	// A response message containing a list of certificates in the key vault along with a link to the next page of certificates.
	Value []*CertificateProperties

	// READ-ONLY; The URL to get the next set of certificates.
	NextLink *string
}

// Contact - The contact information for the vault certificates.
type Contact struct {
	// Email address.
	Email *string

	// Name.
	Name *string

	// Phone number.
	Phone *string
}

// Contacts - The contacts for the vault certificates.
type Contacts struct {
	// The contact list for the vault certificates.
	ContactList []*Contact

	// READ-ONLY; Identifier for the contacts collection.
	ID *string
}

// CreateCertificateParameters - The certificate create parameters.
type CreateCertificateParameters struct {
	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy

	// Specifies whether the certificate chain preserves its original order. The default value is false, which sets the leaf certificate
	// at index 0.
	PreserveCertOrder *bool

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// DeletedCertificate - A Deleted Certificate consisting of its previous id, attributes and its tags, as well as information
// on when it will be purged.
type DeletedCertificate struct {
	// The certificate attributes.
	Attributes *CertificateAttributes

	// CER contents of x509 certificate.
	CER []byte

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string

	// Specifies whether the certificate chain preserves its original order. The default value is false, which sets the leaf certificate
	// at index 0.
	PreserveCertOrder *bool

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string

	// Application specific metadata in the form of key-value pairs
	Tags map[string]*string

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time

	// READ-ONLY; The certificate id.
	ID *string

	// READ-ONLY; The key id.
	KID *string

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy

	// READ-ONLY; The secret id.
	SID *string

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte
}

// DeletedCertificateProperties - The deleted certificate item containing metadata about the deleted certificate.
type DeletedCertificateProperties struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes

	// Certificate identifier.
	ID *string

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// Thumbprint of the certificate.
	X509Thumbprint []byte

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time
}

// DeletedCertificatePropertiesListResult - A list of certificates that have been deleted in this vault.
type DeletedCertificatePropertiesListResult struct {
	// READ-ONLY; The URL to get the next set of deleted certificates.
	NextLink *string

	// READ-ONLY; A response message containing a list of deleted certificates in the vault along with a link to the next page
	// of deleted certificates.
	Value []*DeletedCertificateProperties
}

// ImportCertificateParameters - The certificate import parameters.
type ImportCertificateParameters struct {
	// REQUIRED; Base64 encoded representation of the certificate object to import. This certificate needs to contain the private
	// key.
	Base64EncodedCertificate *string

	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy

	// If the private key in base64EncodedCertificate is encrypted, the password used for encryption.
	Password *string

	// Specifies whether the certificate chain preserves its original order. The default value is false, which sets the leaf certificate
	// at index 0.
	PreserveCertOrder *bool

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// Issuer - The issuer for Key Vault certificate.
type Issuer struct {
	// Attributes of the issuer object.
	Attributes *IssuerAttributes

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails

	// The issuer provider.
	Provider *string

	// READ-ONLY; Identifier for the issuer object.
	ID *string
}

// IssuerAttributes - The attributes of an issuer managed by the Key Vault service.
type IssuerAttributes struct {
	// Determines whether the issuer is enabled.
	Enabled *bool

	// READ-ONLY; Creation time in UTC.
	Created *time.Time

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time
}

// IssuerCredentials - The credentials to be used for the certificate issuer.
type IssuerCredentials struct {
	// The user name/account name/account id.
	AccountID *string

	// The password/secret/account key.
	Password *string
}

// IssuerParameters - Parameters for the issuer of the X509 component of a certificate.
type IssuerParameters struct {
	// Indicates if the certificates generated under this policy should be published to certificate transparency logs.
	CertificateTransparency *bool

	// Certificate type as supported by the provider (optional); for example 'OV-SSL', 'EV-SSL'
	CertificateType *string

	// Name of the referenced issuer object or reserved names; for example, 'Self' or 'Unknown'.
	Name *string
}

// IssuerProperties - The certificate issuer item containing certificate issuer metadata.
type IssuerProperties struct {
	// Certificate Identifier.
	ID *string

	// The issuer provider.
	Provider *string
}

// IssuerPropertiesListResult - The certificate issuer list result.
type IssuerPropertiesListResult struct {
	// READ-ONLY; The URL to get the next set of certificate issuers.
	NextLink *string

	// READ-ONLY; A response message containing a list of certificate issuers in the key vault along with a link to the next page
	// of certificate issuers.
	Value []*IssuerProperties
}

// KeyProperties - Properties of the key pair backing a certificate.
type KeyProperties struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Curve *CurveName

	// Indicates if the private key can be exported. Release policy must be provided when creating the first version of an exportable
	// key.
	Exportable *bool

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32

	// The type of key pair to be used for the certificate.
	KeyType *KeyType

	// Indicates if the same key pair will be used on certificate renewal.
	ReuseKey *bool
}

type KeyVaultErrorError struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The key vault server error.
	InnerError *KeyVaultErrorError

	// READ-ONLY; The error message.
	Message *string
}

// LifetimeAction - Action and its trigger that will be performed by Key Vault over the lifetime of a certificate.
type LifetimeAction struct {
	// The action that will be executed.
	Action *LifetimeActionType

	// The condition that will execute the action.
	Trigger *LifetimeActionTrigger
}

// LifetimeActionTrigger - A condition to be satisfied for an action to be executed.
type LifetimeActionTrigger struct {
	// Days before expiry to attempt renewal. Value should be between 1 and validity_in_months multiplied by 27. If validity_in_months
	// is 36, then value should be between 1 and 972 (36 * 27).
	DaysBeforeExpiry *int32

	// Percentage of lifetime at which to trigger. Value should be between 1 and 99.
	LifetimePercentage *int32
}

// LifetimeActionType - The action that will be executed.
type LifetimeActionType struct {
	// The type of the action.
	ActionType *CertificatePolicyAction
}

// MergeCertificateParameters - The certificate merge parameters
type MergeCertificateParameters struct {
	// REQUIRED; The certificate or the certificate chain to merge.
	X509Certificates [][]byte

	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// OrganizationDetails - Details of the organization of the certificate issuer.
type OrganizationDetails struct {
	// Details of the organization administrator.
	AdminContacts []*AdministratorContact

	// Id of the organization.
	ID *string
}

// RestoreCertificateParameters - The certificate restore parameters.
type RestoreCertificateParameters struct {
	// REQUIRED; The backup blob associated with a certificate bundle.
	CertificateBackup []byte
}

// SecretProperties - Properties of the key backing a certificate.
type SecretProperties struct {
	// The media type (MIME type).
	ContentType *string
}

// SetIssuerParameters - The certificate issuer set parameters.
type SetIssuerParameters struct {
	// REQUIRED; The issuer provider.
	Provider *string

	// Attributes of the issuer object.
	Attributes *IssuerAttributes

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails
}

// SubjectAlternativeNames - The subject alternate names of a X509 object.
type SubjectAlternativeNames struct {
	// Domain names.
	DNSNames []*string

	// Email addresses.
	Emails []*string

	// User principal names.
	UserPrincipalNames []*string
}

// UpdateCertificateOperationParameter - The certificate operation update parameters.
type UpdateCertificateOperationParameter struct {
	// REQUIRED; Indicates if cancellation was requested on the certificate operation.
	CancellationRequested *bool
}

// UpdateCertificateParameters - The certificate update parameters.
type UpdateCertificateParameters struct {
	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// UpdateIssuerParameters - The certificate issuer update parameters.
type UpdateIssuerParameters struct {
	// Attributes of the issuer object.
	Attributes *IssuerAttributes

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails

	// The issuer provider.
	Provider *string
}

// X509CertificateProperties - Properties of the X509 component of a certificate.
type X509CertificateProperties struct {
	// The enhanced key usage.
	EnhancedKeyUsage []*string

	// Defines how the certificate's key may be used.
	KeyUsage []*KeyUsageType

	// The subject name. Should be a valid X509 distinguished Name.
	Subject *string

	// The subject alternative names.
	SubjectAlternativeNames *SubjectAlternativeNames

	// The duration that the certificate is valid in months.
	ValidityInMonths *int32
}
