/**
 * @Author: guobob
 * @Description:
 * @File:  root.go
 * @Version: 1.0.0
 * @Date: 2022/3/25 09:37
 */

package cmd

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRootCmd() *cobra.Command {
	var opts struct {
		logLevel  LogLevel
		logOutput []string
		pprof     string
	}
	var profiler interface{ Stop() }
	cmd := &cobra.Command{
		Use: "generate_data",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
			cfg := zap.NewDevelopmentConfig()
			cfg.Level = zap.NewAtomicLevelAt(opts.logLevel.Level)
			cfg.OutputPaths = opts.logOutput
			cfg.ErrorOutputPaths = opts.logOutput
			cfg.DisableStacktrace = !cfg.Level.Enabled(zap.DebugLevel)
			logger, err := cfg.Build()
			if err != nil {
				fmt.Println("init log fail , " + err.Error())
				os.Exit(-1)
			}
			zap.ReplaceGlobals(logger)
			if len(opts.pprof) > 0 {
				switch opts.pprof {
				case "cpu":
					profiler = profile.Start(profile.CPUProfile, profile.NoShutdownHook)
				case "mem", "mem.heap":
					profiler = profile.Start(profile.MemProfileHeap, profile.NoShutdownHook)
				case "mem.alloc":
					profiler = profile.Start(profile.MemProfileAllocs, profile.NoShutdownHook)
				case "mutex":
					profiler = profile.Start(profile.MutexProfile, profile.NoShutdownHook)
				case "block":
					profiler = profile.Start(profile.BlockProfile, profile.NoShutdownHook)
				default:
					go func() {
						logger.Info("serve pprof", zap.Error(http.ListenAndServe(opts.pprof, nil)))
					}()
				}
			}
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if profiler != nil {
				profiler.Stop()
			}
		},
	}
	opts.logLevel = LogLevel{zapcore.InfoLevel}
	cmd.PersistentFlags().Var(&opts.logLevel, "log-level", "log level")
	cmd.PersistentFlags().StringSliceVar(&opts.logOutput, "log-output", []string{"stderr"}, "log output")
	cmd.PersistentFlags().StringVar(&opts.pprof, "pprof", "", "enable pprof")
	//cmd.PersistentFlags().StringVar()
	cmd.AddCommand(NewTextCommand())
	return cmd
}

type LogLevel struct {
	zapcore.Level
}

func (lv *LogLevel) Type() string {
	return "string"
}
