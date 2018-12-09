---
{
    "title": "How to packet sniff Webconverger",
    "permalink": "/wireshark/"
}
---

Tools:

* [Virtualbox](https://www.virtualbox.org/) for running Webconverger on your host system as a guest
* [Wireshark](https://www.wireshark.org/) for packet sniffing

Tested in Windows. Unfortunately Virtualbox bridged networking under Linux tends to be broken.

# Webconverger settings

<img src=/img/2015/aboutblank.png alt="most network quiet settings">

	noping noconfig homepage=about:blank

* `noping` disables the [machine ping for statistics](http://ping.webconverger.org/) as documented in [privacy](/privacy/)
* `noconfig` disables the [configuration service](https://config.webconverger.com/)

# Virtualbox settings

Using **bridged networking** (it's NAT by default) and filter in Wireshark on the `eth.src` address of your guest.

<img src=/img/2015/wireshark.png alt="Packet sniffing under Windows">

Besides NTP (for accurate time) and the usual router pings and such, you should
see [no automatic
connections](https://support.mozilla.org/en-US/kb/how-stop-firefox-making-automatic-connections).
This actually takes a lot of time and effort to turn off Firefox's automatic
connections since they are so poorly documented and [new ones seem to be
constantly introduced](https://wiki.mozilla.org/Advocacy/heartbeat).

Related project: <https://github.com/pyllyukko/user.js>

Be sure to check out the [Webconverger packet sniffing
videos](https://www.youtube.com/watch?v=7Zmsj5DnQYo) for the [upgrade](/upgrade/)
scenario run through.
