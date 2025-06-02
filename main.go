package main

import (
	_ "GoFrame_Demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"GoFrame_Demo/internal/cmd"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
