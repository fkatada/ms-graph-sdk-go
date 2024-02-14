package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type EntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementAssignmentsFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse()(*EntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse{
        EntitlementManagementAssignmentsFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementAssignmentsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAssignmentsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementManagementAssignmentsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAssignmentsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type EntitlementManagementAssignmentsFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementAssignmentsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
