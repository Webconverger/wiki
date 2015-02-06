[[!meta title="Notes on PC BIOS"]]

BIOS sound scary, but it's simply the system responsible for booting your
Operating System of choice.

Some PC BIOS boot a USB stick by default, so you might need to adjust your PC's
BIOS at all.

* [Video showing how to change the BIOS to boot from the USB stick](https://www.youtube.com/watch?v=-EaYTPbfuhI)
* Video on [configuring Intel Video BIOS](http://youtu.be/WxnmhOLFpjE)

There are different keys of entering the BIOS **boot menu** depending on the BIOS,
it could be [Del], [Esc], [F2], [F8] or [F12] key. The BIOS usually flashes
which key to press on your keyboard that allows you to access the BIOS setup
for the briefest of moments as you turn your PC computer on.

Once you are in the BIOS menu, you need to find the boot menu and bubble up the
USB disk where you copied Webconverger upon. Then Save & Exit, to boot
Webconverger every time you power on that machine.

<a href="http://www.flickr.com/photos/hendry/6869935304/" title="Adjusting boot order on a American megatrends BIOS"><img src="http://farm8.staticflickr.com/7247/6869935304_610027e95b.jpg" width="500" height="375" alt=""></a>

# Intel NUC notes

Sadly the new <https://twitter.com/visualbios> might need you to toggle the
OS selection in order to boot!

<img src="/img/intelNUCbios.jpg" alt="Intel Visual BIOS OS selection">

F2 (BIOS) => Advanced => Boot => Boot Configuration => OS Selection

* for Wheezy based release (current) choose : Windows 7
* for Jessie (development branch), choose: Windows 8.X
