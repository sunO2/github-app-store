package data

type AppSource struct {
	Version string
	Apps    []AppRepositories
}

type AppRepositories struct {
	Name         string
	User         string
	Repositories string
	Des          string
}
