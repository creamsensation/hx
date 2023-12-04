package hx

import "github.com/creamsensation/gox"

func Target(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-target")(value...)
}
