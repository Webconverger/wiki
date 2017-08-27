[[!meta title="Servers and services used by Webconverger"]]

Domain registration handled by Dreamhost.

DNS records hosted with [AWS Route 53](http://aws.amazon.com/route53/), this is mainly for DNS failover, though its performance is pretty damn good.

# webconverger.com

Hosted by AWS CloudFront.

Github Pages hosted test site is at <http://beta.webconverger.com/>.

# nl.webconverger.com

A bottom of the range [Digital Ocean droplet](https://www.digitalocean.com/?refcode=37b3b1850b32) with

* 512MB Ram 20GB SSD Disk Amsterdam 1 Arch Linux 2013.05 x64
* 1 TB pcm bandwidth (that's 3x more than most other VPS providers)
* 5USD a month
* Archlinux (systemd)

Uses:

* <https://config.webconverger.com>, synchronised by a Github Webhook, used as a [secondary failover](http://s.natalian.org/2014-04-21/1398064783_1364x748.png) using [AWS Route 53](https://console.aws.amazon.com/route53/home)
* [filter.webconverger.com service](http://webconverger.org/blog/2014/Fine_grained_filtering/)

# uk.webconverger.com (gb)

Hosted with [Linode](https://www.linode.com/?r=59198f21b93910b67d8e1e4212a3135343941aae)

* Debian 7 (jessie)
* Linode 2048
* London, England, UK
* 2676GB Quota
* 20USD pcm
* [Configuration service](https://config.webconverger.com)
* [Support service](http://support.webconverger.com/)
* [Ping/Statistics service](http://ping.webconverger.org)
* [Nagios](https://github.com/Webconverger/nagios)

# webconverger.org

<https://webconverger.org> is expertly curated by <http://www.branchable.com> for the public wiki.

~60USD a year

# [Webconverger Github](https://github.com/Webconverger)

On a bronze package at 25USD a month.

# SSL

Most sites are now using automagic Lets Encrypt or AWS Certificate Manager on CF.

* [config.webconverger.com](https://www.ssllabs.com/ssltest/analyze.html?d=config.webconverger.com), expires Sat Apr 29 18:57:32 UTC 2017
* <https://wss.webconverger.com> - bought with <https://sslmate.com/>:

# Email

Hosted with [FastMail.fm at 90USD a year](http://www.fastmail.com/?STKI=11542869) with 1 basic & 1 professional account.
