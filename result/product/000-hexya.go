package product

import(
"github.com/hexya-erp/hexya/hexya/server"
)

const MODULE_NAME string = "product"

func init() {
server.RegisterModule(&server.Module{
Name:     MODULE_NAME,
PostInit: func() {},
})
}