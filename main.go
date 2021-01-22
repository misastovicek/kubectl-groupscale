package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	groupscale "github.com/misastovicek/kubectl-groupscale/cmd"

	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String(
			"kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	replicas := flag.Int("replicas", 1, "(optional) number of replicas for each matched deployment")
	label := flag.String("label", "", "Key/Value pair used to identify the kubernetes applications which should be scaled")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	xLabel := strings.SplitN(*label, "=", 2)
	labelKey := xLabel[0]
	labelValue := xLabel[1]

	groupscale.GroupScale(kubeconfig, labelKey, labelValue, *replicas)
}
