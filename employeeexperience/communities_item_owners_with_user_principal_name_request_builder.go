package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder provides operations to manage the owners property of the microsoft.graph.community entity.
type CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters the admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
type CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilderInternal instantiates a new CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userPrincipalName *string)(*CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) {
    m := &CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/communities/{community%2Did}/owners(userPrincipalName='{userPrincipalName}'){?%24expand,%24select}", pathParameters),
    }
    if userPrincipalName != nil {
        m.BaseRequestBuilder.PathParameters["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilder instantiates a new CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get the admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// ToGetRequestInformation the admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
// returns a *RequestInformation when successful
func (m *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder when successful
func (m *CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) WithUrl(rawUrl string)(*CommunitiesItemOwnersWithUserPrincipalNameRequestBuilder) {
    return NewCommunitiesItemOwnersWithUserPrincipalNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
