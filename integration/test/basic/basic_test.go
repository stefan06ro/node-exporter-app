// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/node-exporter-app/integration/key"
)

// TestDaemonSets ensures that node-exporter daemonset are deployed
func TestDaemonSets(t *testing.T) {
	ctx := context.Background()
	var err error

	err = checkDaemonsetExists(ctx)
	if err != nil {
		t.Fatalf("could not get node-exporter: %v", err)
	}
}

func checkDaemonsetExists(ctx context.Context) error {
	var err error

	config.Logger.Debugf(ctx, "waiting for daemonset create")

	o := func() error {
		lo := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", "app.kubernetes.io/name", key.AppName()),
		}
		daemonsets, err := config.AppTest.K8sClient().AppsV1().DaemonSets(metav1.NamespaceSystem).List(ctx, lo)
		if err != nil {
			return microerror.Mask(err)
		} else if len(daemonsets.Items) == 0 {
			return microerror.Maskf(executionFailedError, "daemonset %#q in %#q not found", key.AppName(), metav1.NamespaceSystem)
		}

		daemonset := daemonsets.Items[0]

		if daemonset.Status.DesiredNumberScheduled != daemonset.Status.CurrentNumberScheduled {
			return microerror.Maskf(executionFailedError, "desired number of pods are %d but have %d", daemonset.Status.DesiredNumberScheduled, daemonset.Status.CurrentNumberScheduled)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(config.Logger, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	config.Logger.Debugf(ctx, "daemonsets are ready")

	return nil
}
