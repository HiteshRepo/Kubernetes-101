In this post we shall solve the lightning labs section of `CKAD` course by [KodeKloud](https://kodekloud.com)

## Lab - 1

### Q1

Create a Persistent Volume called `log-volume`. 
It should make use of a storage class name `manual`. 
It should use `RWX` as the access mode and have a size of `1Gi`. 
The volume should use the `hostPath /opt/volume/nginx`.

Next, create a PVC called `log-claim` requesting a minimum of `200Mi` of storage. 
This PVC should bind to log-volume.

Mount this in a pod called logger at the location `/var/www/nginx`. 
This pod should use the image `nginx:alpine`.

<details> <summary> Solution </summary>

The solution will have 3 steps in this sequence:
- create a PersistentVolume based on given specifications
- create a PersistentVolumeClaim, that will bind to above PersistentVolume, based on given specifications
- finally create a Pod with given specification that will bind to the above PersistentVolume via  PersistentVolumeClaim

Lets create the PersistentVolume.
The manifest `pv.yaml` will look like this:

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: log-volume
spec:
  accessModes:
  - ReadWriteMany
  storageClassName: manual
  capacity:
    storage: 1Gi
  hostPath:
    path: /opt/volume/nginx
```

create and validate the PersistentVolume

```sh
# create
kubectl apply -f pv.yaml #persistentvolume/log-volume created

# validate
kubectl get pv

NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
log-volume   1Gi        RWX            Retain           Available           manual                  18s
# notice the `STATUS` column --> `Available`
```

Now lets create the PersistentVolumeClaim
The manifest `pvc.yaml` will look like this:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: log-claim
spec:
  accessModes:
  - ReadWriteMany
  storageClassName: manual
  resources:
    requests:
      storage: 200Mi
```

create and validate the PersistentVolumeClaim
```sh
# create
kubectl apply -f pvc.yaml # persistentvolumeclaim/log-claim created

# validate
kubectl get pvc
NAME        STATUS   VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
log-claim   Bound    log-volume   1Gi        RWX            manual         13s

# notice the `STATUS` column

# re-check the pv's STATUS and CLAIM columns
kubectl get pv

NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM               STORAGECLASS   REASON   AGE
log-volume   1Gi        RWX            Retain           Bound    default/log-claim   manual                  99s
```

Finally lets create the Pod. Generate the manifest imperatively
```sh
kubectl run logger --image=nginx:alpine --dry-run=client -oyaml > pod.yaml
```

Edit the generated manifest and populate volume specs.
The final manifest `pod.yaml` would look like this:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: logger
  labels:
    run: logger
spec:
  containers:
  - image: nginx:alpine
    name: logger
    volumeMounts:
    - name: log-claim
      mountPath: /var/www/nginx
  volumes:
  - name: log-claim
    persistentVolumeClaim:
      claimName: log-claim
```

create and validate the Pod

```sh
# create
kubectl apply -f pod.yaml # pod/logger created

# validate
kubectl get po

NAME           READY   STATUS    RESTARTS   AGE
logger         1/1     Running   0          20s
```

</details>

### Q2

We have deployed a new pod called secure-pod 
and a service called secure-service. 

Incoming or Outgoing connections to this pod are not working.

Troubleshoot why this is happening.

Make sure that incoming connection from the pod webapp-color are successful.

<details> <summary> Solution </summary>

First lets check the resources already created.

Check the pod specifications: `kubectl describe po secure-pod`

Check the service specifications: `kubectl describe svc secure-service`

Is the service correctly attached the pod?
```sh
kubectl describe po secure-pod | grep Labels
# Labels:           run=secure-pod

kubectl describe svc secure-service | grep Selector
# Selector:          run=secure-pod
```

since Labels and Selctor match. The answer is `yes`.

Is the incoming request to the pod `secure-pod` from `webapp-color` pod not working?
```sh
kubectl get po secure-pod -o wide
# note the IP --> 10.244.0.7

kubectl exec -it webapp-color -- /bin/sh -c 'ping 10.244.0.7'
```

since there are no replies for the pings, this confirms no incoming connection for webapp-color.

But what is stopping? Maybe there is a network policy. Lets check.
```sh
kubectl get netpol

NAME           POD-SELECTOR   AGE
default-deny   <none>         11m

kubectl describe netpol default-deny
Name:         default-deny
Namespace:    default
Created on:   2023-08-13 03:55:32 -0400 EDT
Labels:       <none>
Annotations:  <none>
Spec:
  PodSelector:     <none> (Allowing the specific traffic to all pods in this namespace)
  Allowing ingress traffic:
    <none> (Selected pods are isolated for ingress connectivity)
  Not affecting egress traffic
  Policy Types: Ingress
```

The above network policy:
- applies to all pods in default namespace.
- blocks all incoming incoming connections

Let create a new network policy `secure-netpol` to:
- only apply to `secure-po`
- allow incoming traffic `ingress` from `webapp-color` pod

The manifest `netpol.yaml` would look like this:
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: secure-netpol
  namespace: default
spec:
  podSelector:
    matchLabels:
      run: secure-pod
  policyTypes:
  - Ingress
  ingress:
    - from:
      - podSelector:
          matchLabels:
            name: webapp-color
status: {}
```

create the netpol and verify
```sh
kubectl apply -f netpol.yaml

kubectl exec -it webapp-color -- /bin/sh -c 'ping 10.244.0.7' # there are replies now
```

</details>

### Q3

Create a pod called `time-check` in the `dvl1987` namespace. 
This pod should run a container called `time-check` that uses the `busybox` image.
The `time-check` container should run the command: `while true; do date; sleep $TIME_FREQ;done` and write the result to the location `/opt/time/time-check.log`.
The path `/opt/time` on the pod should mount a volume that lasts the lifetime of this pod.

Create a config map called `time-config` with the data `TIME_FREQ=10` in the same namespace.

<details> <summary> Solution </summary>

Lets verify the existence of namespace: 
```sh
kubectl get ns | grep dvl1987 # does not exist

# so lets create it
kubectl create ns dvl1987 # namespace/dvl1987 created
```

Now lets create the config map imperatively: `kubectl create cm time-config -n dvl1987 --from-literal=TIME_FREQ=10`

verify the config map
```sh
kubectl describe cm time-config -n dvl1987

Name:         time-config
Namespace:    dvl1987
Labels:       <none>
Annotations:  <none>

Data
====
TIME_FREQ:
----
10
```

Now lets generate the pod manifest imperatively:
```sh
kubectl run time-check -n dvl1987 --image=busybox --dry-run=client -oyaml > time-check.yaml
```

Edit the manifest to populate commands, add volumes and map the config map `time-config` for `TIME_FREQ` data.
The final manifest `time-check.yaml` would look like this:
```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: time-check
  name: time-check
  namespace: dvl1987
spec:
  containers:
  - image: busybox
    name: time-check
    resources: {}
    env: 
      - name: TIME_FREQ
        valueFrom:
          configMapKeyRef:
            name: time-config
            key: TIME_FREQ
    command: ["/bin/sh", "-c", "while true; do date; sleep $TIME_FREQ; done > /opt/time/time-check.log"]
    volumeMounts:
      - name: log
        mountPath: /opt/time
  volumes:
    - name: log
      emptyDir: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

Now lets create and verify the pod
```sh
kubectl apply -f time-check.yaml # pod/time-check created

kubectl get po time-check -n dvl1987

NAME         READY   STATUS    RESTARTS   AGE
time-check   1/1     Running   0          2m3s

# check if the TIME_FREQ environment variable is set
kubectl exec time-check -n dvl1987 -- /bin/sh -c 'echo $TIME_FREQ' # 10

# check if the time-check.log file is being populated
kubectl exec time-check -n dvl1987 -- /bin/sh -c 'cat /opt/time/time-check.log'
```

</details>

### Q4

Create a new deployment called `nginx-deploy`, 
with one single container called `nginx`, image `nginx:1.16` and `4` replicas.
The deployment should use `RollingUpdate` strategy with `maxSurge=1`, and `maxUnavailable=2`.

Next `upgrade` the deployment to version `1.17`.

Finally, once all pods are updated, `undo the update` and go back to the previous version.

<details> <summary> Solution </summary>

Lets generate the deployment manifest imperatively:
```sh
kubectl create deploy nginx-deploy --image=nginx:1.16 --replicas=4 --dry-run=client -oyaml > nginx-deploy.yaml
```

Edit the manifest and update the RollingUpdate strategy.
The final manifest `nginx-deploy.yaml` would look like this:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx-deploy
  name: nginx-deploy
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 2
  replicas: 4
  selector:
    matchLabels:
      app: nginx-deploy
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx-deploy
    spec:
      containers:
      - image: nginx:1.16
        name: nginx
        resources: {}
status: {}
```

Now lets update the image to `nginx:1.17`: `kubectl set image deploy nginx-deploy nginx=nginx:1.17` --> deployment.apps/nginx-deploy image updated

Check the rollout status: `kubectl rollout status deploy nginx-deploy` --> deployment "nginx-deploy" successfully rolled out

Undo the rollout: `kubectl rollout undo deploy nginx-deploy` --> deployment.apps/nginx-deploy rolled back

Check the rollout history: `kubectl rollout history deploy nginx-deploy`

Check the current image: `kubectl describe deploy nginx-deploy | grep Image` --> Image: nginx:1.16

</details> 


### Q5

Create a redis deployment with the following parameters:
- Name of the deployment should be `redis` using the `redis:alpine` image. It should have exactly `1` replica.
- The container should request for `.2 CPU`. It should use the label `app=redis`.
- It should mount exactly 2 volumes.
  a. An Empty directory volume called `data` at path `/redis-master-data`.
  b. A configmap volume called `redis-config` at path `/redis-master`.
- The container should expose the port `6379`.


The configmap has already been created.

<details> <summary> Solution </summary>

Let generate the deployment imperatively:

```sh
kubectl create deploy redis --image=redis:alpine --replicas=1 --port=6379 --dry-run=client -oyaml > redis.yaml
```

Edit the manifest to populate resource requests and volumes.
The final manifest `redis.yaml` would look like this:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: redis
  name: redis
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: redis
    spec:
      containers:
      - image: redis:alpine
        name: redis
        ports:
        - containerPort: 6379
        resources:
          requests:
            cpu: .2
        volumeMounts:
          - name: data
            mountPath: /redis-master-data
          - name: redis-config
            mountPath: /redis-master
      volumes:
        - name: data
          emptyDir: {}
        - name: redis-config
          configMap:
            name: redis-config
status: {}
```

create and validate the deployment

```sh
# create

kubectl get deploy

NAME           READY   UP-TO-DATE   AVAILABLE   AGE
nginx-deploy   4/4     4            4           14m
redis          1/1     1            1           33s

# validate

kubectl get po

NAME                           READY   STATUS    RESTARTS   AGE
...
redis-6b5b7895c7-nbxj5         1/1     Running   0          47s
...
```

</details> 