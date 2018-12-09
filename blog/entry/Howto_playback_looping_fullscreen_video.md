<iframe width="560" height="315" src="//www.youtube.com/embed/a05Uswl171A" frameborder="0" allowfullscreen></iframe>

Say you want to use Webconverger for **Web digital signage**, to playback a
promotional video in fullscreen in a loop. Such a configuration would look
like:

	homepage=http://signage.webconverger.com
	prefs=http://prefs.webconverger.com/2013/signage.js
	hidecursor
	noblank
	chrome=neon

Notice you need to tweak <http://prefs.webconverger.com/2013/signage.js>
"autoconfig" preferences for your particular homepage's domain in order to give
**fullscreen video permissions** for your domain.

	Components.utils.import("resource://gre/modules/Services.jsm");
	Components.utils.import("resource://gre/modules/NetUtil.jsm");
	Services.perms.add(NetUtil.newURI("http://YOUR_HOMEPAGE_DOMAIN.com/"), "fullscreen", 1);
	pref("full-screen-api.allow-trusted-requests-only", false);

As for view-source:http://signage.webconverger.com/ it's a very simple use of
[HTML
video](http://www.whatwg.org/specs/web-apps/current-work/multipage/the-video-element.html#the-video-element):

	<video id=signage autoplay loop src=demo.webm></video>
	<script>
	video = document.getElementById('signage');
	video.mozRequestFullScreen();
	</script>

You can also re-purpose this technique for a [Screensaver / Attraction loop for
Webconverger](/screensaver/).

"Is this <a href=http://neon.webconverger.com/>Webconverger Neon</a>?" you may
ask. Yes. Neon has converged into Webconverger in every functionality, except
the [realtime
monitoring](http://webconverger.org/blog/entry/Monitoring_kiosk_uptime/) and
the price.

Enjoy simple device management of the [control
panel](http://config.webconverger.com/), not just for your kiosks, but for your
PC signage devices too!
