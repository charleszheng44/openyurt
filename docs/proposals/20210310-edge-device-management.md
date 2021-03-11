---
title: Edge Device Management
authors:
  - "@charleszheng44"
reviewers:
  - "@yixingjia"
  - "@Fei-Guo"
  - "@rambohe-ch"
  - "@kadisi"
  - "@huangyuqi"
  - "@Walnux"
creation-date: 2021-03-10
last-updated: 2021-03-10
status: provisional
---

# Managing Edge Devices using EdgeX Foundry

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a proposal and for highlighting
any additional information provided beyond the standard proposal template.
[Tools for generating](https://github.com/ekalinin/github-markdown-toc) a table of contents from markdown are available.

- [Title](#title)
  - [Table of Contents](#table-of-contents)
  - [Glossary](#glossary)
  - [Summary](#summary)
  - [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals/Future Work](#non-goalsfuture-work)
  - [Proposal](#proposal)
    - [User Stories](#user-stories)
      - [Story 1](#story-1)
      - [Story 2](#story-2)
    - [Requirements (Optional)](#requirements-optional)
      - [Functional Requirements](#functional-requirements)
        - [FR1](#fr1)
        - [FR2](#fr2)
      - [Non-Functional Requirements](#non-functional-requirements)
        - [NFR1](#nfr1)
        - [NFR2](#nfr2)
    - [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    - [Risks and Mitigations](#risks-and-mitigations)
  - [Alternatives](#alternatives)
  - [Upgrade Strategy](#upgrade-strategy)
  - [Additional Details](#additional-details)
    - [Test Plan [optional]](#test-plan-optional)
  - [Implementation History](#implementation-history)

## Glossary

Refer to the [OpenYurt Glossary](docs/proposals/00_openyurt-glossary.md).

If this proposal adds new terms, or defines some, make the changes to the glossary when in PR stage.

## Summary

The `Summary` section is incredibly important for producing high quality user-focused documentation such as release notes or a development roadmap.
It should be possible to collect this information before implementation begins in order to avoid requiring implementors to split their attention between writing release notes and implementing the feature itself.

A good summary is probably at least a paragraph in length.

## Motivation

This section is for explicitly listing the motivation, goals and non-goals of this proposal.

- Describe why the change is important and the benefits to users.
- The motivation section can optionally provide links to [experience reports](https://github.com/golang/go/wiki/ExperienceReports)
to demonstrate the interest in a proposal within the wider Kubernetes community.

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

This is where we get down to the nitty gritty of what the proposal actually is.

- What is the plan for implementing this feature?
- What data model changes, additions, or removals are required?
- Provide a scenario, or example.
- Use diagrams to communicate concepts, flows of execution, and states.

[PlantUML](http://plantuml.com) is the preferred tool to generate diagrams,
place your `.plantuml` files under `images/` and run `make diagrams` from the docs folder.

### User Stories

1. As a vendor, I would like to connect a category of device into the OpenYurt
2. As an end user, 

### Requirements (Optional)

Some authors may wish to use requirements in addition to user stories.
Technical requirements should be derived from user stories, and provide a trace from
use case to design, implementation and test case. Requirements can be prioritised
using the MoSCoW (MUST, SHOULD, COULD, WON'T) criteria.

The difference between goals and requirements is that between an executive summary
and the body of a document. Each requirement should be in support of a goal,
but narrowly scoped in a way that is verifiable or ideally - testable.

#### Functional Requirements

Functional requirements are the properties that this design should include.

##### FR1

##### FR2

#### Non-Functional Requirements

Non-functional requirements are user expectations of the solution. Include
considerations for performance, reliability and security.

##### NFR1

##### NFR2

### Implementation Details/Notes/Constraints

- What are some important details that didn't come across above.
- What are the caveats to the implementation?
- Go in to as much detail as necessary here.
- Talk about core concepts and how they releate.

### Risks and Mitigations

- What are the risks of this proposal and how do we mitigate? Think broadly.
- How will UX be reviewed and by whom?
- How will security be reviewed and by whom?
- Consider including folks that also work outside the SIG or subproject.

## Alternatives

The `Alternatives` section is used to highlight and record other possible approaches to delivering the value proposed by a proposal.

## Upgrade Strategy

If applicable, how will the component be upgraded? Make sure this is in the test plan.

Consider the following in developing an upgrade strategy for this enhancement:
- What changes (in invocations, configurations, API use, etc.) is an existing cluster required to make on upgrade in order to keep previous behavior?
- What changes (in invocations, configurations, API use, etc.) is an existing cluster required to make on upgrade in order to make use of the enhancement?

## Additional Details

### Test Plan [optional]

## Implementation History

- [ ] MM/DD/YYYY: Proposed idea in an issue or [community meeting]
- [ ] MM/DD/YYYY: Compile a Google Doc following the CAEP template (link here)
- [ ] MM/DD/YYYY: First round of feedback from community
- [ ] MM/DD/YYYY: Present proposal at a [community meeting]
- [ ] MM/DD/YYYY: Open proposal PR

