<iframe width="560" height="315" src="https://www.youtube.com/embed/FaBFbkBhnXI" frameborder="0" allowfullscreen></iframe>

Webconverger was started because I couldn't trust those terminals and kiosks in
public spaces to browse the Web.

I trust my work because I had a hand in building it, though there are some
technical features which may convince you to trust it too.

First every commit is made publicly on
[Github](https://github.com/Webconverger/webc/commits/master). You can see me
transparently make changes. You can verify every changed file. Every upstream
file is a binary built by Debian, Mozilla & for flash, Adobe, again you can
verify that using file checksums.

However how can you infer that the head of Webconverger git repository
corresponds to an [[ISO]] release you downloaded which you would use to deploy
to your hardware?

Previously you would have to trust the checksum I provided on the release page,
but if you built Webconverger yourself, you would get a different checksum. A
different checksum should make you suspicious that something nefarious was
inserted in the build pipeline.

But why were the checksums different if you [[built this ISO
yourself|develop]]? Because the build chain would typically use the current
date and when bundled up, the checksum is different. [Chris
Lamb](https://chris-lamb.co.uk/) has fixed this and now you can too can produce
independently verifiable ISO builds of Webconverger. This [[security]] feature
is called [reproducible builds](https://reproducible-builds.org/).

I must admit we owe the motivation to achieve this was due to [Mozilla
sponsoring
Tails](https://blog.mozilla.org/blog/2016/06/22/mozilla-awards-385000-to-open-source-projects-as-part-of-moss-mission-partners-program/)
to do so. Webconverger achieved this feature before Tails, at a lower cost,
funded by [paying Webconverger subscribers](https://webconverger.com/pricing/).
If you like our work, please tell your {school,office,company} IT department
that this is the very best Web kiosk software.

The build system has been improved to use a third party continuous integration
process at <https://travis-ci.org/Webconverger/webc>. The finished builds are
uploaded to <https://build.webconverger.com/>. Since everything now is
automated, much like the [[automatic updates|upgrade]] do once Webconverger is
installed on read/write medium, there is no need for releases anymore. Hence
the x86 download button on the [Webconverger
homepage](https://webconverger.com/) is set to
<https://build.webconverger.com/latest.iso>.
