package main

import (
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/boot"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/SupenBysz/gf-admin-community"
)

func main() {
	boot.Main.Run(gctx.New())
}
