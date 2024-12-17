package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersFileStorageContainerItemRequestBuilder provides operations to manage the deletedContainers property of the microsoft.graph.fileStorage entity.
type FileStorageDeletedContainersFileStorageContainerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetQueryParameters get deletedContainers from storage
type FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetQueryParameters
}
// FileStorageDeletedContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
// returns a *FileStorageDeletedContainersItemActivateRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Activate()(*FileStorageDeletedContainersItemActivateRequestBuilder) {
    return NewFileStorageDeletedContainersItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilderInternal instantiates a new FileStorageDeletedContainersFileStorageContainerItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) {
    m := &FileStorageDeletedContainersFileStorageContainerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilder instantiates a new FileStorageDeletedContainersFileStorageContainerItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deletedContainers for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Drive provides operations to manage the drive property of the microsoft.graph.fileStorageContainer entity.
// returns a *FileStorageDeletedContainersItemDriveRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Drive()(*FileStorageDeletedContainersItemDriveRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get deletedContainers from storage
// returns a FileStorageContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFileStorageContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable), nil
}
// Lock provides operations to call the lock method.
// returns a *FileStorageDeletedContainersItemLockRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Lock()(*FileStorageDeletedContainersItemLockRequestBuilder) {
    return NewFileStorageDeletedContainersItemLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deletedContainers in storage
// returns a FileStorageContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFileStorageContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable), nil
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *FileStorageDeletedContainersItemPermanentDeleteRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) PermanentDelete()(*FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) {
    return NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.fileStorageContainer entity.
// returns a *FileStorageDeletedContainersItemPermissionsRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Permissions()(*FileStorageDeletedContainersItemPermissionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecycleBin provides operations to manage the recycleBin property of the microsoft.graph.fileStorageContainer entity.
// returns a *FileStorageDeletedContainersItemRecycleBinRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) RecycleBin()(*FileStorageDeletedContainersItemRecycleBinRequestBuilder) {
    return NewFileStorageDeletedContainersItemRecycleBinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *FileStorageDeletedContainersItemRestoreRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Restore()(*FileStorageDeletedContainersItemRestoreRequestBuilder) {
    return NewFileStorageDeletedContainersItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deletedContainers for storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get deletedContainers from storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deletedContainers in storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FileStorageContainerable, requestConfiguration *FileStorageDeletedContainersFileStorageContainerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unlock provides operations to call the unlock method.
// returns a *FileStorageDeletedContainersItemUnlockRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) Unlock()(*FileStorageDeletedContainersItemUnlockRequestBuilder) {
    return NewFileStorageDeletedContainersItemUnlockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder when successful
func (m *FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersFileStorageContainerItemRequestBuilder) {
    return NewFileStorageDeletedContainersFileStorageContainerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
