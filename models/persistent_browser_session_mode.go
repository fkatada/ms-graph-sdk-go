package models
import (
    "errors"
)
// Provides operations to manage the collection of agreement entities.
type PersistentBrowserSessionMode int

const (
    ALWAYS_PERSISTENTBROWSERSESSIONMODE PersistentBrowserSessionMode = iota
    NEVER_PERSISTENTBROWSERSESSIONMODE
)

func (i PersistentBrowserSessionMode) String() string {
    return []string{"always", "never"}[i]
}
func ParsePersistentBrowserSessionMode(v string) (interface{}, error) {
    result := ALWAYS_PERSISTENTBROWSERSESSIONMODE
    switch v {
        case "always":
            result = ALWAYS_PERSISTENTBROWSERSESSIONMODE
        case "never":
            result = NEVER_PERSISTENTBROWSERSESSIONMODE
        default:
            return 0, errors.New("Unknown PersistentBrowserSessionMode value: " + v)
    }
    return &result, nil
}
func SerializePersistentBrowserSessionMode(values []PersistentBrowserSessionMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
