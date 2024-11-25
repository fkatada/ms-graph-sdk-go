package security
type PurgeType int

const (
    RECOVERABLE_PURGETYPE PurgeType = iota
    UNKNOWNFUTUREVALUE_PURGETYPE
    PERMANENTLYDELETE_PURGETYPE
)

func (i PurgeType) String() string {
    return []string{"recoverable", "unknownFutureValue", "permanentlyDelete"}[i]
}
func ParsePurgeType(v string) (any, error) {
    result := RECOVERABLE_PURGETYPE
    switch v {
        case "recoverable":
            result = RECOVERABLE_PURGETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PURGETYPE
        case "permanentlyDelete":
            result = PERMANENTLYDELETE_PURGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePurgeType(values []PurgeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PurgeType) isMultiValue() bool {
    return false
}
