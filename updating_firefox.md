---
{
    "title": "Updating Firefox for a Webconverger developer",
    "permalink": "/updating_firefox/"
}
---

We now use [ESR because of the painful signing process of
extensions](https://www.mozilla.org/en-US/firefox/organizations/all/)

Get the 32bit version and for the `ver=` value, see what
<https://www.mozilla.org/en-US/security/known-vulnerabilities/firefox-esr/> or
<https://www.mozilla.org/en-US/firefox/organizations/notes/> redirects to.

	cd /tmp
	ver=52.2; wget "https://download.mozilla.org/?product=firefox-$ver.0esr-SSL&os=linux&lang=en-US" -O firefox-$ver.tar.bz2
	open https://www.mozilla.org/en-US/firefox/$ver/releasenotes/

	tar xvf firefox-$ver.tar.bz2

	sudo mv firefox $WEBC_CHECKOUT/tmp/
	/opt# mv firefox/ firefox-old
	/opt# mv /tmp/firefox/ .

Using the **firefox-i18n** packaging from Archlinux:

	curl -O https://git.archlinux.org/svntogit/packages.git/plain/firefox-i18n/trunk/PKGBUILD

Update `pkgver=52.0esr`

	makepkg -o
	mkdir /tmp/exts
	/tmp/firefox-i18n$ for i in *.xpi; do unzip -d /tmp/exts/$(unzip -p $i install.rdf | grep -m 1 -H em:id  | sed 's,.*em:id=",,' | sed 's,"$,,') $i; done
	cd /tmp/exts
	rsync -a --delete . $WEBC_CHECKOUT/opt/firefox/langpacks/

Check for invalid XML: `find -name '*.xml' -o -name '*.rdf' | xargs xmlstarlet val | grep inv`

Delete searchplugins since they are not needed, except DDG.

# Doing an update

	/opt# cp -ra firefox-old/browser/extensions/ firefox/browser/
	/opt# cp -ra firefox-old/browser/defaults/ firefox/browser/
	/opt# cp -ra firefox-old/distribution/ firefox/

Remove crash reporter

	vim firefox/application.ini

Make sure nothing is deleted / missing. Otherwise you might get a strange error
message, like please contact your system administrator.

	git checkout firefox/mozilla.cfg

# Remove useless items from langpacks

	find /opt/firefox/langpacks/ -name searchplugins -type d | xargs rm -rf
