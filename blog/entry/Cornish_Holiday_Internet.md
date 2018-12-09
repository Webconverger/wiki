Happy new year 2009!

I'm staying at my parent's [cottage in Cornwall](http://prazefarm.co.uk/) and
enjoying slow Internet connectivity.

<img src="http://www.speedtest.net/result/383433002.png">

I moved the ADSL splitter from the back room to plug directly into the [BT
NTE5](http://www.telephonesuk.co.uk/images/NTE5_BT_front.jpg). I wish these
things had integrated [ADSL](http://en.wikipedia.org/wiki/ADSL) splitters and
basically [BT](http://en.wikipedia.org/wiki/British_telecom) was responsible
for them.

I then emailed (wrong move) [Devon based Eclipse
Internet](http://www.eclipse.net.uk/) to retrain the connection, since the
connection seemed locked to half a meg... I should have called Eclipse support
on 0845 1224 111 and actually got something done, instead of waiting for a week
for a silly email reply.

Another annoying element to Eclipse ISP is that they've dropped their unlimited
"Fair use" Evolution packages, ushering customers onto a limited bandwidth (3G
a month!) plans. Welcome to 2009 in the UK. Unbelievable. What's worse is that
seem to encourage you to prioritize traffic with their [traffic
controller](https://www.eclipse.net.uk/index.cfm?id=connectionmanager)... oh
pretty please [KISS](http://en.wikipedia.org/wiki/KISS_principle)!

I've reworked the home networking now too, which was a little tricky since the
[Grade II listed house](http://en.wikipedia.org/wiki/Listed_building) has
**very** thick walls (especially on the ground level).

1. Master socket direct to upper level Linksys WAG354G Firmware Version: 1.01.11
* WAG354G to [wall power socket](http://en.wikipedia.org/wiki/Power_line_communication) [ZyXEL PLA401](http://flickr.com/photos/hendry/3106453789/)
* ZyXEL PLA401 to downstairs WRT54GL running the latest greatest [Tomato firmware](http://www.polarcloud.com/tomato)
* A second WRT54GL operating on [WDS](http://en.wikipedia.org/wiki/Wireless_Distribution_System) extending both 'downstairs WRT54GL' and WAG354G wireless access point, towards the cottage

So the connection can hop from my [X61 laptop](http://dabase.com/x61/)
'wirelessly to' WRT54GL (WDS) 'wirelessly to'  WRT54GL 'cabled to' PLA401
'powerlined to' PLA401 'cabled to' WAG354G 'ADSL to' Eclipse... Yowsers.

Retraining the ADSL connection according to BT (via Eclipse telephone support)
can take several days. By that time, I'll be back in the South East of England.
I'm really hoping Praze Farm's connection can increase from half a meg to
something like the national UK average of 2 meg... shameful for such a small
developed country.

UPDATE: The connection at [Praze Farm](http://prazefarm.co.uk/) should be up
from ~500 Kbps (half a megabit) to 8 mbit!!! (8128 Kbps)

<img src="http://www.speedtest.net/result/385243946.png">

	DSL Status:      Up
	DSL Modulation Mode:     GDMT
	DSL Path Mode:           INTERLEAVED
	Downstream Rate:         8128 Kbps
	Upstream Rate:           448 Kbps
	Downstream Margin:       6 db
	Upstream Margin:         25 db
	Downstream Line Attenuation:     44
	Upstream Line Attenuation:       24
	Downstream Transmit Power:       0
	Upstream Transmit Power:         0


	PVC Connection
	Encapsulation:           RFC 2364 PPPoA
	Multiplexing:    VC
	QoS:     UBR
	PCR Rate:        0
	SCR Rate:        0
	Autodetect:      Disable
	VPI:     0
	VCI:     38
	Enable:          Yes
	PVC Status:      Applied

However the complex wireless over PNA in the house seems a tad unstable
according to my folks. Two steps forward, one step back. :)

