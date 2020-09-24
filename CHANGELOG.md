# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.1] - 2020-09-24

### Fixed

- Do not use priorityClass in KVM

## [1.4.0] - 2020-09-24

### Changed
- Updated helm chart to use resource helpers, more precise labels.
- Deploy `node-exporter-app` on installations as part of app collection.

## [1.3.0] - 2020-07-23

### Added

- Added workflows for automatic releases.

### Updated

- Updated node-exporter version to 1.0.1.
- Newly disabled collectors: powersupplyclass, schedstat, udp_queues.

## [1.2.0] 2020-01-08

### Changed

- Change Priority Class to `system-node-critical`

## [1.1.1] 2019-12-18

### Changed

- Updated dependencies to support Kubernetes 1.16.

## [1.1.0] 2019-11-20

Note: Version number 1.0.0 was skipped to align the version with other default apps.

### Changed

- Updated to node-exporter version 0.18.1.

## [0.6.0] 2019-10-02

### Changed

- Migrated to be deployed via an app CR not a chartconfig CR.

## [0.5.1] 2019-07-17

### Changed

- Tolerations changed to tolerate all taints.
- Change prioty class from to `giantswarm-critical`.
- Run node-exporter as root user to get proper permissions.
- Remove CPU limits, decrease CPU requests.

## [0.4.1] 2019-06-28

### Fixed

- Fix systemd collector D-Bus connection. https://github.com/giantswarm/kubernetes-node-exporter/pull/44

## [0.4.0] 2019-06-14

### Changed

- Disabled ipvs collector.

### Fixed

- Fix monitored file system mount points.

## [0.3.0] 2019-05-24

### Changed

- Upgrade to node-exporter 0.18.0.

## [0.2.0] 2019-05-17

### Added

- Separate pod security policy for node-exporter and node-exporter-migration workloads.
- Security context with non-root user (`nobody`) for running node-exporter inside container.

[Unreleased]: https://github.com/giantswarm/node-exporter-app/compare/v1.4.1...HEAD
[1.4.1]: https://github.com/giantswarm/node-exporter-app/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/giantswarm/node-exporter-app/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/node-exporter-app/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/giantswarm/node-exporter-app/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/node-exporter-app/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/node-exporter-app/compare/v0.6.0...v1.1.0
[0.6.0]: https://github.com/giantswarm/node-exporter-app/releases/tag/v0.6.0
[0.5.1]: https://github.com/giantswarm/kubernetes-node-exporter/compare/v0.4.1...v0.5.1
[0.4.1]: https://github.com/giantswarm/kubernetes-node-exporter/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/kubernetes-node-exporter/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/kubernetes-node-exporter/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/kubernetes-node-exporter/releases/tag/v0.2.0
