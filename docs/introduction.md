# ClickHouse Operator Introduction

## Prerequisites
You may need to have the following items in order to have ClickHouse installation in k8s

1. Persistent Volumes
1. Zookeeper

### Persistent Volumes
ClickHouse needs disk space to keep data. Kubernetes provides [Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
for this purpose.
As it is stated
> A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator.

This means that we have to do some homework in order to provide Persistent Volumes to ClickHouse installation.
PVs can be provided by:

1. system administrator, who is in charge of k8s installation, can prepare required number of PVs
1. Persistent Volume Provisioner, which may be set up in k8s installation in order to [provision volumes dynamically](https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/)

When ClickHouse required some disk storage, in places [Persistent Volume Claim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)
which specifies desired storage class and size. Each Persistent Volume has class assigned and size provisioned.
So the main bond between software and disk to be provisioned is [Storage Class](https://kubernetes.io/docs/concepts/storage/storage-classes/).

### Zookeeper

In case we'd like to have [data replication](https://clickhouse.yandex/docs/en/operations/table_engines/replication/) in ClickHouse,
we need to have [Zookeeper](https://zookeeper.apache.org/) instance accessible by ClickHouse.
There is no requirement to have Zookeeper instance dedicated to serve ClickHouse replication, we just need to have access to running Zookeeper.
However, in case we'd like to have high-available ClickHouse installation, we need to have Zookeeper cluster of at least 3 nodes.
So, we can either use
1. Already existing Zookeeper instance, or
1. Setup our own Zookeeper - in most cases inside the same k8s installation.