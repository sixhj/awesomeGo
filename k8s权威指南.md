《kubernetes权威指南 从docker到kubernetes实践全接触》

k8s是一个完备的分布式系统支撑平台

- 多层安全防护和准收机制
- 多租户应用支撑能力
- 透明的服务注册和服务发现
- 内建智能负载均衡
- 故障发现和自我修复
- 在线扩容、自动调度

---



Service 是分布式集群架构的核心（虚拟Cluster IP 和service port）

- 拥有一个唯一指定的名字
- 拥有一个虚拟IP和端口号
- 能够提供某种远程服务能力
- 被映射到了提供这种服务能力的一组容器应用上



基于Socket通信方式对外提供服务，或者是实现了某个具体业务的一个特定的TCP server进程



容器提供了强大的隔离功能，所以有必要把为Service提供服务的这组进程放入容器中进行隔离。

为此设计了 **Pod对象**，将每个服务进程包装到相应的Pod中，使其成为Pod中运行的一个容器Container。

建立Service和Pod键的关联关系：

1、先给Pod贴上标签Label

2、然后给相应的Service 定义标签选择器 Label Selector



Pod运行在一个我们称之为节点Node的环境中，

这个节点可以是物理机，也可以是私有云或者公有云中的一个虚拟机

一个节点上可以运行几百个Pod

每个Pod里运行着一个特殊的被称之为Pause的容器，

其他容器则为业务容器，这些业务容器共享Pause容器的网络栈和Volume挂载卷

设计时可以将一组密切相关的服务进程放入同一个pod中

只有那些提供服务的一组Pod才会被映射程一个服务





```yaml
apiVersion: v1
kind: ReplicationController # 副本控制器  表明资源对象的类型
metadata:
  name: mysql # RC的名称 全局唯一
spec:  # RC的相关定义 
  replicas: 1 # pod副本期待数量
  selector:
    app: mysql  # 符合目标的Pod拥有此标签
  template: # 根据此模版创建pod的副本（实例）
    metadata:
      name: mysql
      labels: # 指定pod的标签 == 8:的selector
        app: mysql #pod副本拥有的标签，对应RC的selector
    spec:
      containers:
        - name: mysql #容器名称
          image: mysql # 镜像名称
          ports:
          - containerPort: 3306
          env:
          - name: MYSQL_ROOT_PASSWORD
            value: "root"


# kubectl get rc  查看创建的rc
# kubectl get pods    查看pods 创建时的情况
# kubectl get svc  查看刚刚产检的service ，可以通过CLUSTER-IP 来访问

#CLUSTER-IP  是service创建后由k8s系统自动分配的，k8s巧妙的使用linux环境变量来解决这个问题
# 根据service的唯一名字，容器可以从环境变量中获取到service对应的CLUSTER-IP 和端口号
```



```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: myweb
spec:
  replicas: 2
  selector:
    app: myweb
  template:
    metadata:
      name: myweb
      labels:
        app: myweb
    spec:
      containers:
        - name: myweb
          image: kubeguide/tomcat-app:v1
          ports:
            - containerPort: 8080

# kubectl create -f myweb-rc.yaml  创建pod
```



```yaml
kind: Service
apiVersion: v1
metadata:
  name:  myweb
spec:
  selector:
    app:  myweb
  type:  NodePort # 外网访问模式
  ports:
  - name:  name-of-the-port
    port:  8080
    nodePort: 30001 # 外网可以通过该端口访问此服务
  
```







# 基本概念

集群的两种管理角色 Master 和Node

## Master

集群控制节点，负责整个集群的管理和控制，基本上所有控制命令都发给它，它来负责具体的执行过程，master通常会占据一个独立的服务器，master可以布置为集群。master如果宕机，则k8s集群容器应用的管理都将失效。



运行的关键进程：

- Kubernetes API Server （kube-apiserver）：提供http rest接口的关键服务进程，k8s里面所有资源的crud唯一入口，也是集群控制的入口进程。
- Kubernetes Controller Manager（kube-controller-manager）：k8s所有资源对象的自动化控制中心，可以理解为资源对象的大总管。
- Kubernets Scheduler（kube-scheduler）：负责资源调度（pod调度）的进程，相当于公交的调度室

etcd服务，k8s里面所有资源对象的数据全部保存在etcd中的



## Node

集群中的其他机器被称为node节点，可以是一台物理主机，也可以是一台虚拟机

node节点才是k8s集群中的工作负载节点

每个node都会被master分配一些工作负载，宕机时，会被master自动转移到其他节点



运行的关键进程：

- kubelet：负责pod对于的容器的创建、启动等任务，与master协作，实现集群管理的基本功能。
- kube-proxy：实现kubernetes service 的通讯与负载均衡机制
- Docker Engine（docker）：docker引擎，负责本级的容器创建和管理



node可以在运行期间动态增加到k8s集群中，前提是这个节点正确安装和配置上面的进程，

在默认情况下kubelet会向master注册自己，会定时向Master汇报自身的情报



```yaml
kubectl get nodes # 查看节点
kubectl describe node  <NODENAME>  # 查看NODENAME节点详细信息
```





```
~ $ kubectl describe node docker-deskto
Name:               docker-desktop
Roles:              master
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=docker-desktop
                    kubernetes.io/os=linux
                    node-role.kubernetes.io/master=
Annotations:        kubeadm.alpha.kubernetes.io/cri-socket: /var/run/dockershim.sock
                    node.alpha.kubernetes.io/ttl: 0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
CreationTimestamp:  Fri, 27 Sep 2019 23:09:41 +0800
Taints:             <none>
Unschedulable:      false
Conditions:
  Type             Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----             ------  -----------------                 ------------------                ------                       -------
  MemoryPressure   False   Sun, 20 Oct 2019 22:13:46 +0800   Wed, 09 Oct 2019 12:21:51 +0800   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure     False   Sun, 20 Oct 2019 22:13:46 +0800   Wed, 09 Oct 2019 12:21:51 +0800   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure      False   Sun, 20 Oct 2019 22:13:46 +0800   Wed, 09 Oct 2019 12:21:51 +0800   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready            True    Sun, 20 Oct 2019 22:13:46 +0800   Wed, 09 Oct 2019 12:21:51 +0800   KubeletReady                 kubelet is posting ready status
Addresses:
  InternalIP:  192.168.65.3
  Hostname:    docker-desktop
Capacity:
 cpu:                2
 ephemeral-storage:  61255492Ki
 hugepages-1Gi:      0
 hugepages-2Mi:      0
 memory:             2047132Ki
 pods:               110
Allocatable:
 cpu:                2
 ephemeral-storage:  56453061334
 hugepages-1Gi:      0
 hugepages-2Mi:      0
 memory:             1944732Ki
 pods:               110
System Info:
 Machine ID:                 5c10876a-44ae-4724-937c-8d1e91ba76f8
 System UUID:                6F9B4FAF-0000-0000-88A8-FFE00D266CA4
 Boot ID:                    6d06f7f5-74d0-4873-819a-64ccc9d3720c
 Kernel Version:             4.9.184-linuxkit
 OS Image:                   Docker Desktop
 Operating System:           linux
 Architecture:               amd64
 Container Runtime Version:  docker://19.3.1
 Kubelet Version:            v1.14.3
 Kube-Proxy Version:         v1.14.3
Non-terminated Pods:         (16 in total)
  Namespace                  Name                                         CPU Requests  CPU Limits  Memory Requests  Memory Limits  AGE
  ---------                  ----                                         ------------  ----------  ---------------  -------------  ---
  default                    apple-app                                    0 (0%)        0 (0%)      0 (0%)           0 (0%)         20d
  default                    mysql-z57r5                                  0 (0%)        0 (0%)      0 (0%)           0 (0%)         106m
  default                    myweb-ksqjg                                  0 (0%)        0 (0%)      0 (0%)           0 (0%)         30m
  default                    myweb-x2ghw                                  0 (0%)        0 (0%)      0 (0%)           0 (0%)         30m
  docker                     compose-6c67d745f6-dlwf9                     0 (0%)        0 (0%)      0 (0%)           0 (0%)         22d
  docker                     compose-api-57ff65b8c7-w5tnw                 0 (0%)        0 (0%)      0 (0%)           0 (0%)         22d
  ingress-nginx              nginx-ingress-controller-75d5587cb9-qrt9b    0 (0%)        0 (0%)      0 (0%)           0 (0%)         20d
  kube-system                coredns-fb8b8dccf-8vchk                      100m (5%)     0 (0%)      70Mi (3%)        170Mi (8%)     22d
  kube-system                coredns-fb8b8dccf-qx7m4                      100m (5%)     0 (0%)      70Mi (3%)        170Mi (8%)     22d
  kube-system                etcd-docker-desktop                          0 (0%)        0 (0%)      0 (0%)           0 (0%)         22d
  kube-system                kube-apiserver-docker-desktop                250m (12%)    0 (0%)      0 (0%)           0 (0%)         22d
  kube-system                kube-controller-manager-docker-desktop       200m (10%)    0 (0%)      0 (0%)           0 (0%)         22d
  kube-system                kube-proxy-s6rxb                             0 (0%)        0 (0%)      0 (0%)           0 (0%)         22d
  kube-system                kube-scheduler-docker-desktop                100m (5%)     0 (0%)      0 (0%)           0 (0%)         22d
  kube-system                kubernetes-dashboard-5f7b999d65-b2hfj        0 (0%)        0 (0%)      0 (0%)           0 (0%)         20d
  kube-system                tiller-deploy-57fb6f4785-4hl5k               0 (0%)        0 (0%)      0 (0%)           0 (0%)         22d
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource           Requests    Limits
  --------           --------    ------
  cpu                750m (37%)  0 (0%)
  memory             140Mi (7%)  340Mi (17%)
  ephemeral-storage  0 (0%)      0 (0%)
Events:              <none>

```


## pod

每个pod都有一个特殊的被称为"根容器"的Pause容器

还包含一个或多个用户业务容器

引入业务无关的根容器代表整个容器组的状态，解决真整个容器的状态问题

pod里面的多个业务容器共享Pause容器的Ip、Volume，简化了业务容器之间的通信、文件共享问题

每个pod都分配了唯一的IP，一个pod里的多个容器共享podip，底层网络支持集群内任意两个pod之间的TCP/IP直接通信，通常采用虚拟两层网络技术来实现。

在k8s里，一个pod里的容器与另外主机上的pod容器能够直接通信。

静态pod并不存放在etcd存储里面，而是放在某个具体的node上的一个具体文件中，并且只在此node上启动允许

普通的pod一但被创建，就会被放入到etcd中存储

pod里面的某个容器停止时，k8s会自动检测到这个问题并且重新启动这个pod里所有的容器

pod所在的node宕机，则会将这个node上的所有pod重新调度到其他节点上

 pod的ip 加上容器端口，就组成了Endpoint，代表此pod里的一个服务进程对外通信地址，
 hello-pod.yml
 
 
 pod volume 是定义在pod之上，然后被各个容器挂载到自己的文件系统中的。
 
 Event 是一个事件的记录，记录了事件的相关信息。
 通常会关联到某个具体的资源上，是排查故障的重要参考信息。

千分之一的cpu配额为最小单位，用m表示，通常是500m，表示0.5个cpu

Memory配额也是一个绝对值，单位是内存字节数


Request：该资源的最小申请量，系统必须满足要求。

Limits：该资源最大允许使用的量，不能被突破，试图超过时，可能会被kill并重启




## label

是一个key=value的键值对

可以附加到各种资源对象上

一个资源对象可以定义任意数量的label

给指定的资源对象捆绑一个或多个不同的label来实现多维度的资源分组管理功能。

然后可以通过 label 
selector 查询和筛选

还可以用表达式来表示
- name = 
- env !=
- name in(label2,label3)
- name ont in (label4)

matchLabels 用于定义一组label




## Peplication Controller

RC其实是定义了一个期望的场景，即声明某种pod的副本数量在任意时刻都符合某个预期值，包含：
- pod期待的副本数量replicas
- 用于筛选目标pod的label 
selector
- 当pod 副本数量小于预期数量时,用于创建新pod的pod模版

例： mysql-rc.yaml

master节点上的controller manager组件就会得到通知，定期巡检，确定目标pod实例的数量刚好等于rc的期望值

删除rc并不会影响通过该rc已创建好的pod，删除所有pod，可以设置replicas的值为0，然后更新该rc

kubectl提供了stop、delete命令来一次性删除rc和rc控制的全部pod


Replication Set：rs  与rc的唯一区别是支持基于集合的label 
selector

主要被Deployment这个更高层的资源对象所使用，从而形成一套pod create\update\del 的编排机制


rs与deployment这两个重要资源对象逐步替换了之前的rc的作用，自动 伸缩 功能的基础

改变rc里的pod模版中的镜像版本，可以实现pod的滚动升级



## Deployment

引入的目的是为了更好地解决pod的编排问题，

内部使用rs来实现

hello-deployment.yaml

`kubectl get deployments`

- desired：副本数量的期望值
- current：当前replica的值
- up-to-date：最新版本的pod副本数量，用于指示多个pod副本已经升级成功
- available：当前集群中可用的pod副本数量

 


## Horizontal Pod Autoscaler

**kubectl scale** 可以实现pod扩容或所容


pod智能扩容的特性，横向自动扩容，简称HPA

也属于一种资源对象

通过追踪分析rc控制的所有目标pod的负载变化情况，来确定是否需要针对性的调整目标pod的副本数

指标：
- cpu utilization percentage：目标pod所有副本自身的cpu利用率的平均值，通常是1min内的
- 应用程序自定义的度量指标，



## StatefulSet
有很多服务是有状态的，特别是一些复杂的中间件

本质上可以看做是deployment/rc的一个特殊变种，
- statefulset里的每个pod都有稳定、唯一的网络标识，可以发现集群内的其他成员
- statefulSet控制的pod副本的启动顺序是受控的，前n个已经是运行且准备好的状态
- statefulSet里面的pod 采用稳定持久化存储卷，通过pv/pvc来实现，删除pod时默认不会删除statefulSet相关的存储卷

每个satefulset的定义中要声明它属于哪个Headless service （没有ip）

如果解析headless service的DNS域名，则返回的是该
service对应的全部pod的Endpoint列表


## Service
每个service其实就是微服务中的一个"微服务"

之前的pod、rc等资源其实都是为该
service做的前提准备


service 定义了一个服务的访问入口地址，

前端的应用pod通过这个入口地址访问背后的一组由pod副本组成的集群实例


service与其后端的pod副本集群则是通过label 
selector 来实现无缝对接

rc的作用就是保证service的服务能力和服务质量始终处于预期的标准



---


每个pod都会被分配一个单独的IP地址

每个pod 都提供了一个独立的Endpoint（pod ip + container port）以被客户端访问，

多个pod副本组成一个集群，部署一个lb为这组pod开启一个对外的服务端口，

将这些pod的endpoint列表加入8000端口的转发列表中

客户端就可以通过lb的对外ip地址+服务端口来访问此服务


每个node 上的kube-proxy进程其实就是一个lb，负责吧对
service的请求转发到厚度俺的某个pod实例上，**并在内部实现服务的lb与会话保持机制**

serivce不是共用一个lb的ip地址，而是每个
service 分配了全局唯一的虚拟ip地址，这个虚拟的ip被称为 cluster ip，让每个服务就成为了具备唯一IP地址的通信节点，服务调用就变成了最基础的tcp网络通信问题




service 一但被创建，就会自动分配一个可用的cluster ip，在
service的整个生命周期内，ip不会发生改变，


只要用
service的name与
service的clusterIP地址做一个dns域名映射就可以解决服务发现的问题


默认targetPort与port相同



**服务发现机制**

每个servic都有唯一的 cluster ip和唯一的名字

每个servi生成一些对应的linux环境变量env，并且在每个pod的容器启动时，自动注入这些环境变量


后来通过add-on增值包的方式引入了dns系统，把服务名作为dns域名，这样以来程序就可以直接使用服务名来建立通信连接了


三种ip
- Node IP ；node 节点的IP地址，每个节点的物理网卡的IP地址，所有属于这个网络的服务器之间都能通过这个网络直接通信，
- pod IP ：是docker 根据docker0网桥的IP地址段进行分配的，通常是一个虚拟的二层网络，真实的tcp/ip流量则是通过node ip所在的物理网卡流出
- Cluster IP ：service的ip，虚拟ip。
   
  
Cluster Ip 仅仅作用于service对象，并由k8s管理和分配IP地址

Cluster IP 无法被ping ，因为没有一个实体网络对象来响应

CLuster IP 只能结合Service Port 组成一个具体的通信端口，单独的ClusterIp 不具备tcp/ip通信的基础，并且它们属于k8s这样一个封闭的空间，

ks集群内，Node IP  、Pod IP 、 Cluster Ip 网之间的通信，采用的是ks自己设计的一种编程方式的特殊的路由规则，与现有的ip路由有很大的不同

myweb-svc.yaml  

通过 type：NodePort 手动指定端口，然后http://<nodePort IP >:nodePort  进行访问


## volume 
volume 是pod中能够被多个容器访问的共享目录，

volume定义在pod上，然后被一个pod里的多个容器挂载到具体的文件目录下

生命周期与pod相公，但与容器的生命周期不相关，容器终止或者重启volume中的数据也不会丢失

- emptyDir：是pod分配到node时创建的，临时空间
- hostPath：挂载宿主机上的文件或目录，永久保存，
- gcePersistentDisk：表示使用谷歌公有云提供的永久磁盘
- awsElasticBlockStore：亚马逊公有云
- NFS：网络文件系统提供的共享目录存储数据时，在系统中部署一个nfs server





## Persistent Volume
可以理解成ks集群中的某个网络存储中对应的一块存储
- 只能是网络存储，不属于任何node，但可以在每个node上访问
- 独立于pod之外定义


pv状态

- Available：空闲状态
- Bound：已经绑定到某个pvc上
- Released：对应的pvc已经删除，但资源还没有被集群收回
- Failed：pv自动回收失败

## Namespace

实现多租户的资源隔离

## Annotation
与label类似，也使用k/v键值对的形式进行定义，

定义的是ks对象的元数据Metadata，并且用于Label Selector



