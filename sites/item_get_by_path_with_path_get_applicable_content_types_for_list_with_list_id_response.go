package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse struct {
    ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdGetResponse
}
// NewItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse instantiates a new ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse and sets the default values.
func NewItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse()(*ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse) {
    m := &ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse{
        ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdGetResponse: *NewItemGetByPathWithPathGetApplicableContentTypesForListWithListIdGetResponse(),
    }
    return m
}
// CreateItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdResponseable interface {
    ItemGetByPathWithPathGetApplicableContentTypesForListWithListIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
