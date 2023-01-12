package main

import (
	_ "github.com/SupenBysz/gf-admin-company-modules/internal/boot"
	"github.com/SupenBysz/gf-admin-company-modules/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/SupenBysz/gf-admin-community"
)

func main() {
	cmd.Main.Run(gctx.New())
}
