## Tricks to work with kubectl faster

### Use aliases to switch contexts faster

```bash
alias devkube "kubectl config use-context kube-dev-context"
alias stgkube "kubectl config use-context kube-stg-context"
alias prdkube "kubectl config use-context kube-prd-context"
```

### Setup auto completion

1. Check if `bash-completion` is installed: `type _init_completion`

2. Install `bash-completion` if not already installed: `apt-get install bash-completion`

3. Set-up auto completion for kubectl
    - `source <(kubectl completion bash)`
    - For all system users: `kubectl completion bash | sudo tee /etc/bash_completion.d/kubectl > /dev/null`

### Use aliases for common kubectl commands

1. alias k='kubectl'
2. alias kcfg='k config'
3. alias kdp='kubectl describe pod'
4. alias kd='kubectl describe'
5. alias ke='kubectl explain'
6. alias kc='kubectl create'
7. alias kg='kubectl get'
8. alias krep='kubectl replace'
9. alias krem='kubectl delete'
10. alias kns='kubectl get namespaces'
11. alias ksvc='kubectl get svc'
12. alias kdep='kubectl get deploy'
13. alias kpod='kubectl get pod'
14. alias ksetimg='kubectl set image deploy'
15. alias kgpa='k get pod --all-namespaces'
16. alias kgaa='kubectl get all --show-labels'