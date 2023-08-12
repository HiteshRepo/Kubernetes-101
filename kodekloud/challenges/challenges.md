In this post we shall solve the challenges section of `CKAD` course by [KodeKloud](https://kodekloud.com)

## Challenge - 1

Below is the architecture asked to deploy

![Architecture](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/architecture.png)

Since the question says that the PersistentVolume `jekyll-site` has already created, so lets verify:

```sh
kubectl get pv

NAME          CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS    REASON   AGE
jekyll-site   1Gi        RWX            Delete           Available           local-storage            5m39s
```

Next since the whole architecture has been create under namespace `development`, let verify is the namespace exists or not:

```sh
kubectl get ns | grep development

development       Active   7m22s
```

Lets start by deploying/creating other components of the architecture. We will try to use imperative commands wherever necessary to quickly finish the challenge.

Lets deploy PersistentVolumeClaim `jekyll-site` first.
Below are the requirements:

![PersistentVolumeClaim](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/pvc.png)

We will require the `storageClassName` of the PersistentVolume to which we want to attach the PersistentVolumeClaim.
So that would be: 
```sh
kubectl describe pv jekyll-site | grep StorageClass

StorageClass:      local-storage
```

The manifest `pvc.yaml` would look like this:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: jekyll-site
  namespace: development
spec:
  accessModes:
    - ReadWriteMany
  storageClassName: local-storage
  resources:
    requests:
      storage: 1Gi
```

Lets create and validate:
```sh
# create
kubectl apply -f pvc.yaml

# validate
kubectl get pvc -n development

NAME          STATUS   VOLUME        CAPACITY   ACCESS MODES   STORAGECLASS    AGE
jekyll-site   Bound    jekyll-site   1Gi        RWX            local-storage   6m32s

kubectl get pv

NAME          CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                     STORAGECLASS    REASON   AGE
jekyll-site   1Gi        RWX            Delete           Bound    development/jekyll-site   local-storage            21m
```

Both the PersistentVolumeClaim and PersistentVolume have their `STATUS` as `Bound`.
And the `CLAIM` of PersistentVolume says it is bound to `development/jekyll-site` which means it is claimed by PV named `jekyll-site` in `development` namespace.

Lets create and deploy Role `developer-role`
Below are the requirements:

![Role](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/role.png)

The manifest `role.yaml` would look like this:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: development
  name: developer-role
rules:
- apiGroups: [""]
  resources: ["pods", "services", "persistentvolumeclaims"]
  verbs: ["*"]
```

Lets create and validate the role:
```sh
# create
kubectl apply -f role.yaml

role.rbac.authorization.k8s.io/development-role created

# validate
kubectl get role -n development

NAME               CREATED AT
development-role   2023-08-12T06:48:01Z
```

Lets create and deploy RoleBinding `developer-rolebinding`
Below are the requirements:

![Role Binding](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/role-binding.png)

The manifest `binding.yaml` would look like this:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: developer-rolebinding
  namespace: development
subjects:
- kind: User
  name: martin
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: developer-role
  apiGroup: rbac.authorization.k8s.io
```

Lets create and validate the role:
```sh
# create
kubectl apply -f binding.yaml

rolebinding.rbac.authorization.k8s.io/developer-rolebinding created

# validate
kubectl get rolebinding -o wide -n development

NAME                    ROLE                  AGE   USERS    GROUPS   SERVICEACCOUNTS
developer-rolebinding   Role/developer-role   30s   martin
```

Now lets create the Pod `jekyll`
Below are the requirements:

![Jekyll Pod](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/jekyll-pod.png)

Lets use imperative commands to generate the Pod manifest:

```sh
k run jekyll -n development --image=kodekloud/jekyll-serve --dry-run=client -oyaml > pod.yaml
```

Then now lets edit and populate other requirements. The final manifest `pod.yaml` would look like this:

```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: jekyll
  name: jekyll
  namespace: development
spec:
  initContainers:
    - image: kodekloud/jekyll
      name: copy-jekyll-site
      command: ["jekyll", "new", "/site"]
      volumeMounts:
        - name: site
          mountPath: "/site"
  containers:
  - image: kodekloud/jekyll-serve
    name: jekyll
    resources: {}
    volumeMounts:
      - name: site
        mountPath: "/site"
  volumes:
    - name: site
      persistentVolumeClaim:
        claimName: jekyll-site
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

Lets create and validate the pod:
```sh
# create
kubectl apply -f pod.yaml 

pod/jekyll created

# validate
kubectl get po -n development

NAME     READY   STATUS    RESTARTS   AGE
jekyll   1/1     Running   0          33s
```

Now lets create the Service `jekyll`
Below are the requirements:

![Jekyll service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/jekyll-svc.png)

Lets use imperative commands to generate the Service manifest:

```sh
kubectl expose -n development pod jekyll --name=jekyll --type=NodePort --port=8080 --target-port=4000 --dry-run=client -oyaml > svc.yaml
```
Then now lets edit and populate the node port. The final manifest `svc.yaml` would look like this:

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    run: jekyll
  name: jekyll
  namespace: development
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 4000
    nodePort: 30097
  selector:
    run: jekyll
  type: NodePort
status:
  loadBalancer: {}
```

Lets create and validate the pod:
```sh
# create
kubectl apply -f svc.yaml 

service/jekyll created

# validate
kubectl get svc -n development

NAME     TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
jekyll   NodePort   10.109.111.12   <none>        8080:30097/TCP   31s

curl -v http://10.109.111.12:8080 # should give 200 OK
```

Now lets update the `kube-config` file to add user `martin` and a new context `developer`.
Below are the requirements:

![Kube Config](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/kube-config.png)

![Martin](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-1/dist/user-martin.png)

The diff of `kube-config` file would be as below:
```yaml
...
contexts:
- context:
    cluster: kubernetes
    user: martin
  name: developer
  namespace: development
...
users:
- name: martin
  user:
    client-certificate: /root/martin.crt
    client-key: /root/martin.key
...
```

Now lets set the current context to `developer` context

```sh
k config use-context developer

Switched to context "developer".
```

Now verify if we are able to access pods, services and persistent volumes or not.

Finally check your answer.

## Challenge - 3

Below is the architecture asked to deploy

![Architecture](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/architecture.png)

Lets start by deploying/creating other components of the architecture. We will try to use imperative commands wherever necessary to quickly finish the challenge.

since the whole architecture has been create under namespace `vote`, let verify is the namespace exists or not:

```sh
kubectl get ns | grep vote
# no output
```

Since the `vote` namespace does not exist, lets create it

```sh
kubectl create ns vote
```

Lets create `vote-deployment`
Below is the requirement:

![Vote Deployment](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/vote-deploy.png)

Lets use imperative commands to create the deployment:

```sh
kubectl create deploy -n vote vote-deployment --image=kodekloud/examplevotingapp_vote:before
```

Verify the pods are in Running state or not:
```sh
kubectl get po -n vote -l app=vote-deployment

NAME                              READY   STATUS    RESTARTS   AGE
vote-deployment-8d495c5b7-vj82w   1/1     Running   0          103s
```

Lets create `vote-service`
Below is the requirement:

![Vote Service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/vote-service.png)

Lets use imperative commands to generate the manifest of the service:
```sh
kubectl expose -n vote deployment vote-deployment --name=vote-service --type=NodePort --port=5000 --target-port=80 --dry-run=client -oyaml > vote-svc.yaml
```

Now lets populate the node port in the generated manifest.
The final manifest `vote-svc.yaml` would look like:

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: vote-deployment
  name: vote-service
  namespace: vote
spec:
  ports:
  - port: 5000
    protocol: TCP
    targetPort: 80
    nodePort: 31000
  selector:
    app: vote-deployment
  type: NodePort
status:
  loadBalancer: {}
```

Lets create and verify the service:

```sh
# create
kubectl apply -f vote-svc.yaml 

service/vote-service created

kubectl get svc -n vote

NAME           TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
vote-service   NodePort   10.99.180.188   <none>        5000:31000/TCP   9s

curl -v http://10.99.180.188:5000 ## should give 200 OK
```

Now lets create the redis deployment.
Below is the requirement:

![Redis Deployment](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/redis-deploy.png)


Lets use imperative commands to generate the manifest of the deployment:

```sh
kubectl create deploy redis-deployment -n vote --image=redis:alpine --dry-run=client -oyaml > redis-deploy.yaml
```

Now lets edit the manifest and populate the `volumes` and `volumeMounts`.
The final manifest `redis-deploy.yaml` would look like this:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: redis-deployment
  name: redis-deployment
  namespace: vote
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis-deployment
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: redis-deployment
    spec:
      containers:
      - image: redis:alpine
        name: redis
        resources: {}
        volumeMounts:
          - name: redis-data
            mountPath: /data
      volumes:
        - name: redis-data
          emptyDir: {}
status: {}
```

Lets create and verify the deployment:

```sh
# create
kubectl apply -f redis-deploy.yaml

deployment.apps/redis-deployment created

# verify
kubectl get deploy -n vote

NAME               READY   UP-TO-DATE   AVAILABLE   AGE
redis-deployment   1/1     1            1           80s
vote-deployment    1/1     1            1           13m
```

Lets create `redis-service`
Below is the requirement:

![Redis Service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/redis-svc.png)

Lets use imperative commands to create the service:
```sh
kubectl expose -n vote deployment redis-deployment --name=redis --type=ClusterIP --port=6379 --target-port=6379
```

verify the service:
```sh
kubectl get svc -n vote
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
redis          ClusterIP   10.106.60.131   <none>        6379/TCP         26s
vote-service   NodePort    10.99.180.188   <none>        5000:31000/TCP   10m
```

Now lets create the worker deployment.
Below is the requirement:

![Worker Deployment](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/worker-deploy.png)

Lets use imperative commands to create the deployment:
```sh
kubectl create deploy -n vote worker --image=kodekloud/examplevotingapp_worker
```

verify the deployment:
```sh

```

Now lets create the db deployment:
Below is the requirement:

![DB deployment](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/db-deployment.png)

Lets use imperative commands to generate the manifest:
```sh
kubectl create deploy -n vote db-deployment --image=postgres:9.4 --dry-run=client -oyaml > db-deploy.yaml
```

Now lets edit the deployment and populate environment and volumes.
The final manifest `db-deploy.yaml` would look like this:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: db-deployment
  name: db-deployment
  namespace: vote
spec:
  replicas: 1
  selector:
    matchLabels:
      app: db-deployment
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: db-deployment
    spec:
      containers:
      - image: postgres:9.4
        name: postgres
        resources: {}
        env:
          - name: POSTGRES_HOST_AUTH_METHOD
            value: trust
        volumeMounts:
          - name: db-data
            mountPath: /var/lib/postgresql/data
      volumes:
        - name: db-data
          emptyDir: {}
```

lets create and validate the deployment

```sh
# create
kubectl apply -f db-deploy.yaml

deployment.apps/db-deployment created

# validate
kubectl get deploy -n vote

NAME               READY   UP-TO-DATE   AVAILABLE   AGE
db-deployment      1/1     1            1           73s
redis-deployment   1/1     1            1           15m
vote-deployment    1/1     1            1           27m
worker             1/1     1            1           7m42s
```

Now lets create db service.
Below is the requirement:

![DB service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/db-service.png)

Lets use imperative commands to create the service:
```sh
kubectl expose deploy -n vote db-deployment --name=db --type=ClusterIP --port=5432 --target-port=5432
```

verify the service:
```sh
kubectl get svc -n vote
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
db             ClusterIP   10.100.153.37   <none>        5432/TCP         30s
redis          ClusterIP   10.106.60.131   <none>        6379/TCP         16m
vote-service   NodePort    10.99.180.188   <none>        5000:31000/TCP   26m
```

Now lets create the result deployment
Below is the requirement:

![Result deployment](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/result-deployment.png)

Lets use imperative commands to create the deployment:
```sh
kubectl create deploy -n vote result-deployment --image=kodekloud/examplevotingapp_result:before
```

Verify the deployment
```
# validate
kubectl get deploy -n vote

NAME                READY   UP-TO-DATE   AVAILABLE   AGE
db-deployment       1/1     1            1           7m42s
redis-deployment    1/1     1            1           22m
result-deployment   1/1     1            1           45s
vote-deployment     1/1     1            1           34m
worker              1/1     1            1           14m
```

Now lets create result service.
Below is the requirement:

![DB service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-3/dist/result-svc.png)

Lets use imperative commands to generate the service manifest:
```sh
kubectl expose deploy -n vote result-deployment --name=result-service --type=NodePort --port=5001 --target-port=80 --dry-run=client -oyaml > result-svc.yaml
```

Now lets edit and populate the node port.
The final manifest `result-svc.yaml` would look like this:
```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: result-deployment
  name: result-service
  namespace: vote
spec:
  ports:
  - port: 5001
    protocol: TCP
    targetPort: 80
    nodePort: 31001
  selector:
    app: result-deployment
  type: NodePort
status:
  loadBalancer: {}
```

create and verify the service:
```sh
# create
kubectl apply -f result-svc.yaml

service/result-service created


kubectl get svc -n vote
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
db               ClusterIP   10.100.153.37   <none>        5432/TCP         7m19s
redis            ClusterIP   10.106.60.131   <none>        6379/TCP         22m
result-service   NodePort    10.108.38.244   <none>        5001:31001/TCP   16s
vote-service     NodePort    10.99.180.188   <none>        5000:31000/TCP   32m
```

We have deployed all the components.

Finally check your answer.

## Challenge - 4

Below is the architecture asked to deploy

![Architecture](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/architecture.png)

Lets create the persistent volumes.

There are 6 PersistentVolume resources that needs to deployed and all of them similar specifications except for name and path.

Lets create a single manifest to create all the PVs together.
Below are the requirements:

![Redis-1](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis01.png)

![Redis-2](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis02.png)

![Redis-3](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis03.png)

![Redis-4](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis04.png)

![Redis-5](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis05.png)

![Redis-6](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis06.png)


The manifest `pvs.yaml` would like this:

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis01
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis01

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis02
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis02

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis03
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis03

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis04
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis04

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis05
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis05

---

apiVersion: v1
kind: PersistentVolume
metadata:
  name: redis06
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 1Gi
  hostPath:
    path: /redis06
```

lets create and verify the persistent volumes:

```sh
# create
kubectl apply -f pvs.yaml 

persistentvolume/redis01 created
persistentvolume/redis02 created
persistentvolume/redis03 created
persistentvolume/redis04 created
persistentvolume/redis05 created
persistentvolume/redis06 created

# verify
kubectl get pv

NAME      CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
redis01   1Gi        RWO            Retain           Available                                   52s
redis02   1Gi        RWO            Retain           Available                                   52s
redis03   1Gi        RWO            Retain           Available                                   52s
redis04   1Gi        RWO            Retain           Available                                   52s
redis05   1Gi        RWO            Retain           Available                                   52s
redis06   1Gi        RWO            Retain           Available                                   52s
```

Now lets create the `redis-cluster-service` service before stateful set because stateful set requires a service name.
Below are the requirements:

![Redis Cluster Service](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/redis-service.png)

The manifest `redis-svc.yaml` would look like this:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: redis-cluster-service
spec:
  clusterIP: None
  ports:
  - name: client
    targetPort: 6379
    port: 6379
  - name: gossip
    targetPort: 16379
    port: 16379
```

create and validate the service

```sh
# create
kubectl apply -f redis-svc.yaml 

service/redis-cluster-service created

# validate
kubectl get svc

NAME                    TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)              AGE
kubernetes              ClusterIP   10.96.0.1    <none>        443/TCP              3h22m
redis-cluster-service   ClusterIP   None         <none>        6379/TCP,16379/TCP   6s
```

Now lets create the `redis-cluster` stateful-set

Below are the requirements:

![Redis Cluster](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/statefulset.png)

The manifest `statefulset.yaml` would look like this:

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: redis-cluster
spec:
  selector:
    matchLabels:
      app: redis-cluster
  serviceName: redis-cluster-service
  replicas: 6
  template:
    metadata:
      name: redis-cluster
      labels:
        app: redis-cluster
    spec:
      containers:
      - image: redis:5.0.1-alpine
        name: redis
        command: ["/conf/update-node.sh", "redis-server", "/conf/redis.conf"]
        env:
        - name: POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        ports:
        - containerPort: 6379
          name: client
        - containerPort: 16379
          name: gossip
        volumeMounts:
        - name: conf
          mountPath: /conf
          readOnly: false
        - name: data
          mountPath: /data
          readOnly: false
      volumes:
      - name: conf
        configMap:
          name: redis-cluster-configmap
          defaultMode: 0755
  volumeClaimTemplates:
    - metadata:
       name: data
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi

```

create and validate the stateful set

```sh
# create
kubectl apply -f statefulset.yaml

statefulset.apps/redis-cluster created

# validate
kubectl get po
NAME              READY   STATUS    RESTARTS   AGE
redis-cluster-0   1/1     Running   0          23s
redis-cluster-1   1/1     Running   0          18s
redis-cluster-2   1/1     Running   0          15s
redis-cluster-3   1/1     Running   0          12s
redis-cluster-4   1/1     Running   0          9s
redis-cluster-5   1/1     Running   0          6s
```

execute command as mentioned in `redis-cluster-config`

![Redis Cluster Config](https://github.com/HiteshRepo/Kubernetes-101/blob/master/kodekloud/challenges/challenge-4/dist/validate.png)

```sh
kubectl exec -it redis-cluster-0 -- redis-cli --cluster create --cluster-replicas 1 $(kubectl get pods -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 {end}')

>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 10.244.192.4:6379 to 10.244.192.1:6379
Adding replica 10.244.192.5:6379 to 10.244.192.2:6379
Adding replica 10.244.192.6:6379 to 10.244.192.3:6379
M: 17c25a0bd004f2372c74cf8d9e56b1c70946fd54 10.244.192.1:6379
   slots:[0-5460] (5461 slots) master
M: 98c0922e0a13d3efa9377f3f76c878da843f8a1e 10.244.192.2:6379
   slots:[5461-10922] (5462 slots) master
M: 37ff6f3eb0aca87682296e4bd09d96cd4b88cacb 10.244.192.3:6379
   slots:[10923-16383] (5461 slots) master
S: 86ae291916b863e0cd5b132579eaf9f33343fa59 10.244.192.4:6379
   replicates 17c25a0bd004f2372c74cf8d9e56b1c70946fd54
S: 83e503793a6845e776211df61d274b98bc1a5837 10.244.192.5:6379
   replicates 98c0922e0a13d3efa9377f3f76c878da843f8a1e
S: d889b408104efba53a9800d8eb8944e193008a04 10.244.192.6:6379
   replicates 37ff6f3eb0aca87682296e4bd09d96cd4b88cacb
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
........
>>> Performing Cluster Check (using node 10.244.192.1:6379)
M: 17c25a0bd004f2372c74cf8d9e56b1c70946fd54 10.244.192.1:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: d889b408104efba53a9800d8eb8944e193008a04 10.244.192.6:6379
   slots: (0 slots) slave
   replicates 37ff6f3eb0aca87682296e4bd09d96cd4b88cacb
M: 98c0922e0a13d3efa9377f3f76c878da843f8a1e 10.244.192.2:6379
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: 83e503793a6845e776211df61d274b98bc1a5837 10.244.192.5:6379
   slots: (0 slots) slave
   replicates 98c0922e0a13d3efa9377f3f76c878da843f8a1e
S: 86ae291916b863e0cd5b132579eaf9f33343fa59 10.244.192.4:6379
   slots: (0 slots) slave
   replicates 17c25a0bd004f2372c74cf8d9e56b1c70946fd54
M: 37ff6f3eb0aca87682296e4bd09d96cd4b88cacb 10.244.192.3:6379
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

We have deployed all the components.

Finally check your answer.