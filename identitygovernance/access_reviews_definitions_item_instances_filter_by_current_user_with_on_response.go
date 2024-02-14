package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse()(*AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse{
        AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDefinitionsItemInstancesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
