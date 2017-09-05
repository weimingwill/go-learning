package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"docker.io/go-docker"
	"hue-ec-swat-golang/common/dockerutil"
	"hue-ec-swat-golang/harmonia/registration"
)

func main() {
	var cmd string
	var name string
	flag.StringVar(&cmd, "cmd", "", "cmd to run")
	flag.StringVar(&name, "name", "", "scheduler name")
	flag.Parse()
	cmds := strings.Split(cmd, " ")

	imageName := "registry.docker.workslan/swat/ares"
	containerName := "ares"
	internalPort := "3000/tcp"
	hostIP := "0.0.0.0"

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containerID, err := dockerutil.StartContainer(ctx, cli, imageName, containerName, internalPort, hostIP, cmds, true)
	if err != nil {
		panic(err)
	}

	// Todo: figure out how to output results only once.
	defer func() {
		timeout := 0 * time.Second
		err = cli.ContainerStop(ctx, containerID, &timeout)
		if err != nil {
			panic(err)
		}
		//err = dockerutil.RemoveContainer(ctx, cli, containerID)
		//if err != nil {
		//	panic(err)
		//}
	}()

	//defer dockerutil.ContainerWait(ctx, cli, containerID)

	port, err := dockerutil.GetContainerPort(ctx, cli, internalPort, containerID)
	if err != nil {
		panic(err)
	}

	err = accepter.Register(port, name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(port)
}
