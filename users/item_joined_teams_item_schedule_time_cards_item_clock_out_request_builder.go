package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder provides operations to call the clockOut method.
type ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderInternal instantiates a new ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) {
    m := &ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/schedule/timeCards/{timeCard%2Did}/clockOut", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder instantiates a new ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action clockOut
// returns a TimeCardable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemScheduleTimeCardsItemClockOutPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeCardable, error) {
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
// ToPostRequestInformation invoke action clockOut
// returns a *RequestInformation when successful
func (m *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemScheduleTimeCardsItemClockOutPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder when successful
func (m *ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleTimeCardsItemClockOutRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
