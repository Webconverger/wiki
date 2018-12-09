<a href="http://git-scm.herokuapp.com/downloads/logos">
<img src="/img/2color-lightbg.png">
</a>

Previously we presented a good example of how [[downgrading]] is now manageable
with our [rootfs maintained in
git](https://github.com/Webconverger/webc/commits/master) innovation pioneered
with Webconverger.

## Git maintained rootfs would help debugging

Tracking down bugs with the powerful `git bisect` tool is now possible. Imagine
an elusive bug is introduced and you do not know what is causing it. You could
binary chop through previous revisions and now have a reasonably good chance to
find it.

In traditional package maintained distros, this is really hard to do. Perhaps
several dependencies were at play.
