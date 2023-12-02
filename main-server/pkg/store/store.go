package store

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
)

func ResolvePath(dir string) (string, error) {
	home, err := os.UserHomeDir()

	if dir == "~" {
		if err != nil {
			return "", err
		}
		return home, nil
	}

	absPath := filepath.Join(home, dir)
	return absPath, nil
}

func StartDockerContainer(rootDir string, team int, portInt int) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	port := strconv.Itoa(portInt)
	// Docker image to use (replace with the actual image)
	imageName := "codercom/code-server"
	zap.S().Infof("Docker container on port : %v", port)

	// Container name based on the team
	containerName := fmt.Sprintf("team%d-code-server", team)

	// HostConfig for volume mount and port binding

	//hostConfig := &container.HostConfig{
	//	Binds: []string{fmt.Sprintf("%s:/home/coder/project", rootDir)},
	//	PortBindings: map[nat.Port][]nat.PortBinding{
	//		"8080/tcp": {{HostIP: "0.0.0.0", HostPort: "8080"}},
	//	},
	//}

	tcpStr := port + "/tcp"
	hostConfig := &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:/home/coder/project", rootDir)},
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(tcpStr): {{HostIP: "0.0.0.0", HostPort: port}},
		},
	}
	//hostConfig := &container.HostConfig{
	//	Binds: []string{fmt.Sprintf("%s:/home/coder/project", rootDir)},
	//	PortBindings: map[nat.Port][]nat.PortBinding{
	//		nat.Port(fmt.Sprintf("%s/tcp", port)): {{HostIP: "0.0.0.0", HostPort: fmt.Sprintf("%s", port)}},
	//	},
	//}

	// ContainerConfig with environment variable for code-server authentication
	containerConfig := &container.Config{
		Image: imageName,
		//Cmd:   []string{"--auth=none", "--disable-telemetry", fmt.Sprintf("--port=%s", port)},
		Cmd: []string{"--auth=none", "--disable-telemetry"},
		Env: []string{"PASSWORD=password"}, // Replace with your desired authentication method
	}

	// Create the container
	resp, err := cli.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		return "", err
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}