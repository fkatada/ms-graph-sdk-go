package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
