package main

import (
	_ "github.com/SupenBysz/gf-admin-community"
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/boot"
	_ "github.com/SupenBysz/gf-admin-company-modules/example/internal/consts"
	_ "github.com/SupenBysz/gf-admin-company-modules/internal/logic"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	boot.Main.Run(gctx.New())
}
