---
{
  "title": "Webconverger networking",
  "permalink": "/networking/"
}
---


::: warning
Webconverger does not support static/fixed IP addresses.
:::


Ideally Webconverger **should** be deployed in a **wired network** with a
**DHCP server**, instead of a potentially unstable [wireless](/wireless/) connection.

Webconverger's default  [network/interfaces](https://github.com/Webconverger/webc/blob/master/etc/network/interfaces)

* Allows hotplug of wired interfaces including [USB dongles](/blog/entry/usb_tethering/)
* Configures a firewall with this [iptables configuration](https://github.com/Webconverger/Debian-Live-config/blob/master/webconverger/config/includes.chroot/etc/iptables.conf)

Please also see [wireless](/wireless/) and [testing](/testing/).

### Overriding DNS via [API](/api/)

Preferably you set this in your DHCPD, however you can override DNS for example for
[filtering](/filtering/) like so:

	dns=8.8.8.8,8.8.4.4

This is implemented by [resolvconf script](https://github.com/Webconverger/webc/blob/master/etc/webc/network-up.d/resolvconf)

### Static IP networking via [boot](/boot/) command line on the Live version

We defer to using
[live-boot](http://manpages.debian.net/cgi-bin/man.cgi?query=live-boot) in this
case. This is unsupported on the Install version currently.

	boot=live ip=10.0.0.1::10.0.0.254:255.255.255.0::

Notice `skipconfig` is removed and the `ip=` API is a little tricky(!)
