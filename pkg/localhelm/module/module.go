package module

type Capabilities struct {
	// KubeVersion is the Kubernetes version.
	KubeVersion KubeVersion
}

type KubeVersion struct {
	Version string // Kubernetes version
	Major   string // Kubernetes major version
	Minor   string // Kubernetes minor version
}
