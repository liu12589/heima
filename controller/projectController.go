package controller

import (
	"AIAssistServer/constants"
	"context"
	"fmt"
	v1 "k8s.io/api/admissionregistration/v1"
	"os"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateProject(namespace string, projectid string) string {
	clientset := GetClient()
	port := constants.GetPort()
	podName := "front-" + projectid

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:            podName,
					Image:           "heima-front-product:latest",
					ImagePullPolicy: corev1.PullPolicy(v1.NeverReinvocationPolicy),
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8501,
							HostPort:      port,
						},
					},
				},
			},
		},
	}

	_, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Error creating Pod:", err.Error())
		os.Exit(1)
	}

	return "http://10.69.70.21:" + strconv.Itoa(int(port))
}
