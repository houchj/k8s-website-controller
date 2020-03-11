package v1

type Metadata struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

type WebsiteSpec struct {
	GitRepo string `json:"gitRepo"`
}

type Website struct {
	Metadata Metadata `json:"metadata"`
	Spec WebsiteSpec `json:"spec"`
}

type WebsiteWatchEvent struct {
	Kind string `json:"type"`
	MyObject Website `json:"object"`
}
