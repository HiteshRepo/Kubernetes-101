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