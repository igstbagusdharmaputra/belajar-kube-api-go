package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfig string
	namespace  string
)

func init() {
	flag.StringVar(&kubeConfig, "kubeConfig", "./config", "Give the file config (optional)")
	flag.StringVar(&namespace, "namespace", "", "*Give the namespace name (default namespace)")
}

func loadKubeconfig() (*rest.Config, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
func main() {
	flag.Parse()

	fmt.Println("Loading kube config")
	config, err := loadKubeconfig()
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Println("Pod name -> ", pod.Name)
	}
}
