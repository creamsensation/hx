package hx

import "github.com/creamsensation/gox"

func Indicator(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-indicator")(value...)
}
