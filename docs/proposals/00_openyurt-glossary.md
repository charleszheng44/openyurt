# Table of Contents

This document lists terms for the OpenYurt implementation.

[C](#c) | [E](#e) | [N](#n) | [U](#u) | [Y](#y)

# C
---

### CloudNode

The node that are located on the cloud. The control-plane and other cluster management components are usually running on the CloudNode.

# E
---

### EdgeNode

The node that are accessible to the edge device. The EdgeNode are usually located in sub-optimal network environment and may be disconnected to the cloud node at any time.

### End User

Represents a user of the OpenYurt cluster.

# N
---

### NodePool

A CRD that represents a pool of edge nodes in the same network region.

### NodePool Controller

The controller of the NodePool CRD, which reconciles the actual status with the desired state of the NodePool.

# U
---

### UnitedDeployment

TODO

### UnitedDeployment Controller

TODO

# Y
---

### YurtControllerManager

A replacement of the nodelifecycle controller that prevents the APIServer from evicting disconnected EdgeNodes.

### YurtHub

The local cache running on each EdgeNode, which periodically synchronizes the cluster state from the APIServer running on the cloud.

### YurtTunnel

The network tunnel that helps CloudNodes to send http requests to EdgeNodes located in an isolated network.

### YurtTunnel Server

The server of the YurtTunnel that run on each CloudNode and redirect http requests to corresponding agents.

### YurtTunnel Agent

The agent of the YurtTunnel that run on each EdgeNode, receive requests from the YurtTunnel Server and send requests to destination hosts.   

### YurtAppManager

The controller manager that includes the NodePool Controller and the UnitedDeployment Controller.
