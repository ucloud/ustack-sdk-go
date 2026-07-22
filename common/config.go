package common

import "third_party/platform-sdk-go/common/config"

// Config is the config of ucloud sdk, use for setting up client
type Config = config.Config

// NewConfig will return a new client config with default options.
var NewConfig = config.NewConfig
