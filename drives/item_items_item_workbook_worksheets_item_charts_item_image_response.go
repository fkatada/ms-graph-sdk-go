package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemImageGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemImageGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
