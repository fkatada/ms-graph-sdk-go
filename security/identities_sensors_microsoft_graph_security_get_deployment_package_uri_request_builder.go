package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder provides operations to call the getDeploymentPackageUri method.
type IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal instantiates a new IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    m := &IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/identities/sensors/microsoft.graph.security.getDeploymentPackageUri()", pathParameters),
    }
    return m
}
// NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder instantiates a new IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder and sets the default values.
func NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the sensor deployment package URL and version. You can use this URL to download the installer to install the sensor on a server.
// returns a SensorDeploymentPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SensorDeploymentPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateSensorDeploymentPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SensorDeploymentPackageable), nil
}
// ToGetRequestInformation get the sensor deployment package URL and version. You can use this URL to download the installer to install the sensor on a server.
// returns a *RequestInformation when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder when successful
func (m *IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) WithUrl(rawUrl string)(*IdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder) {
    return NewIdentitiesSensorsMicrosoftGraphSecurityGetDeploymentPackageUriRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
