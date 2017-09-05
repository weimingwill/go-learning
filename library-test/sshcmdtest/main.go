package main

import (
	"fmt"
	"hue-ec-swat-golang/common/sshcmd"
)

func main() {
	sshClient, err := sshcmd.NewSSHClient("/home/weiming/.ssh/id_rsa", "localhost:22")
	if err != nil {
		fmt.Println(err)
	}
	//err = sshClient.ExecuteCmd("docker run -w /go/src/hue-ec-swat-golang/ares -p 3000 ares -url=http://172.26.158.201:17050/api/v2/api-docs -target=/v1/products/{productId} -method=get -payload=payload.json -rate=1 -requests=20 -server=true -scheduler=load")
	res, err := sshClient.ExecuteCmd("docker container inspect ares")
	fmt.Println(res)
	//err = sshClient.ExecuteCmd("cd /home/weiming/go/src/hue-ec-swat-golang/ares", "./ares -url=http://172.26.158.201:17050/api/v2/api-docs -target=/v1/products/{productId} -method=get -payload=payload.json -rate=1 -requests=5 -server=true -scheduler=load")
	//err = sshClient.ExecuteCmd("pwd")
	if err != nil {
		fmt.Println("error:", err)
	}
}
