package services

import (
	"bytes"
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetServiceLogs(name string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "Error creating Docker client: " + err.Error()
	}

	out, err := cli.ContainerLogs(context.Background(), name, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Tail:       "50",
	})

	if err != nil {
		return "Error fetching logs: " + err.Error()
	}
	defer out.Close()

	var buf bytes.Buffer
	io.Copy(&buf, out)

	return buf.String()
}
