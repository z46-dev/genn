package shared

import "github.com/z46-dev/golog"

// Shared logger for all modules, shared doesn't import it so it can live here safely when initialized elsewhere.
var L *golog.Logger = nil
