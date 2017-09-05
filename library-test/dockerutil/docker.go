package dockerutil

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/filters"
	"github.com/docker/go-connections/nat"
)

type Options struct {
	Cmd  string
	Name string
}

func Cmd() *Options {
	opts := &Options{}
	flag.StringVar(&opts.Cmd, "cmd", "", "cmd to run")
	flag.StringVar(&opts.Name, "name", "", "scheduler name")
	flag.Parse()
	return opts
}

func Image(ctx context.Context, cli *docker.Client, imageName string) error {
	existImage := false

	// Pull image if does not exist
	f := filters.NewArgs()
	f.Add("dangling", "false")
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     true,
		Filters: f,
	})

	if err != nil {
		fmt.Println(err)
	}

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
	return err
}

func StartContainer(ctx context.Context, cli *docker.Client, imageName string, containerName string, internalPort string, hostIP string, cmds []string, log bool) (containerID string, err error) {
	err = Image(ctx, cli, imageName)
	if err != nil {
		return
	}

	port := nat.Port(internalPort)
	existContainer := false
	var out io.ReadCloser

	config := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			port: struct{}{},
		},
		Cmd: cmds,
	}

	// Use random port for binding
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			port: []nat.PortBinding{
				{
					HostIP: hostIP,
				},
			},
		},
	}

	// Create container if does not exist
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name == "/"+containerName {
				existContainer = true
				break
			}
		}
	}

	if !existContainer {
		// Create container
		var resp container.ContainerCreateCreatedBody
		resp, err = cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
		if err != nil {
			return
		}
		containerID = resp.ID
	} else {
		containerID = containers[0].ID
	}

	if log {
		options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Details: true}
		out, err = cli.ContainerLogs(ctx, containerID, options)
		if err != nil {
			panic(err)
		}
	}

	// Start container
	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return
	}

	if log {
		io.Copy(os.Stdout, out)
	}

	return
}

func GetContainerPort(ctx context.Context, cli *docker.Client, innerPort string, containerID string) (port int, err error) {
	// Get exposed port
	res, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return
	}

	binding := res.NetworkSettings.NetworkSettingsBase.Ports[nat.Port(innerPort)][0]

	port, err = strconv.Atoi(binding.HostPort)

	return
}

func ContainerWait(ctx context.Context, cli *docker.Client, containerID string) {
	statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
}

func RemoveContainer(ctx context.Context, cli *docker.Client, containerID string) error {
	return cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
}
