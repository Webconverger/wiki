---
{
    "title": "Developing Webconverger",
    "permalink": "/develop/"
}
---

* [Report and view bugs](https://github.com/Webconverger/webc/issues)
* [Core concepts](/design/)
* [Webconverger Debian Live configuration](https://github.com/Webconverger/Debian-Live-config)
* [Webconverger daily builds](https://build.webconverger.com/)
* [Webconverger in your locale](/i18n/)
* [Artwork](/artwork/)
* [Firefox extension development](/kiosk/)
* [Debian Live project](http://live.debian.net/) and [Live helper packages](http://live.debian.net/)
* [Howto add Java](/java/)
* [Howto integrate nvidia drivers](/nvidia/)
* [Customising cronjob](/cronjob/)
* [Debug mode](/debug/)
* [Using PXEBOOT for deployment](/pxeboot/)

If you are new to the project, **Welcome**!! Please start out by [testing](/testing/).

We might be the first distribution to use [git](http://git-scm.com/) for
[upgrades](/upgrade/) and managing just about everything we do. :)

<div class="notebox"> The Webconverger distribution is an MIT licensed
production (notable exceptions are non-free wireless firmware and
[flash](/adobe/)). You <b>must</b> understand that the <a
href="https://webconverger.com/images/logo.svg">branding (logo and the name
"Webconverger")</a> are <b>proprietary</b>.  Only Webconverger built images
should display the logo and can be sold as "Webconverger". Thank you.</div>

## Build with a Docker image

Follow the [README](https://github.com/Webconverger/Debian-Live-config) for
tips how to build and run the
[Dockerfile](https://github.com/Webconverger/Debian-Live-config/blob/master/Dockerfile),
which we use for building Webconverger in the Debian environment.

**WEBC_CHECKOUT** environment variable should be the clone of
<https://github.com/Webconverger/webc>. Though you could try get away with a
shallow clone (~2GB versus ~6GB for full) to save bandwidth and disk space,
e.g. `git clone --depth 1 https://github.com/Webconverger/webc.git` but it's
not guaranteed to work!

For tweaks, take a look at
[auto/config](https://github.com/Webconverger/Debian-Live-config/blob/master/webconverger/auto/config).

`make` is run from within the Docker container to build the ISO image, which is
named `live-image-i386.hybrid.iso` by default.

Next use this [testing guide](/testing/) on your freshly built image.

To work within the chroot, you shouldn't require the Docker image. `chroot` should be sufficient:

	alias c="sudo chroot $WEBC_CHECKOUT env -i PATH=/bin:/usr/bin:/sbin:/usr/sbin /bin/bash"

Typical work as you can see from the
[commits](https://github.com/Webconverger/webc/commits/master) is keeping the
distro upto date with the latest packages.

Please share your forks on github! :) Contributions must be [MIT
licensed](http://en.wikipedia.org/wiki/MIT_License). For patches to be
accepted, you must agree to transfer copyright to Webconverger (the company) to
make things manageable. We do [acknowledge](/acknowledgements/) contributions!

## Tips for getting your patch out a running Webconverger

In [debug](/debug/) mode:

	root@webconverger:/etc/systemd/system# git status | less
	root@webconverger:/etc/systemd/system# git commit
	[systemd fa18ba0] Upgrade service
	 create mode 120000 etc/systemd/system/multi-user.target.wants/upgrade.service
	 create mode 120000 etc/systemd/system/upgrade.service
	 create mode 100644 lib/systemd/system/upgrade.service
	root@webconverger:~# git format-patch HEAD~1
	0001-Upgrade-service.patch
	root@webconverger:~# cat 0001-Upgrade-service.patch | p
	http://ix.io/41O

And provide the link <http://ix.io/41O> in a bug report! :)

## If your build doesn't work

1. Check with the latest [daily build](http://build.webconverger.org/) logs. Compare them.
* [Submit a bug report](https://github.com/Webconverger/Debian-Live-config/issues)
