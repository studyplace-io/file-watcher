### file-watcher
<a href="./README.md">English</a> | <a href="./README-zh.md">简体中文</a>
### Introduction
- `file-watcher` is a file listener implemented based on golang. When there are any changes to the file, an event will be generated. 

- It supports `Create` `Rename` `Write` `Remove` events and generates corresponding k8s events.

![](https://github.com/studyplace-io/file-watcher/blob/main/image/%E6%97%A0%E6%A0%87%E9%A2%98-2023-08-10-2343.png?raw=true)
### Project support
- Customize monitoring of multiple files (separated by spaces)
- Generate k8s event (event)(`Create` `Rename` `Write` `Remove`event)

### start up
- usage
```bash
# go run cmd/main.go <filepath1> <filepath2>
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

- Generate the event type
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

### Deploy
Currently supports binary, docker deployment, k8s cluster deployment
- docker image
```bash
docker build -t file-watcher:v1 .
```
- k8s cluster
```bash
[root@VM-0-16-centos yaml]# cd ..
[root@VM-0-16-centos file-watcher]# kubectl apply -f yaml/rbac.yaml
serviceaccount/file-watcher-sa unchanged
clusterrole.rbac.authorization.k8s.io/file-watcher-clusterrole unchanged
clusterrolebinding.rbac.authorization.k8s.io/file-watcher-ClusterRoleBinding unchanged

[root@VM-0-16-centos file-watcher]# kubectl apply -f yaml/deploy.yaml
deployment.apps/file-watcher unchanged
```

P.S.：When deploying yaml/deploy.yaml, you need to pay special attention to **the mounting problem**
