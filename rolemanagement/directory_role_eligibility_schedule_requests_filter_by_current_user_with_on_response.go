package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse struct {
    DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse
}
// NewDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse instantiates a new DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse()(*DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse) {
    m := &DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse{
        DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse: *NewDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable interface {
    DirectoryRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
