package module

type ReleaseOptions struct {
	Name      string
	Namespace string
	Revision  int
	IsUpgrade bool
	IsInstall bool
}
