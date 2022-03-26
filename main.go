/**
 * @Author: guobob
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/3/24 22:57
 */

package main

import (
	"github.com/generate_data/cmd"
	"go.uber.org/zap"
	"os"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		zap.L().Error("error exit: "+err.Error(), zap.Error(err))
		os.Exit(1)
	}
}
