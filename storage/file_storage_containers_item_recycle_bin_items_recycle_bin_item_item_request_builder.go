package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.recycleBin entity.
type FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetQueryParameters list of the recycleBinItems deleted by a user.
type FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetQueryParameters
}
// FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderInternal instantiates a new FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) {
    m := &FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/recycleBin/items/{recycleBinItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder instantiates a new FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemRecycleBinItemsItemCreatedByUserRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) CreatedByUser()(*FileStorageContainersItemRecycleBinItemsItemCreatedByUserRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsItemCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of the recycleBinItems deleted by a user.
// returns a RecycleBinItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRecycleBinItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable), nil
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) LastModifiedByUser()(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a RecycleBinItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRecycleBinItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable), nil
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of the recycleBinItems deleted by a user.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecycleBinItemable, requestConfiguration *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsRecycleBinItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
