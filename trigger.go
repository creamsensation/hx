package hx

import "github.com/creamsensation/gox"

func Trigger(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-trigger")(value...)
}
