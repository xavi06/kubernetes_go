package kubernetes

import (
	"fmt"
	"testing"
)

func TestLogs(t *testing.T) {
	baseurl := "https://172.16.151.70:6443"
	namespace := "default"
	pod := "app-microservice-web-8fb8ffddb-x4x5h"
	token := "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyLXRva2VuLXRicGtzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI1YjY2MmQ2Zi1iZjRhLTExZTktYWY2Ny1mYTE2M2U2NTk4M2QiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06YWRtaW4tdXNlciJ9.bhoGq45idcHbwCgrjlNWGpzEuwghqne8z9asLwAsX3NIJV3W44HbFq4cyCHPc_3GXl8NKcebE5w64v6bRheM4qqYkzU-GIHMVkH2HtfADv5GQ8-XcE5ngu5cnnARATU4mYaUMsAqqlgyS1gtQePcsg4qWxgzgpc2VnxwRsG96zHiegm7jUlfXzLncBdOXmoHiBkyjao7LoAotdRcUcpHQjcUY2Ba7XFdTmDODxoMHvfyz14GChhAmXwytpRmFODucSPa0v0BtIMSqQRn4DT7zyrq0FTbWgsq0_dTk52tgkdx13TRVNArp3-hVHbAo7sexFx9xfOojbS1vhhj5qmT7A"
	line := 20
	data, err := GetPodLogs(baseurl, namespace, pod, token, line)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}
