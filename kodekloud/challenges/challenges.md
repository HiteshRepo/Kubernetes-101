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