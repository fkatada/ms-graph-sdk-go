package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse struct {
    PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse
}
// NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse instantiates a new PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse()(*PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse) {
    m := &PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse{
        PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse: *NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreatePrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponseable
}
