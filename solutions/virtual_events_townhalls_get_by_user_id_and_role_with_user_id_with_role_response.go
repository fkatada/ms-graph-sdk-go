package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable instead.
type VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse struct {
    VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponse
}
// NewVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse instantiates a new VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse and sets the default values.
func NewVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse()(*VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse) {
    m := &VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse{
        VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponse: *NewVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponse(),
    }
    return m
}
// CreateVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponse(), nil
}
// Deprecated: This class is obsolete. Use VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable instead.
type VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsTownhallsGetByUserIdAndRoleWithUserIdWithRoleGetResponseable
}
