package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse()(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse{
        AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
