## High Level Arch
1. Master Node
    1. Api server - face of k8 master, every c’ happens via api server
    2. Schedulers - schedule workloads to worker nodes
    3. Control manager - compare state mentioned in request [desired] and actual state, then act accordingly
    4. Etcd - distributed key value store - only stateful component - source of truth
2. Worker Nodes
    1. Kubelet - take request from master and fulfil them, reports to master node
    2. Docker runtime - to run containers - OCI compliant container engine, deals with container abstraction
    3. Kube proxy - manage n/w b/w worker nodes, Assigns IP to eat pod with the help of CNI provider
    4. Pods
3. Flow
    1. Client sends a request - To keep infra in a particular state
    2. Api server receives request and save it to Etcd
    3. Ctrl manager keeps looking at Etcd to notice any differences b/w current state and desired state
    4. Once decision has been made on what needs to be changed in pods, scheduler assign actual pod configuration to worker node
    5. Kubelet in worker node keeps listening to the api server in Master node
    6. Kubelet uses docker runtime to spin up new pods with mentioned configuration
    7. The new IPs of pods and routes definition are done by Kube proxy - IP table route
4. Even Kubernetes components like api server, controller, scheduler, kubeproxy, etc run as pods


## Kubectl
1. CLI to communicate with k8 api server
2. Restful communication
3. kubectl [command] [type] [name] [flags]
4. Commands - get, patch, delete
5. Type - pods, services, jobs
6. Flags - -o (wide)
7. Connects to API server of K8 master node
8. Use rest apis to do that
9. Kubeconfig - info related to :
    1. Cluster info
    2. User info
    3. Namespace
10. Default loc of kubeconfig - $HOME/.kube/config
11. KUBECONFIG env var

## k8 commands

* kubectl version
* kubectl version —short - client version and server version
* kubectl get nodes
* kubectl get nodes -o wide
* kubectl config view - to get cluster info, user info and namespace
* kubectl get config get-contexts
* kubectl get pods
* kubectl get pods -A -o wide
* kubectl apply -f file.yaml
* kubectl delete <object-type>/<object-name>
* kubectl describe <object-type>/<object-name>
* kubectl get pods —show-labels
* kubectl get svc
* kubectl get endpoints
* kubectl describe endpoints svc-name
* kubectl rollout history deployment/<deploy-name>
* kubectl rollout undo deployment/<deploy-name> --to-revision=1
* kubectl cordon node-name -> no further pod will be scheduled here -> STATUS: SchedulingDisabled
* kubectl replace -f <file name> -> Replaces existing configuration with latest, works same as apply
* kubectl scale --replicas=6 <type> <name>
* minikube ip
* minikube ssh - to connect to minikube
* eval $(minikube docker-env) - to make docker point to minikube docker context

## Formatting o/p
1. -o json -> in json formatted API object
2. -o name -> only name of the resource
3. -o wide -> additional info in plain-text format
4. -o yaml -> YAML formatted API object

## Minikube Objects
1. Persistent entity in K8s system and rep state of system
2. Includes:
    1. Spec - desired/requested state
    2. Status - current state
3. Also called API resources
4. Smallest deployable unit - pods
5. Abstraction on top of pods - replica-set, stateful-set, daemon-set, job and cron-job, services and ingress
6. Abstraction on top of Replicaset - Deployment
7. Volumes, PVC,PV, Storage Class
8. ConfigMap and Secrets
9. Object descriptor YAML - to communicate our desired state
10. Parts of object descriptor file:
    1. apiVersion,
    2. kind [of object],
    3. metadata [info about object, name - unique identifier, labels]
    4. Spec - actual specification of the object to be created
11. Replication controller (same purpose as Replica Set) -
    1. Replica set is recommended,
    2. Replication controller is an older concept
    3. Replication controller does not have 'selector' under spec, but Replica Set has
    4. Selector helps Replica Set to attach any already running pods to itself or any other pods that can be started individually in future

## Pods
1. Smallest unit
2. Run inside nodes
3. Can run multiple pods in 1 node
4. Pods are a wrapper over containers
5. Multiple containers in a pod is possible and they share the same container env, but best practice is to run 1 container/pod unless other containers are monitoring/tracking apps
6. Ring-fenced env
    1. Network stack
    2. Volume mounts
    3. Kernel namespace
7. High level Pod lifecycle -
    1. Kubectl -> API server
    2. API server -> Etcd
    3. Scheduler reads from Etcd -> Node [kubelet/worker]
    4. Pod - pending
    5. Pod - Running / Failed
    6. Pod - Success
8. Intra pod communication
    1. Containers within pod talk to each other via localhost
    2. Share same n/w namespace, hence same IP and Port
    3. Container within Pod to avoid same port, use to avoid port binding error
9. Inter pod communication
    1. Each pod gets own private IP from k8 cluster vpn
10. Container specs tags
    1. name
    2. image
    3. command
    4. args
    5. workingDir
    6. ports
    7. env
    8. resources
    9. volumeMounts
    10. livenessProbe
    11. readinessProbe
    12. lifecycle
    13. terminationMessagePath
    14. imagePullPolicy
    15. securityContext
    16. stdin
    17. stdinOnce
    18. tty

## Replica sets
    1. Abstraction over pods, which ensures that a particular no. of pods is always running in the cluster
    2. Uses Reconciliation control loop -> Current state - Desired State - Observe-Diff-Act
    3. Ensures that a pod or homogeneous set of pods are always available
    4. Maintains desired no. of pods:
        1. Excess pods - killed
        2. Launch new pod - in case of fail/deleted/terminated
    5. Associated with pods via matching labels
    6. Labels: Key-Value pair attached to objects like pod - user defined
    7. Selectors: Help identify objects in cluster - equality based / set based
    8. apiVersion - apps/v1
    9. kind - ReplicaSet
    10. metadata - name, labels…
    11. spec - 
        1. replicas
        2. selector - matchLabels - app
        3. template - pod specification - prevents specifying separate pod yaml
    12. Distributes pods evenly across nodes
    13. Deleting replica set -> deletes associated pods as well

## Health check probes for containers:
These diagnostics are performed periodically - in template section of replicaset/deployments - httpGet [path] /exec [command] - initialDelaySeconds and periodSeconds
1. readinessProbe - indicates if container is ready to serve requests, halts sending new requests until probe succeed - in template section of replicaset/deployments - httpGet/exec - initialDelaySeconds and periodSeconds
2. livenessProbe - indicates whether the container is running healthy, if fails, declares container unhealthy and restarts container
3. startupProbe - protect slow starting containers with startup probes

### Supported check types
1. httpGet - /health endpoint
2. exec - shell script or command to exit successfully with return code 0
3. tcpSocket - open a socket to container on specified port successfully


## Services
1. Pods are ephemeral
2. They are recreated and not resurrected
3. Services are abstraction of a way to expose an app running on a set of pods by reliable network svc.
4. Exposes pod over a reliable IP, Port, DNS
5. Associated with pods via matching labels
6. Also used for inter pod communication
7. Client -> service [DNS/IP] -> Endpoint object [list of all pod IP address associated with svc, keeps getting updated]
8. Types:
    1. ClustedIP - default - cluster-internal IP only access within n/w
    2. NodePort - exposes node on a static port - NodeIP:NodePort
    3. LoadBalancer - Exposes service publicly
9. apiVersion - v1
10. kind - service
11. metadata - name
12. spec - type, selector - app [same as replicaset/template/metadata/name or pod/metadata/name]
13. ports - protocol, port, targetPort
14. Deleting pod or replica sets does not affect svc but just removes them from endpoints. Upon new spin ups, services will update the endpoints based on label-selector
15. Readiness and Liveliness probe also affect the endpoints

## Deployments
1. How to deploy a new version of app?
2. How to roll back?
3. Is replica set good enough?
4. Change in rs.pod spec - no effect
5. Delete and re-deploy rs - change effected
6. Updates with zero downtime
7. Rollbacks
8. A higher level of abstraction over replica set, provides declarative way of upgrading and rollbacks to pods
9. Flow:
    1. Current state - RS 1
    2. Client -> Revision 2 -> API server
    3. Scheduler + Control Manager -> spin up RS 2, pods created
    4. Terminate pods in RS1
    5. RS 1 still persists -> so that during rollback, the can be used
10. The diff b/w replica-set and deployment is the kind
11. Default strategy - RollingUpdate - maxSurge, maxUnavailable
12. Recreate strategy -> downtime

## Volumes
1. Containers are ephemeral
2. We require persistent storage
3. Types:
    1. emptyDir -
        1. No data at start,
        2. created when pods get created,
        3. mounted and accessible across all containers in the pod
        4. Help sharing data across containers
        5. spec -> volumes/name : html, volumes/emptyDir: {}
        6. spec/containers -> volumeMounts/name : html, volumeMounts/mountPath: <path-inside-container>
        7. Good option to share data b/w container but data is lost once pod goes down
    2. hostPath -
        1. Storage from backing Node [Host] is mounted inside container [Pod]
        2. Data retained on Node even after Pod goes down
        3. Data not available if Pod is scheduled on another Node
        4. Cant save data from Node outage
        5. spec -> volumes/name : html, volumes/hostPath/path: <existing-worker-node-path>, volumes/hostPath/type: Directory
        6. spec/containers -> volumeMounts/name : html, volumeMounts/mountPath: <path-inside-container>
        7. Good option to shared data across pods in a Node
    3. Cloud volume type -
        1. awsEBS
        2. gcePersistentDisk
        3. azureDisk
    4. Nfs

## PV and PVC
1. Abstracts how storage is provided and how storage is consumed
2. PV
    1. Represent actual volume
    2. Provisioned by Admin or dynamically provisioned using StorageClass
    3. Lifecycle <-> Pod
3. PVC
    1. Represent request for volume by user
    2. Abstract the storage resource without exposing details how those volumes are implemented
    3. Claims are fulfilled by PV hence PVC is linked with PV
4. Retain - Actual volume is retained even after PV and PVC is deleted
5. Delete - Actual physical storage is deleted, default
6. Recycle - Deprecated
7. Access modes
    1. ReadWriteOnce - RWO - volume can mounted by read-write by single node
    2. ReadOnlyMany - ROX - read-only by many nodes
    3. ReadWriteMany - RWX - read-write by many nodes

## Storage Class
1. Provisioning
    1. Static:
        1. Admin creates a number of PVC
        2. Cluster matches one of the PV for a PVC
        3. Only one PVC can be attached for a PV
    2. Dynamic:
        1. Allows storage volumes to be created on-demand as per the request
        2. Claims are fulfilled by PV, hence PVC are linked to PV
2. Helps create dynamic on-demand PVs
3. PVC refers storage class, Storage class provisions PVC on demand, Deployment/ReplicaSet/Pod mount the PV via PVC
4. Basically storage class are template for PVs
5. Provisioners - cloud service providers
6. Parameters - specific to provisioners
7. If PVC is deleted, PV is also gone, id reclaim policy is not set to ‘retain’

## Other sources
1. Link to K8 commands compilation: https://www.evernote.com/shard/s645/sh/18a2e56b-3451-90a2-75b5-2f91ec5ac6ef/3e5b88d59f5bb686d5fb7350cf823e63

## Namespaces
1. resource address format: <resource-name>.<namespace-name>.<resource-type-domain>.cluster.local
2. kubectl create -f <file-name> --namespace=<namespace-name>
3. Also, namespace can be mentioned in metadata of the resource
4. kubectl create namespace <namespace-name>
5. kubectl config set-context $(kubectl config current-context) --namespace=<namespace-name>

## Resource Quota
```
apiVersion: v1
kind: ResourceQuota
metadata:
    name: compute-quota
    namespace: dev
spec:
    hard:
        pods: "10"
        requests.cpu: "4"
        requests.memory: 5Gi
        limits.cpu: "10"
        limits.memory: 10Gi
```

### Imperative commands
1. --dry-run=client -> resource won't be created, instead will tell if resource would be created or not
2. -o yaml -> resource definition in YAML format
3. kubectl run nginx --image=nginx --dry-run=client -o yaml : will not create the resource 'pod' but will give pod declarative definition
4. kubectl create deployment --image=nginx nginx --dry-run -o yaml : will not create the resource 'deployment' but will give deployment declarative definition
5. kubectl create deployment nginx --image=nginx--dry-run=client -o yaml > nginx-deployment.yaml : saves definition to a file
6. kubectl expose pod redis --port=6379 --name redis-service --dry-run=client -o yaml : will not create the resource 'service' but will give service declarative definition

### Commands in Docker/Kubernetes
1. CMD vs EntryPoint - command line args replace CMD while it gets appended in EntryPoint
2. Default can be specified by having both CMD and EntryPoint - CMD instructions are appended to EntryPoint
3. ENTRYPOINT (docker) -> command (k8)
4. CMD (docker) -> args (k8)

### Editing properties of a running Pod
1. Specifications of an existing POD, CANNOT be edited other than the below:
   1. spec.containers[*].image
   2. spec.initContainers[*].image
   3. spec.activeDeadlineSeconds
   4. spec.tolerations
2. The environment variables, service accounts, resource limits of a running pod cannot be edited
3. There are 2 options to achieve though:
   1. Approach 1:
      1. kubectl edit pod <pod name> -> This will open up pod specification in a vi editor
      2. Change the specifications and try to save it -> will through error but will save the changed specifications in a temp file
      3. delete the existing pod: `kubectl delete pod <pod-name>`
      4. create the changed pod: `kubectl create -f <tmp file path>`
   2. Approach 2:
      1. Extract the pod definition in YAML format to a file using the command: `kubectl get pod <pod-name> -o yaml > my-new-pod.yaml`
      2. vi my-new-pod.yaml: changes specifications and save
      3. kubectl delete pod <pod-name>
      4. kubectl create -f my-new-pod.yaml
4. For deployments: `kubectl edit deployment my-deployment`, the new changes will be applied to the pods (running pods will be terminated and new pods with latest specifications will be created)

### Environment variables
1. In pod specifications, under 'env' attribute. This is an array of (Key value pair) name & value.
2. Other ways of specifying env vars are: ConfigMap and Secrets
3. Example of direct key-value pair under 'env'
```
env:
    - name: APP_COLOR
      value: pink
```
4. Example of config-map under 'env'
```
env:
    - name: APP_COLOR
      valueFrom:
        configMapKeyRef: <config-map-name>
```
5. Example of secret under 'env'
```
env:
    - name: APP_COLOR
      valueFrom:
        secretKeyRef: <secret-name>
```

### ConfigMaps
1. Centralized way of configuring configuration data in the form of key-value pairs.
2. When pods are created, these configuration data are injected to the apps inside the container inside the pod for usage
3. Phases: Create config map, inject them into pod
4. Imperative ways of creating a config map
```
kubectl create configmap <config-map-name> --from-literal=key1=value1 --from-literal=key2=value2
kubectl create configmap <config-map-name> --from-file=<path-to-file> 
```
5. Declarative way of creating a config map: apiVersion, kind, metadata, data (ke-value pairs)
```
kubectl apply -f <config-map-definition-file-path>
```
6. kubectl get configmaps
7. kubectl describe configMaps
8. Map config map to pod definition/template
```
envFrom:
    - configMapRef:
        name: <config-map-name>
```

```
volumes:
- name: <volume-name>
  configMap:
    name: <config-map-name>
```

### Secrets
1. Imperative way to create a secret:
```
kubectl create secret generic <secret-name> --from-literal=<key>=<value>
kubectl create secret generic <secret-name> --from-file=<path-to-file>
```

2. Declarative way to create a secret
```
kubectl create -f <secret-file-name>
```
3. Encoded data values in secret definition. Although just encoding is not enough, so it is better to use some KMS decryption
4. kubectl get secrets
5. kubectl describe secrets
6. kubectl describe secrets <secret-name> -o yaml : to view the hashed secrets
7. Map secret to pod definition/template
```
envFrom:
    - secretRef:
        name: <secret-name>
```

```
volumes:
- name: <volume-name>
  secret:
    secretName: <secret-name>
```

If secret is used as volume mount, each attribute in secret is creates its own file and with vaue as contents in it

### Docker Security
1. Host itself runs a set of processes, docker daemon, ssh-server, etc.
2. Docker containers unlike VMs share same linux kernel as the hosts' but they are separated by namespaces
3. Container has its own namespace and host has its own
4. All processes run on container in fact run on host itself but in a different namespace (namespace of container)
5. Docker container can see only see its own processes only
6. Listing processes in a container (ps aux) will only show processes within container
7. Listing processes in the host (ps aux) will show all processes within and out of container(s)
8. Docker container has a set of users root users and a set of non-root users
9. By default, docker runs processes within container as root users
10. User can be changed, user can be set using while running docker using --user flag: `docker run --user=1000 ubuntu sleep 1000`
11. Another way to set user is creating a custom image from an existing image and setting used in the docker file itself
Example dockerfile:
```
FROM ubuntu
USER 1000
```
building the above custom image
```
docker build -t my-ubuntu-image .
```
run the image w/o specifying the user
12. If we run container as a root user, is it not dangerous?
    1. Docker implements the set of security features that limits the capability of the root user within the container
    2. Root user within the container is not really same as root user on host
    3. Docker uses linux capabilities to achieve this
    4. Root user is the most powerful user in a system and can do set of these ops: CHOWN, DAC, KILL, SETGID, SETUID, NET_ADMIN, KILL, etc.
    5. The process running as a root user too has unrestricted access of the system
    6. Docker's root user by default has limited capabilities, they do not have all the privilleges
    7. We can add more capabilities to the container's user while running it: `docker run --cap-add KILL ubuntu`
    8. We can drop capabilities of the container's user while running it: `docker run --cap-drop MAC_ADMIN ubuntu`
    9. We can run container with all privileges as well: `docker run --privilleged ubuntu`

### Security contexts
1. Configuring user id of a container, adding/removing privileges of a user in a k8 is also possible
2. Security settings can be configured at container/pod level
3. If we set at pod level the settings will be applied to all containers within pod
4. If we set at both pod and container level, then settings of container level will take precedence over pod settings
5. Configuration
```
apiVersion: v1
kind: Pod
metadata:
    name: web-app
spec:
    securityContext:
        runAsUser: 1000 #all conatainers within this pod will run with user id 1000
    containers:
        - name: ubuntu
          image: ubuntu
          command: ["sleep", "1000"]
          securityContext:
            runAsUser: 2000 #the user id for this container would be 2000 overrinding 1000
            capabilities: 
                add: ["MAC_ADMIN", "KILL"]
```

### Service Accounts
1. Two types of account in K8: User a/c and Service a/c.
2. User account: used by humans, Service account: for automated tasks(by machines)
3. User account types (not limited to): Admin (to perform admin tasks), Developer(to access the cluster and deploy apps)
4. Service account types are used my an app to interact with k8 cluster, examples:
   1. A monitoring app like Prometheus uses service a/c to poll k8 metrics/logs to come up with performance metrics
   2. An automated build tool like Jenkins uses service a/c to deploy app on the cluster
5. To create a service a/c: `kubectl create serviceaccount <account-name>`
6. To view all service a/c: `kubectl get service a/c`
7. On creation of service a/c a token is created automatically: `kubectl descrive serviceaccount <acocunt-name>` - see Tokens
8. The above token can be used by the external apps for authentication of kube-api as a bearer token.
9. Token is stored as a secret object.
10. To view the secret object: `kubectl describe secret <secret-name>`
11. Steps:
    1. create a service a/c
    2. assign role based permissions/access control mechanisms
    3. export the token
    4. use it in external app while making kube api requests
12. If the external app itself is hosted in K8 cluster, the exporting can be made simpler by mounting the secret as a volume to the application.
13. To view the secret files in the pod (which has secret mounted as volume):
    1. exec into the pod: kubectl exec -it <pod-name>
    2. ls /var/run/secrets/kubernetes.io/serviceaccount -> ca.crt, namespace, token
    3. cat /var/run/secrets/kubernetes.io/serviceaccount/token
14. Default service accounts are mounted automatically to every pods, which has limited permissions.
15. To assign a service account: spec/serviceAccountName: <service a/c name>
16. To prevent k8 from automatically mounting default service a/c : spec/automountServiceAccountToken: false

### Resource Requirements
1. Scheduler decides which node the pod goes to.
   1. Scheduler takes into consideration: the amount of resources by a pod and availability of it in node.
2. If there is no sufficient resources available on any of the nodes, K8 keeps the po in pending state with event reason as insufficient CPU/memory/disk
3. Default CPU: 0.5, MEM: 256 Mi, Disk: (Resource Request)
4. spec/conatiners:
   ```
   resources:
    requests:
        memory: "1Gi"
        cpu: 1
   ```
5. cpu 0.1 means 100m (m -> milli)
6. cpu can be requested as low as 1m
7. 1 cpu equivalent to
   1. 1 AWS vCPU
   2. 1 GCP core
   3. 1 Azure core
   4. 1 Hyperthread
8. 1Gi memory means 1 Gibibyte while 1G means 1 Gigabyte
9. set limits under spec/conatiners/resources, to prevent pod from consuming too much resources and suffocating other pods
   ```
   limits:
    memory: "2Gi"
    cpu: 2
   ```
10. when pod tries to go beyond the limit cpu, k8 tries to throttle the cpu so that pod will not be able to consume more cpu
11. when pod tries to go beyond the limit mem, k8 terminates the pod
12. The status OOMKilled indicates that it is failing because the pod ran out of memory. Identify the memory limit set on the POD

### Taints and Tolerations
1. Taints and tolerations are used to set restrictions on what pods can be scheduled on which node.
2. They have nothing to do with security.
3. Lets' take a use case:
   1. We have 4 pods: A, B, C, D
   2. We have 3 nodes: Node1, Node2, Node3
   3. Now if there are no taints and tolerations configured, then A, B, C, D will be placed on nodes via load balancing/resource management
   4. But suppose we want to place pods like D (running same as in D) to be scheduled only on Node1
   5. Then we apply a taint on Node1, so since until now none of the pods have any sort of tolerations configured, none of the pods will be scheduled in Node1
   6. Now we can enable pod D to be placed on Node1, by adding a toleration on pod D.
4. Taints are placed on nodes and Tolerations are placed on pods.
5. Apply Taints to nodes: `kubectl taint node <node-name> <key>=<value>:<taint-effect>`
6. Taint-Effect determine what happens to the pod if they DO NOT TOLERATE this taint, there are 3 taint-effects
   1. NoSchedule: Pods will not be scheduled
   2. PreferNoSchedule: K8 will try not to schedule pods but with no guarantee
   3. NoExecute: New pods will not be scheduled, but if already there are few pods in the node they will be evicted.
7. Apply Tolerations to pods (@ spec/containers):
    ```
    tolerations:
        - key: "app"
          operator: "Equal"
          value: "blue"
          effect: "NoSchedule"
    ```
8. Taints and tolerations do not guarantee that certain pods will be scheduled on certain nodes only. They enable nodes to accept certain pods but those pods can very well be placed on other nodes. as well.
9. Scheduler does not place any pod on master node: because when K8 cluster is first set up a taint is applied on the master node automatically that prevents placing of other pods on master node.
10. To see the above taint in master node: `kubectl describe node kubemaster | grep Taint`

### Node Selectors
1. There might be use cases where we will require placing certain pods only certain nodes.
2. For example, 
   1. There are 3 nodes (2 nodes with low resources and 1 node with high resources).
   2. We would like to place pods running high processing apps in node with higher resources.
3. The default setup places pods in nodes based on load balancing and resource availability strategy.
4. Also, with taints and tolerations, we can guarantee nodes to accept certain pods but not guarantee placing pods on certain nodes.
5. A simple way to achieve this is using Node Selectors.
6. An example of Pod configuration using node selector 
```
apiVersion: v1
kind: Pod
metadata:
    name: myapp-pod
spec:
    containers:
        - name: data-processor
          image: data-processor
    nodeSelector:
        size: Large
```
7. The key value pair (size: Large) are in fact labels assigned to nodes. Scheduler uses these to assign pods to specific Nodes.
8. To label a node: `kubectl label nodes <node-name> <key>:<value>`
9. Limitations:
   1. Cannot serve complex requirements: if we want to place pod on a large or medium nodes instead of small.
10. Node affinity is the solution here.

### Node Affinity
1. Complex requirements can be executed in Node Affinity.
2. The example used in Node Selectors can be re-defined as this:
```
apiVersion: v1
kind: Pod
metadata:
    name: myapp-pod
spec:
    containers:
        - name: data-processor
          image: data-processor
    affinity:
        nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
                nodeSelectorsTerms:
                - matchExpressions:
                  - key: size
                    operator: In #NotIn, Exists,...
                    values:
                    - Large
```
3. If node affinity does not match any of the rules: 
4. Node affinity types:
   1. requiredDuringSchedulingIgnoredDuringExecution: Pod will not be scheduled if rules do not match (Pods remain in pending state), but pods already running are ignored (irrespective of the rules).
   2. preferredDuringSchedulingIgnoredDuringExecution: Pod will be scheduled in available node if rules do not match, and pods already running are ignored (irrespective of the rules).
   3. requiredDuringSchedulingRequiredDuringExecution: Pod will not be scheduled if rules do not match (Pods remain in pending state), and pods already running are evicted if rules do not match.

### Node Affinity vs Taints and Toleration
1. Lets' take a use case
   1. There are 3 nodes: Red, blue and green. There are other nodes as well.
   2. There are 3 pods: Red, blue and green. There are other pods as well.
   3. Our aim is to put red pod in red node, green pod in green node and blue pod in blue node.
   4. We also do not want any other pods to be placed in our (red, green and blue) nodes.
   5. We also do not want our pods to be placed on other nodes.
2. How to achieve this:
   1. Lets' try with Taints and Toleration first
      1. We apply taints red, blue and green to nodes.
      2. Then we apply tolerations red, blue and green to pods.
      3. This will help in placing pods with appropriate tolerance end up in corresponding tainted node but this does nt guarantee pod ending up in nodes that do not have taints.
   2. Lets' try with Node Affinity
      1. We apply key-value pair labels on nodes.
      2. We then configure nodes with appropriate affinity.
      3. This will help us in placing pods in appropriate nodes but other pods also might end up in our nodes.
3. So a combination of both Taints and Toleration and node affinity is used.

### Multi Container Pods
1. Microservices enable us to develop small, independent, reusable code.
2. Also, it helps us in scaling them.
3. However, at times two services are required to work together such as a web server and a log agent.
4. We want a web server and a log agent paired together, we do not want to merge them and bloat the code though.
5. So we need multi-container pods that share same lifecycle, network space and storage volumes.
6. An example of multi-container setup looks something like below:
```
apiVersion: v1
kind: Pod
metadata:
    name: simple-webapp
    labels:
        name: simple-webapp
spec:
    containers:
    - name: simple-webapp
      image: simple-webapp
      ports:
        - containerPort: 8080
    - name: log-agent
      image: log-agent
```
7. Common design patterns:
   1. SIDECAR: we can run a logging agent along with the main app that will push logs on to a centralized logs-storage  
   2. ADAPTER: sometimes each application produces different format of logs and hence we need to format them before pushing them to centralized system
   3. AMBASSADOR: very often, it is required to connect to different databases based on env. So based on the env we connect to that DB instance. This logic can be extracted out to an ambassador container which can act as a proxy.

### Readiness and Liveness probe
1. A pod has a pod status.
2. The pod status states where is the pod in its lifecycle.
3. If pod is first created, it is in pending state. This is when the scheduler tries to figure out where to place the pod.
4. If scheduler cannot find a node to place the pod, then it remains in pending state.
5. Once the pod is scheduled, it goes into containercreating status, it is when the image is pulled and containers are created.
6. Once all the containers in the pod starts, pod status changes to running state.
7. The pod status remains in running state, unless program in the container is completed or the pod is terminated. 
8. So complete and terminating are the other pod statuses.
9. Pod conditions
   1. PodScheduled
   2. Initialized
   3. ContainersReady 
   4. Ready - indicate app inside the pod is running and ready to accept requests
10. Container could be running various apps within them
    1. A Simple script performing a job, a db service, or a large web server serving end users.
    2. The script may take few milliseconds to get ready
    3. The db service may tale few milliseconds to connect to db and run migration scripts
    4. The webserver might require some seconds to powerup before serving requests
    5. So the apps are not yet ready for those milliseconds to serve any requests
11. W/o readiness probe, the pod continues to indicate being ready even though the underlying containers are powering up
12. So readiness probes are important to let k8s know of the actual state of the containers
13. If Pod is not ready k8s service will not divert request on to it because k8s service relies on pod's ready state to route traffic
14. As developers, we know that when exactly the app is ready to serve requests
15. So we need a way to tie up the actual app's ready state with k8s status indicating ready or not
16. There are a few ways to do so:
    1. HTTP test: /api/ready is responding with correct status code or not
    2. TCP test: TCP socket is up or not
    3. exec command: if command gets executed successfully or not
17. Example of HTTP test readiness probe:
```
apiVersion: v1
kind: Pod
metadata:
    name: simple-webapp
    labels:
        name: simple-webapp
spec:
    conatiners:
    - name: simple-webapp
      image: simple-webapp
      ports:
        - containerPort: 8080
      readinessProbe:
        httpGet: 
            path: /api/ready
            port: 8080
```
18. Example of TCP test:
```
readinessProbe:
    tcpSocket:
        port: 3306
```
19. Example of Exec Command test:
```
readinessProbe:
    exec:
        command:
            - cat
            - /app/is_ready
```
20. We can add additional delay to the probe considering that app might take a few more time to start and hence requires readiness probe to be tested after that time. This can be achieved by 'initialDelaySeconds':
```
readinessProbe:
    httpGet: 
        path: /api/ready
        port: 8080
    initialDelaySeconds: 10
```
21. If we wish to run the probe periodically and change the state of the container based on it. We cam achieve it by 'periodSeconds':
```
readinessProbe:
    httpGet: 
        path: /api/ready
        port: 8080
    initialDelaySeconds: 10
    periodSeconds: 5
```
22. By default, if app is not ready after 3 attempts, the probe will stop and pod will not be sent request to. But we can configure the number of fail attempts by 'failureThreshold':
```
readinessProbe:
    httpGet: 
        path: /api/ready
        port: 8080
    initialDelaySeconds: 10
    periodSeconds: 5
    failureThreshold: 8
```
23. Liveness probe is very much similar as in readiness probe. But in this case the pod is killed upon failing and new instance of the pod is respawned.
24. The configurations stay similar to readiness probe
25. HTTP test
```
apiVersion: v1
kind: Pod
metadata:
    name: simple-webapp
    labels:
        name: simple-webapp
spec:
    conatiners:
    - name: simple-webapp
      image: simple-webapp
      ports:
        - containerPort: 8080
      livenessProbe:
        httpGet: 
            path: /api/healthy
            port: 8080
        initialDelaySeconds: 10
        periodSeconds: 5
        failureThreshold: 8
```
26. TCP test
```
livenessProbe:
    tcpSocket:
        port: 3306
    initialDelaySeconds: 10
    periodSeconds: 5
    failureThreshold: 8
```
27. Exec command test
```
livenessProbe:
    exec:
        command:
            - cat
            - /app/is_ready
    initialDelaySeconds: 10
    periodSeconds: 5
    failureThreshold: 8
```

### Container logging
1. to view logs of a container: `kubectl logs -f <pod-name>` (f option is to stream logs live).
2. if multiple containers are running in a single pod, it would ask for the container name, else it would fail: `kubectl logs -f <pod-name> <container-name>`.

### Monitoring
1. What to Monitor:
    1. Count of nodes in cluster
    2. Healthy nodes count
    3. Performance metrics: CPU usage, memory, n/w and disk utilization
    4. Pod level metrics: number of them and performance metrics of each pod
2. Tools to integrate with k8s:
    1. Metrics server
    2. Prometheus
    3. Elastic stack
    4. Datadog
    5. Dynatrace
3. Heapster - Original project to enable monitoring and analytics on k8s objects - deprecated
4. Metrics server
    1. A trimmed down version of it
    2. 1 Metrics server per cluster
    3. Gets metrics from each node, pods, aggregates them and stores them
    4. In memory monitory solution - no historical data
5. Kubelet runs on each node
    1. it has a sub-component called cAdvisor
    2. cAdvisor is responsible for retrieving performance metrics and put them to kubelet API
6. minikube enable addons metrics-server
7. git clone https://github.com/kubernetes-incubator/metrics-server.git - download the deployment binaries
    1. kubectl create -f deploy/1.8+/ - creates set of pods, services and roles to enable metric server to poll for performance metrics of cluster
8. kubectl top node - to view the metrics of nodes
9. kubectl top pod  - to view the metrics of pods

### Labels, Selectors and Annotations
1. Ability to group kubernetes objects together and filter them based on needs is achieved using labels and selectors.
2. Labels are basically properties attached to each item.
3. Selectors help us filter kubernetes objects based on the attached properties (labels).
4. An example of labels and selectors would be:
   1. When we create pods, we attach some labels.
   2. And then when we create service to redirect requests to the pods, we create selectors and matchLabels to link service and pods
5. An example of a pod with labels is as below (here app: mock-app and function: backend are the labels):
```
apiVersion: v1
kind: Pod
metadata:
    name: simple-webapp
    labels:
        app: mock-app
        function: backend
spec:
    containers:
    - name: simple-webapp
      image: simple-webapp
      ports:
        - containerPort: 8080
```
6. After creating a pod with certain labels, we can filter it by: `kubectl get pods --selector app=mock-app`
7. An example of a service using selector to attach itself to pods (here app: mock-app and function: backend under spec/selector are the selectors)
```
apiVersion: v1
kind: Service
metadata:
    name: my-service
spec:
    selector:
        app: mock-app
        function: backend
    ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```
8. Having one selector is enough unless further nested filtering is required.
9. Annotations are used to record other details for informatory purposes. For details like build information, name or contact.
