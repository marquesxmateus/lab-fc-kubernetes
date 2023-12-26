# Kubernetes 

As informações de sobre os cluster ficam no kubeconfig. ex: 
```
\\wsl.localhost\Ubuntu\home\mateus\.kube
```

## Kind K8S
```bash
# create cluster 
kind create cluster --config=k8s/kind.yaml  --name=kubelearn   
# set context
kubectl cluster-info --context kind-kubelearn
```
## Kubernetes commands
```bash
# list cluster 
kubectl config get-clusters
# get node from context
kubectl get nodes 
# switch context
kubectl config use-context cluster-name
# apply 
kubectl apply -f k8s/pod.yaml 
# get pods 
kubectl get po   
# access container 
kubectl port-forward pod/goserver 8081:8081
# access service 
kubectl port-forward svc/goserver-service 8081:8081
# delete pod 
kubectl delete pod goserver
# history 
kubectl rollout history deployment goserver
# rollback to last version 
kubectl rollout undo deployment goserver 
kubectl rollout undo deployment goserver --to-revision=2
# describe
kubectl describe pod goserver-5bf6c594b9-5hl25
# get services
kubectl get svc
# delete service
kubectl delete svc goserver-service
# acesso a api do kubernetes
kubectl proxy --por=8081  
# acessar dentro do pod 
kubectl exec -it goserver-68594bd665-cfzl2 -- bash 
# acessar logs do pod 
kubectl logs goserver-68594bd665-cfzl2
```

## Docker commands 
```bash
# create image
docker build -t marquesxmateus/hello-go .
# run 
docker run --rm -p 8081:8081 marquesxmateus/hello-go
# push 
docker push marquesxmateus/hello-go
```

## Extra 
```bash
echo "mateus" | base64
```