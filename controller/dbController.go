package controller

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

func CreateMilvus(namespace, user, password string) (string, string, error) {
	// 配置集群客户端
	clientset := GetClient()

	serviceName := user + "-milvus"
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: namespace,
			Labels: map[string]string{
				"app": serviceName,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "milvus-container",
					Image:           "soulteary/milvus:embed-2.1.0",
					ImagePullPolicy: v1.PullIfNotPresent,
				},
			},
		},
	}

	// 在Kubernetes中创建Pod
	_, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Pod: %v\n", err)
	}

	//创建Service对象
	serviceClient := clientset.CoreV1().Services(namespace)
	//实例化一个数据结构
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": serviceName,
			},
			Ports: []v1.ServicePort{
				{
					Protocol:   v1.ProtocolTCP,
					Port:       19530,
					TargetPort: intstr.FromInt(19530),
				},
			},
		},
	}
	_, err = serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	service, err = clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	port := strconv.Itoa(int(service.Spec.Ports[0].Port))
	return service.Spec.ClusterIP, port, nil
}
