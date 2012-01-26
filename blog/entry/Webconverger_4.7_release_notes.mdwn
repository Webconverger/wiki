# New features and fixes

<a href="http://www.flickr.com/photos/hendry/3476173455/" title="Adobe Acroread 8 on Webconverger 4.7 by Kai Hendry, on Flickr"><img src="http://farm4.static.flickr.com/3558/3476173455_00755c9052.jpg" width="500" height="375" alt="Adobe Acroread 8 on Webconverger 4.7" /></a>

* Added [iptables](http://git.debian.org/?p=debian-live/config-webc.git;a=blob;f=webconverger/config/includes.chroot/etc/iptables.conf;) firewall
* [Fix UA string for Hotmail](http://groups.google.com/group/webc-users/browse_thread/thread/fc8171407b40a733)
* [file:/// disabled](http://groups.google.com/group/webc-users/browse_thread/thread/60d62ae72fcff707) in the latest [[kiosk]] extension. Try it out from [here](http://webconverger.com/xpis/).
* Remove previous [[wireless]] _default_ to join any open network.
* [Iceweasel 3.0.9](http://www.mozilla.com/en-US/firefox/3.0.9/releasenotes/)
* CUPS [[printing]] support re-instated with a Firewall to rule to allow for printer sharing broadcasts
* 2.6.29 kernel backport, which means even better hardware support
* Xpdf dropped in favour of a working printing dialog with [[Acroread|adobe]] -- there are some embedding issues when you first run it, [Ctrl+q] is needed to close the PDF viewer. Also saving allows file browsing.

# Known issues

<a href="http://www.flickr.com/photos/hendry/3476862244/" title="Wrong default locale dictionary by Kai Hendry, on Flickr"><img src="http://farm4.static.flickr.com/3303/3476862244_99848f4d65.jpg" width="500" height="375" alt="Wrong default locale dictionary" /></a>

* Spelling language has to be manually chosen and doesn't respect chosen [[locale|i18n]] [[boot]] options
* The **debian installer** which only worked on the [[ISO]] version is **temporarily disabled** until the debian installer uses 2.6.29 which supports [squashfs 4](http://en.wikipedia.org/wiki/Squashfs).
* USB version seems to trigger _Unknown video mode_ at least one of our test systems
* [CJK support disabled](http://lists.debian.org/debian-live/2009/04/msg00151.html) until new code proven and accepted in [live-initscripts](http://git.debian.net/?p=debian-live/live-initscripts.git)

[Available first on the German mirror](http://download.webconverger.org/). Please [comment on webc-users](http://groups.google.com/group/webc-users/browse_thread/thread/8b4a477db0a4fe04), thank you.
