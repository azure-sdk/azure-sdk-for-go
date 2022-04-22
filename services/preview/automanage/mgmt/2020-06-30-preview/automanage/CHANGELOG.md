# Unreleased

## Breaking Changes

### Removed Funcs

1. AccountsClient.Update(context.Context, string, string, AccountUpdate) (Account, error)
1. AccountsClient.UpdatePreparer(context.Context, string, string, AccountUpdate) (*http.Request, error)
1. AccountsClient.UpdateResponder(*http.Response) (Account, error)
1. AccountsClient.UpdateSender(*http.Request) (*http.Response, error)
1. ConfigurationProfilePreferencesClient.Update(context.Context, string, string, ConfigurationProfilePreferenceUpdate) (ConfigurationProfilePreference, error)
1. ConfigurationProfilePreferencesClient.UpdatePreparer(context.Context, string, string, ConfigurationProfilePreferenceUpdate) (*http.Request, error)
1. ConfigurationProfilePreferencesClient.UpdateResponder(*http.Response) (ConfigurationProfilePreference, error)
1. ConfigurationProfilePreferencesClient.UpdateSender(*http.Request) (*http.Response, error)
