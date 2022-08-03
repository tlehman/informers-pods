package main

import (
	"context"
	"fmt"
	"flag"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/tobi/etc/local.yaml", "Kubeconfig path")
	flag.Parse()
	fmt.Println("Starting Pod informer:")
	// get the kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// use the Watch() interface to listen for changes in the k8s resources
	watcher, err := clientset.CoreV1().Pods(corev1.NamespaceAll).
		Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for event := range watcher.ResultChan() {
		pod := event.Object.(*corev1.Pod)
		fmt.Printf("%v pod with name %s\n", event.Type, pod.Name)
	}
}
