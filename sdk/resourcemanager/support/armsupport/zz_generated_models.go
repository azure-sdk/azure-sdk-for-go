//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import "time"

// CheckNameAvailabilityInput - Input of CheckNameAvailability API.
type CheckNameAvailabilityInput struct {
	// REQUIRED; The resource name to validate.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The type of resource.
	Type *Type `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput - Output of check name availability API.
type CheckNameAvailabilityOutput struct {
	// READ-ONLY; The detailed error message describing why the name is not available.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Indicates whether the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; The reason why the name is not available.
	Reason *string `json:"reason,omitempty" azure:"ro"`
}

// CommunicationDetails - Object that represents a Communication resource.
type CommunicationDetails struct {
	// Properties of the resource.
	Properties *CommunicationDetailsProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource 'Microsoft.Support/communications'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CommunicationDetailsProperties - Describes the properties of a communication resource.
type CommunicationDetailsProperties struct {
	// REQUIRED; Body of the communication.
	Body *string `json:"body,omitempty"`

	// REQUIRED; Subject of the communication.
	Subject *string `json:"subject,omitempty"`

	// Email address of the sender. This property is required if called by a service principal.
	Sender *string `json:"sender,omitempty"`

	// READ-ONLY; Direction of communication.
	CommunicationDirection *CommunicationDirection `json:"communicationDirection,omitempty" azure:"ro"`

	// READ-ONLY; Communication type.
	CommunicationType *CommunicationType `json:"communicationType,omitempty" azure:"ro"`

	// READ-ONLY; Time in UTC (ISO 8601 format) when the communication was created.
	CreatedDate *time.Time `json:"createdDate,omitempty" azure:"ro"`
}

// CommunicationsClientBeginCreateOptions contains the optional parameters for the CommunicationsClient.BeginCreate method.
type CommunicationsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommunicationsClientCheckNameAvailabilityOptions contains the optional parameters for the CommunicationsClient.CheckNameAvailability
// method.
type CommunicationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsClientGetOptions contains the optional parameters for the CommunicationsClient.Get method.
type CommunicationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsClientListOptions contains the optional parameters for the CommunicationsClient.List method.
type CommunicationsClientListOptions struct {
	// The filter to apply on the operation. You can filter by communicationType and createdDate properties. CommunicationType
	// supports Equals ('eq') operator and createdDate supports Greater Than ('gt') and
	// Greater Than or Equals ('ge') operators. You may combine the CommunicationType and CreatedDate filters by Logical And ('and')
	// operator.
	Filter *string
	// The number of values to return in the collection. Default is 10 and max is 10.
	Top *int32
}

// CommunicationsListResult - Collection of Communication resources.
type CommunicationsListResult struct {
	// The URI to fetch the next page of Communication resources.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Communication resources.
	Value []*CommunicationDetails `json:"value,omitempty"`
}

// ContactProfile - Contact information associated with the support ticket.
type ContactProfile struct {
	// REQUIRED; Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string `json:"country,omitempty"`

	// REQUIRED; First name.
	FirstName *string `json:"firstName,omitempty"`

	// REQUIRED; Last name.
	LastName *string `json:"lastName,omitempty"`

	// REQUIRED; Preferred contact method.
	PreferredContactMethod *PreferredContactMethod `json:"preferredContactMethod,omitempty"`

	// REQUIRED; Preferred language of support from Azure. Support languages vary based on the severity you choose for your support
	// ticket. Learn more at Azure Severity and responsiveness
	// [https://azure.microsoft.com/support/plans/response]. Use the standard language-country code. Valid values are 'en-us'
	// for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French,
	// 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for
	// Chinese and 'de-de' for German.
	PreferredSupportLanguage *string `json:"preferredSupportLanguage,omitempty"`

	// REQUIRED; Time zone of the user. This is the name of the time zone from Microsoft Time Zone Index Values [https://support.microsoft.com/help/973627/microsoft-time-zone-index-values].
	PreferredTimeZone *string `json:"preferredTimeZone,omitempty"`

	// REQUIRED; Primary email address.
	PrimaryEmailAddress *string `json:"primaryEmailAddress,omitempty"`

	// Additional email addresses listed will be copied on any correspondence about the support ticket.
	AdditionalEmailAddresses []*string `json:"additionalEmailAddresses,omitempty"`

	// Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// Engineer - Support engineer information.
type Engineer struct {
	// READ-ONLY; Email address of the Azure Support engineer assigned to the support ticket.
	EmailAddress *string `json:"emailAddress,omitempty" azure:"ro"`
}

// ExceptionResponse - The API error.
type ExceptionResponse struct {
	// The API error details.
	Error *ServiceError `json:"error,omitempty"`
}

// Operation - The operation supported by Microsoft Support resource provider.
type Operation struct {
	// The object that describes the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationDisplay - The object that describes the operation.
type OperationDisplay struct {
	// READ-ONLY; The description of the operation.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Service provider: Microsoft Support.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsListResult - The list of operations supported by Microsoft Support resource provider.
type OperationsListResult struct {
	// The list of operations supported by Microsoft Support resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// ProblemClassification resource object.
type ProblemClassification struct {
	// Properties of the resource.
	Properties *ProblemClassificationProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource 'Microsoft.Support/problemClassification'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ProblemClassificationProperties - Details about a problem classification available for an Azure service.
type ProblemClassificationProperties struct {
	// Localized name of problem classification.
	DisplayName *string `json:"displayName,omitempty"`
}

// ProblemClassificationsClientGetOptions contains the optional parameters for the ProblemClassificationsClient.Get method.
type ProblemClassificationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProblemClassificationsClientListOptions contains the optional parameters for the ProblemClassificationsClient.List method.
type ProblemClassificationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProblemClassificationsListResult - Collection of ProblemClassification resources.
type ProblemClassificationsListResult struct {
	// List of ProblemClassification resources.
	Value []*ProblemClassification `json:"value,omitempty"`
}

// QuotaChangeRequest - This property is required for providing the region and new quota limits.
type QuotaChangeRequest struct {
	// Payload of the quota increase request.
	Payload *string `json:"payload,omitempty"`

	// Region for which the quota increase request is being made.
	Region *string `json:"region,omitempty"`
}

// QuotaTicketDetails - Additional set of information required for quota increase support ticket for certain quota types,
// e.g.: Virtual machine cores. Get complete details about Quota payload support request along with
// examples at Support quota request [https://aka.ms/supportrpquotarequestpayload].
type QuotaTicketDetails struct {
	// Required for certain quota types when there is a sub type, such as Batch, for which you are requesting a quota increase.
	QuotaChangeRequestSubType *string `json:"quotaChangeRequestSubType,omitempty"`

	// Quota change request version.
	QuotaChangeRequestVersion *string `json:"quotaChangeRequestVersion,omitempty"`

	// This property is required for providing the region and new quota limits.
	QuotaChangeRequests []*QuotaChangeRequest `json:"quotaChangeRequests,omitempty"`
}

// Service - Object that represents a Service resource.
type Service struct {
	// Properties of the resource.
	Properties *ServiceProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource 'Microsoft.Support/services'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServiceError - The API error details.
type ServiceError struct {
	// The error code.
	Code *string `json:"code,omitempty"`

	// The error message.
	Message *string `json:"message,omitempty"`

	// The target of the error.
	Target *string `json:"target,omitempty"`

	// READ-ONLY; The list of error details.
	Details []*ServiceErrorDetail `json:"details,omitempty" azure:"ro"`
}

// ServiceErrorDetail - The error details.
type ServiceErrorDetail struct {
	// The target of the error.
	Target *string `json:"target,omitempty"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ServiceLevelAgreement - Service Level Agreement details for a support ticket.
type ServiceLevelAgreement struct {
	// READ-ONLY; Time in UTC (ISO 8601 format) when the service level agreement expires.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" azure:"ro"`

	// READ-ONLY; Service Level Agreement in minutes.
	SLAMinutes *int32 `json:"slaMinutes,omitempty" azure:"ro"`

	// READ-ONLY; Time in UTC (ISO 8601 format) when the service level agreement starts.
	StartTime *time.Time `json:"startTime,omitempty" azure:"ro"`
}

// ServiceProperties - Details about an Azure service available for support ticket creation.
type ServiceProperties struct {
	// Localized name of the Azure service.
	DisplayName *string `json:"displayName,omitempty"`

	// ARM Resource types.
	ResourceTypes []*string `json:"resourceTypes,omitempty"`
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListOptions contains the optional parameters for the ServicesClient.List method.
type ServicesClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesListResult - Collection of Service resources.
type ServicesListResult struct {
	// List of Service resources.
	Value []*Service `json:"value,omitempty"`
}

// TechnicalTicketDetails - Additional information for technical support ticket.
type TechnicalTicketDetails struct {
	// This is the resource Id of the Azure service resource (For example: A virtual machine resource or an HDInsight resource)
	// for which the support ticket is created.
	ResourceID *string `json:"resourceId,omitempty"`
}

// TicketDetails - Object that represents SupportTicketDetails resource.
type TicketDetails struct {
	// Properties of the resource.
	Properties *TicketDetailsProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource 'Microsoft.Support/supportTickets'.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TicketDetailsProperties - Describes the properties of a support ticket.
type TicketDetailsProperties struct {
	// REQUIRED; Contact information of the user requesting to create a support ticket.
	ContactDetails *ContactProfile `json:"contactDetails,omitempty"`

	// REQUIRED; Detailed description of the question or issue.
	Description *string `json:"description,omitempty"`

	// REQUIRED; Each Azure service has its own set of issue categories, also known as problem classification. This parameter
	// is the unique Id for the type of problem you are experiencing.
	ProblemClassificationID *string `json:"problemClassificationId,omitempty"`

	// REQUIRED; This is the resource Id of the Azure service resource associated with the support ticket.
	ServiceID *string `json:"serviceId,omitempty"`

	// REQUIRED; A value that indicates the urgency of the case, which in turn determines the response time according to the service
	// level agreement of the technical support plan you have with Azure. Note: 'Highest
	// critical impact', also known as the 'Emergency - Severe impact' level in the Azure portal is reserved only for our Premium
	// customers.
	Severity *SeverityLevel `json:"severity,omitempty"`

	// REQUIRED; Title of the support ticket.
	Title *string `json:"title,omitempty"`

	// Time in UTC (ISO 8601 format) when the problem started.
	ProblemStartTime *time.Time `json:"problemStartTime,omitempty"`

	// Additional ticket details associated with a quota support ticket request.
	QuotaTicketDetails *QuotaTicketDetails `json:"quotaTicketDetails,omitempty"`

	// Indicates if this requires a 24x7 response from Azure.
	Require24X7Response *bool `json:"require24X7Response,omitempty"`

	// Service Level Agreement information for this support ticket.
	ServiceLevelAgreement *ServiceLevelAgreement `json:"serviceLevelAgreement,omitempty"`

	// Information about the support engineer working on this support ticket.
	SupportEngineer *Engineer `json:"supportEngineer,omitempty"`

	// System generated support ticket Id that is unique.
	SupportTicketID *string `json:"supportTicketId,omitempty"`

	// Additional ticket details associated with a technical support ticket request.
	TechnicalTicketDetails *TechnicalTicketDetails `json:"technicalTicketDetails,omitempty"`

	// READ-ONLY; Time in UTC (ISO 8601 format) when the support ticket was created.
	CreatedDate *time.Time `json:"createdDate,omitempty" azure:"ro"`

	// READ-ONLY; Enrollment Id associated with the support ticket.
	EnrollmentID *string `json:"enrollmentId,omitempty" azure:"ro"`

	// READ-ONLY; Time in UTC (ISO 8601 format) when the support ticket was last modified.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty" azure:"ro"`

	// READ-ONLY; Localized name of problem classification.
	ProblemClassificationDisplayName *string `json:"problemClassificationDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Localized name of the Azure service.
	ServiceDisplayName *string `json:"serviceDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; Status of the support ticket.
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; Support plan type associated with the support ticket.
	SupportPlanType *string `json:"supportPlanType,omitempty" azure:"ro"`
}

// TicketsClientBeginCreateOptions contains the optional parameters for the TicketsClient.BeginCreate method.
type TicketsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TicketsClientCheckNameAvailabilityOptions contains the optional parameters for the TicketsClient.CheckNameAvailability
// method.
type TicketsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// TicketsClientGetOptions contains the optional parameters for the TicketsClient.Get method.
type TicketsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TicketsClientListOptions contains the optional parameters for the TicketsClient.List method.
type TicketsClientListOptions struct {
	// The filter to apply on the operation. We support 'odata v4.0' filter semantics. Learn more [https://docs.microsoft.com/odata/concepts/queryoptions-overview].
	// Status, ServiceId, and
	// ProblemClassificationId filters can only be used with Equals ('eq') operator. For CreatedDate filter, the supported operators
	// are Greater Than ('gt') and Greater Than or Equals ('ge'). When using both
	// filters, combine them using the logical 'AND'.
	Filter *string
	// The number of values to return in the collection. Default is 25 and max is 100.
	Top *int32
}

// TicketsClientUpdateOptions contains the optional parameters for the TicketsClient.Update method.
type TicketsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// TicketsListResult - Object that represents a collection of SupportTicket resources.
type TicketsListResult struct {
	// The URI to fetch the next page of SupportTicket resources.
	NextLink *string `json:"nextLink,omitempty"`

	// List of SupportTicket resources.
	Value []*TicketDetails `json:"value,omitempty"`
}

// UpdateContactProfile - Contact information associated with the support ticket.
type UpdateContactProfile struct {
	// Email addresses listed will be copied on any correspondence about the support ticket.
	AdditionalEmailAddresses []*string `json:"additionalEmailAddresses,omitempty"`

	// Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string `json:"country,omitempty"`

	// First name.
	FirstName *string `json:"firstName,omitempty"`

	// Last name.
	LastName *string `json:"lastName,omitempty"`

	// Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Preferred contact method.
	PreferredContactMethod *PreferredContactMethod `json:"preferredContactMethod,omitempty"`

	// Preferred language of support from Azure. Support languages vary based on the severity you choose for your support ticket.
	// Learn more at Azure Severity and responsiveness
	// [https://azure.microsoft.com/support/plans/response/]. Use the standard language-country code. Valid values are 'en-us'
	// for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French,
	// 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for
	// Chinese and 'de-de' for German.
	PreferredSupportLanguage *string `json:"preferredSupportLanguage,omitempty"`

	// Time zone of the user. This is the name of the time zone from Microsoft Time Zone Index Values [https://support.microsoft.com/help/973627/microsoft-time-zone-index-values].
	PreferredTimeZone *string `json:"preferredTimeZone,omitempty"`

	// Primary email address.
	PrimaryEmailAddress *string `json:"primaryEmailAddress,omitempty"`
}

// UpdateSupportTicket - Updates severity, ticket status, and contact details in the support ticket.
type UpdateSupportTicket struct {
	// Contact details to be updated on the support ticket.
	ContactDetails *UpdateContactProfile `json:"contactDetails,omitempty"`

	// Severity level.
	Severity *SeverityLevel `json:"severity,omitempty"`

	// Status to be updated on the ticket.
	Status *Status `json:"status,omitempty"`
}
