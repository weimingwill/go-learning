package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/filters"
	"github.com/docker/go-connections/nat"
)

func main() {

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	count := 0

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name == "/ares" {
				count++
			}
		}
	}
	fmt.Println(count)
}

func dockertest() {
	var containerID string
	existImage := false
	existContainer := false
	imageName := "registry.docker.workslan/swat/ares"

	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// Pull image if does not exist
	f := filters.NewArgs()
	f.Add("dangling", "false")
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     true,
		Filters: f,
	})

	for _, img := range imageList {
		for _, tag := range img.RepoTags {
			if strings.Contains(tag, imageName) {
				existImage = true
				break
			}
		}
	}

	if !existImage {
		out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, out)
	}

	config := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"3000/tcp": struct{}{},
		},
		Cmd: []string{"-url=http://172.26.158.201:17050/api/v2/api-docs"},
	}

	// Use random port for binding
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3000/tcp": []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
				},
			},
		},
	}

	// Create container if does not exist
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name == "ares" {
				existContainer = true
				break
			}
		}
	}

	if !existContainer {
		// Create container
		resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, "ares")
		if err != nil {
			fmt.Println(err)
		}
		containerID = resp.ID
	} else {
		containerID = containers[0].ID
	}

	// Start container
	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// Get exposed port
	res, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		fmt.Println(err)
	}

	binding := res.NetworkSettings.NetworkSettingsBase.Ports["3000/tcp"][0]
	fmt.Println("port", binding.HostPort)
}
