package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttackSimulationLandingPagesItemDetailsCountRequestBuilder provides operations to count the resources in the collection.
type AttackSimulationLandingPagesItemDetailsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetQueryParameters get the number of the resource
type AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetQueryParameters
}
// NewAttackSimulationLandingPagesItemDetailsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewAttackSimulationLandingPagesItemDetailsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationLandingPagesItemDetailsCountRequestBuilder) {
    m := &AttackSimulationLandingPagesItemDetailsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/landingPages/{landingPage%2Did}/details/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewAttackSimulationLandingPagesItemDetailsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewAttackSimulationLandingPagesItemDetailsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationLandingPagesItemDetailsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationLandingPagesItemDetailsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *AttackSimulationLandingPagesItemDetailsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *AttackSimulationLandingPagesItemDetailsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationLandingPagesItemDetailsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AttackSimulationLandingPagesItemDetailsCountRequestBuilder) WithUrl(rawUrl string)(*AttackSimulationLandingPagesItemDetailsCountRequestBuilder) {
    return NewAttackSimulationLandingPagesItemDetailsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
