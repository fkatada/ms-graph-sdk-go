package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder provides operations to count the resources in the collection.
type ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetQueryParameters get the number of the resource
type ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetQueryParameters
}
// NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) {
    m := &ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/outcomes/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder) {
    return NewClassesItemAssignmentsItemSubmissionsItemOutcomesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
