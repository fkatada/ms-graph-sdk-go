package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemSitesItemGetApplicableContentTypesForListWithListIdResponse struct {
    ItemSitesItemGetApplicableContentTypesForListWithListIdGetResponse
}
// NewItemSitesItemGetApplicableContentTypesForListWithListIdResponse instantiates a new ItemSitesItemGetApplicableContentTypesForListWithListIdResponse and sets the default values.
func NewItemSitesItemGetApplicableContentTypesForListWithListIdResponse()(*ItemSitesItemGetApplicableContentTypesForListWithListIdResponse) {
    m := &ItemSitesItemGetApplicableContentTypesForListWithListIdResponse{
        ItemSitesItemGetApplicableContentTypesForListWithListIdGetResponse: *NewItemSitesItemGetApplicableContentTypesForListWithListIdGetResponse(),
    }
    return m
}
// CreateItemSitesItemGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemGetApplicableContentTypesForListWithListIdResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemSitesItemGetApplicableContentTypesForListWithListIdResponseable interface {
    ItemSitesItemGetApplicableContentTypesForListWithListIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
