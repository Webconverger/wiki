---
{
    "title": "How to connect Webconverger wirelessly to your access point",
    "permalink": "/wireless/"
}
---

# Wireless

<iframe width="560" height="315" src="https://www.youtube.com/embed/f8xR42iDQ2c" frameborder="0" allowfullscreen></iframe>

Note only the subscribed "install version" of Webconverger can persistently
keep your wireless configuration between reboots.

Wireless setup works by [directly mapping
wpa](https://github.com/Webconverger/webc/blob/master/etc/webc/wireless)
[boot](/boot/) configuration names onto [wpasupplicant for
Debian](http://anonscm.debian.org/viewvc/pkg-wpa/wpasupplicant/trunk/debian/README.Debian?view=markup).

### Example 0, essid 'home' with no security

Simply appending `wpa-ssid=home wpa-key-mgmt=NONE` to your [boot](/boot/) command
line should set your machine up for these wireless settings.

### Example 1, hidden essid 'home', no/disabled security

Appending `wpa-ssid=home wpa-key-mgmt=NONE wpa-ap-scan=1 wpa-scan-ssid=1` will create the following configuration in [/etc/network/interfaces](https://github.com/Webconverger/Debian-Live-config/blob/master/webconverger/config/includes.chroot/etc/network/interfaces).

	iface wlan0 inet dhcp
		   pre-up iptables-restore < /etc/iptables.conf
		   wpa-driver wext
		   wpa-ssid home
		   wpa-key-mgmt NONE
		   wpa-ap-scan 1
		   wpa-scan-ssid 1

WARNING: Hidden essids are trivially detected with tools like [Wifi
Analyzer](https://market.android.com/details?id=com.farproc.wifi.analyzer).
Setup your Access Point to use WPA please.

### Example 2, broadcasting essid 'home', with WEP HEX key BBE54998315E7E1616B8462B45

Append `wpa-ssid=home wpa-key-mgmt=NONE wpa-wep-key0=BBE54998315E7E1616B8462B45`

WARNING: WEP is deprecated and insecure technology. Please use WPA instead.

### Example 3, broadcasting essid 'home', with WPA key uiopzxcv

Append `wpa-ssid=home wpa-psk=uiopzxcv`

### Example 4 "Spaces in the ESSID", broadcasting essid 'Hopstock Gjestenett', with WPA key uiopzxcv

Please avoid spaces in ESSIDs. In this case we workaround with a `encodeURI('Hopstock
Gjestenett')`, to get the following [boot](/boot/) recipe:

	wpa-ssid=Hopstock%20Gjestenett wpa-psk=uiopzxcv

# Wireless router (recommended)

<http://routerboard.com/RBmAP2n>
