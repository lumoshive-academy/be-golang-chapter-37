package main

import (
	"golang-chapter-37/implem-gin-shipping/infra"
	"golang-chapter-37/implem-gin-shipping/router"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("Error", zap.Error(err))
		return
	}

	router.SetupReouter(ctx)
}
