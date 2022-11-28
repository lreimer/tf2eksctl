package model

// the ClusterConfig manifest
type ClusterConfig struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Region  string `yaml:"region"`
	}
	IAM struct {
		WithOIDC        bool             `yaml:"withOIDC"`
		ServiceAccounts []ServiceAccount `yaml:"serviceAccounts"`
	}
}

// the Service Account
type ServiceAccount struct {
	Metadata struct {
		Name      string
		Namespace string
	}
	WellKnownPolicies map[string]bool `yaml:"wellKnownPolicies"`
}

// Creates a new ClusterConfig manifest wut Kind and ApiVersion set
func NewClusterConfig() ClusterConfig {
	return ClusterConfig{
		ApiVersion: "eksctl.io/v1alpha5",
		Kind:       "ClusterConfig",
	}
}
