package translator

import (
	"github.com/erlangs/okoo/src/server"
)

const MODULE_NAME string = "translator"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
}
