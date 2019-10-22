package balancers

import "github.com/google/wire"

// Set of providers for balancers components.
var Providers = wire.NewSet(NewStore, HttpHandler)
