package hx

import "github.com/creamsensation/gox"

func Sync(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-sync")(value...)
}
