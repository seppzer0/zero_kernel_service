# zero_kernel_service — an update tracking service for the zero_kernel

## Contents
- [zero\_kernel\_service — an update tracking service for the zero\_kernel](#zero_kernel_service--an-update-tracking-service-for-the-zero_kernel)
  - [Contents](#contents)
  - [zero\_kernel](#zero_kernel)
  - [Usage](#usage)
    - [Build](#build)
    - [Launch](#launch)

## zero_kernel

[zero_kernel](https://github.com/seppzer0/zero_kernel) is a project dedicated to adding numerous enhancements to an Android kernel.

The project has a number of different types of dependencies: ROM builds, kernel source, KernelSU distributions, RTL8812AU driver updates etc.

The purpose of this service is to track all of such dependencies and take the appropriate actions in an event of an update, whether it's triggering a pipeline or a sending a notification.

Currently, this service is <u>in development</u>.

## Usage

### Build

To build the service, use:

```sh
docker build backend/ -t zero_kernel_service
```

### Launch

To launch the service in an autonomous (==detached) state, use:

```sh
docker run --rm -d zero_kernel_service
```

Depending on your case and scenario, you can alternate used Docker flags in any way desired.
