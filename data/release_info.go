package data

type Release struct {
	Name   string
	Assets []ReleaseAssets
}

type ReleaseAssets struct {
	Name                 string
	Browser_download_url string
}
