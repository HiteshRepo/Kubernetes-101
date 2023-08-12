package controllers

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	appInformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appListers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const FlightBookingQueueName = "fbqueue"

type Controller struct {
	clientSet     kubernetes.Interface
	depLister     appListers.DeploymentLister
	depCachedSync cache.InformerSynced
	queue         workqueue.RateLimitingInterface
}

func NewDeploymentController(clientSet kubernetes.Interface, depInformer appInformers.DeploymentInformer) *Controller {
	c := &Controller{
		clientSet:     clientSet,
		depLister:     depInformer.Lister(),
		depCachedSync: depInformer.Informer().HasSynced,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "fbqueue"),
	}

	depInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDel,
		},
	)

	return c
}

func (c *Controller) Run(ch <-chan struct{}) {
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.depCachedSync) {
		fmt.Println("waiting for cache to be synced")
	}

	wait.Until(c.worker, 1*time.Second, ch)

	<-ch
}

func (c *Controller) worker() {
	for c.processItem() {
		fmt.Println("item processed")
	}
}

func (c *Controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Forget(item)

	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Printf("error while getting key from cache")
		return false
	}

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Printf("error while splitting key into namespace and name")
		return false
	}

	err = c.syncDeployment(ns, name)
	if err != nil {
		fmt.Printf("error while syncing deployment")
		return false
	}

	return true
}

func (c *Controller) syncDeployment(ns, name string) error {
	fmt.Println(ns, name)
	return nil
}

func (c *Controller) handleAdd(obj interface{}) {
	fmt.Println("add was called")
	c.queue.Add(obj)
}

func (c *Controller) handleDel(obj interface{}) {
	fmt.Println("del was called")
	c.queue.Add(obj)
}
