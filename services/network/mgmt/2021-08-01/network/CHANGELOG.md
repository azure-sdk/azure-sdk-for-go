# Unreleased

## Additive Changes

### New Funcs

1. *AzureFirewallsListLearnedPrefixesFuture.UnmarshalJSON([]byte) error
1. AzureFirewallsClient.ListLearnedPrefixes(context.Context, string, string) (AzureFirewallsListLearnedPrefixesFuture, error)
1. AzureFirewallsClient.ListLearnedPrefixesPreparer(context.Context, string, string) (*http.Request, error)
1. AzureFirewallsClient.ListLearnedPrefixesResponder(*http.Response) (IPPrefixesList, error)
1. AzureFirewallsClient.ListLearnedPrefixesSender(*http.Request) (AzureFirewallsListLearnedPrefixesFuture, error)

### Struct Changes

#### New Structs

1. AzureFirewallsListLearnedPrefixesFuture
1. IPPrefixesList

#### New Struct Fields

1. FirewallPolicySNAT.AutoLearnPrivateRanges
