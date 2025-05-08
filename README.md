# Homelab Information

## ProxMox

### ProxMox Reference

<https://proxmox.com/en/products/proxmox-virtual-environment/get-started>

### ProxMox Steps

For installing ProxMox on a bare metal server:

[ProxMox](https://proxmox.com/en/products/proxmox-virtual-environment/get-started)

The steps are easy:

1. Download the ProxMox installer onto a USB
2. Boot on the bare metal server
3. Follow the instructions that it provides

## Ubuntu Server

### Ubuntu Server Steps

For installing the Ubuntu server on ProxMox:

[Ubuntu Server Download](https://releases.ubuntu.com/24.04.2/ubuntu-24.04.2-live-server-amd64.iso)

> Ensure that you are installing the "headless" version of Ubuntu, which will
> be a version without a desktop

#### Steps

##### Create a New VM

1. Log in to the Proxmox web interface and navigate to the "Datacenter" view.
2. Click on "Create VM" to start the process of creating a new virtual machine.
3. You will need to specify the VM's name, ID, and other basic settings.

##### Configure VM Settings

1. Set the CPU and memory resources for the VM.
2. Ensure that the CPU and memory allocation is sufficient for your needs.
3. For example, you might allocate 2 CPUs and 4GB of RAM.

##### Add Disks

1. Add one or more disks to the VM. This can be done by specifying the size of
   the virtual disk and choosing the appropriate storage backend (e.g., local
   storage, ZFS, or LVM).

##### Network Configuration

1. Configure the network settings for the VM.
2. You can choose to use a bridged network or a virtual network, depending on
   your requirements.

##### Install Ubuntu

1. Once the VM is created, start it and boot from an Ubuntu ISO image.
2. You can download the ISO from the official Ubuntu website and add it as an
   installation media in Proxmox.
   - You can also use the link above
3. During the installation process, follow the on-screen instructions to
   install Ubuntu on the newly created VM.

##### Post-Installation Configuration

1. After the installation is complete, you can log in to the Ubuntu VM and
   perform any necessary post-installation configurations, such as setting up
   networking, installing updates, and configuring services.

### Installing Dependencies

To install some dependencies, use the following commands:

#### Update

> ==In Linux, you should **ALWAYS** update before installing packages!==

In order to do so, run the following command:

```bash
sudo apt update -y
```

> Note that the `-y` portion is optional, and included to skip manually
> accepting installation

#### Java

To install the default Java Runtime Environment, run the following command:

```bash
sudo apt install default-jre-headless -y
```

In order to test and ensure that the correct Java version was installed, run the following command:

```bash
java --version
```

##### Java Example Output

```bash
openjdk version "17.0.9" 2023-10-17
OpenJDK Runtime Environment (build 17.0.9+9-Debian-1deb12u1)
OpenJDK 64-Bit Server VM (build 17.0.9+9-Debian-1deb12u1, mixed mode, sharing)
```
