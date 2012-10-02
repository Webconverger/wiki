[[!meta title="Adding a Firefox Addon to Webconverger"]]

To install an extension you need to:


1. Place the extracted extension directory into the [[chroot]], such as [/opt/firefox/extensions](https://github.com/Webconverger/webc/tree/master/opt/firefox/extensions)
* Ensure the extension is correctly named as its `install.rdf` **id**

	root@x220:/opt/firefox/extensions# grep 'kiosk@webconverger.com' kiosk\@webconverger.com/install.rdf
	em:id="kiosk@webconverger.com"

You can verify this manually with something like:

	for i in */; do grep -m 1 -H em:id "${i}install.rdf"; done
