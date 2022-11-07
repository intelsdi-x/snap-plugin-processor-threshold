DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


[![Build Status](https://travis-ci.org/intelsdi-x/snap-plugin-processor-threshold.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-processor-threshold)

# Snap processor plugin - threshold
This plugin filtering metrics with task defined threshold

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.6+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download threshold processor plugin binary:
You can get the pre-built binaries for your OS and architecture under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-processor-threshold/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-processor-threshold

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-processor-threshold.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

### Examples
Examplary task manifest with threshold set up for core 0

```

---
  version: 1
  schedule:
    type: "simple"
    interval: "5s"
  max-failures: 10
  workflow:
    collect:
      metrics:
        /intel/procfs/cpu/*/utilization_percentage: {}
      process:
        - plugin_name: "threshold"
          config:
            /intel/procfs/cpu/0/utilization_percentage : 95
          publish:
            - plugin_name: "file"
              config:
                file: "/tmp/cpu_threshold.log"
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-processor-threshold/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-processor-threshold/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@Patryk Matyjasek](https://github.com/PatrykMatyjasek/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
