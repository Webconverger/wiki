<img src="/img/webc-26.png" alt="Revised boot menu">

Highlights of this 26.0 signed & tagged snapshot:

* Revised boot menu; helping you get started with [Neon](http://neon.webconverger.com/), our [Web signage](http://www.linux.com/community/blogs/130-distributions/782621-what-is-web-signage) product
* [Firefox 32](https://www.mozilla.org/en-US/firefox/32.0/releasenotes/)
* Basic proxy authentication. A customer wanted this to fit into their complex Windows deployment, so now you have it too.
* Tab right click menu removed to make user interface simpler
* Bug fixes to the print button and the job scheduler API (cron=)
* The usual stable security updates and Adobe Flash, with an additional font to make [Flash video text render correctly](https://github.com/Webconverger/webc/issues/190)

Please ask your Web developers to switch to HTML video. If you host your video
on Youtube, you can make embeds use HTML5 with a `html5=1` argument.

[Detailed changelog between 25 & 26](https://github.com/webconverger/webc/compare/25.0...26.0).

The sha256sum is `f4374d183bbc4b897f423d482690ebf128f3f9b65484aee81e7aef180c8afe1e  webc-26.0.iso`, quickly download from our CDN <http://dl.webconverger.com/latest.iso>.

# Known issue and future work

There is a known issue where setting up external screens on new Intel hardware
might not be possible with the currently installed stable Intel driver. We have
the solution which we commercially support, an upgraded driver though we didn't
roll this into this release since it causes choppy Flash playback.

We are already working towards to smoothly upgrade to the next Debian release,
"jessie", which will have more current video drivers et al.

# Device management trials are now free of cost

Our commercial service now is now **free to trial**. No Paypal required. This
[screencast demonstrates the new signup flow](http://youtu.be/aAbIYAbXf5E). We
hope you will find our configuration management and support service worthy of
[supporting via payment](http://webconverger.com/pricing).

Thanks to the support of our customers; banks, cafes, governments, NGOs,
schools, universites, stores, small and large we look forward to supplying you
a stable, current and **open** Web client platform for running your Web
application upon long into the future. Please help spread the word!
