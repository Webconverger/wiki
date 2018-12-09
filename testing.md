---
{
    "title": "Howto test Webconverger",
    "permalink": "/testing/"
}
---

<iframe width="560" height="315" src="//www.youtube.com/embed/vpAJybQrbjc?list=UUH9ZMgfESFakRPvm6QRFNMQ" frameborder="0" allowfullscreen></iframe>

Please help test the very latest [daily build
images](https://build.webconverger.com). Report bugs on
[webc-users](http://groups.google.com/group/webc-users/topics) or on
[github](https://github.com/Webconverger/Debian-Live-config/issues).

# [Virtualbox](https://www.virtualbox.org) >= 4.2

<a href="http://www.flickr.com/photos/hendry/6509326697/" title="Webconverger Virtualbox by Kai Hendry, on Flickr"><img src="http://farm8.staticflickr.com/7166/6509326697_acce79b28f.jpg" width="500" height="281" alt="Webconverger Virtualbox"></a>

**NOTE**: You need to <a href="/img/pae.png">enable PAE</a> and if your VM fails to boot try give it more RAM

**NOTE**: If Webconverger fails to boot, make sure your BIOS has Virtualization options turned on as noted in [blog/2016/Webconverger 35 release](/blog/2016/webconverger_35_release/).

* [Archlinux Virtualbox setup guide](https://wiki.archlinux.org/index.php/VirtualBox)

Setting a bridged network interface is straightforward, you just need to ensure:

	sudo modprobe -a vboxnetadp vboxnetflt

## Testing branches from the install version

With `fetch-revision=` from [upgrade](/upgrade/) we can test experimental [branches](https://github.com/Webconverger/webc/branches), before they are merged to master.

Here is a [video explaining how to get an installed Webconverger to switch branches for testing](http://r2d2.webconverger.org/2012-11-05/branches.html).

## Testing network failure

<a href="http://www.flickr.com/photos/hendry/7013543725/" title="Right click Network Adaptors by Kai Hendry, on Flickr"><img src="http://farm8.staticflickr.com/7278/7013543725_109efd128c_m.jpg" width="240" height="240" alt="Right click Network Adaptors"></a>

<a href="http://www.flickr.com/photos/hendry/7013543723/" title="Cable connected togglebox by Kai Hendry, on Flickr"><img src="http://farm8.staticflickr.com/7053/7013543723_1f0be8a9f7_m.jpg" width="240" height="166" alt="Cable connected togglebox"></a>

Network use cases:

1. Boot with network - should work normally
* Boot with network and later become disconnected - [should retry](https://github.com/Webconverger/iceweasel-webconverger/blob/master/content/netError.xhtml#L415)
* Boot without network - should show no-net.png
* Boot without network and become connected - should come back online
* Boot connected to router, but router has no internet connectivity - should show no-net.png
* Boot connected to router, with router regaining internet connectivity after 2 minutes

Think of cases where network comes online but configuration is failed to be fetched.

"debug mode" will bypass no-net.png

# Important elements to test before delivery

* SSL <https://config.webconverger.com/>
* Test DDG search engine works

## [wireless](/wireless/) can only be realistically tested on physical hardware (i.e. not under virtualisation)

## [YouTube](https://www.youtube.com/)

* [Adobe](/adobe/) flash test <http://www.adobe.com/software/flash/about/>
* Sound
* [PDF](http://hendry.iki.fi/cv/resume.pdf)
