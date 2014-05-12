[[!meta title="When Webconverger idles"]]

# How Webconverger clients "blank" (save monitor power, by turning off or going black)

Webconverger does not currently take active measures to sleep the PC hardware
as it is assumed to be deployed on mains connected PCs. So when we talk about
power saving and sleeping, we are referring to the screen monitor.

By default the screen monitor is set to [sleep after 10
minutes](https://github.com/Webconverger/webc/blob/master/home/webc/.xinitrc#L38).
Sleeping at best uses hardware level
[DPMS](http://en.wikipedia.org/wiki/VESA_Display_Power_Management_Signaling)
features of the monitor to power down to an idle state. At worse it sets the
screen to a blank black screen to save monitor burn which is only really a
problem with [older CRT
monitors](http://en.wikipedia.org/wiki/Cathode_ray_tube).

Customers who don't want the screen to blank at all should use the boot option
`noblank`.

Customers who want to adjust the default 10 minutes, can use `blank=` API to
set the minutes. E.g. a configuration with `blank=60` will only blank a monitor
after 1 hour of non-use.

Caveat: A [bug suggests there is an uppler limit](https://github.com/Webconverger/webc/issues/186)

Use case: Client may want to show an "attraction loop" when kiosks are not in
use. To realise this set `noblank` in the configuration and setup a [Web based
screensaver](https://github.com/Webconverger/Screensaver) from their homepage.

Breaking out of sleep simply involves a user pressing any key on the keyboard
or moving the mouse. It certain environments a system may never go to sleep
since a sensitive mouse keeps getting spurious wake events from its noisy
environment.

# Kiosk reset timer

Use case: Member of public uses Webconverger without closing last tab and
leaves it on a non-default homepage. We need to set a timeout of non-use to
reset the station. This is implemented with **kioskresetstation**.

The cmdline option `kioskresetstation` resets the Webconverger terminal after a
specified number of minutes of inactivity.

For example, appending `kioskresetstation=10` will reset the kiosk after ten minutes of inactivity.

* [Source of kioskresetstation](https://github.com/Webconverger/webc/blob/master/usr/bin/kioskresetstation)

Caveat: Your homepage must have a title for `kioskresetstation` to work currently. [Logged issue](https://github.com/Webconverger/webc/issues/24).

# Webconverger hides the mouse cursor after a period of inactivity

[Unclutter](http://packages.qa.debian.org/u/unclutter.html) is run by [default
in
Webconverger](https://github.com/Webconverger/webc/blob/master/home/webc/.xinitrc).

To override, you can add the [[API]] parameter `showcursor`.
