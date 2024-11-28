package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder provides operations to call the regenerateDeploymentAccessKey method.
type IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderInternal instantiates a new IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) {
    m := &IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors/microsoft.graph.security.regenerateDeploymentAccessKey", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder instantiates a new IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate a new deployment access key that can be used to install a sensor associated with the workspace.
// returns a DeploymentAccessKeyTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DeploymentAccessKeyTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateDeploymentAccessKeyTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DeploymentAccessKeyTypeable), nil
}
// ToPostRequestInformation generate a new deployment access key that can be used to install a sensor associated with the workspace.
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftGraphSecurityRegenerateDeploymentAccessKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
