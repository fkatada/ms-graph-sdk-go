package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters user experience analytics device Startup Process Performance
type UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    m := &UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/{userExperienceAnalyticsDeviceStartupProcessPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get user experience analytics device Startup Process Performance
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/{userExperienceAnalyticsDeviceStartupProcessPerformance%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics device Startup Process Performance
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/{userExperienceAnalyticsDeviceStartupProcessPerformance%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
