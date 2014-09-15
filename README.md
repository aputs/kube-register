# kube-register

Register Kubernetes Kubelet machines with the Kubernetes API server using Fleet.

## Usage

```
kube-register -metadata="kubelet=true" -fleet-endpoint="http://127.0.0.1:4002" -apiserver-endpoint="http://127.0.0.1:8080"
```
