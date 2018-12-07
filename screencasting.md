[[!meta title="Screen casting" ]]

For <https://www.youtube.com/webconverger/> we've made a couple of videos to
show users around Webconverger.

<a href="http://www.flickr.com/photos/hendry/7272704214/" title="Tripod with Iphone mount by Kai Hendry, on Flickr"><img src="http://farm8.staticflickr.com/7102/7272704214_402ce1f037.jpg" width="500" height="375" alt="Tripof with Iphone mount"></a>

Our video camera is a Iphone 4S with a generic plastic iphone4 and tripod to
keep the "video camera" steady. The [tripod was
18GBP](http://www.flickr.com/photos/hendry/7181199520) from ALDI and the iphone
mount was ~5GBP from Ebay.

For editing we are using iMovie '11 (9.0.4) on a [MBP](http://buyersguide.macrumors.com/#MacBook_Pro), which came with it.

# Screencasting

Surprisingly screencasting is a dark art.

# MacOSX

Although [vlc](http://www.videolan.org/vlc/) can capture `screen://` we found it
unstable, so we forked out 20USD for <http://shinywhitebox.com/ishowu-v1/>

The tricky thing here was setting up the preset as the same size as the full
size screen.

# Windows

Unfortunately [vlc](http://www.videolan.org/vlc/) was also unstable when capturing `screen://`.

I did try [CamStudio](http://camstudio.org) but it had some bugs like not being
able to stop a record once minimised and not detecting that I dropped to a
lower 800x600 resolution.

After [studying
Wikipedia](http://en.wikipedia.org/wiki/Comparison_of_screencasting_software) I
found [Active Presenter Free edition](http://atomisystems.com/) and downloaded it.

"Smart Capture" looks cool, but it did some weird things so I ended with Full
Motion with 1024x768.

Next problem is importing the Exported AVI file to iMovie. By default the
exported ActivePresenter MP4/AVI files are greyed out, since they are not in
the correct format, whatever that is. Workaround here was to copy the exported
file to my Linux machine and run [HandBrake](http://handbrake.fr/) over it:

	HandBrakeCLI -i export.avi -o export-able-to-be-imported.mp4

Now the resulting file can be imported into iMovie on the MBP. A joint Windows,
Linux & Mac effort!

# recordmydesktop 2.0

[recordmydesktop](http://en.wikipedia.org/wiki/RecordMyDesktop) has been a reliable tool for creating screencasts in OGG
video format. Unfortunately uploading a OGG video to Youtube results in a green
mess (no kidding). So instead we've been using:
<https://github.com/kaihendry/recordmydesktop2.0>

# iMovie

When importing videos, the Inspector tool's Normalize Clip Volume in the Audio
tab and Enhance to reduce background noise are its killer features.
