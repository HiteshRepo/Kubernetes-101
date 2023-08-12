package informers

import (
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

func GetCoreInformer(clientSet *kubernetes.Clientset) informers.SharedInformerFactory {
	informerFactory := informers.NewSharedInformerFactory(clientSet, 10*time.Minute)
	return informerFactory
}
