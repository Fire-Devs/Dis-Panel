package controllers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerStruct struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Ports  []string `json:"ports"`
	Image  string   `json:"image"`
}

func GetALLContainers() []ContainerStruct {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var containerStruct []ContainerStruct

	for _, container := range containers {
		var ports []string
		for _, port := range container.Ports {
			ports = append(ports, fmt.Sprintf("%d", port.PublicPort))
		}
		// append with struct key
		containerStruct = append(containerStruct, ContainerStruct{
			ID:     container.ID,
			Name:   container.Names[0],
			Status: container.State,
			Ports:  ports,
			Image:  container.Image,
		})

	}
	return containerStruct

}

func StopContainer(id string) (string, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return "Error", err
	}
	err = cli.ContainerStop(context.Background(), id, nil)
	if err != nil {
		return "Error", err
	}

	return "Success", nil
}

func StartContainer(id string) (string, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return "Error", err
	}
	err = cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		return "Error", err
	}

	return "Success", nil
}