The explicit need to allow popups is for homepages that are frankly poorly
engineered to use popups as an interaction model. Popups in Webconverger are
configured to forcefully open as new tabs.

Using the new improved `prefs=` API from
<https://github.com/Webconverger/webc/commit/a1b18d281e1b4858564b38fdda12b36352eed7bf>
we can use Firefox's powerful "autoconfig" commands to setup these permissions.

Lets walk through an example using <http://popuptest.com/> as a source for popups.

We [configure our Webconverger](http://config.webconverger.com/) using
`prefs=http://prefs.webconverger.com/2015/autoconfigfile.js` and this
"autoconfig" file contains:

	$ curl -s http://prefs.webconverger.com/2015/autoconfigfile.js
	Components.utils.import("resource://gre/modules/Services.jsm");
	Components.utils.import("resource://gre/modules/NetUtil.jsm");
	Services.perms.add(NetUtil.newURI("http://popuptest.com"), "popup", Services.perms.ALLOW_ACTION);
	Services.prompt.alert(null, "", "Popup test");

You should host `http://prefs.webconverger.com/2015/autoconfigfile.js` yourself,
though Webconverger as a service to its customers can reliably host it for you.

As you can see in the `Services.perms.add` invocation, with `prefs=` we can now
issue commands that add <http://popuptest.com> into a popup whitelist.

You can verify this works by visiting <http://popuptest.com/popuptest4.html> or
by examining the whitelist using `chrome=debug` in
Edit→Preferences→Content→Exceptions.

# More on "Netscape Mission Control Desktop" aka Firefox's autoconfig

Webconverger's `prefs=` [[API]] uses **autoconfig** files.

* <http://mike.kaply.com/2012/03/16/customizing-firefox-autoconfig-files/>
* <http://mike.kaply.com/2012/03/22/customizing-firefox-advanced-autoconfig-files/>
* [Canonical Mozilla documentation on the PermissionManager](https://developer.mozilla.org/en-US/docs/XPCOM_Interface_Reference/nsIPermissionManager)

Many [[thanks|acknowledgements]] to [Mike Kaply](http://mike.kaply.com) for
taking the time to document this almost hidden Enterprise feature of the
Firefox browser.
