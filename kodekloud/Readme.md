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