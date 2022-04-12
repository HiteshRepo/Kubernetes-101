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