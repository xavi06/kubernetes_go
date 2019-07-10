# kubernetes_go
kubernetes golang module, written in go

## Usage

```bash
go get github.com/xavi06/kubernetes_go
```

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
