package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder provides operations to call the count method.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/count()", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function count
// Deprecated: This method is obsolete. Use GetAsCountGetResponse instead.
// returns a FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseable), nil
}
// GetAsCountGetResponse invoke function count
// returns a FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) GetAsCountGetResponse(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponseable), nil
}
// ToGetRequestInformation invoke function count
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
