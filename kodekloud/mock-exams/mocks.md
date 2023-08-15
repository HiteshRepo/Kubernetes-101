In this post we shall solve the lightning labs section of `CKAD` course by [KodeKloud](https://kodekloud.com)

## Mock - 1

### Q1

Deploy a pod named `nginx-448839` using the `nginx:alpine` image.

<details> <summary> Solution </summary>

```sh
kubectl run nginx-448839 --image=nginx:alpine

kubectl get po

NAME           READY   STATUS    RESTARTS   AGE
nginx-448839   1/1     Running   0          4s
```

</details>

### Q2

Create a namespace named `apx-z993845`

<details> <summary> Solution </summary>

```sh
kubectl create ns apx-z993845

kubectl get ns

NAME              STATUS   AGE
apx-z993845       Active   75s
```

</details>

### Q3

Create a new Deployment named `httpd-frontend` with `3` replicas using image `httpd:2.4-alpine`

<details> <summary> Solution </summary>

```sh
kubectl create deploy httpd-frontend --image=httpd:2.4-alpine --replicas=3

kubectl get deploy

NAME             READY   UP-TO-DATE   AVAILABLE   AGE
httpd-frontend   3/3     3            3           40s
```

</details>

### Q4

Deploy a `messaging` pod using the `redis:alpine` image with the labels set to `tier=msg`.

<details> <summary> Solution </summary>

```sh
kubectl run messaging --image=redis:alpine

kubectl label po messaging tier=msg

kubectl get po -l tier=msg

NAME        READY   STATUS    RESTARTS   AGE
messaging   1/1     Running   0          94s
```

</details>

### Q5

A replicaset `rs-d33393` is created. 
However the pods are not coming up. Identify and fix the issue.
Once fixed, ensure the ReplicaSet has `4 Ready` replicas.

<details> <summary> Solution </summary>

spot the issue: invalid image
```sh
kubectl describe rs rs-d33393 | grep Image

# Image:      busyboxXXXXXXX
```

get the replica set manifest
```sh
kubectl get rs rs-d33393 -o yaml > rs.yaml
```

delete the old replica set
```sh
kubectl delete rs rs-d33393
```

change the image and apply
```sh
kubectl apply -f rs.yaml
```

</details>

### Q6

Create a service `messaging-service` to expose the `redis` deployment in the `marketing` namespace within the cluster on port `6379`.

<details> <summary> Solution </summary>

```sh
kubectl expose -n marketing deploy redis --port 6379 --name=messaging-service

kubectl get svc

NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   90m
```

</details>

### Q7

Update the environment variable on the pod `webapp-color` to use a green background.

<details> <summary> Solution </summary>

```sh
kubectl get po webapp-color -oyaml > webapp-color.yaml

# update the environment variable APP_COLOR to green

# delete the po
kubectl delete po webapp-color

# re-create the po
kubectl apply -f webapp-color.yaml
```

</details>

### Q8

Create a new ConfigMap named cm-3392845. Use the spec given on the below.

ConfigName Name: cm-3392845
- Data: DB_NAME=SQL3322
- Data: DB_HOST=sql322.mycompany.com
- Data: DB_PORT=3306

<details> <summary> Solution </summary>

the manifest `cm.yaml` would look like this
```yaml
apiVersion: v1
kind: ConfigMap
metadata: 
  name: cm-3392845
data:
  DB_NAME: "SQL3322"
  DB_HOST: "sql322.mycompany.com"
  DB_PORT: "3306"
```

directly using imperative command:
```sh
kubectl create cm cm-3392845 --from-literal=DB_NAME=SQL3322 --from-literal=DB_HOST=sql322.mycompany.com --from-literal=DB_PORT=3306
```

create and validate
```sh
# create
kubectl apply -f cm.yaml

# validate
kubectl describe cm cm-3392845

Name:         cm-3392845
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
DB_HOST:
----
sql322.mycompany.com
DB_NAME:
----
SQL3322
DB_PORT:
----
3306
```

</details>

### Q9

Create a new Secret named db-secret-xxdf with the data given (on the below).
Secret Name: db-secret-xxdf
- Secret 1: DB_Host=sql01
- Secret 2: DB_User=root
- Secret 3: DB_Password=password123

<details> <summary> Solution </summary>

```sh
kubectl create secret generic db-secret-xxdf --from-literal=DB_User=root --from-literal=DB_Host=sql01 --from-literal=DB_Password=password123
```

validate
```sh
kubectl describe secret db-secret-xxdf

controlplane ~ ➜  kubectl describe secret db-secret-xxdf
Name:         db-secret-xxdf
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
DB_Password:  11 bytes
DB_User:      4 bytes
DB_Host:      5 bytes
```

</details>

### Q10

Update pod app-sec-kff3345 to run as Root user and with the SYS_TIME capability.
Pod Name: app-sec-kff3345
Image Name: ubuntu
SecurityContext: Capability SYS_TIME

<details> <summary> Solution </summary>

get the pod spec:
```sh
kubectl get po app-sec-kff3345 -oyaml > app-sec-kff3345.yaml
```

update the container:
```yaml
...
securityContext:
  runAsUser: 0
  capabilities:
    add: ["SYS_TIME"]
...
```

delete the pod
```sh
kubectl delete po app-sec-kff3345
```

re-create the pod
```sh
kubectl apply -f app-sec-kff3345.yaml
```

validate
```sh
kubectl get po app-sec-kff3345

NAME              READY   STATUS    RESTARTS   AGE
app-sec-kff3345   1/1     Running   0          44s
```

</details>


### Q11

Export the logs of the `e-com-1123` pod to the file `/opt/outputs/e-com-1123.logs`

<details> <summary> Solution </summary>

check the pod namespace
```sh
kubectl get po -A | grep e-com-1123

e-commerce    e-com-1123                             1/1     Running   0              49m
```

export the logs
```sh
kubectl logs e-com-1123 -n e-commerce > /opt/outputs/e-com-1123.logs
```

</details>


### Q12

Create a Persistent Volume with the given specification.
- Volume Name: pv-analytics
- Storage: 100Mi
- Access modes: ReadWriteMany
- Host Path: /pv/data-analytics

<details> <summary> Solution </summary>

the manifest `pv.yaml` would look like this

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-analytics
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 100Mi
  hostPath:
    path: /pv/data-analytics
```

</details>

### Q13

Create a redis deployment using the image `redis:alpine` with `1` replica and label `app=redis`. 
Expose it via a ClusterIP service called redis on port `6379`. 
Create a new Ingress Type NetworkPolicy called `redis-access` which allows only the pods with label `access=redis` to access the deployment.

<details> <summary> Solution </summary>

create the pod
```sh
kubectl create deploy redis --image=redis:alpine --replicas=1
```

create the service
```sh
kubectl expose deploy redis --port=6379
```

create a network policy manifest and deploy it
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: redis-access
spec:
  podSelector:
    matchLabels:
      app: redis
  policyTypes:
    - Ingress
  ingress:
    - from:
        - podSelector:
            matchLabels:
              access: redis
      ports:
        - protocol: TCP
          port: 6379
```

</details>

### Q14

Create a Pod called `sega` with two containers:
- Container 1: Name `tails` with image `busybox` and command: `sleep 3600`.
- Container 2: Name `sonic` with image `nginx` and Environment variable: `NGINX_PORT` with the value `8080`.

<details> <summary> Solution </summary>

create the pod manifest and deploy it
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: sega
spec:
  containers:
  - name: tails
    image: busybox
    command: ["sleep", "3600"]
  - name: sonic
    image: nginx
    env:
    - name: NGINX_PORT
      value: "8080"
```

</details>

## Mock - 2

### Q1

Create a deployment called `my-webapp` with image: `nginx`, label `tier:frontend` and `2` replicas. 
Expose the deployment as a `NodePort` service with name `front-end-service` , port: `80` and NodePort: `30083`

<details> <summary> Solution </summary>

create the deployment
```sh
kubectl create deploy my-webapp --image=nginx --replicas=2
```

label the deployment
```sh
kubectl label deploy my-webapp tier=frontend
```

create the service
```sh
kubectl expose deploy my-webapp --name=front-end-service --type=NodePort --port=80 --dry-run=client -oyaml > my-webapp.svc.yaml

# edit the service manifest to add node port
nodePort: 30083

# create the svc
kubectl apply -f my-webapp.svc.yaml
```

</details>

### Q2

Add a taint to the node `node01` of the cluster. Use the specification below:
- key: app_type
- value: alpha
- effect: NoSchedule

Create a pod called `alpha`, image: `redis` with toleration to `node01`.

<details> <summary> Solution </summary>

taint the node
```sh
kubectl taint node node01 app_type=alpha:NoSchedule # node/node01 tainted
```

verify the taint
```sh
kubectl describe node node01 | grep Taints # Taints: app_type=alpha:NoSchedule
```

pod manifest `pod.yaml` with toleration:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: alpha
spec:
  containers:
  - image: redis
    name: alpha
  tolerations:
    - key: app_type
      value: alpha
      effect: NoSchedule
```

verify
```sh
kubectl get po -o wide

kubectl get po -o wide
NAME                         READY   STATUS    RESTARTS   AGE   IP             NODE           NOMINATED NODE   READINESS GATES
alpha                        1/1     Running   0          27s   10.244.192.3   node01         <none>           <none>
```

</details>

### Q3

Apply a label `app_type=beta` to node `controlplane`. 
Create a new deployment called `beta-apps` with image: `nginx` and replicas: `3`. 
Set Node Affinity to the deployment to place the PODs on `controlplane` only.

<details> <summary> Solution </summary>

label the node
```sh
kubectl label node controlplane  app_type=beta
```

create the deployment
```sh
kubectl create deploy beta-apps --image=nginx --replicas=3 --dry-run=client -oyaml > beta-apps.yaml
```

update the manifest to add node affinity
```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
      - matchExpressions:
        - key: app_type
          operator: In
          values: ["beta"]
```

apply the manifets `beta-apps.yaml`
```sh
kubectl apply -f beta-apps.yaml
```

validate
```sh
kubectl get po -owide

NAME                         READY   STATUS    RESTARTS   AGE    IP             NODE           NOMINATED NODE   READINESS GATES
beta-apps-574fd8858c-gpk2s   1/1     Running   0          18s    10.244.0.6     controlplane   <none>           <none>
beta-apps-574fd8858c-nqpxd   1/1     Running   0          18s    10.244.0.7     controlplane   <none>           <none>
beta-apps-574fd8858c-xsq8r   1/1     Running   0          18s    10.244.0.5     controlplane   <none>           <none>
```

</details>

### Q4

Create a new Ingress Resource for the service `my-video-service` to be made available at the URL: http://ckad-mock-exam-solution.com:30093/video.

To create an ingress resource, the following details are: -
- annotation: nginx.ingress.kubernetes.io/rewrite-target: /
- host: ckad-mock-exam-solution.com
- path: /video

<details> <summary> Solution </summary>

get the service details
```sh
kubectl get svc my-video-service -o yaml

apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2023-08-15T07:39:48Z"
  name: my-video-service
  namespace: default
  resourceVersion: "5038"
  uid: 67a6728b-f6c5-43c8-8c11-171050c1807d
spec:
  clusterIP: 10.103.124.68
  clusterIPs:
  - 10.103.124.68
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: webapp-video
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```


the ingress manifest `ingress.yaml` would look like:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: video-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: ckad-mock-exam-solution.com
    http:
      paths:
      - path: /video
        pathType: Prefix
        backend:
          service:
            name: my-video-service
            port:
              number: 8080
```

deploy the ingress
```sh
kubectl apply -f ingrss.yaml
```

validate
```sh
curl http://ckad-mock-exam-solution.com:30093/video
<!doctype html>
<title>Hello from Flask</title>
<body style="background: #30336b;">

<div style="color: #e4e4e4;
    text-align:  center;
    height: 90px;
    vertical-align:  middle;">
    <img src="https://res.cloudinary.com/cloudusthad/image/upload/v1547052431/video.jpg">

</div>

</body>
```

</details>

### Q5

We have deployed a new pod called `pod-with-rprobe`. 
This Pod has an `initial delay` before it is Ready. 
Update the newly created pod pod-with-rprobe with a `readinessProbe` using the given spec:
- httpGet path: /ready
- httpGet port: 8080

<details> <summary> Solution </summary>

fetch the pod manifest
```sh
kubectl get po pod-with-rprobe -oyaml > pod-with-rprobe.yaml
```

check the pod spec to get the delay
```yaml
...
- env:
    - name: APP_START_DELAY
      value: "180"
...
```

update the manifest with readiness probe
```yaml
readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 180
```

delete the pod
```sh
kubectl delete po pod-with-rprobe
```

apply the updated manifest
```sh
kubectl apply -f pod-with-rprobe.yaml 
```

validate after 180 seconds
```sh
NAME              READY   STATUS    RESTARTS   AGE
pod-with-rprobe   1/1     Running   0          185s
```

</details>

### Q6

Create a new pod called `nginx1401` in the default namespace with the image `nginx`. 
Add a livenessProbe to the container to restart it if the command `ls /var/www/html/probe` fails. 
This check should start after a `delay of 10 seconds` and `run every 60 seconds`.

<details> <summary> Solution </summary>

generate the pod manifest
```sh
kubectl run nginx1401 --image=nginx --dry-run=client -oyaml > nginx1401.yaml
```

update the manifest to add liveness probe
```yaml
livenessProbe:
  initialDelaySeconds: 10
  periodSeconds: 60
  exec:
    command: ["ls", "/var/www/html/probe"]
```

apply the manifest
```sh
kubectl apply -f nginx1401.yaml
```

validate
```sh
kubectl get po nginx1401 

NAME        READY   STATUS    RESTARTS   AGE
nginx1401   1/1     Running   0          24s
```

</details>

### Q7

Create a job called `whalesay` with image `docker/whalesay` and command `cowsay I am going to ace CKAD!`.
- completions: 10
- backoffLimit: 6
- restartPolicy: Never

<details> <summary> Solution </summary>

the manifest `whalesay.yaml` would look like this

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: whalesay
spec:
  completions: 10
  backoffLimit: 6
  template:
    metadata:
      name: whalesay
    spec:
        containers:
        - image: docker/whalesay
          name: whalesay
          command: ["/bin/sh", "-c", "cowsay I am going to ace CKAD!"]
        restartPolicy: Never

```

apply and validate
```sh
kubectl apply -f whalesay.yaml

kubectl get job

NAME       COMPLETIONS   DURATION   AGE
whalesay   5/10          44s        44s

kubectl get po 

whalesay-6dcvg                  0/1     Completed           0               22s
whalesay-cvhpf                  0/1     Completed           0               27s
whalesay-dfhcq                  0/1     Completed           0               37s
whalesay-f2xzk                  0/1     Completed           0               17s
whalesay-h722v                  0/1     Completed           0               32s
whalesay-jm8zv                  0/1     Completed           0               42s
whalesay-rfzlf                  0/1     Completed           0               62s
whalesay-wd92n                  0/1     ContainerCreating   0               1s
whalesay-wjw9z                  0/1     Completed           0               7s
whalesay-xprkx                  0/1     Completed           0               12s
```

</details>

### Q8

Create a pod called multi-pod with two containers.
- Container 1:
  name: jupiter, image: nginx

- Container 2:
  name: europa, image: busybox
  command: sleep 4800

Environment Variables:
- Container 1:
  type: planet

- Container 2:
  type: moon

<details> <summary> Solution </summary>

the pod manifest would look like this

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: multi-pod
spec:
  containers:
  - name: jupiter
    image: nginx
    env:
    - name: type
      value: "planet"
  - name: europa
    image: busybox
    command: ["sleep", "4800"]
    env:
    - name: type
      value: "moon"
```

</details>

### Q9

Create a PersistentVolume called `custom-volume` with:
 - size: `50MiB` 
 - reclaim policy:`retain`
 - Access Modes: `ReadWriteMany`
 - hostPath: `/opt/data`

<details> <summary> Solution </summary>

the pv manifest would look like this

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: custom-volume
spec:
  capacity:
    storage: 50Mi
  persistentVolumeReclaimPolicy: Retain
  accessModes:
  - ReadWriteMany
  hostPath:
    path: /opt/data
```

</details>