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
7. 