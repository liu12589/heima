package controller

import (
	"AIAssistServer/constants"
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
)

var Client *kubernetes.Clientset

func GetClient() *kubernetes.Clientset {
	return Client
}

func NewClient() {
	configPath := homedir.HomeDir() + "/.kube/config"
	kubeConfig := flag.String(constants.KubeConfig, configPath, "(optional) absolute path to the kubeConfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Fatal(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	Client = clientSet
}
