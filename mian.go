package main

import (
	"fmt"
	"github-app-store/store"
)

func main() {

	if store, err := store.NewStore("Android"); nil != err {
		fmt.Println(err.Error())
	} else {
		apps := store.AppSource()
		app := apps.Apps[0]
		if release, err := store.AppRelease(app); nil != err {
			fmt.Println(err)
		} else {
			for _, v := range release.Assets {
				fmt.Println(v.Name)
			}
		}
	}

}
