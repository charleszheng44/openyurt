---
title: Edge Device Management
authors:
  - '@charleszheng44'
reviewers:
  - '@yixingjia'
  - '@Fei-Guo'
  - '@rambohe-ch'
  - '@kadisi'
  - '@huangyuqi'
  - '@Walnux'
creation-date: 2021-03-10T00:00:00.000Z
last-updated: 2021-03-10T00:00:00.000Z
status: provisional
---
# Managing Edge Devices using EdgeX Foundry

## Table of Contents

- [Managing Edge Devices using EdgeX Foundry](#managing-edge-devices-using-edgex-foundry)
  * [Table of Contents](#table-of-contents)
  * [Glossary](#glossary)
  * [Summary](#summary)
  * [Motivation](#motivation)
    + [Goals](#goals)
    + [Non-Goals/Future Work](#non-goalsfuture-work)
  * [Proposal](#proposal)
    + [User Stories](#user-stories)
    + [Components](#components)
    + [Interaction with EdgeX Foundry](#interaction-with-edgex-foundry)
      - [Setting up EdgeX Foundry](#setting-up-edgex-foundry)
      - [Interaction with EdgeX services](#interaction-with-edgex-services)
  * [Upgrade Strategy](#upgrade-strategy)
  * [Implementation History](#implementation-history)

## Glossary

Refer to the [OpenYurt Glossary](docs/proposals/00_openyurt-glossary.md) and the [EdgeX Foundry Glossary](https://docs.edgexfoundry.org/1.3/general/Definitions/).

## Summary

## Motivation

### Goals

- To design a new custom resource definition(CRD), DeviceProfile, to represent different categories of devices
- To design a new CRD, Device, that represents a physical edge devices
- To design a new CRD, DeviceService, that defines the way of how to connect to a specific device
- To support device management using DeviceProfile, Device, DeviceService and EdgeX Foundry 
- To support automatically setting up of the EdgeX Foundry on the OpenYurt
- To support declartive device state modification, i.e., modifying the device's properties by changing fields of the device CRs

### Non-Goals/Future Work

Non-goals are limited to the scope of this proposal, these features may evolve in the future.

- To implement DeviceService for any specific protocol
- To support data transmission between edge devices and external services
- To support edge data processing   

## Proposal

### User Stories

1. As a vendor, I would like to connect a category of device into the OpenYurt.
2. As an end user, I would like to define how to connect a device, which belongs to a supported DeviceProfile, into the OpenYurt.
3. As an end user, I would like to connect a new device, which belongs to a supported DeviceProfile, into the OpenYurt.
4. As an end user, I would like to modify the states of devices by changing the values of device properties defined in corresponding Device CRs.
5. As an end user, I would like to disconnect a device by deleting the corresponding Device CR.

### Components

### Interaction with EdgeX Foundry 

#### Setting up EdgeX Foundry

#### Interaction with EdgeX services

## Upgrade Strategy

In the first implementation, we will support the EdgeX Foundry [Hanoi](), and would not support upgrading/downgrading to other versions.

## Implementation History

- [ ] 03/15/2021: Proposed idea in an issue or [community meeting](https://us02web.zoom.us/j/82828315928?pwd=SVVxek01T2Z0SVYraktCcDV4RmZlUT09)
- [ ] MM/DD/YYYY: First round of feedback from community
- [ ] MM/DD/YYYY: Present proposal at a [community meeting](https://us02web.zoom.us/j/82828315928?pwd=SVVxek01T2Z0SVYraktCcDV4RmZlUT09)
- [ ] MM/DD/YYYY: Open proposal PR


