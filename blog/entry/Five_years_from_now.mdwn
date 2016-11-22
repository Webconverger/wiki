<a href="http://www.flickr.com/photos/hendry/5625226358/" title="London Hackspace Webcam by Kai Hendry, on Flickr"><img src="http://farm6.static.flickr.com/5142/5625226358_bee1ffb8ac.jpg" width="500" height="313" alt="London Hackspace Webcam"></a>

The [[last 2 year key|Security_and_privacy]] (7EEFD8EC) expired before I should
have updated it, before it broke the [builds](https://build.webconverger.com/)
with "WARNING: The following packages cannot be authenticated!"

[Apologies](https://groups.google.com/d/topic/webc-users/rU1KkDqMK0g/discussion) for that. :/

Now we have a [new key](http://debian.webconverger.com/archive-key.asc) that expires in 5 years.

	:r !wget -q http://debian.webconverger.com/archive-key.asc -O /dev/stdout | gpg
	pub  2048R/AB9E8DA9 2011-04-16 Webconverger Archive Automatic Signing Key <key@webconverger.com>
	sub  2048R/B8E27CCA 2011-04-16 [expires: 2016-04-14]

The [webconverger-archive-keyring package has been
updated](http://git.webconverger.org/?p=webconverger-archive-keyring.git;) and
so has the live-helper
[configuration](http://git.debian.org/?p=debian-live/config-webc.git;a=commitdiff;h=e2a729e2add2a9a9f9b0ea982f0d473b3fa15063).

This message comes from the [London Hackspace](http://london.hackspace.org.uk/).
