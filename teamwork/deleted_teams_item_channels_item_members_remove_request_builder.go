package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder provides operations to call the remove method.
type DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilderInternal instantiates a new DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/members/remove", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilder instantiates a new DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// Deprecated: This method is obsolete. Use PostAsRemovePostResponse instead.
// returns a DeletedTeamsItemChannelsItemMembersRemoveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-1.0
func (m *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) Post(ctx context.Context, body DeletedTeamsItemChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(DeletedTeamsItemChannelsItemMembersRemoveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemMembersRemoveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemMembersRemoveResponseable), nil
}
// PostAsRemovePostResponse remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a DeletedTeamsItemChannelsItemMembersRemovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmember-remove?view=graph-rest-1.0
func (m *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) PostAsRemovePostResponse(ctx context.Context, body DeletedTeamsItemChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(DeletedTeamsItemChannelsItemMembersRemovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemMembersRemovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemMembersRemovePostResponseable), nil
}
// ToPostRequestInformation remove multiple members from a team in a single request. The response provides details about which memberships could and couldn't be removed.
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletedTeamsItemChannelsItemMembersRemovePostRequestBodyable, requestConfiguration *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemMembersRemoveRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemMembersRemoveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
