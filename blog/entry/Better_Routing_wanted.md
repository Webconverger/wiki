Many customers ask for Webconverger to have advanced networking features.

For example, a seemingly simple requirement:

	Either accept Verizon wireless connection (default) or wired broadband connection (preference would be for this to be automatic).

Wireless is actually really complex. Getting Wireless cards working in
Webconverger or Linux is a bit of a nightmare. Switching between different
Internet connections depending on their quality is also tricky.

So I usually beg customers to use something like a [Linksys
WRT54G](http://en.wikipedia.org/wiki/WRT54G) to act as a wireless
[[networking]] client.

Then some customers want more. :) A WRT54G isn't quite plug and play. You need
to (often flash it and then) configure it, which most customers aren't prepared to do. Nevermind the
difficult part of getting them to buy it and to ensure to get the more
expensive WRT54G-L version...

Then customers want Webconverger to be able to do [[filter|filtering]] or be
able to cache [[upgrades|upgrade]]. Implementing such features in Webconverger
would be a big detour from the initial [[design_goals|design]]. They are once
again, better fulfilled by a clever router.

However what clever router can do?

1. Advanced [[networking]] like being able to provide a Web interface for Wireless, 3G and other broadband options
* Squid caching proxy for making [[upgrading|upgrade]] less painful
* [[Filtering]]

The [dd-wrt](http://www.dd-wrt.com) firmware shows the most promise. It has
support for [keyword
filtering](http://www.dd-wrt.com/wiki/index.php/Access_Restrictions), though
not a **caching** proxy (it needs a much bigger disk!). dd-wrt does support quite a [lot of
hardware](http://www.dd-wrt.com/wiki/index.php/Supported_Devices). Though can
it be bought off the shelf from Amazon, Ebuyer or Dabs? Not yet. :(
