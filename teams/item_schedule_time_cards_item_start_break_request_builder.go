package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemScheduleTimeCardsItemStartBreakRequestBuilder provides operations to call the startBreak method.
type ItemScheduleTimeCardsItemStartBreakRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleTimeCardsItemStartBreakRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleTimeCardsItemStartBreakRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemScheduleTimeCardsItemStartBreakRequestBuilderInternal instantiates a new ItemScheduleTimeCardsItemStartBreakRequestBuilder and sets the default values.
func NewItemScheduleTimeCardsItemStartBreakRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimeCardsItemStartBreakRequestBuilder) {
    m := &ItemScheduleTimeCardsItemStartBreakRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/timeCards/{timeCard%2Did}/startBreak", pathParameters),
    }
    return m
}
// NewItemScheduleTimeCardsItemStartBreakRequestBuilder instantiates a new ItemScheduleTimeCardsItemStartBreakRequestBuilder and sets the default values.
func NewItemScheduleTimeCardsItemStartBreakRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimeCardsItemStartBreakRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleTimeCardsItemStartBreakRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action startBreak
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemScheduleTimeCardsItemStartBreakRequestBuilder) Post(ctx context.Context, body ItemScheduleTimeCardsItemStartBreakPostRequestBodyable, requestConfiguration *ItemScheduleTimeCardsItemStartBreakRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeCardable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTimeCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeCardable), nil
}
// ToPostRequestInformation invoke action startBreak
// returns a *RequestInformation when successful
func (m *ItemScheduleTimeCardsItemStartBreakRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemScheduleTimeCardsItemStartBreakPostRequestBodyable, requestConfiguration *ItemScheduleTimeCardsItemStartBreakRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemScheduleTimeCardsItemStartBreakRequestBuilder when successful
func (m *ItemScheduleTimeCardsItemStartBreakRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleTimeCardsItemStartBreakRequestBuilder) {
    return NewItemScheduleTimeCardsItemStartBreakRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
