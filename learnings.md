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
* minikube ip
* minikube ssh - to connect to minikube
* eval $(minikube docker-env) - to make docker point to minikube docker context

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