package dockerutil

import (
	"context"
	"strings"
	"testing"

	"docker.io/go-docker"
	"fmt"
)

func TestDocker(t *testing.T) {
	// Todo: how to mock and test docker?
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := "registry.docker.workslan/swat/ares"
	containerName := "ares"
	internalPort := "3000/tcp"
	hostIP := "0.0.0.0"
	cmds := strings.Split(" ", " ")

	containerID, err := StartContainer(ctx, cli, imageName, containerName, internalPort, hostIP, cmds)
	if err != nil {
		t.Fatal(err)
	}

	port, err := GetContainerPort(ctx, cli, internalPort, containerID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(port)
}
