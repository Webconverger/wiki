[[!meta title="Customising Webconverger with your branding"]]

**Boot artwork** is **not** customisable. Ideally users should never see Webconverger boot.

The background image on a white background (black for neon) seen between
browser restarts can be customised by:

	bgurl=https://webconverger.com/img/logo.png

Currently the background image URL is checked after every browser session, here
is the relevant code and search for "bgurl":
<https://github.com/Webconverger/webc/blob/master/home/webc/webc.sh>

# Turning off Webconverger branding / [throbber](https://en.wikipedia.org/wiki/Throbber)

<img src=/img/2014/sq-with-tooltip.png>

The Webconverger brand mark can be turned off with `nobrand` in your [configuration](https://config.webconverger.com/) or setting `pref("extensions.webconverger.nobrand", true);`, using the `prefs=` API.

Please see [[blog/2014/New_logo]] for more information on the new logo.

# Themes

From June 2016 we have introduced
[extensions.webconverger.themeURL](https://github.com/Webconverger/webconverger-addon/commit/c1a066f24729d902514577d3db4a6bfa59ba376d). There is a [video introducing the feature](https://youtu.be/49-JMDpvZ2w).

<https://developer.mozilla.org/en-US/Add-ons/Themes/Lightweight_themes> lists
all the options, though after testing only the following settings seem
applicable:

	{
		"id": "webc",
		"name": "webc",
		"headerURL": "http://s.natalian.org/2016-06-06/headereg.png",
		"textcolor": "red"
	}

The headerURL dimensions are 3000x200 pixels.
