package data

type Release struct {
	Name   string
	Assets []ReleaseAssets
	Body   string
}

type ReleaseAssets struct {
	Name                 string
	Browser_download_url string
}
