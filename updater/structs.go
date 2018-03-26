package updater

type GitHubStatues struct {
	URL         string `json:"url"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type Updater struct {
	Statues []GitHubStatues
}
