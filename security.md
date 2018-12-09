---
{
    "title": "Webconverger has better out of the box security",
    "permalink": "/security/"
}
---

Webconverger unlike Windows for example does not require:

* a Firewall
* anti-spyware
* anti-malware
* anti-virus

Webconverger is [designed](/design/) for [privacy](/privacy/) and [security](/security/), based on:

* Transparency with [opensource software](https://github.com/webconverger/webc)
* [Reproducble builds](/blog/2016/Webconverger_has_reproducible_builds)
* Always clears private data once a session is ended (last tab is closed or by inactivity timeout or by poweroff)
* Does not require any signup or registration
* Does not snoop on users

Once installed, Webconverger has [automatic updates](/upgrade/) that keep you
deployment security patched. Note that fetched updates only apply at boot.

We can support Enterprise IT administrators looking to finely tune update roll
outs.

# Peer review

Webconverger welcomes security review. Please email **security** AT
webconverger.com with your findings. Valid bugs are typically kindly
[acknowledged](/acknowledgements/) building up your online reputation as an expert!

Could security reviewers please practice [responsible disclosure](https://en.wikipedia.org/wiki/Responsible_disclosure).

* Webconverger is a member of the [Open Source Computer Emergency Response Team](https://ocert.org/team_and_members.html)
* [Review by Andrew Patrick, Ph.D. of the Information Security Group](http://www.andrewpatrick.ca/security-and-privacy/quest-for-a-good-boot-cd-for-internet-banking/) of the [National Research Council Canada](http://www.nrc-cnrc.gc.ca/)
* Webconverger has been reviewed by "kiosk hacker" Paul Craig of <http://ikat.ha.cked.net/>

Please be vigilant: [Packet sniff Webconverger](/wireshark/) and [monitor changes](https://github.com/Webconverger/webc/commits/master).
