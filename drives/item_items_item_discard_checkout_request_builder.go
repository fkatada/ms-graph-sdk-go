package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemDiscardCheckoutRequestBuilder provides operations to call the discardCheckout method.
type ItemItemsItemDiscardCheckoutRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemDiscardCheckoutRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemDiscardCheckoutRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemDiscardCheckoutRequestBuilderInternal instantiates a new ItemItemsItemDiscardCheckoutRequestBuilder and sets the default values.
func NewItemItemsItemDiscardCheckoutRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemDiscardCheckoutRequestBuilder) {
    m := &ItemItemsItemDiscardCheckoutRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/discardCheckout", pathParameters),
    }
    return m
}
// NewItemItemsItemDiscardCheckoutRequestBuilder instantiates a new ItemItemsItemDiscardCheckoutRequestBuilder and sets the default values.
func NewItemItemsItemDiscardCheckoutRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemDiscardCheckoutRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemDiscardCheckoutRequestBuilderInternal(urlParams, requestAdapter)
}
// Post discard the check out of a driveItem. This action releases a driveItem resource that was previously checked out. Any changes made to the item while it was checked out are discarded.  The same user that performed the checkout must discard it. Another alternative is to use application permissions.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-discardcheckout?view=graph-rest-1.0
func (m *ItemItemsItemDiscardCheckoutRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemDiscardCheckoutRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation discard the check out of a driveItem. This action releases a driveItem resource that was previously checked out. Any changes made to the item while it was checked out are discarded.  The same user that performed the checkout must discard it. Another alternative is to use application permissions.
// returns a *RequestInformation when successful
func (m *ItemItemsItemDiscardCheckoutRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemDiscardCheckoutRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemDiscardCheckoutRequestBuilder when successful
func (m *ItemItemsItemDiscardCheckoutRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemDiscardCheckoutRequestBuilder) {
    return NewItemItemsItemDiscardCheckoutRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
