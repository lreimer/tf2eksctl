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
	CloudWatch CloudWatch `yaml:"cloudWatch"`
}

// the Service Account
type ServiceAccount struct {
	Metadata struct {
		Name      string
		Namespace string
	}
	WellKnownPolicies map[string]bool `yaml:"wellKnownPolicies"`
}

type CloudWatch struct {
	ClusterLogging ClusterLogging `yaml:"clusterLogging"`
}

type ClusterLogging struct {
	EnableTypes []string `yaml:"enableTypes,flow"`
}

// Creates a new ClusterConfig manifest with Kind and ApiVersion set
func NewClusterConfig() ClusterConfig {
	return ClusterConfig{
		ApiVersion: "eksctl.io/v1alpha5",
		Kind:       "ClusterConfig",
	}
}
