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

	l.LogCtx(ctx, "level", "debug", "message", "waiting for daemonset create")

	o := func() error {
		lo := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", "app.kubernetes.io/name", app),
		}
		daemonsets, err := appTest.K8sClient().AppsV1().DaemonSets(metav1.NamespaceSystem).List(ctx, lo)
		if err != nil {
			return microerror.Mask(err)
		} else if len(daemonsets.Items) == 0 {
			return microerror.Maskf(executionFailedError, "daemonset %#q in %#q not found", app, metav1.NamespaceSystem)
		}

		daemonset := daemonsets.Items[0]

		if daemonset.Status.DesiredNumberScheduled != daemonset.Status.CurrentNumberScheduled {
			return microerror.Maskf(executionFailedError, "desired number of pods are %d but have %d", daemonset.Status.DesiredNumberScheduled, daemonset.Status.CurrentNumberScheduled)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "daemonsets are ready")

	return nil
}
