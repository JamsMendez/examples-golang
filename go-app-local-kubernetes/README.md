```
> kind load docker-image [image:tag] --name [cluster-name]
> minikube image load [image:tag]
```
```
> kubectl cluster-info --context [cluster-name]
```

```
> docker container exec -it [container-cluster-name] crictl images
```

```
> kubectl apply -f deployment.yaml
> kubectl apply -f service.yaml

```