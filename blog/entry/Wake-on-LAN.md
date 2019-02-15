---
{
       "title": "Wake on LAN",
       "permalink": "/blog/entry/Wake-on-LAN/"
}
---

Background: [Wake on Lan
support](https://groups.google.com/forum/#!topic/webc-users/7VvFvf8yq2k/discussion)
was asked for sometime ago and it almost delayed Webconverger 13.0 until it was
moved to target 14.0 and it's one of those trivial new features that has cost
at least 3 days worth of man hours of development time.

Why isn't wake-on-lan, simply not just the default?

I think it is the default on some machines but not others. Webconverger
shipping
[ethtool](https://github.com/Webconverger/webc/blob/d678eed4e9b2024b41eb1e28a5fdb7d034cec252/etc/network/interfaces#L9)
seems like just a **waste of space**, to configure silly machines to flip that
bit.

For example after installing ethtool for the first time on a Lenovo Thinkpad
X220 and I get on a `sudo ethtool eth0` invocation:

	Supports Wake-on: pumbg
	Wake-on: g

So we are fine. So by default my X201 was "Wake-on-Lan" _ready_, if there is such terminology.

# BIOS override

I noticed on a machine that refused to wake up, that it was in fact disabled in
the BIOS. DOH. Is this reflected on the Linux user space tool `ethtool`? NO.
Very annoying.

# Shutdown now not working

I noticed that once enabling Wake-On-LAN in the BIOS, a certain Intel based
machine refuses to cleanly shutdown for whatever reason. :(

# When it does work, it can do strange things

I managed to Wake-On-LAN my X201 and then it tried to netboot for about a minute before booting. Why ....

WHY OVERRIDE DEFAULT BOOT OPTIONS?

# Which port ?!?

Oh btw, there seems to be a million ways
[Wake-on-LAN](http://en.wikipedia.org/wiki/Wake_on_lan) is
implemented/triggered. I've chosen the sanest [Magic
Packet](http://en.wikipedia.org/wiki/Magic_packet#Magic_packet) flavour which
needs a mac address in order to work.

Is it port 40000 or port 9? I can't find a definitive answer for [Magic
Packet](http://en.wikipedia.org/wiki/Wake-on-LAN) since the normal `wol`
utility goes on port 40000 by default for whatever reason.

Update: Port doesn't matter, the hardware just looks for _any_ packet with the magic sequence

# And security?

Tbh in Enterprise network level security is usually well done, so I really
don't see the need for passwords to be set. Especially the way most client
connections are behind a NAT nowadays. Daft.

# Horrible usability of `wol`

Say you are testing from your laptop, sharing internet on eth0 from 192.168.0.1
You need to do an invocation like:

	wol -v -p 9 00:04:5f:04:bf:dd -h 192.168.0.255

Not many people know "255" is a broadcast address. _Sigh_

# Is Wake-on-LAN the right way to manage machines power usage?

In "Enterprise" environments having timed shutdowns and ways to wake the
machines in the morning seems like the "old" approach for Energy saving. Timed
shutdowns for example are obviously **NOT** very desirable if you actually working on a
Webconverger machine after hours.

For example Apple Mac machines aggressively save power at all times unless for example
you are using the machines or prefix commands with
[caffeinate](http://opensource.apple.com/source/PowerManagement/PowerManagement-271.1/caffeinate/caffeinate.c).

This "sleep always until actually used" approach seems better for _real_ energy
savings for kiosks.  Nonetheless lots of Webconverger customers want `noblank`
option so that their kiosk users see and know the machines are ready for use.
Just imagine walking up to a blank screened library kiosk. Would you know that
pressing the keyboard would wake it up? Not everyone would know this without a
sign...

For [Web signage](http://neon.webconverger.com/), a different use case, a
different approach needs to be taken again to keep the sign visible at all
times. In our deployment we dim the screens from midnight to morning and the
machine is **never halted** and currently this is controlled on a hardware
level.
