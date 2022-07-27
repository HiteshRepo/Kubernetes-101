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

## Define, build and modify container images
1. Steps:
   1. OS
   2. Update source repository
   3. Install dependencies
   4. Install code dependencies using dependency management files
   5. Copy source code to docker
   6. Build/Compile the code to generate binaries
   7. Copy binaries from a builder container to a new image
   8. Copy the config files
   9. Commands to start the app
2. Build: docker build <Dockerfile path> -t <account-name/image-name> .
3. Push: docker push <account-name/image-name>
4. Format - <Instruction> <Argument>

## Security Primitives
1. Host
   1. Root access disabled
   2. Password based authentication disabled
   3. SSK key based authentication
2. kube-apiserver
   1. Who can access the cluster? - Authentication
      1. Files - Usernames and Passwords
      2. Files - Usernames and Tokens
      3. Certificates
      4. External Authentication providers - LDAP
      5. Service Accounts
   2. What can they do? - Authorization
      1. RBAC 
      2. ABAC
      3. Node
      4. Webhook
   3. Communication - TLS encryption
   4. Pod communication - Network policies

## Authentication
1. Types of users accessing K8s cluster
   1. Admin: To perform administrative tasks 
   2. Developer: To deploy applications
   3. End User: To access the applications deployed on cluster
   4. Bots: For integration purposes
2. Challenge: How to secure cluster by securing communication b/w internal components and management access to cluster through Authentication and Authorization mechanisms.
3. End User Authentication is managed by applications themselves.
4. Admin, Developers : User, Bots: ServiceAccounts
5. Users in K8s cannot be created/managed
   1. This cannot be done: kubectl create user user1
   2. Nor this: kubectl get users
6. However, this can be done:
   1. kubectl create serviceaccount sa1
   2. kubectl get serviceaccount
7. User access is managed by kube-apiserver
8. kube-apiserver authenticates the kubectl requests before processing it.
9. Ways:
   1. Static Password File: List usernames and passwords
   2. Static Token File: List usernames and tokens
   3. Certificates
   4. Identity Services: LDAP
10. Static Password File: 
    1. .csv file with 3 columns: password, username, userId and group name. The file name is passed as an argument to kube-apiserver: --basic-auth-file=user-details.csv
    2. Authenticate user by curl command: curl ..... -u "<username>:<password>"
11. Static Token File
    1. .csv file with 3 columns: token, username, userId and group name. The file name is passed as an argument to kube-apiserver: --token-auth-file=user-details.csv
    2. Authenticate user by curl command: curl ..... -header "Authorization: Bearer <token>"
12. Above mechanisms are not recommended for usage in actual envs.

## Kubeconfig
1. Use certificate in curl commands: 
   ```
      curl https://<cluster-url>:6443/api/v1/pods --key admin.key --cert admin.crt --cacert ca.crt
   ```
2. User certificate in kubectl commands:
   ```
      kubectl get pods --server <cluster-url>:6443 --client.key admin.key --client-certificate admin.crt --certificate-authority ca.crt
   ```
3. Prepare kubeconfig file and use in kubectl:
   ```
      kubectl get pods --kubeconfig <config-file-path>
   ```
4. By default, it looks for file at: $HOME/.kube/config. To change it set $KUBECONFIG env var.
5. Kube config file format
   1. Clusters: Dev, Prod, UAT, Stage, etc.
   2. Users: Admin, Dev, Prod, etc.
   3. Contexts: Admin@Dev, Dev@Prod, etc.
6. Example
   ```
      apiVersion: v1
      kind: Config
      current-context: my-kube-admin@my-kube-palyground
      clusters:
      - name: my-kube-palyground
        cluster:
         cluster-authority: ca.crt [filepath] or certificate-authority-data: <base-64 encoded format of ca.crt contents>
         server: https://my-kube-palyground:6443
      contexts:
      - name: my-kube-admin@my-kube-palyground
        context:
          cluster: my-kube-palyground
          user: my-kube-admin
          namespace: <namespace-name>
      users:
      - name: my-kube-admin
        user:
         client-certificate: admin.crt
         client-key: admin.key
   ```
7. command to show available kubeconfig: `kubectl config view`
8. command to show available kubeconfig of a specified file: `kubectl config view --kubeconfig=my-custom-config`
9. command to change current context: `kubectl config use-context user@cluster`

## API Groups
1. We interact with api-server via kubectl or REST
2. REST:
   1. curl https://<master-cluster-url>:6443/version
   2. curl https://<master-cluster-url>:6443/api/v1/pods
3. Groups:
   1. /metrics - integrate 3rd party apps with k8 metrics
   2. /healthz - monitor health of cluster
   3. /version - version of the cluster
   4. /api
   5. /apis
   6. /logs - integrate 3rd party apps with k8 logs
4. Core API groups (/api)
   1. /v1 - namespaces, pods, rc, services, secrets, configmaps, PV, PVC, endpoints
5. Named API groups (/apis) - more organized, newer feature
   1. /apps
      1. /v1 - /deployments, /replicasets, /statefulsets - /list, /get, /create, /watch, /delete, /update
   2. /extensions
   3. /networkin.k8s.io
      1. /v1 - /networkpolicies
   4. /storage.k8s.io
   5. /authentication.k8s.io
   6. /certificates.k8s.io
      1. /v1 - /certificatesigningrequests
6. To know which actions apis you have access to:
   1. curl http://<cluster-url>:6443 -k
   2. curl http://<cluster-url>:6443/apis -k
7. If unaccessible
   1. Specify certificate in curl
   2. Start kube proxy

## Authorization
1. --authorization-mode flag is used.
   1. Multiple modes can be set, --authorization-mode=Node,RBAC,Webhook
   2. The modes are used in sequence. If a request is denied by Node here, it is forwarded to RBAC
2. Admins:
   1. Viewing K8s objects.
   2. Creating/Deleting K8s objects.
3. Developers, ServiceAccounts also need authorized.
4. But not all of them should have same types of access.
5. Requirement: Varied level of access.
6. Types of Authorization Mechanisms:
   1. Node
   2. ABAC - Attribute
   3. RBAC - Role based
   4. Webhook
7. Node Authorizer:
   1. Users use 'KubeAPi' [or Master node's API serve] for Management Purposes.
   2. Similarly, Kubelets also use 'KubeAPi' for Management Purposes within cluster like
      1. Read info about: Service, Endpoints, Nodes, Pods, etc
      2. Write info about: Node status, Pod status, events, etc
   3. Kubelets should be part of 'System:Nodes' group. So any requests coming from kubelets will be authorized by Node Authorizor.
8. ABAC
   1. Associate a user/set of user with a set of permissions.
   2. Policy definition files needs to be created and passed to API Server.
   3. Manual editing of files becomes a norm and restarting of API Server is a necessity.
   4. These are difficult to manage.
9. RBAC
   1. Instead of associating a user/set of user with permissions a Role is defined.
   2. These roles are then associated with user/set of users.
   3. We simply modify the role to reflect on the associated users.
10. Webhook
    1. To manage authorization outside the cluster [API server].
    2. KubeAPI is going to rely on a thrid party component to verify the access/authorization
11. AlwaysAllow
    1. Allows all requests w/o auth check
    2. Default Auth mode
12. AlwaysDeny
    1. Denies all requests w/o auth check

## RBAC
1. Role definition file
   ```
   apiVersion: rbac.authorization.k8s.io/v1
   kind: Role
   metadata: 
      name: developer
   rules:
   - apiGroups: [""] // blank apiGroups indicates core group
     resources: ["pods"]
     verbs: ["list", "get", "create", "update", "delete"]
   - apiGroups: [""]
     resources: ["configMap"]
     verbs: ["create"]
   ```
2. kubectl create -f <filepath>
3. Link role and user
4. Role binding object
   ```
   apiVersion: rbac.authorization.k8s.io/v1
   kind: RoleBinding
   metadata: 
      name: devuser-developer-binding
   subjects:
   - kind: User
     name: dev-user
     apiGroup: rbac.authorization.k8s.io
   roleRef:
      kind: Role
      name: developer
      apiGroup: rbac.authorization.k8s.io
   ```
5. kubectl create -f <filepath>
6. View roles: kubectl get roles, kubectl describe role developer
7. View role bindings: kubectl get rolebindings, kubectl describe rolebinding devuser-developer-binding
8. Check Access: kubectl auth can-i create deployments / kubectl auth can-i delete nodes
9. Check Access by Admin: kubectl auth can-i create deployments --as dev-user / kubectl auth can-i delete nodes --as dev-user
10. Specific access:
```
   apiVersion: rbac.authorization.k8s.io/v1
   kind: Role
   metadata:
      name: developer
   rules:
   - apiGroups: [""] // blank apiGroups indicates core group
     resources: ["pods"]
     resourceNames: ["blue", "green"]
     verbs: ["list", "get", "create", "update", "delete"]
   - apiGroups: [""]
     resources: ["configMap"]
     verbs: ["create"]
```
11. How to check kube-apiserver configurations?
12. Get roles in all namespaces: `kubectl get roles --all-namespaces`

## Cluster roles
1. Role and Role bindings can be created only within a namespace.
2. But for resources like Nodes, PV, certificatesigningrequests, namespaces, etc cannot be associated with a namespace, they are cluster wide resources.
3. Cluster Scoped
4. View namespaced resources: `kubectl api-resources --namespaced=true`
5. View cluster scoped resources: `kubectl api-resources --namespaced=false`
6. ClusterRoles - similar to roles
   ```
   apiVersion: rbac.authorization.k8s.io/v1
   kind: ClusterRole
   metadata:
      name: cluster-administrator
   rules:
   - apiGroups: [""]
     resources: ["nodes"]
     verbs: ["list", "get", "create", "update", "delete"]
   ```
7. ClusterRoleBindings - similar to role bindings
   ```
   apiVersion: rbac.authorization.k8s.io/v1
   kind: ClusterRoleBinding
   metadata: 
      name: cluster-admin-binding
   subjects:
   - kind: User
     name: cluster-admin
     apiGroup: rbac.authorization.k8s.io
   roleRef:
      kind: ClusterRole
      name: cluster-administrator
      apiGroup: rbac.authorization.k8s.io
   ```
8. Cluster roles can also be used for namespaced resources as well but it will have access to all resources across namespaces

## Admission Controllers
1. kubectl -> authentication (mostly via certificates in kubeconfig) -> authorization (mostly RBACs) -> Actions (like Create Pods)
2. RBACs are on API level but not beyond it.
3. What if you would want to?
   1. Restrict images from certain registry
   2. Do not permit container running as rootuser - `containers/0/securityContext/runAsUser: 0`
   3. Allow certain capabilities only
   4. Metadata must conatins labels
4. Admission controllers helps us better security measures
5. kubectl -> authentication -> authorization -> Admission Controllers -> Actions (like Create Pods)
6. In built admission controllers
   1. AlwaysPullImages - always pulls image on Pod creation
   2. DefaultStorageClass - observes creation of PVCs and assigns SCs even if it is not specified
   3. EventRateLimit - requests api server can handle at a time
   4. NamespaceExists - rejects requests to namespaces that do not exist
   5. NamespaceAutoProvision - not enabled by default
7. `kube-apiserver -h | grep enable-admission-plugins` - list of admission controllers enabled
8. To add admission controllers - /etc/kubernetes/manifests/kube-apiserver.yaml - kubeadm
9. --disable-admission-plugins, --enable-admission-plugins flags
10. Which admission controller is enabled in this cluster which is normally disabled? : `kubectl -n kube-system describe po kube-apiserver-controlplane`
11. Note that the NamespaceExists and NamespaceAutoProvision admission controllers are deprecated and now replaced by NamespaceLifecycle admission controller.
    The NamespaceLifecycle admission controller will make sure that requests to a non-existent namespace is rejected and that the default namespaces such as default, kube-system and kube-public cannot be deleted.
12. Since the kube-apiserver is running as pod you can check the process to see enabled and disabled plugins: ps -ef | grep kube-apiserver | grep admission-plugins
13. Types of admission controllers:
    1. Validating admission controller: NamespaceExists, etc - They just accept/reject requests
    2. Mutating admission controller: DefaultStorageClass, etc - They change requests
14. Some admission controllers can do both validation and mutation
15. Sequence of admission controllers is also important - NamespaceAutoProvision, NamespaceExists
16. MutatingAdmissionWebhook, ValidatingAdmissionWebhook
17. Admission webhook server - contains our logic to be deployed and then configure k8s cluster with webhook config object
18. Steps
    1. Deploy webhook server - validate and mutate functions
    2. Containerize with k8s cluster - webhook service
    3. configure k8s cluster with our webhook server via ''
       ```
       apiVersion: admissionregistration.k8s.io/v1
       kind: ValidatingWebhookConfiguration
       metadata:
         name: "pod-policy.example.com"
       webhooks: 
       - name: "pod-policy.example.com"
         clientConfig: 
            url: "https://external-server.example.com" // if deployed externally
            service: // if deployed internally
               namespace: "webhook-namespace"
               name: "webhook-service"
            caBundle: "njjl;kd;k"  // for tls
         rules: // when to call
         - apiGroups: [""]
           apiVersions: ["v1"]
           operations: ["CREATE"]
           resources: ["pods"]
           scope: "Namespaced"
       ```
19. Create TLS secret webhook-server-tls for secure webhook communication in webhook-demo namespace.
    We have already created below cert and key for webhook server which should be used to create secret.
    Certificate : /root/keys/webhook-server-tls.crt
    Key : /root/keys/webhook-server-tls.key
    1. Refer -> admission-controllers/webhook-server-tls-secret.yaml
    2. cat <cert file path> | base64
       cat <key file path> | base64
    3. Also can be done w/o above to steps: kubectl create secret tls webhook-server-tls --key="<cert file path>" --cert="<key file path>"

## Api versions
1. /v1 - GA/stable, reliable, all users
2. /v1beta1 - has e2e, minor bugs, commitment to move to GA
3. /v1alpha1 - lacks e2e tests, may have bugs, no commitment, users - feedback users

## Api Deprecations
1. Support multiple versions at a time.
2. Sequence
   1. X -> /v1alpha1
   2. X+1 -> /v1alpha2 (#1 -> API elements may only be removed by incrementing the version of the API group) (Release notes - Mention breaking changes, notification for migration, etc)
   3. X+2 -> /v1beta1
   4. X+3 -> /v1beta2 and /v1beta1 [deprecated but preferred] (#2 -> API objects must be able to round trip between API versions in a given release w/o info loss, except whole REST resources that do not exist in some versions)
   5. X+4 -> /v1beta2 [preferred] and /v1beta1 [deprecated] (#3 -> Support - GA (12 months), Beta(9 months), Alpha(0 months))
   6. X+5 -> /v1 and /v1beta2 [deprecated but preferred] and /v1beta1 [deprecated]
   7. X+6 -> /v1 [preferred] and /v1beta2 [deprecated]
   8. X+7 -> /v1 [preferred] and /v1beta2 [deprecated]
   9. X+8 -> /v1 [preferred]
   10. X+8 -> /v1 [preferred] and /v2alpha1 (#4 -> API version in a given track may not be deprecated until a new stabel version is released)
3. `kubectl convert` command is a separate utility to be installed
   1. curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl-convert"
   2. validation of downlaoded binary [optional]:
      1. curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl-convert.sha256"
      2. echo "$(cat kubectl-convert.sha256) kubectl-convert" | sha256sum --check
   3. sudo install -o root -g root -m 0755 kubectl-convert /usr/local/bin/kubectl-convert
   4. kubectl convert --help
4. Enable the v1alpha1 version for rbac.authorization.k8s.io API group on the controlplane node.
   1. As a good practice, take a backup of that apiserver manifest file before going to make any changes.
      In case, if anything happens due to misconfiguration you can replace it with the backup file: `cp -v /etc/kubernetes/manifests/kube-apiserver.yaml /root/kube-apiserver.yaml.backup`
   2. Now, open up the kube-apiserver manifest file in the editor of your choice. It could be vim or nano: `vi /etc/kubernetes/manifests/kube-apiserver.yaml`
   3. Add the --runtime-config flag in the command field as follows : `--runtime-config=rbac.authorization.k8s.io/v1alpha1`
   4. usage: `kubectl convert -f ingress-old.yaml --output-version networking.k8s.io/v1`


## Custom Resource Definitions
1. kubectl commands -> Create, list and modify in ETCD data store
2. Who is responsible for converting these deployments into actual pods and replicasets? -> Controller
3. Controller -> is built-in -> continuously monitor resource states
4. Suppose we want to create a resource like below:
   ```
   apiVersion: flights.com/v1
   kind: FlightTicket
   metadata: 
      name: my-flight-ticket
   spec:
      from: Mumbai
      to: London
      number: 2
   ```
   We would like to manage the resource, like this:
   - kubectl create -f flightticket.yaml
   - kubectl get flightticket
   - kubectl delete -f flightticket.yaml
5. How to achieve?
  - Create a custom flight ticket controller that keeps checking the ETCD for any flight ticket request.
  - CRD
    ```
    apiVersion: apiextensions.k8s.io/v1
    kind: CustomResourceDefinition
    metadata: 
      name: flighttickets.flights.com
    spec:
      scope: Namespaced
      group: flights.com
      names:
         kind: FlightTicket
         singular: flightticket
         plural: flighttickets
         shortnames: 
            - ft
      versions:
         - name: v1
           served: trye
           storage: true
      schema:
         openAPIV3Schema: 
            type: object
            properties:
               spec:
                  type: object
                  properties:
                     from:
                        type: string
                     to:
                        type: string
                     number:
                        type: integer
                        minimum: 1
                        maximum: 10
    ```

## Custom Controllers
1. Deploy a code that monitors ETCD for FlightTicket resource requests and manages.
2. Get started:
   1. Clone repo: https://github.com/kubernetes/sample-controller
   2. cd sample-controller
   3. write your custom logic in controller.go
   4. go build -o sample-controller .
   5. ./sample-controller -kubeconfig=$HOME/.kube/config
   6. package the sample-controller and run it as a pod

## Operator framework
1. CRD and Custom controllers can be packaged together as a Operator framework and deployed.
2. Popular operator: ETCD operator -> EtcdCluster CRD + ETCD controller , EtcdBackup + Backup Operator, EtcdRestore + Restore operator
3. Operator available at: operatorhub.io


## Deployment strategies: Blue Green and Canary
1. Already discussed: Recreate and RollingUpdate(default)
2. Blue Green:
   1. Newer version (Green) is deployed along older version (Blue)
   2. 100% traffic on Blue
   3. Switch traffic to Green all at once
   4. Better inplemented with service mesh like ISTIO
   5. Steps
      1. Already have Deployment1 (label = version:v1) and Service1 (selector = version:v1)
      2. New deployment Deployment2 (label = version:v2) and Service1 (selector = version:v2)
3. Canary:
   1. Newer version (Green) is deployed along older version (Blue)
   2. Route small % of traffic - majority traffic to older
   3. Run tests on new
   4. Upgrade original with newer version
   5. Steps:
      1. Already have Deployment1 (label = version:v1) and Service1 (selector = version:v1)
      2. New deployment Deployment2 (label = version:v2 and app:FE)
      3. Update older deployment Deployment1 (label = version:v1 and app:FE)
      4. Update Service1 (selector = app:FE) -> But traffic is equally divided
      5. Reduce the Newer deployment pods to as required, so that traffic is divided as required
      6. After tests, Update older deployment Deployment1 with new image version (label = version:v2)
      7. Update Service1 (selector = version:v2)
      8. Caveat: Traffic split is governed by number of pods
      9. Service Mesh is solution for this caveat
4. Update replicas of an existing deployment: `kubectl scale deployment --replicas=1 frontend-v2`


