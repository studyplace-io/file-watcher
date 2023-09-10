package events

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/practice/file-watcher/pkg/common"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"log"
	"strings"
	"time"
)

// EventGenerator generates events.
type EventGenerator struct {
	//kubeConfig string
	//
	//kind      string
	//name      string
	//namespace string

	client    *kubernetes.Clientset
}

const (
	defaultNamespace = "default"
	defaultEventType = v1.EventTypeWarning
)

// NewGenerator creates a generator.
func NewGenerator() *EventGenerator {
	r := common.K8sRestConfig()
	clientSet, err := kubernetes.NewForConfig(r)
	if err != nil {
		fmt.Println("dkjsl;djflsd: ", err)
		return nil
	}
	eg := &EventGenerator{
		client: clientSet,
	}
	return eg
}

// 获取文件事件类型
func getEventType(event fsnotify.Event) string {
	switch {
	case event.Op&fsnotify.Create == fsnotify.Create:
		return "Created"
	case event.Op&fsnotify.Write == fsnotify.Write:
		return "Modified"
	case event.Op&fsnotify.Remove == fsnotify.Remove:
		return "Removed"
	case event.Op&fsnotify.Rename == fsnotify.Rename:
		return "Renamed"
	case event.Op&fsnotify.Chmod == fsnotify.Chmod:
		return "PermissionChanged"
	default:
		return "Unknown"
	}
}

// SendKubernetesEvent 生成事件
func (g *EventGenerator) SendKubernetesEvent(ee fsnotify.Event) error {

	eventType := getEventType(ee)
	message := fmt.Sprintf("File %s %s", ee.Name, eventType)
	fmt.Println(message)
	ss := strings.Split(ee.Name, "/")
	lastFileName := ss[len(ss)-1]

	now := time.Now()
	// 使用客户端创建模拟事件
	event, err := g.client.CoreV1().Events(defaultNamespace).CreateWithEventNamespace(&v1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%v-%s", lastFileName, time.Now().String()),
			Namespace: defaultNamespace,
		},
		FirstTimestamp:      metav1.NewTime(now),
		LastTimestamp:       metav1.NewTime(now),
		EventTime:           metav1.NewMicroTime(now),
		ReportingController: "k8s-events-generator",
		ReportingInstance:   "k8s-events-generator",
		Action: "Test",
		InvolvedObject:      v1.ObjectReference{
			Kind: "File",
			Name: lastFileName,
			Namespace: defaultNamespace,
		},
		Reason:              "Watch file change",
		Type:                eventType,
		Message:             message,
	})
	if err != nil {
		klog.Errorf("events error...: %s", err)
		return err
	}

	klog.Infof("Event generated successfully: %v", event.Name)

	return nil
}

// HandleEvent 处理文件系统事件
func (g *EventGenerator) HandleEvent(ee fsnotify.Event) {
	switch {
	case ee.Op&fsnotify.Create == fsnotify.Create:
		log.Printf("File modified: %s\n", ee.Name)
	case ee.Op&fsnotify.Write == fsnotify.Write:
		log.Printf("File modified: %s\n", ee.Name)
	case ee.Op&fsnotify.Remove == fsnotify.Remove:
		log.Printf("File removed: %s\n", ee.Name)
	case ee.Op&fsnotify.Rename == fsnotify.Rename:
		log.Printf("File renamed: %s\n", ee.Name)
	case ee.Op&fsnotify.Chmod == fsnotify.Chmod:
		//log.Printf("File permission changed: %s\n", ee.Name)
		return
	}


	// 创建 k8s 事件并发送
	_ = g.SendKubernetesEvent(ee)
}