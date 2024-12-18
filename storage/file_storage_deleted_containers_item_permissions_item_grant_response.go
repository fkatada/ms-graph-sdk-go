package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemPermissionsItemGrantPostResponseable instead.
type FileStorageDeletedContainersItemPermissionsItemGrantResponse struct {
    FileStorageDeletedContainersItemPermissionsItemGrantPostResponse
}
// NewFileStorageDeletedContainersItemPermissionsItemGrantResponse instantiates a new FileStorageDeletedContainersItemPermissionsItemGrantResponse and sets the default values.
func NewFileStorageDeletedContainersItemPermissionsItemGrantResponse()(*FileStorageDeletedContainersItemPermissionsItemGrantResponse) {
    m := &FileStorageDeletedContainersItemPermissionsItemGrantResponse{
        FileStorageDeletedContainersItemPermissionsItemGrantPostResponse: *NewFileStorageDeletedContainersItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemPermissionsItemGrantPostResponseable instead.
type FileStorageDeletedContainersItemPermissionsItemGrantResponseable interface {
    FileStorageDeletedContainersItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
