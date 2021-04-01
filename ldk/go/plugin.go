package ldk

import (
	"github.com/hashicorp/go-plugin"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
)

// ServeLoopPlugin serves a specific loop plugin.
func ServeLoopPlugin(logger *utils.Logger, loop Loop) {
	plugin.Serve(&plugin.ServeConfig{
		Logger:          logger.Logger(),
		HandshakeConfig: utils.Handshake,
		Plugins: map[string]plugin.Plugin{
			"loop": &LoopPlugin{
				Impl:   loop,
				logger: logger,
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
