package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"

	yaml "gopkg.in/yaml.v2"

	"bytes"
	"hue-ec-swat-golang/common/logger"
	"hue-ec-swat-golang/common/swatutils"
	"strings"

	"github.com/ericchiang/k8s"
)

func main() {
	leaderPod := getPods("", "leader-node=true")
	fmt.Println(leaderPod)
	//os.Setenv("KUBERNETES_SERVICE_HOST", "172.26.158.167")
	//os.Setenv("KUBERNETES_SERVICE_PORT", "32005")
	//
	//labelNewLeaderPod("", "load-tester-v8868")

	////client, err := loadClient("/home/weiming/swat/kube/ngix/nginx.yaml")
	//client, err := k8s.NewInClusterClient()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var nodes corev1.NodeList
	//if err := client.List(context.Background(), "", &nodes); err != nil {
	//	log.Fatal(err)
	//}
	//for _, node := range nodes.Items {
	//	fmt.Printf("name=%q schedulable=%t\n", *node.Metadata.Name, !*node.Spec.Unschedulable)
	//}
	//
	//label := new(k8s.LabelSelector)
	//label.Eq("name", "load-tester")
	//
	//var pods corev1.PodList
	//ctx := context.Background()
	//err = client.List(ctx, "custom-namespace", &pods, label.Selector())
	//if err != nil {
	//	fmt.Println(err)
	//}

}

func getPods(k8sHost string, selector string) []string {
	c1 := exec.Command("kubectl", "get", "pods", "--selector="+selector, "--namespace=load-tester")
	c2 := exec.Command("awk", "NR > 1 {print $1}")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var podsByte bytes.Buffer
	c2.Stdout = &podsByte

	err := c1.Start()
	if err != nil {
		logger.Error.Printf("failed to start command to get load tester pods, error: %s", err)
	}
	err = c2.Start()
	if err != nil {
		logger.Error.Printf("failed to start command to filter pods name, error: %s", err)
	}

	err = c1.Wait()
	if err != nil {
		logger.Error.Printf("failed to wait for value return from command to get load tester pods, , error: %s", err)
	}

	err = w.Close()
	if err != nil {
		logger.Error.Printf("failed to close write pipe, error: %s", err)
	}

	err = c2.Wait()
	if err != nil {
		logger.Error.Printf("failed to wait for value return from command to filter pods name, error: %s", err)
	}

	podsStr := podsByte.String()
	pods := strings.Split(podsStr, "\n")
	pods = swatutils.RemoveStrElem(pods, len(pods)-1)
	return pods
}

func labelNewLeaderPod(k8sHost string, hostName string) {
	var stdout, stderr bytes.Buffer
	c := exec.Command("kubectl", "label", "pods", hostName, "leader-node=true", "--namespace=load-tester", "--overwrite")
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		logger.Error.Printf("failed to label leader-node for new leader pod, error: %s", err)
		logger.Error.Printf(stderr.String())
		return
	}
	logger.Info.Printf(stdout.String())
}

func loadClient(kubeconfigPath string) (*k8s.Client, error) {
	data, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("read kubeconfig: %v", err)
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	return k8s.NewClient(&config)
}
