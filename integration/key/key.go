// +build k8srequired

package key

func AppName() string {
	return "node-exporter"
}

func CRName() string {
	return "node-exporter-app"
}

func DefaultCatalogName() string {
	return "default"
}

func DefaultTestCatalogName() string {
	return "default-test"
}

func Namespace() string {
	return "giantswarm"
}
