package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

const (
	EtcdELB      = "internal-tokyo-main-2-elb-etcd-1309004264.ap-northeast-1.elb.amazonaws.com"
	EtcdPort     = 4001
	K8SELB       = "internal-tokyo-main-2-elb-k8s-839152653.ap-northeast-1.elb.amazonaws.com"
	K8SPort      = 8080
	NameSpace    = "load-tester"
	NameSelector = "name=load-tester"
)

func main() {
	var stdout, stderr bytes.Buffer

	//host := fmt.Sprintf("http://%s:%s", K8SELB, K8SPort)

	//cmd := exec.Command("echo", "$FOO")
	//cmd := exec.Command("kubectl", "-s", "http://${K8S_ELB}:${K8S_PORT} get pods --selector=${POD_SELECTOR} --namespace=${LOAD_TESTER_NAMESPACE} | awk 'NR > 1 {print $1}'")

	//os.Setenv("POD_SELECTOR", "\"name=load-tester\"")

	//c1 := exec.Command("kubectl", "get", "pods", "--selector=name=load-tester", "--namespace=load-tester")
	c1 := exec.Command("kubectl", "get", "pods", "--selector=app=nginx")
	c2 := exec.Command("awk", "NR > 1 {print $1}")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var podsByte bytes.Buffer
	c2.Stdout = &podsByte

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	podsStr := podsByte.String()
	pods := strings.Split(podsStr, "\n")
	for _, pod := range pods {
		// To skip last index of pods slice, which is an empty string
		if pod == "" {
			continue
		}

		// remove leader-node label
		//c3 := exec.Command("kubectl", "-s", host, "label", "pods", pod, "leader-node-", "--namespace="+NameSpace)
		//new := fmt.Sprintf("leader-node=\"true\"")
		c3 := exec.Command("kubectl", "label", "pods", pod, `leader-node="true"`)
		c3.Stdout = &stdout
		c3.Stderr = &stderr
		err := c3.Run()
		if err != nil {
			fmt.Printf("failed to remove leader-node from pod %s, error: %s\n", pod, err)
		}
	}


	// label leader-node for new leader pod
	//c4 := exec.Command("kubectl", "-s", host, "label", "pods", "$LOAD_TESTER_LEADER_POD", "leader-node=\"true\"", "--namespace="+NameSpace)
	//c4.Stdout = &stdout
	//c4.Stderr = &stderr
	//err := c4.Run()
	//if err != nil {
	//	fmt.Printf("failed to label leader-node for new leader pod, error: %s", err)
	//}

	//io.Copy(os.Stdout, &pods)
	//fmt.Println()

	//cmd.Env = append(os.Environ(),
	//	"LOAD_TESTER_RC=load-tester",
	//	"LOAD_TESTER_NAMESPACE=load-tester",
	//	"LOAD_TESTER_LEADER_KEY=\"HostName\"",
	//	"ETCD_ELB=internal-tokyo-main-2-elb-etcd-1309004264.ap-northeast-1.elb.amazonaws.com",
	//	"ETCD_PORT=4001",
	//	"K8S_ELB=internal-tokyo-main-2-elb-k8s-839152653.ap-northeast-1.elb.amazonaws.com",
	//	"K8S_PORT=8080",
	//	"POD_SELECTOR=\"name=load-tester\"",
	//)

	//cmd := exec.Command("ls", "-lah")
	//var stdout, stderr bytes.Buffer
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Printf("cmd.Run() failed with %s\n", err)
	//}
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	//var out bytes.Buffer
	//cmd.Stdout = &out
	//
	//if err := cmd.Run(); err != nil {
	//	fmt.Println(err)
	//	log.Fatal(err)
	//}
	//fmt.Println(out.String())

	//out, err := exec.Command("date").Output()
	//if err != nil {
	//}
	//fmt.Printf("The date is %s\n", out)

	//cmd := exec.Command("ls")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//}
	//fmt.Printf("in all caps: %q\n", out.String())

	//cmdStr := "trigger-load-tester-svc.sh"
	//cmd := exec.Command("/bin/sh", cmdStr)
	//_, err := cmd.Output()
	//
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
}
