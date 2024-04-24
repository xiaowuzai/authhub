package service

import "github.com/google/wire"

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(NewService)
