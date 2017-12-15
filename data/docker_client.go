package data

import (
	"io"
	"os"
	"fmt"
	// "net/http"
	// "net"
	// "time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	_ "github.com/joho/godotenv/autoload"
)

func listLocalImages() (images []types.ImageSummary) {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, nil)
	// cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err = cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	return
}

func listRemoteImages() (images []types.ImageSummary) {
	// var headers map[string]string

	// tr := &http.Transport{}

	// tr.Dial = func(proto, addr string) (net.Conn, error) {
	// 	fmt.Println("Dial called")
	// 	conn, err := net.DialTimeout(proto, addr, time.Minute)
	// 	if err != nil {
	// 			fmt.Println("There was  an err", err)
	// 	}
	// 	return conn, err
	// }

	// cl := &http.Client{Transport: tr}

	// cli, err := client.NewEnvClient("https://0.0.0.0:5000", "v1.22", nil, nil)
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err = cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Printf("%+v \n", image)
	}

	return
}

func listRunningContainers() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, nil)
	// cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%+v \n", container)
	}
}

func pullImage() {
	ctx := context.Background()
	//cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, nil)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, "lesson4-comp1", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

func pushImage() {
	ctx := context.Background()
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, nil)
	// cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePush(ctx, "alpine", types.ImagePushOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

func runContainer() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	_, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
