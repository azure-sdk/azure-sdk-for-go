package botserviceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/botservice/mgmt/2021-03-01/botservice"
	"github.com/Azure/go-autorest/autorest"
)

// BotsClientAPI contains the set of methods on the BotsClient type.
type BotsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, parameters botservice.Bot) (result botservice.Bot, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result botservice.Bot, err error)
	GetCheckNameAvailability(ctx context.Context, parameters botservice.CheckNameAvailabilityRequestBody) (result botservice.CheckNameAvailabilityResponseBody, err error)
	List(ctx context.Context) (result botservice.BotResponseListPage, err error)
	ListComplete(ctx context.Context) (result botservice.BotResponseListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result botservice.BotResponseListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result botservice.BotResponseListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters botservice.Bot) (result botservice.Bot, err error)
}

var _ BotsClientAPI = (*botservice.BotsClient)(nil)

// ChannelsClientAPI contains the set of methods on the ChannelsClient type.
type ChannelsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, channelName botservice.ChannelName, parameters botservice.BotChannel) (result botservice.BotChannel, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (result botservice.BotChannel, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, resourceName string) (result botservice.ChannelResponseListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, resourceName string) (result botservice.ChannelResponseListIterator, err error)
	ListWithKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName botservice.ChannelName) (result botservice.ListChannelWithKeysResponse, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, channelName botservice.ChannelName, parameters botservice.BotChannel) (result botservice.BotChannel, err error)
}

var _ ChannelsClientAPI = (*botservice.ChannelsClient)(nil)

// DirectLineClientAPI contains the set of methods on the DirectLineClient type.
type DirectLineClientAPI interface {
	RegenerateKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName botservice.RegenerateKeysChannelName, parameters botservice.SiteInfo) (result botservice.BotChannel, err error)
}

var _ DirectLineClientAPI = (*botservice.DirectLineClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result botservice.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result botservice.OperationEntityListResultIterator, err error)
}

var _ OperationsClientAPI = (*botservice.OperationsClient)(nil)

// BotConnectionClientAPI contains the set of methods on the BotConnectionClient type.
type BotConnectionClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters botservice.ConnectionSetting) (result botservice.ConnectionSetting, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result botservice.ConnectionSetting, err error)
	ListByBotService(ctx context.Context, resourceGroupName string, resourceName string) (result botservice.ConnectionSettingResponseListPage, err error)
	ListByBotServiceComplete(ctx context.Context, resourceGroupName string, resourceName string) (result botservice.ConnectionSettingResponseListIterator, err error)
	ListServiceProviders(ctx context.Context) (result botservice.ServiceProviderResponseList, err error)
	ListWithSecrets(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result botservice.ConnectionSetting, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters botservice.ConnectionSetting) (result botservice.ConnectionSetting, err error)
}

var _ BotConnectionClientAPI = (*botservice.BotConnectionClient)(nil)

// HostSettingsClientAPI contains the set of methods on the HostSettingsClient type.
type HostSettingsClientAPI interface {
	Get(ctx context.Context) (result botservice.HostSettingsResponse, err error)
}

var _ HostSettingsClientAPI = (*botservice.HostSettingsClient)(nil)
