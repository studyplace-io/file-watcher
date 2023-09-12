### file-watcher 文件监听器

### 介绍
`file-watcher`是基于golang实现的文件监听器，支持`Create` `Rename` `Write` `Remove`事件，并生成对应的k8s事件。
![](https://github.com/studyplace-io/file-watcher/blob/main/image/%E6%97%A0%E6%A0%87%E9%A2%98-2023-08-10-2343.png?raw=true)
### 项目功能
- 自定义监听多个文件(使用空格分隔)
- 生成k8s event(事件)(`Create` `Rename` `Write` `Remove`事件)


### 项目启动
- 使用方法
```bash
# go run cmd/main.go <文件路径1> <文件路径2>
➜  file-watcher git:(main) go run cmd/main.go test.txt test11.yaml 
I0910 11:55:06.031217   55434 init_k8s_config.go:33] run outside the cluster
I0910 11:55:06.033137   55434 watcher.go:36] Start watching files: /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt
I0910 11:55:06.033161   55434 watcher.go:36] Start watching files: /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml
2023/09/10 11:55:09 File modified: /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml
File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml Modified
I0910 11:55:09.641927   55434 event_generator.go:100] Event generated successfully: test11.yaml-2023-09-10 11:55:09.607918 +0800 CST m=+3.591877917
2023/09/10 11:55:15 File modified: /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt
File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt Modified
I0910 11:55:15.879443   55434 event_generator.go:100] Event generated successfully: test.txt-2023-09-10 11:55:15.874743 +0800 CST m=+9.858852501
2023/09/10 11:55:34 File renamed: /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml
File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml Renamed
I0910 11:55:34.411631   55434 event_generator.go:100] Event generated successfully: test11.yaml-2023-09-10 11:55:34.40578 +0800 CST m=+28.3903301

```

- 生成对应的event事件类型
```bash
➜  .kube kubectl get event
LAST SEEN   TYPE       REASON              OBJECT             MESSAGE
58m         Modified   Watch file change   file/test.txt      File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt Modified
55m         Modified   Watch file change   file/test.txt      File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt Modified
51m         Modified   Watch file change   file/test.txt      File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt Modified
26s         Modified   Watch file change   file/test.txt      File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test.txt Modified
32s         Modified   Watch file change   file/test11.yaml   File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml Modified
7s          Renamed    Watch file change   file/test11.yaml   File /Users/zhenyu.jiang/go/src/golanglearning/new_project/file-watcher/test11.yaml Renamed

```
