package common

import "github.com/ucloud/ustack-sdk-go/common/config"

// Config is the config of ucloud sdk, use for setting up client
type Config = config.Config

// NewConfig will return a new client config with default options.
var NewConfig = config.NewConfig
