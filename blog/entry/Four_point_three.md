4.3 release notably fixes the  `xvideomode=widthxheight` boot parameter. So in
the cases where Webconverger incorrectly detected your for example 320x200
display you can append:

	xvideomode=320x200

<a href="http://www.flickr.com/photos/hendry/3244661865/" title="xvideomode=320x200 by Kai Hendry, on Flickr"><img src="http://farm4.static.flickr.com/3107/3244661865_6af1dc48ea_o.png" width="320" height="200" alt="xvideomode=320x200" /></a>

And as promised [[last_week|blog/entry/Web_Crunch/]] the [Compose key is bound
to the Right-Alt
key](http://git.webconverger.org/?p=webconverger-base.git;a=blob;f=home/webc/.Xmodmap)
(113). This feature is of especial interest to French users who care about
accented characters. Superb.

Developers, builders and integrators might be interested to note some improved
[[networking_documentation|networking]] and the fact that I have migrated the
[Webconverger configs](http://git.debian.org/?p=debian-live/config-webc.git) to
take advantage of Daniel's new autoconfig script/ style in [Debian
Live](http://debian-live.alioth.debian.org/). If you noticed any problems,
please let me know.

# Maxi does not install

Maxi, the version of Webconverger with [[foreign_language|i18n]] support like
Japanese, Korean and Chinese is just a Live CD. You cannot install it to your
hard drive.

The way you switch locales with the boot parameter `locale=` does not work once
you've installed Webconverger, due to the fact live-initramfs is not part of
the bootup sequence on installations.

This needs to be fixed. The plan I have is to roughly have a Web interface to
change locales and maintain an installed copy of Webconverger on the hard
drive. This Web interface needs to be protected by a password ideally,
otherwise people could disrupt the kiosk, which is certainly not a [[design]]
goal of Webconverger. Hence this feature demands some thought still.

So you might notice that [[maxi]] and [[mini]] and roughly the same size now.
That's because mini has extra bits to install to the hard drive, whilst maxi
does not. In future [[maxi]] and the [[mini]] concepts will be dropped and
things will get simpler with just one version that can meet most people's kiosk
needs. :)

# Other notes and issues

Pushing the upload now. The source tar ball will come during the week. It's quite large. [Download now from Germany](http://download.webconverger.org/).

The USB images seem to have the wrong bootloader graphic. This is probably due
to the config/ changes. I am not going through the expense of regenerating the
USB images, until the next release.

Also the video mode could be wrong at least on my machine.
