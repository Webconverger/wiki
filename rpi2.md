---
{
    "title": "Webconverger for Raspberry PI 2 & 3",
    "permalink": "/rpi2/"
}
---

<div class="notebox">We also support the <a href=https://twitter.com/webconverger/status/709405833709297664>Rpi3</a> with our Rpi2 binaries!</div>

# Webconverger for Raspberry PI 2/3 port

A curated
[auto-updating](https://github.com/Webconverger/rpi2/blob/master/usr/bin/webc-upgrade)
Archlinux Arm (alarm) rootfs <https://github.com/webconverger/rpi2> using
[webkit2gtk](http://webkitgtk.org/) based surf2 and a new [Websocket based
updater](https://github.com/Webconverger/wsc) allowing for instant updates to
the configuration and better visibility if the device is actually online.

Currently this port is designed for <abbr title="Digital Out of
Home">DOOH</abbr> / Signage applications. It can be used as a kiosk, though it
will not have a "chrome", i.e. [user interface options](/kiosk/) is set to `chrome=webcfullscreen`.

* [Installation instructions](https://github.com/Webconverger/rpi2/blob/master/README.md)
* [Video of preparing the Webconverger microSD from Windows 10](https://www.youtube.com/watch?v=X4B0emkwvaw)

A >= 4GB microSD card is recommended.

# [API](/api/) notes

<iframe width="560" height="315" src="https://www.youtube.com/embed/K54BDqJcYsA?rel=0" frameborder="0" allowfullscreen></iframe>

Only a limited subset of the [API](/api/) is implemented. The
[webc-config](https://github.com/Webconverger/rpi2/blob/master/usr/bin/webc-config)
is the best documentation.

`instantupdate` is enabled by default, see this [video demo with two
rpi2s](https://www.youtube.com/watch?v=SgaqofU1zgc). This allows instant
updates to the homepage without rebooting.

You can override the browser binary `surf2` with the `webck` variable. For
example `webck=omxplayer%20-r` with
`homepage=http://techslides.com/demos/sample-videos/small.mp4` makes
Webconverger loop and play back an MP4 file.

You can power down the rpi2 with `poweroff` and reboot with `reboot`. This is
actually good practice since unfortunately the rpi2 boots from a [vfat
partition](http://en.wikipedia.org/wiki/File_Allocation_Table) that can easily
get corrupted.

To rotate the screen, you can use
`rpi2config=http://prefs.webconverger.com/2015/90-degrees.txt`.
[xrandr](/display/) commands do not work with the rpi2's framebuffer powered
display. You may need to patient rebooting the device a couple of times as this
configuration is only embedded at boot via
[webc-upgrade](https://github.com/Webconverger/rpi2/blob/master/usr/bin/webc-upgrade).

Begin with the [canonical
config.txt](https://raw.githubusercontent.com/Webconverger/rpi2/master/boot/config.txt)
when making your changes. Other examples:

* `rpi2config=http://prefs.webconverger.com/2015/disable-overscan.txt` if your screen isn't fullscreen

# Caveats

You must have your HDMI screen on as the rpi2 boots since it does some sort of
HDMI negotiation only at boot.

# For hardware tinkerers

[Video showing the possibilities of hardware motion detector integration](https://youtu.be/O0jdE7A-vbg?list=PLECEw2eFfW7hYMucZmsrryV_9nIc485P1&t=3721)

* [/etc/systemd/system/dpms.service](http://ix.io/iSq)
* [/root/dpms.sh](http://ix.io/iSp)

Useful for saving power when no one is in the vicinity.

# Wireless

Here is an example configuration I use for testing:

	boot_append=debug
	homepage=https://google.com
	wpa-ssid=SINGTEL-F726
	wpa-psk=0011466534
	instantupdate

This is naively implemented at the bottom of
<https://github.com/Webconverger/rpi2/blob/master/usr/bin/webc-config>. On rpi2
your WEBCID will stay on the eth0's mac address.

You need to have an initial wired internet setup to provision the wireless
settings.

# Debug mode

Triggered with `boot_append=debug`. In debug mode [sshd is
enabled](https://github.com/Webconverger/rpi2/blob/master/etc/systemd/system/sshd.service.d/override.conf)
and you can log in with a ssh public configured like so `ssh=http://hendry.iki.fi/hendry.pub`.

# Experimental shell API

Removing the API because of [blog/2015/Why a shell API would be insecure](/blog/2015/why_a_shell_api_would_be_insecure/).
