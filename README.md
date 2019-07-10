# kubernetes_go

go get github.com/xavi06/kubernetes_go

## Usage
```go
import (
    kubernetes "github.com/xavi06/kubernetes_go"
)

func main() {
    kubeconfig := "/root/.kube/config"
    clientset, err := kubernetes.GetClientset(kubeconfig)
    if err != nil {

    }
    ...

}
```
