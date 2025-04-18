package services

func ListRunningContainers() string {
    return "container_1, container_2"
}

func StartContainer(name string) string {
    return "Container " + name + " started"
}

func StopContainer(name string) string {
    return "Container " + name + " stopped"
}
