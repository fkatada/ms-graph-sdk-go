package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse()(*EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse{
        EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementAssignmentRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
