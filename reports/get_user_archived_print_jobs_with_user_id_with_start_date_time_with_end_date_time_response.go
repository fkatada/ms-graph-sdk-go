package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse struct {
    GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse
}
// NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse instantiates a new GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse and sets the default values.
func NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse()(*GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse) {
    m := &GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse{
        GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse: *NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponse(),
    }
    return m
}
// CreateGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeResponseable interface {
    GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
