package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters the set of vertical columns in this section.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder instantiates a new ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property columns for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of vertical columns in this section.
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable), nil
}
// Patch update the navigation property columns in groups
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable), nil
}
// ToDeleteRequestInformation delete navigation property columns for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the set of vertical columns in this section.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property columns in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Webparts provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
// returns a *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Webparts()(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
