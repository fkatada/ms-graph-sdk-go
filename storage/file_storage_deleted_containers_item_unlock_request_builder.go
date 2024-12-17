package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemUnlockRequestBuilder provides operations to call the unlock method.
type FileStorageDeletedContainersItemUnlockRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemUnlockRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemUnlockRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemUnlockRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemUnlockRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemUnlockRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemUnlockRequestBuilder) {
    m := &FileStorageDeletedContainersItemUnlockRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/unlock", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemUnlockRequestBuilder instantiates a new FileStorageDeletedContainersItemUnlockRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemUnlockRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemUnlockRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemUnlockRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unlock
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemUnlockRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemUnlockRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action unlock
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemUnlockRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemUnlockRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemUnlockRequestBuilder when successful
func (m *FileStorageDeletedContainersItemUnlockRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemUnlockRequestBuilder) {
    return NewFileStorageDeletedContainersItemUnlockRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
