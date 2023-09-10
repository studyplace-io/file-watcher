package file_watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/practice/file-watcher/pkg/events"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func StartWatcher(cmd *cobra.Command, args []string) {
	// 创建一个新的文件系统通知器
	g := events.NewGenerator()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// 获取要监听的文件列表
	files, err := resolveFiles(args)
	if err != nil {
		log.Fatal(err)
	}

	// 添加要监听的文件到通知器
	for _, file := range files {
		err = watcher.Add(file)
		if err != nil {
			log.Fatal(err)
		}
		klog.Infof("Start watching files: %s", file)
	}

	// 退出信号
	done := make(chan bool)

	// 启动一个goroutine来处理文件系统事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				g.HandleEvent(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				klog.Errorf("watcher error: %s", err)
			case <-done:
				klog.Errorf("exit watcher")
				return
			}
		}
	}()


	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)
		<-sigCh
		close(done)
	}()

	<-done

	// 等待goroutine退出
	time.Sleep(time.Second * 3)
	klog.Info("Stop watching")
}

// resolveFiles 解析文件路径为绝对路径
func resolveFiles(files []string) ([]string, error) {
	var resolvedFiles []string
	for _, file := range files {
		absPath, err := filepath.Abs(file)
		if err != nil {
			return nil, err
		}
		resolvedFiles = append(resolvedFiles, absPath)
	}
	return resolvedFiles, nil
}


