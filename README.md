![Easee Image](https://d2csxpduxe849s.cloudfront.net/media/A2D1552F-F177-4B01-9E863028C1D46C3F/71D770BF-9E4D-4579-8C70BD1E4DDEC6DC/webimage-577038BF-2EB9-43FA-B2921FF789B05FAA.png)

# Easee-CLI

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

A CLI written in go, for controlling a Easee charger using the [Easee.cloud](https://easee.cloud) Api

## Getting Started <a name = "getting_started"></a>

Get your credentials for [Easee.cloud](https://easee.cloud)
and run

```
easee-cli config
```

to add them to the applications config file.

The config file will be places in `$home/.easee-cli/config.yaml`

If you have problems finding your Charger's Id, then run

```
easee-cli config getchargerid
```

To get the id of your charger.
Note we currently only support one charger.

### Prerequisites

To be able to run any commands, you must have a valid token.
Obtain one, that works for 24 hours, by running

```
easee-cli login
```

The token will be saved in the config file, and used by all commands you run.

### Installing

TBD

## Usage <a name = "usage"></a>

Run

```
easee-cli -h 
```

to see the commands available.
