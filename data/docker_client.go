package data

import (
	"fmt"
	"os/exec"
)

func StopLocalContainer(name string) (out []byte, err error) {
	fmt.Println("stop local container " + name)
	cmdStr := "docker stop " + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", out)
	}

	return
}

func PauseLocalContainer(name string) (out []byte, err error) {
	fmt.Println("pause local container " + name)
	cmdStr := "docker pause " + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", out)
	}

	return
}

func StartLocalContainer(name string) (out []byte, err error) {
	fmt.Println("start local container " + name)
	cmdStr := "docker start " + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", out)
	}

	return
}

func RestartLocalContainer(name string) (out []byte, err error) {
	fmt.Println("restart local container " + name)
	cmdStr := "docker restart " + name
	out, err = exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", out)
	}

	return
}

// func DownloadLocalContainer(name string) (out []byte, err error) {
// 	cmdStr := "docker restart " + name
// 	out, _ = exec.Command("/bin/sh", "-c", cmdStr).Output()
// 	fmt.Printf("%s", out)
// 	return
// }

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
