<img src="/img/webc-28.png" alt="28 booted up" width=640 height=480>

2015 has been a busy year. New [[Android 5|Web_kiosk_and_signage_for_Android]]
and [[Raspberry PI2|rpi2]] ports give customers cheaper hardware deployment
options. All these different device architectures managed consistently via
<https://config.webconverger.com/>.

The PC x86 version however is our mainstay and this is our last release based
on [Debian Wheezy](https://www.debian.org/releases/wheezy/).  Debian's
**stable** platform lived up to its name. It has been rock solid.  Thank you to
the Debian project, its developers and users. The next version "jessie" should
run better on new Intel hardware.

28.0's most exciting **new feature** is _opted in_ with the `instantupdate`
[[API]] option which allows you to change the homepage remotely for one or many
of your devices in an instant. **No reboot required!** See the [instantupdate
demo](https://www.youtube.com/watch?v=VJehjCZg02k). The technology behind this
is [Websockets](http://en.wikipedia.org/wiki/WebSocket) and the service runs
upon <https://wss.webconverger.com/>.

The [detailed changelog between 27.1 and
28.0](https://github.com/webconverger/webc/compare/27.1...28.0) otherwise shows
a raft of security updates and minor bug fixes. If you are Live user, you
really must update for the [Flash fixes
alone](https://twitter.com/darkuncle/status/590240595563585536). For install
users... you should have nothing to worry about. Do double check the version on
`about:` in case something is amiss.. it's never a bad idea to re-install (it only
takes a couple of minutes)!

Coming soon: A **major update** where we will upgrade from Debian codename
wheezy to [jessie](https://www.debian.org/releases/jessie/) smoothly. Work in
our [testing branch](https://github.com/Webconverger/webc/tree/testing) is well
underway and we hope to release the next version of "stable" soon after Debian
makes the switch.

The sha256sum is `ce422a67f18b8fbf97303bf421c0010f7d8af87b0c09fd78b9c281dca5cb9cbb  webc-28.0.iso`, please <a href=http://dl.webconverger.com/latest.iso><strong>download</strong> from our <abbr title="Content Delivery Network">CDN</abbr></a>, where we are playing with some <abbr title="Amazon Web Services">AWS</abbr> credits for this release.
