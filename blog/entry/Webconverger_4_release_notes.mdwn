This new major 4.x series release of Webconverger is marked by initial
[[wireless]] driver support. Webconverger also has a new super fast [download
mirror](http://download.webconverger.org/mini) hosted by a new machine
[[hetty]] located somewhere in Germany. Try it out!

Historically [[blog/entry/Webconverger_3]] before celebrated Firefox 3.

# Webconverger 4 release notes

The Debian installer does not contain wireless firmware support. So **if** you
want to install Webconverger to your hard drive from CD ISO
[[grub_boot_menu|boot]] or [[USB|blog/entry/Webconverger_install_from_USB]],
please do so from a [[wired_connection|networking]].

There is a [video describing Webconverger USB wireless boot
option](http://uk.youtube.com/watch?v=CkryMFlBotg). Please post [[testing]]
feedback publicly on the
[forums](http://groups.google.com/group/webc-users/topics).

Last, but not least. For production deployments you **must** use [WPA
security](http://en.wikipedia.org/wiki/Wi-Fi_Protected_Access) on your Wireless
access point otherwise attackers can too easily sniff the sites you are
surfing. Webconverger is designed to be [[secure|security]]!

## UPDATE FLASH does not seem to work with 4.0

4.1 fixes the problem. Enjoy :-)

## About non-English [[locales|i18n]]

[[Maxi]] will be updated when people are happy with [[wifi_support|wireless]].
I am tinkering with the idea of setting up locale dynamically from a control
panel once Webconverger mini has booted up. This would download the necessary
locale files instead of distributing a huge distribution such as [[maxi]] with
everything bundled in.
