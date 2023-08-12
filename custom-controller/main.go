package main

import (
	"flag"
	"fmt"

	"github.com/hiteshrepo/Kubernetes-101/custom-controller/controllers"
	"github.com/hiteshrepo/Kubernetes-101/custom-controller/informers"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const DefaultKubeConfigPath = "/home/hitesh/.kube/config"

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/hitesh/.kube/config", "location of your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error while building config from kubeconfig file location : %s\n", err.Error())
		fmt.Println("fetching config within cluster")
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error while getting inclusterconfig : %s\n", err.Error())
			return
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error while creating client : %s\n", err.Error())
		return
	}

	client := dynamic.NewForConfigOrDie(config)

	gvr := schema.GroupVersionResource{
		Group:    "flights.com",
		Version:  "v1",
		Resource: "flighttickets",
	}

	// List all resources of your custom kind
	resources, err := client.Resource(gvr).Namespace("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, res := range resources.Items {
		namespace := res.GetNamespace()
		name := res.GetName()

		obj, err := client.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Resource Name: %s\n", obj.GetName())
		fmt.Printf("Resource Namespace: %s\n", obj.GetNamespace())
		fmt.Println("-----------------------------------")
	}

	informerFactory := informers.GetCoreInformer(clientSet)

	ch := make(chan struct{})

	depC := controllers.NewDeploymentController(clientSet, informerFactory.Apps().V1().Deployments())
	informerFactory.Start(ch)
	depC.Run(ch)
}
