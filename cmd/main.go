package main

import (
	_ "github.com/practice/file-watcher/pkg/events"
	"github.com/practice/file-watcher/pkg/file_watcher"
	"github.com/spf13/cobra"
	"log"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "file-watcher",
		Short: "run file-wather",
		Run: func(cmd *cobra.Command, args []string) {
			// 启动事件监听器
			file_watcher.StartWatcher(cmd, args)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
