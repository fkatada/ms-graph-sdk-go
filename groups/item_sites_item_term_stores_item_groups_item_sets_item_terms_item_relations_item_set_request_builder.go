package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
type ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters the [set] in which the relation is relevant.
type ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    m := &ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did1}/sets/{set%2Did}/terms/{term%2Did}/relations/{relation%2Did}/set{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder instantiates a new SetRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the [set] in which the relation is relevant.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}
// ToGetRequestInformation the [set] in which the relation is relevant.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    return NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
