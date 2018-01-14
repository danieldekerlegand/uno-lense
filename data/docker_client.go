package data

import (
	"fmt"
	"os/exec"
)

func RunLocalImage() (out []byte, err error) {
	cmdStr := "docker start hello-world"
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	return
}

func ListRunningContainers() (out []byte, err error) {
	cmdStr := "curl --unix-socket /var/run/docker.sock http:/containers/json"
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	return
}

func ListImages() (out []byte, err error) {
	cmdStr := "curl --unix-socket /var/run/docker.sock http:/images/json"
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	return
}

func ListRemoteImages(repo string, username string, password string) (out []byte, err error) {
	cmdStr := "curl -X GET https://" + username + ":" + password + "@" + repo + "/v2/_catalog"
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	return
}

func PullImage(name string) (out []byte, err error) {
	cmdStr := "docker tag " + name + " registry.cs.uno.edu/" + name
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker pull registry.cs.uno.edu/" + name
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	return
}

func PushImage(name string) (out []byte, err error) {
	cmdStr := "docker login -u daniel -p pass https://registry.cs.uno.edu"
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker tag " + name + " registry.cs.uno.edu/" + name
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker push registry.cs.uno.edu/" + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	fmt.Printf("%s", err)
	return
}

func startRemoteImage(ip string, name string) (out []byte, err error) {
	cmdStr := "docker login -u daniel -p pass https://registry.cs.uno.edu"
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker tag " + name + " registry.cs.uno.edu/" + name
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker push registry.cs.uno.edu/" + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	fmt.Printf("%s", err)
	return
}

func stopRemoteImage(ip string, name string) (out []byte, err error) {
	cmdStr := "docker login -u daniel -p pass https://registry.cs.uno.edu"
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker tag " + name + " registry.cs.uno.edu/" + name
	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

	cmdStr = "docker push registry.cs.uno.edu/" + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
	fmt.Printf("%s", err)
	return
}