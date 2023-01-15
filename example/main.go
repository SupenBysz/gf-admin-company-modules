package main

import (
	_ "github.com/SupenBysz/gf-admin-community"
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/boot"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/SupenBysz/gf-admin-company-modules/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
