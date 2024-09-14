package store

import (
	"fmt"
	"github-app-store/data"
	"io"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"gopkg.in/yaml.v2"
)

type Store struct {
	osType       string
	gitHubClient *api.RESTClient
}

func NewStore(ostype string) (*Store, error) {
	if client, err := api.NewRESTClient(api.ClientOptions{
		AuthToken: os.Getenv("GITHUB_TOKEN"),
	}); nil != err {
		return nil, err
	} else {
		return &Store{
			osType:       ostype,
			gitHubClient: client,
		}, nil
	}
}

func (store *Store) AppSource() data.AppSource {
	var appsConf string
	switch store.osType {
	case "Anroid":
		appsConf = "android_apps.yaml"
	case "Windows":
		appsConf = "windows_apps.yaml"
	case "MacOS":
		appsConf = "mac_os_apps.yaml"
	default:
		appsConf = "android_apps.yaml"
	}
	pwd, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/asstes/%s", pwd, appsConf))
	if nil != err {
		fmt.Println(err.Error())
	}
	yamlBytes, _ := io.ReadAll(file)
	apps := data.AppSource{}
	yaml.Unmarshal(yamlBytes, &apps)
	return apps
}

// / 获取app最新的 Release 版本数据
func (store *Store) AppRelease(appDes data.AppRepositories) (*data.Release, error) {
	if releases, err := store.AppReleases(appDes, 1); nil != err {
		return nil, err
	} else {
		return &releases[0], nil
	}
}

// / 获取app部分 Release 版本数据
func (store *Store) AppReleases(appDes data.AppRepositories, pages int) ([]data.Release, error) {
	client := store.gitHubClient

	response := []data.Release{}
	fmt.Println(appDes)
	if err := client.Get(fmt.Sprintf("repos/%s/%s/releases?per_page=1", appDes.User, appDes.Repositories), &response); err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		return response, nil
	}
}
