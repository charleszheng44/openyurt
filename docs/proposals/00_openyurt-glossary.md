# Table of Contents

This document lists terms for the OpenYurt implementation.

[C](#c) | [E](#e) | [N](#n) | [U](#u) | [Y](#y)

# C
---

### CloudNode

The node that runs on the cloud. The control-plane and other cluster management components are usually running on the CloudNode.

# E
---

### EdgeNode

The node that is accessible to the edge device. The EdgeNodes are usually located in a sub-optimal network environment. They may be disconnected from the cloud node at any time.

### End User

Represents a user of the OpenYurt cluster.

# N
---

### NodePool

The CRD represents a pool of edge nodes in the same network region.

# U
---

### UnitedDeployment

The CRD defines the way of deploying homogeneous workloads with different versions/configurations by NodePools.

# Y
---

### YurtControllerManager

A replacement of the nodelifecycle controller that prevents the APIServer from evicting disconnected EdgeNodes.

### YurtHub

The local cache running on each EdgeNode, which periodically synchronizes the cluster state from the APIServer running on the cloud.

### YurtTunnel

The network tunnel helps CloudNodes to send HTTP requests to EdgeNodes located in an isolated network.

### YurtTunnel Server

The server of the YurtTunnel runs on each CloudNode and redirects HTTP requests to corresponding agents.

### YurtTunnel Agent

The agent of the YurtTunnel that runs on each EdgeNode receives requests from the YurtTunnel Server and sends requests to destination hosts.

### YurtAppManager

The controller manager includes the NodePool Controller and the UnitedDeployment Controller.
