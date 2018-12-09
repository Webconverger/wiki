---
{
    "title": "Webconverger chrome UI options",
    "permalink": "/kiosk/"
}
---

Webconverger has its own [Firefox kiosk
extension](https://addons.mozilla.org/en-US/firefox/addon/webconverger) to
tweak the **user interface** elements such as the URL bar and icons.

Github: <https://github.com/Webconverger/webconverger-addon/>

# Default

	chrome=webconverger

The goal with this default to have a **stable**, simple & familiar user
interface for using the Web.

<img src="/img/chrome/webconverger-showprintbutton.png" alt="webconverger" />

`showprintbutton` API is enabled on this screenshot. Normally there would be no print icon.

# No address bar

	chrome=webcnoaddressbar

Customization option, perfect for [filtered](/filtering/) setups. For customers
who want patrons to browse, but not enter in URLs to navigate away.

<img src="/img/chrome/webcnoaddressbar-showprintbutton.png" alt="webcnoaddressbar" />

`showprintbutton` API is enabled on this screenshot. Normally there would be no print icon.

# Fullscreen

	chrome=webcfullscreen

Customization option. For customers who want to do basic [Digital
signage](http://en.wikipedia.org/wiki/Digital_Signage) or just [show off some
pictures](http://www.subitophoto.net/slideshow/machupicchu/) somewhere.

[Neon](http://neon.webconverger.com/) has further tweaks for the digital
signage use case, such as a black background.

<img src="/img/chrome/webcfullscreen.png" alt="webcfullscreen" />

If you're concerned about the scroll bar, alter the **overflow** CSS element.

# Source code

* [Webconverger extension source code](https://github.com/Webconverger/webconverger-addon/)

## Preferences locations

Locking down Firefox is a [project on its own](https://github.com/pyllyukko/user.js). In Webconverger logic to disable various Firefox anti-features reside in:

* <https://github.com/Webconverger/webc/blob/master/opt/firefox/browser/defaults/preferences/webconverger.js>
* <https://github.com/Webconverger/webconverger-addon> for example to disable <https://support.mozilla.org/en-US/kb/refresh-firefox-reset-add-ons-and-settings>
* <https://github.com/Webconverger/webc/blob/master/opt/firefox/distribution/distribution.ini>

## Notes on keeping the extension free of cruft

Look for ids upon the
[browser.xul](http://mxr.mozilla.org/mozilla-central/source/browser/base/content/browser.xul)
that introduce new unwanted UI elements. Then block them in the <https://github.com/Webconverger/webconverger-addon/blob/master/src/webconverger.css>.
