package ldk

import "github.com/open-olive/loop-development-kit/ldk/go/v2/client"

// Loop is an interface that defines the methods required of all loop plugins
type Loop interface {
	LoopStart(client.Sidekick) error
	LoopStop() error
}
