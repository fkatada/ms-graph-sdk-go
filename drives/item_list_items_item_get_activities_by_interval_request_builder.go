package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListItemsItemGetActivitiesByIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type ItemListItemsItemGetActivitiesByIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetQueryParameters
}
// NewItemListItemsItemGetActivitiesByIntervalRequestBuilderInternal instantiates a new ItemListItemsItemGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemListItemsItemGetActivitiesByIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemGetActivitiesByIntervalRequestBuilder) {
    m := &ItemListItemsItemGetActivitiesByIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/list/items/{listItem%2Did}/getActivitiesByInterval(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemListItemsItemGetActivitiesByIntervalRequestBuilder instantiates a new ItemListItemsItemGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewItemListItemsItemGetActivitiesByIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemGetActivitiesByIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListItemsItemGetActivitiesByIntervalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ItemListItemsItemGetActivitiesByIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemGetActivitiesByIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemListItemsItemGetActivitiesByIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListItemsItemGetActivitiesByIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalGetResponse invoke function getActivitiesByInterval
// returns a ItemListItemsItemGetActivitiesByIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemGetActivitiesByIntervalRequestBuilder) GetAsGetActivitiesByIntervalGetResponse(ctx context.Context, requestConfiguration *ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(ItemListItemsItemGetActivitiesByIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemListItemsItemGetActivitiesByIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemListItemsItemGetActivitiesByIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *ItemListItemsItemGetActivitiesByIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListItemsItemGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListItemsItemGetActivitiesByIntervalRequestBuilder when successful
func (m *ItemListItemsItemGetActivitiesByIntervalRequestBuilder) WithUrl(rawUrl string)(*ItemListItemsItemGetActivitiesByIntervalRequestBuilder) {
    return NewItemListItemsItemGetActivitiesByIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
