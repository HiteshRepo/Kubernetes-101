## Core Concepts Revision Tasks

### Pods

1. Get number of pods running in default namespace: `kubectl get pods`
2. Create pod running nginx container: `refer kodekloud/pods/pod-nginx.yaml`
3. Get number of pods running in default namespace now (some more pods were spun up by kodekloud): `kubectl get pods`
4. Check image used to create the new pods: `kubectl describe pod <pod-name>`
5. Check node on which the new pods are running: `kubectl get pods -o wide` - check NODE column
6. Number of containers part of new pod 'webapp' (a new pod was spun up by kodekloud): `kubectl get pods` - check READY column (containers in ready state/total containers in the pod)
7. Image names of containers in pod webapp: `kubectl describe pod webapp` - check under containers[i]/Image
8. Check state of container 'agentx' in pod webapp: `kubectl describe pod webapp` - check under containers/agentx/State
9. Reason for container 'agentx' in pod webapp being in error state: `kubectl describe pod webapp` - check under events section
10. Indication of READY column: `kubectl get pods` - it denotes containers in ready state/total containers in the pod
11. Delete webapp pod: `kubectl delete pod webapp`
12. Create pod running redis container: `refer kodekloud/pods/pod-redis.yaml`

### Replicaset
1. Get number of pods in default namespace: `kubectl get pods`
2. Get number of replica set in default namespace: `kubectl get replicaset`
3. Desired number of pods in a replica set: `kubectl get replicaset` - check under DESIRED column
4. Check image used to create pods under a replica set: `kubectl describe replicaset <replica-set-name>` - PodTemplate/Containers/<conatiner-name>/Image
5. Number of pods in READY state in a replica set: `kubectl get replicaset` - check under READY column
6. Check the reason why some pods under a replica set is not ready/running: `kubectl get pods` - check pods under the replica set - describe the pod that is not ready/running and check the events section
7. Delete a pod: `kubectl delete pod <pod-name>`
8. Reason to why number of pods remain same if deleted: Replica set ensures the desired number of pods are always running/created
9. Create a replica set: `kubectl create -f <replica set file name>`
10. Reason for error `'Unable recognize "<file anme>": no matches for kind "ReplicaSet" in version "v1"'`: ReplicaSet are present under apiVersion 'apps/v1'
11. Reason for error `The ReplicaSet "replicaset-2" is invalid: spec.template.metadata.labels: Invalid value: map[string]string{"tier":"nginx"}: 'selector' does not match template 'labels'`: pod/metadata/label should match with replicaset/spec/selector/matchLabels
12. Scale a replica set to more/less pods: `kubectl scale --replicas=6 replicaset <replica-set-name>`
13. Scale a replica set to more/less pods: `kubectl edit replicaset <replica-set-name>` -> vi editor -> change replicas count

### Deployment
1. Get number of pods in default namespace: `kubectl get pods`
2. Get number of replica set in default namespace: `kubectl get replicaset`
3. Get number of deployment in default namespace: `kubectl get deployment`
4. Check image used to create pods under a deployment: `kubectl describe deployment <deployment-name>` - PodTemplate/Containers/<conatiner-name>/Image
5. Check the reason why some pods under a deployment is not ready/running: `kubectl get pods` - check pods under the deployment - describe the pod that is not ready/running and check the events section

### Namespace
1. Get number of namespace: `kubectl get namespace`
2. Get number of pods in 'research' namespace: `kubectl get pods --namespace=research`
3. Create a pod in 'finance' namespace: `refer kodekloud/pods/pod-redis-ns.yaml`

### Imperative commands
1. Deploy a pod named nginx-pod using the nginx:alpine image: `kubectl run nginx --image=nginx:alpine  --dry-run=client -o yaml` -> copied the declarative YAML result and changed the metadata/name to nginx-pod
2. Deploy a redis pod using the redis:alpine image with the labels set to tier=db: `kubectl run redis --image=redis:alpine  --dry-run=client -o yaml`
3. Create a service redis-service to expose the redis application within the cluster on port 6379: Refer kodekloud/imperative/commands.yaml
4. Create a deployment named webapp using the image kodekloud/webapp-color with 3 replicas: Refer kodekloud/imperative/commands.yaml
5. Create a new pod called custom-nginx using the nginx image and expose it on container port 8080: Refer kodekloud/imperative/commands.yaml
6. Create a new namespace called dev-ns: `kubectl create namespace dev-ns`
7. Refer kodekloud/imperative/commands.yaml last deployment
8. Refer kodekloud/imperative/commands.yaml last

### Commands and arguments
1. What is the command used to run the pod ubuntu-sleeper?: `kubectl describe pod ubuntu-sleeper`: see under Containers/ubuntu/Command
2. Create a pod with the ubuntu image to run a container to sleep for 5000 seconds. Modify the file ubuntu-sleeper-2.yaml: Refer kodekloud/commandsAndArguments/ubuntu-sleeper-2.yaml
3. Create a pod using the file named ubuntu-sleeper-3.yaml. There is something wrong with it. Try to fix it!: Refer kodekloud/commandsAndArguments/ubuntu-sleeper-3.yaml (section 1 & 2)
4. Update pod ubuntu-sleeper-3 to sleep for 2000 seconds: Refer kodekloud/commandsAndArguments/ubuntu-sleeper-3.yaml (section 3)
5. Inspect the file Dockerfile given at /root/webapp-color. What command is run at container startup?: python app.py, Refer kodekloud/commandsAndArguments/webapp/Dockerfile
6. Inspect the file Dockerfile2 given at /root/webapp-color. What command is run at container startup?: python app.py --color red, Refer kodekloud/commandsAndArguments/webapp/Dockerfile2
7. Inspect the two files under directory webapp-color-2. What command is run at container startup?: --color green Refer kodekloud/commandsAndArguments/webapp-2
8. Inspect the two files under directory webapp-color-3.  What command is run at container startup?: python app.py --color pink Refer kodekloud/commandsAndArguments/webapp-color-3
9. Create a pod with the given specifications. By default it displays a blue background. Set the given command line arguments to change it to green: Refer Refer kodekloud/commandsAndArguments/webapp-green.yaml
10. Create a pod with the given specifications. By default it displays a blue background. Set the given command line arguments to change it to green: Refer Refer kodekloud/commandsAndArguments/webapp-green.yaml

### Config Maps
1. What is the environment variable name set on the container in the pod?: `kubectl describe pod webapp-color` -> Containers/webapp-color/Environment
2. What is the environment variable name set on the container in the pod?: `kubectl describe pod webapp-color` -> Containers/webapp-color/Environment/APP_COLOR
3. Change environment variable APP_COLOR to green: Refer kodekloud/configmaps/webapp-color-env.yaml
4. How many ConfigMaps exists in the default namespace? : `kubectl get configmaps`
5. Identify the database host from the config map db-config: `kubectl describe configmap db-config`
6. Create a new ConfigMap for the webapp-color POD. Use the spec given below. ConfigName Name: webapp-config-map, Data: APP_COLOR=darkblue: Refer kodekloud/configmaps/webapp-config-map.yaml
7. Update the environment variable on the POD to use the newly created ConfigMap: Refer kodekloud/configmaps/webapp-color-configmap.yaml

### Secrets
1. How many Secrets exist on the system?: `kubectl get secrets`
2. How many secrets are defined in the default-token secret?: `kubectl describe secrets default-token-xcwxh` : check Data section
3. What is the type of the default-token secret?: `kubectl describe secrets default-token-xcwxh` : check Type section
4. Which of the following is not a secret data defined in default-token secret?: `kubectl describe secrets default-token-xcwxh` : check Data section
5. Create a new secret named db-secret with the data given below. (Refer kodekloud/secrets/db-secret.yaml)
   Secret Name: db-secret
   Secret 1: DB_Host=sql01
   Secret 2: DB_User=root
   Secret 3: DB_Password=password123
6. Configure webapp-pod to load environment variables from the newly created secret: Refer kodekloud/secrets/webapp-pod.yaml

### Docker Security
1. What is the user used to run ubuntu-sleeper? : `kubectl describe pod ubuntu-sleeper`
2. Edit the pod ubuntu-sleeper to run the sleep process with user ID 1010: Refer kodekloud/securitycontexts/ubuntu-sleeper.yaml
3. A Pod definition file named multi-pod.yaml is given. With what user are the processes in the web container started?: Refer kodekloud/securitycontexts/multi-pod.yaml: 1002
4. With what user are the processes in the sidecar container started?: Refer kodekloud/securitycontexts/multi-pod.yaml: 1001
5. Update pod ubuntu-sleeper to run as Root user and with the SYS_TIME capability: Refer kodekloud/securitycontexts/ubuntu-sleeper-root.yaml
6. Now update the pod to also make use of the NET_ADMIN capability: Refer kodekloud/securitycontexts/ubuntu-sleeper-root.yaml

### Service accounts
1. How many Service Accounts exist in the default namespace: `kubectl get serviceaccount`
2. What is the secret token used by the default service account?: `kubectl describe serviceaccount default`: See Tokens
3. We just deployed the Dashboard application. Inspect the deployment. What is the image used by the deployment?: `kubectl describe pod <pod-name>`
4. At what location is the ServiceAccount credentials available within the pod?: /var/run/secrets
5. Create a new ServiceAccount named 'dashboard-sa': `kubectl create serviceaccount dashboard-sa`
6. Enter the access token in the UI of the dashboard application. Click Load Dashboard button to load Dashboard:
   1. `kubectl describe serviceaccount dashboard-sa`
   2. `kubectl describe secret dashboard-sa-token-qh5jz`
7. Update the deployment to use the newly created ServiceAccount: Refer kodekloud/serviceaccounts/web-dashboard.yaml

### Resource Requirements
1. A pod called rabbit is deployed. Identify the CPU requirements set on the Pod: `kuebctl describe pod <pod-name>`: check Containers/Requests/cpu
2. Another pod called elephant has been deployed in the default namespace. It fails to get to a running state. Inspect this pod and identify the Reason why it is not running:  `kuebctl describe pod <pod-name>`: check Containers/Last State/Reason: The status OOMKilled indicates that it is failing because the pod ran out of memory. Identify the memory limit set on the POD
3. The elephant pod runs a process that consume 15Mi of memory. Increase the limit of the elephant pod to 20Mi.: Refer kodekloud/resources/elephant.yaml

### Taints and Tolerations
1. How many nodes exist on the system: `kubectl get nodes`
2. Do any taints exist on node01 node?: `kubectl describe node node01 | grep Taints`
3. Create a taint on node01 with key of spray, value of mortein and effect of NoSchedule: `kubectl taint node node01 spray=mortein:NoSchedule`
4. Create a new pod with the nginx image and pod name as mosquito: Refer kodekloud/taintsandtolerations/mosquito.yaml
5. What is the state of the POD (mosquito)?: `kubectl describe pod mosquito`: Pending
6. Why do you think the pod is in a pending state?: POD mosquito cannot tolerate taint mortein
7. Create another pod named bee with the nginx image, which has a toleration set to the taint mortein: Refer kodekloud/taintsandtolerations/bee.yaml
8. Do you see any taints on controlplane node?: `kubectl describe node controlplane | grep Taints`: NoSchedule
9. Remove the taint on controlplane, which currently has the taint effect of NoSchedule: `kubectl taint nodes controlplane node-role.kubernetes.io/master:NoSchedule-`
10. What is the state of pod mosquito now?: Running
11. Which node is pod mosquito placed now?: controlplane

### Node Affinity

1. How many Labels exist on node node01?: `kubectl describe nodes node01`: Check labels section.
2. Apply a label color=blue to node node01: `kubectl label nodes node01 color=blue`.
3. Create a new deployment named blue with the nginx image and 3 replicas: Refer kodekloud/nodeaffinity/blue-deployment.yaml.
4. Which nodes can the pods for the blue deployment be placed on?: `kubectl describe nodes <node-name> | grep Taints`
5. Set Node Affinity to the deployment to place the pods on node01 only.: Refer kodekloud/nodeaffinity/blue-deployment-with-affinity.yaml
6. Which nodes are the pods placed on now?: `kubectl get pods -o wide`
7. Create a new deployment named red with the nginx image and 2 replicas, and ensure it gets placed on the controlplane node only.: Refer kodekloud/nodeaffinity/red-deployment-with-affinity.yaml.

### Multi Container Pods
1. Identify the number of containers created in the red pod: `kubectl describe pod red`
2. Identify the name of the containers running in the blue pod: `kubectl describe pod blue`
3. Create a multi-container pod with 2 containers.
4. Use the spec given below, If the pod goes into the crashloopbackoff then add sleep 1000 in the lemon container: Refer kodekloud/multicontainer/yellow.yaml
5. The application outputs logs to the file /log/app.log. View the logs and try to identify the user having issues with Login.:  `kubectl exec app -- cat /log/app.log   -n elastic-stack`
6. Edit the pod to add a sidecar container to send logs to Elastic Search. Mount the log volume to the sidecar container
`kubectl -n elastic-stack get pod -o yaml > app.yaml`
Refer kodekloud/multiontainer/elastic-stack/app.yaml

### Readiness and Liveness probe
1. Update the newly created pod 'simple-webapp-2' with a readinessProbe using the given spec: Refer kodekloud/readinessNliveliness/simple-webapp-2.yaml
2. What would happen if the application inside container on one of the PODs crashes?: The crashed container inside the pod is restarted
3. What would happen if the application inside container on one of the PODs freezes?: New Users are impacted
4. Update both the pods with a livenessProbe using the given spec: Refer kodekloud/readinessNliveliness/simple-webapp-2.yaml and kodekloud/readinessNliveliness/simple-webapp-1.yaml

### Container logging
1. A user - USER5 - has expressed concerns accessing the application. Identify the cause of the issue.
   Inspect the logs of the POD: `kubectl logs -f webapp-1` - USER5 Failed to Login as the account is locked due to MANY FAILED ATTEMPTS.
2. A user is reporting issues while trying to purchase an item. Identify the user and the cause of the issue.
   Inspect the logs of the webapp in the POD: `kubectl logs webapp-2 simple-webapp`: USER30 Order failed as the item is OUT OF STOCK

### Monitoring
1. Let us deploy metrics-server to monitor the PODs and Nodes. Pull the git repository for the deployment files.:
   1. git clone https://github.com/kodekloudhub/kubernetes-metrics-server.git
   2. cd kubernetes-metrics-server/
   3. kubectl create -f .
2. It takes a few minutes for the metrics server to start gathering data.: `kubectl top node`
3. Identify node consuming most CPUs: control plane
4. Identify node consuming most Memory: control plane
5. Identify the POD that consumes the most Memory.:`kubectl top pod` - rabbit
6. Identify the POD that consumes the least CPU: lion

### Labels, Selectors and Annotations
1. We have deployed a number of PODs. They are labelled with tier, env and bu. How many PODs exist in the dev environment?
   Use selectors to filter the output: `kubectl get pods --selector env=dev` - 7
2. How many PODs are in the finance business unit (bu)?: `kubectl get pods --selector bu=finance` - 6
3. How many objects are in the prod environment including PODs, ReplicaSets and any other objects?: `kubectl get all --selector env=prod` - 7
4. Identify the POD which is part of the prod environment, the finance BU and of frontend tier?: `kubectl get pods --selector env=prod --selector bu=finance --selector tier=frontend` - 5 (app-1-zzxdf)
5. A ReplicaSet definition file is given replicaset-definition-1.yaml. Try to create the replicaset. There is an issue with the file. Try to fix it.: Refer kodekloud/poddesign/labelsAndSelectors.yaml

### Rolling updates & rollbacks in Deployments
1. Inspect the deployment and identify the number of PODs deployed by it: `kubectl describe deployment frontend` - 4
2. What container image is used to deploy the applications?: `kubectl describe deployment frontend` - kodekloud/webapp-color:v1
3. Inspect the deployment and identify the current strategy: `kubectl describe deployment frontend` - RollingUpdate
4. Upgrade the application by setting the image on the deployment to kodekloud/webapp-color:v2: `kubectl set image deployment frontend simple-webapp=kodekloud/webapp-color:v2`
5. Up to how many PODs can be down for upgrade at a time Consider the current strategy settings and number of PODs 4: `kubectl describe deployment frontend` - RollingUpdateStrategy:  25% max unavailable, 25% max surge
6. Change the deployment strategy to Recreate. Delete and re-create the deployment if necessary. Only update the strategy type for the existing deployment.: Refer kodekloud/poddesign/simple-webapp-deployment.yaml
7. Upgrade the application by setting the image on the deployment to kodekloud/webapp-color:v3 - `kubectl set image deployment frontend simple-webapp=kodekloud/webapp-color:v3`

### Job, Cron Job
1. Create a Job using this POD definition file or from the imperative command and look at how many attempts does it take to get a '6'.: Refer kodekloud/throw-dice-job.yaml
2. Update the job definition to run as many times as required to get 3 successful 6's: Refer kodekloud/throw-dice-job-2.yaml
3. Update the job definition to run 3 jobs in parallel.: Refer kodekloud/throw-dice-job-3.yaml
4. Create a CronJob for the same to be scheduled at: 21.30: Refer kodekloud/throw-dice-cronjob.yaml

### Services
1. How many Services exist on the system? in the current(default) namespace: `kubectl get services`
2. What is the type of the default kubernetes service?: ClusterIP
3. What is the targetPort configured on the kubernetes service?: `kubectl describe service kubernetes`: 6443
4. How many labels are configured on the kubernetes service?: component=apiserver and provider=kubernetes - 2
5. How many Endpoints are attached on the kubernetes service?: 10.10.33.3:6443 - 1
6. How many Deployments exist on the system now? in the current(default) namespace: `kubectl get deployments` - 1
7. What is the image used to create the pods in the deployment?: `kubectl describe deployment simple-webapp-deployment` - kodekloud/simple-webapp:red
8. Create a new service to access the web application using the service-definition-1.yaml file: Refer kodekloud/services/simple-webapp-service.yaml

### Ingress Resources
1. Which namespace is the Ingress Controller deployed in? - `kubectl get all -n ingress-nginx` : ingress-nginx
2. What is the name of the Ingress Controller Deployment? - deployment.apps/ingress-nginx-controller
3. Which namespace are the applications deployed in?: `kubectl get all -n app-space`
4. How many applications are deployed in the app-space namespace?: No. of deployments: 3
5. Which namespace is the Ingress Resource deployed in?: app-space
6. What is the name of the Ingress Resource?: `kubectl get ingress -n app-space`: ingress-wear-watch
7. What is the Host configured on the Ingress Resource?: `kubectl describe ingress ingress-wear-watch -n app-space`: Rules/Host: *
8. What backend is the /wear path on the Ingress configured with?: wear-service:8080
9. At what path is the video streaming application made available on the Ingress: /watch
10. If the requirement does not match any of the configured paths what service are the requests forwarded to?: default-http-backend:80
11. You are requested to change the URLs at which the applications are made available. Make the video application available at /stream.: Refer kodekloud/ingressnetworking/ingress-wear-watch.yaml
12. You are requested to add a new path to your ingress to make the food delivery application available to your customers. Make the new application available at /eat.: Refer kodekloud/ingressnetworking/ingress-wear-watch.yaml
13. Identify the namespace in which the new application (webapp-pay) is deployed.: `kubectl get all -n critical-space`
14. What is the name of the deployment of the new application?: webapp-pay
15. You are requested to make the new application available at /pay: Refer kodekloud/ingressnetworking/ingress-pay.yaml.

### Ingress Controller
1. create a namespace called ingress-space: `kubectl create namespace ingress-space`
2. Create a ConfigMap object in the ingress-space: Refer kodekloud/ingressnetworking/ingress-configmap.yaml
3. Create a ServiceAccount in the ingress-space namespace: Refer kodekloud/ingressnetworking/ingress-serviceaccount.yaml
4. We have created the Roles and RoleBindings for the ServiceAccount. Check it out!!: `kubectl get rolebindings,clusterrolebindings -n ingress-space`
5. Create a deployment using the file given.: Refer kodekloud/ingressnetworking/ingress-deployment.yaml
6. Create a service following the given specs: Refer kodekloud/ingressnetworking/ingress-service.yaml
7. Create the ingress resource to make the applications available at /wear and /watch on the Ingress service.: Refer kodekloud/ingressnetworking/ingress-wear-watch.yaml(Previous)

### Network Policies
1. Traffic Flow & Rules: Ingress & Egress
   1. Ingress Rule : For incoming request into the component.
   2. Egress Rule : For outgoing request from the component.
2. Network Security:
   1. 'All allow' rule: That allows traffic from any pod to any other pod/service within cluster. Default.
3. Solutions that support Network Policies:
   1. Kube-router
   2. Calico
   3. Romana
   4. Weave-net
4. Solutions that do not support Network Policies:
   1. Flannel
5. Show all network policy: `kubectl get networkpolicy`
6. What type of traffic is this Network Policy configured to handle?: `kubectl describe networkpolicy <name>`
7. Create a network policy to allow traffic from the Internal application only to the payroll-service and db-service.
   Use the spec given below. You might want to enable ingress traffic to the pod to test your rules in the UI.: `Refer network_policy_payroll_db.yaml`

## Volumes
1. Like dockers, pods too are transient/ephemeral.
2. So if we store data within the pods, the data will not be available later if pod is killed.
3. In order to have persistence, we can map pod's storage to hosts' storage in the form of volume.
4. Example:
   ```
   apiVersion: v1
   kind: Pod
   metadata:
      name: random-number-generator
   spec:
      containers:
      - image: alpine
        name: alpine
        command: ["bin/sh", "-c"]
        args: ["shuf -i 0-100 -n 1 >> /opt/number.out;"]
        volumeMounts: 
        - mountPath: /opt
          name: data-volume
      volumes:
      - name: data-volume
        hostPath: 
            path: /data
            type: Directory     
   ```
5. The above setup is ok for a single node.
6. But if there are multiple nodes, it is not recommended having the above setup.
7. Instead, usage of external cloud storage would be better choice. For example if you use EBS:
   ```
   volumes:
      - name: data-volume
        awsElasticBlockStore: 
            volumeID: <volume-id>
            fsType: ext4 
   ```

## Persistent Volumes
1. When there are lots of Pods, a user has to configure volumes on all Pod definition file on all envs.
2. Everytime changes are made, the user has to make the changes in all pods.
3. Instead, it would be great to have a centralized management of this.
4. Such that an admin can configure a large pool of storage and users/pods can carve out pieces from it as required.
5. PVs are cluster-wide pool of storage volumes configured by admin.
6. users/pods can use these volumes using PVCs.
7. Example of PV:
   ```
   apiVersion: v1
   kind: PersistentVolume
   metadata: 
      name: pv-vol1
   spec:
      accessModes:
         - ReadWriteOnce
      capacity:
         storage: 1Gi
      awsElasticBlockStore: 
            volumeID: <volume-id>
            fsType: ext4
   ```

## Persistent Volume Claims
1. Admin creates a set of PVs and users create a set of PVCs.
2. K8s binds the PVC requests and volumes based on the properties of PVCs and PVs.
3. PVC and PV have 1-1 mapping.
4. During the binding process, K8s tries to find PV that has same capacity, access modes, volume modes and storage class as the PVC.
5. If there are multiple matches for a single claim, labels and selectors can be used to pinpoint.
6. A smaller claim may get bound to larger volumes if there are no better matches and all properties match.
7. If no PV matches the PVC, the PVC remains in pending state until newer volumes are created.
8. Example of PVC:
   ```
   apiVersion: v1
   kind: PersistentVolumeClaim
   metadata: 
      name: myClaim
   spec:
      accessModes:
         - ReadWriteOnce
      resources:
         requests:
            storage: 500 Mi
   ```
9. What happens if the PVC is deleted? we can choose from the following to configure 'persistentVolumeReclaimPolicy':
   1. Retain - Default. Remain until manually deleted by Admin. Unavailable for reuse.
   2. Delete.
   3. Recycle - Data is scrubbed but available for reuse.
10. Command to view a file in a pod: `kubectl exec <pod-name> -- cat <file-path>`

## Storage Class
1. What happens with PVC and PV setup is that Admin creates big volume of storage so that those can be chunked into PV which thereby can be claimed via PVCs.
2. The above process is called Static provisioning.
3. Admin has to manually create the storage and keep tabs on it to expand when required.
4. Storage class helps to automate this process by provisioning on-demand.
5. Storage class helps with Dynamic provisioning.
6. Example of SC:
   ```
   apiVersion: storage.k8s.io/v1
   kind: StorageClass
   metadata:
      name: google-storage
   provisioner: kubernetes.io/gce-pd
   ```
7. PVC can now use the storage class instead of PV.
8. Example of PVC using SC:
   ```
   apiVersion: v1
   kind: PersistentVolumeClaim
   metadata: 
      name: myClaim
   spec:
      accessModes:
         - ReadWriteOnce
      storageClassName: google-storage 
      resources:
         requests:
            storage: 500 Mi
   ```
9. What is the name of the Storage Class that does not support dynamic volume provisioning? : local-storage

## Stateful Sets
1. Why do we need StatefulSet? one of the use case (Master-Slave DB replica setup):
   1. Setup master first and then slaves
   2. Clone data from master to slave1
   3. Enable continuous replication from master to slave1
   4. Wait for slave1 to be ready
   5. Clone data from slave1 to slave2
   6. Enable continuous replication from master to slave2
   7. Configure master address on slaves
2. How do we achieve this in k8s?
   1. We need the order of creation of these components (master and slaves)
   2. First master, then slave1 and finally slave2
   3. We cannot ensure ordering using Deployments
   4. Also in order to enable replication and cloning of data we need to be able to differentiate master and the salves
   5. Also in order to configure master address in slaves, we need a static address. But since pods in k8s are ephemeral, it is harder to achieve as IPs and Pod names would change.
3. Solution (StatefulSets)
   1. Very similar to Deployments
   2. Pods are created in sequential manner
   3. Assign unique ordinal number to each pod [name = <image-name>-<ordinal-number>]
   4. If pod goes down and comes up, the name remains the same
4. StatefulSet specifications are almost same as Deployment, except for that StatefulSet requires a service name (headless) specified
5. Example of a StatefulSet
   ```
   apiVersion: v1
   kind: StatefulSet
   metadata: 
      name: mysql
      labels:
         app: mysql
   spec:
      template: 
         metadata:
            labels:
               app: mysql
         spec:
            containers:
            - name: mysql
              image: mysql
      replicas: 3
      selector:
         matchLabels:
            app: mysql
      serviceName: mysql-h 
   ```
6. StatefulSet scales up and down in order.
7. But the above behavior can be overridden by: podManagementPolicy: Parallel

## Headless service
1. The way to point from one pod to another is via service.
2. When we create a headless service, it creates DNS entries for each pod in the format: <pod-name>.<headless-service-name>.<namespace>.<svc>.<cluster-domain>
3. Headless service differentiates from a normal service by configuring ClusterIP as None
4. Example of headless service:
   ```
   apiVersion: v1
   kind: Service
   metadata:
      name: mysql-h
   spec:
      ports:
         - port: 3306
      selector:
         app: mysql
      clusterIP: None
   ```
5. Pod Definition to create DNS service via headless service
   ```
    apiVersion: v1
    kind: Pod
    metadata:
       name: myapp-pod
       labels:
         app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql
      subdomain: mysql-h
      hostname: mysql-pod 
   ```
6. Deployment Definition to create DNS service via headless service
   ```
   apiVersion: v1
   kind: Deployment
   metadata:
       name: mysql-deployment
       labels:
         app: mysql
   spec:
      replicas: 3
      matchLabels:
         app: mysql
      template:
         metadata:
            name: mysql
            labels:
               app: mysql
         spec:
           containers:
           - name: mysql
             image: mysql
           subdomain: mysql-h
           hostname: mysql-pod 
   ```
7. StatefulSet Definition to create DNS service via headless service
   ```
   apiVersion: v1
   kind: StatefulSet
   metadata:
       name: mysql-statefulset
       labels:
         app: mysql
   spec:
      serviceName: mysql-h
      replicas: 3
      selector:
          matchLabels:
             app: mysql
      template:
         metadata:
            name: mysql
            labels:
               app: mysql
         spec:
           containers:
           - name: mysql
             image: mysql 
   ```
8. If PVC is used in a Pod, K8s creates a PV via SC and maps it to the Pod.
9. If PVC is used in a Deployment or StatefulSet, k8s creates the configured replicas and maps each pod to same PV via SC.
10. So in order to have different PV for each pod in a Deployment or StatefulSet, used 'volumeClaimTemplate' directly under spec of a Deployment or StatefulSet. The configuration remains same as a PVC
11. Example of above:
    ```
   apiVersion: v1
   kind: StatefulSet
   metadata:
      name: mysql-statefulset
      labels:
         app: mysql
   spec:
      replicas: 3
      selector:
         matchLabels:
            app: mysql
      template:
         metadata:
            labels:
               app: mysql
         spec:
            containers:
            - name: mysql
              image: mysql
              volumeMounts:
               - mountPath: /var/lib/mysql
                  name: data-volume
      volumeClaimTemplate:
      - metadata:
         name: data-volume
        spec:
         accessModes:
            - ReadWriteOnce
         storageClassName: google-storage
         resources: 
            requests: 
               storage: 500Mi
```