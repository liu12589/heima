![image](https://github.com/liu12589/heima/assets/46433529/337a9083-6201-4240-a92c-6b2389aad3ed)

- dbController 在k8s集群指定命名空间内创建一个 milvus 向量数据库pod，并创建 services 服务，返回集群内访问的IP：port
- namespaceController 在集群内创建一个命名空间
- projectController 在集群指定命名空间内拉取基础服务，并绑定 NodeIP 端口。
