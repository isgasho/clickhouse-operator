Release 0.7.0

New features:
* Separated 'podVolumeClaimTemplate' and 'logVolumeClaimTemplate' to retain ClickHouse logs. 
* Sidecar clickhouse-logs container to view logs
* Significantly cleaned up templates model and 'useTemplates' extension
* new system_replicas_is_session_expired monitoring metric (#187 by @teralype)

Bug fixes:
* Fixed bug with installation name being truncated to 15 chars. The current limit is 60 chars. Cluster name is limited to 15.
* General stability improvements in corner cases

Release 0.6.0

New features:
* Added spec.stop property to start/stop all ClickHouse pods in installation for maintenance
* ClickHouseInstallationTemplate custom resource definition
* ClickHouseOperatorConfiguration custom resource definition

Improvements:
* Split operator into two binaries/containers - operator and monitor
* Added 10s timeout to ClickHouse connection
* Improved create/update logic
* Operator now looks at its own namespace if not specified explicitly
* Enhance multi-thread support for concurrent operations
