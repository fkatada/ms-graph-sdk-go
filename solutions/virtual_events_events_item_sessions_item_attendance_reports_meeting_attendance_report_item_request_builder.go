package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters
}
// VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceRecords provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
// returns a *VirtualEventsEventsItemSessionsItemAttendanceReportsItemAttendanceRecordsRequestBuilder when successful
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) AttendanceRecords()(*VirtualEventsEventsItemSessionsItemAttendanceReportsItemAttendanceRecordsRequestBuilder) {
    return NewVirtualEventsEventsItemSessionsItemAttendanceReportsItemAttendanceRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderInternal instantiates a new VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) {
    m := &VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports/{meetingAttendanceReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder instantiates a new VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property attendanceReports for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the attendance reports of an online meeting. Read-only.
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable), nil
}
// Patch update the navigation property attendanceReports in solutions
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable), nil
}
// ToDeleteRequestInformation delete navigation property attendanceReports for solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports/{meetingAttendanceReport%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the attendance reports of an online meeting. Read-only.
// returns a *RequestInformation when successful
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attendanceReports in solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports/{meetingAttendanceReport%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) {
    return NewVirtualEventsEventsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
