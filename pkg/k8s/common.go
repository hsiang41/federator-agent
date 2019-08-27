package k8sInquirer

type k8sNode struct {
	Name    string
	Namespaces string
	Kind    string
	Role    string
	Region  string
	Zone    string
	Size    string
	InstanceType string
	Os      string
	Provider string
	InstanceID string
	StorageSize int64
	Cpu         int64
}
