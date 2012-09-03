[Failed to open connection to session bus](http://ix.io/2VW) means dbus isn't installed.

	apt-get install dbus dbus-x11

Then

	dbus-launch --autolaunch `cat /var/lib/dbus/machine-id`

Then to test dbus use

	dbus-monitor

To monitor the bus and:

	dbus-send --print-reply --dest=org.freedesktop.DBus /org/freedesktop/DBus org.freedesktop.DBus.ListNames

To see stuff whizzing by.

[dbus](http://en.wikipedia.org/wiki/D-Bus) is required by [[ibus|blog/entry/Restoring_Japanese_support]] for CJK
input. Otherwise dbus is to **AVOIDED** _please_. Thanks,
