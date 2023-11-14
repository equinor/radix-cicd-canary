package kubernetes

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetKubernetesClient Gets a kubernetes client using the config of the running pod
func GetKubernetesClient() kubernetes.Interface {
	config := getKubernetesClientConfig()
	return getKubernetesClientFromConfig(config)
}

func getKubernetesClientConfig() *rest.Config {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			err := errors.WithStack(err)
			log.Fatal().Stack().Err(err).Msg("getClusterConfig InClusterConfig")
		}
	}

	return config
}

func getKubernetesClientFromConfig(config *rest.Config) kubernetes.Interface {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		err := errors.WithStack(err)
		log.Fatal().Stack().Err(err).Msg("getClusterConfig k8s client")
	}

	return client
}
