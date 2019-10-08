[![CircleCI](https://circleci.com/gh/giantswarm/node-exporter-app.svg?style=shield&circle-token=0a5aafcebabaed6f39a57293a96427f907674276)](https://circleci.com/gh/giantswarm/node-exporter-app)

# node-exporter

Helm Chart for Node Exporter running in Tenant Clusters.

* Installs the [node-exporter service].

## Deployment

* Managed by [app-operator].
* Production releases are stored in the [default-catalog].
* WIP releases are stored in the [default-test-catalog].

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/node-exporter-app.git
$ cd node-exporter-app
$ helm install helm/node-exporter-app
```

Provide a custom `values.yaml`:

```bash
$ helm install node-exporter-app -f values.yaml
```

 ## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release which will push the tarball to the [default-catalog].
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/cluster-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[node-exporter service]: https://github.com/prometheus/node_exporter
