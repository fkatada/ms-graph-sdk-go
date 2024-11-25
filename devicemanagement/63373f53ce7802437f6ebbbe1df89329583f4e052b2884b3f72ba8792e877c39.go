package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder provides operations to manage the assignedUsers property of the microsoft.graph.cloudPcProvisioningPolicyAssignment entity.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetQueryParameters the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderInternal instantiates a new VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userPrincipalName *string)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) {
    m := &VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/assignments/{cloudPcProvisioningPolicyAssignment%2Did}/assignedUsers(userPrincipalName='{userPrincipalName}'){?%24expand,%24select}", pathParameters),
    }
    if userPrincipalName != nil {
        m.BaseRequestBuilder.PathParameters["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder instantiates a new VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
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
// ToGetRequestInformation the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
// returns a *RequestInformation when successful
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder when successful
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersWithUserPrincipalNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
