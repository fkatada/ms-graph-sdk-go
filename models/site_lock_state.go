package models
type SiteLockState int

const (
    UNLOCKED_SITELOCKSTATE SiteLockState = iota
    LOCKEDREADONLY_SITELOCKSTATE
    LOCKEDNOACCESS_SITELOCKSTATE
    LOCKEDNOADDITIONS_SITELOCKSTATE
    UNKNOWNFUTUREVALUE_SITELOCKSTATE
)

func (i SiteLockState) String() string {
    return []string{"unlocked", "lockedReadOnly", "lockedNoAccess", "lockedNoAdditions", "unknownFutureValue"}[i]
}
func ParseSiteLockState(v string) (any, error) {
    result := UNLOCKED_SITELOCKSTATE
    switch v {
        case "unlocked":
            result = UNLOCKED_SITELOCKSTATE
        case "lockedReadOnly":
            result = LOCKEDREADONLY_SITELOCKSTATE
        case "lockedNoAccess":
            result = LOCKEDNOACCESS_SITELOCKSTATE
        case "lockedNoAdditions":
            result = LOCKEDNOADDITIONS_SITELOCKSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SITELOCKSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSiteLockState(values []SiteLockState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SiteLockState) isMultiValue() bool {
    return false
}
