package ecore

import (
	_ "crypto/md5"
	"github.com/teris-io/shortid"
)

func E取短Id() string {
	s, _ := shortid.Generate()
	return s
}
