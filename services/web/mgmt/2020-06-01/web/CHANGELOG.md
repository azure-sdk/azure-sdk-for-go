# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. AppServiceCertificateOrderPatchResourceProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason
1. AppServiceCertificateOrderProperties.AppServiceCertificateNotRenewableReasons changed type from *[]string to *[]ResourceNotRenewableReason

## Additive Changes

### New Constants

1. ResourceNotRenewableReason.ExpirationNotInRenewalTimeRange
1. ResourceNotRenewableReason.RegistrationStatusNotSupportedForRenewal
1. ResourceNotRenewableReason.SubscriptionNotActive

### New Funcs

1. PossibleResourceNotRenewableReasonValues() []ResourceNotRenewableReason

### Struct Changes

#### New Struct Fields

1. SiteConfig.AcrUseManagedIdentityCreds
1. SiteConfig.AcrUserManagedIdentityID
