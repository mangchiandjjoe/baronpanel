# Switch Panel controller for Beechcraft Baron 58
Beechcraft Baron 58P Switch panel controller based on Beaglebone Black


## Setup environment

### Enable all 64 GPIO ports
Most of GPIO ports are used for eMMC and display. We need to disable eMMC and display override to use whole 64 GPIO ports.

1. Boot beaglebone black using sd card to disable eMMC
2. Edit /boot/uEnv.txt 
3. Uncomment following two lines and save.
```
disable_uboot_overlay_emmc=1
disable_uboot_overlay_video=1
```
4. Reboot
5. Run show-pins tool to check all GPIO ports are enabled.

### [show-pins](https://github.com/mvduin/bbb-pin-utils)
This perl script gives a nicely formatted overview of the current pin configuration of your BeagleBone Black.
```
cd /usr/local/sbin
sudo wget -N https://raw.githubusercontent.com/mvduin/bbb-pin-utils/master/show-pins
sudo chmod a+x show-pins
```

## Build and run
To access GPIO ports, application requires sudo permission.
```
# go build baron.go

# ./baron
Port: 15 State: false
Port: 15 State: true
```
