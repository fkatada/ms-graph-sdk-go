package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder provides operations to manage the resource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetQueryParameters get resource from identityGovernance
type EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderInternal instantiates a new ResourceRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/resourceRoleScopes/{accessPackageResourceRoleScope%2Did}/scope/resource/roles/{accessPackageResourceRole%2Did}/resource{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder instantiates a new ResourceRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resource for identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Environment provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) Environment()(*EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceEnvironmentRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get resource from identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable), nil
}
// Patch update the navigation property resource in identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable), nil
}
// Scopes provides operations to manage the scopes property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) Scopes()(*EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resource for identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get resource from identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resource in identityGovernance
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemResourceRoleScopesItemScopeResourceRolesItemResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
