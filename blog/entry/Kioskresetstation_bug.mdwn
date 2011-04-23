<a href="http://www.flickr.com/photos/hendry/2852122333/" title="The kioskresetstation bug by Kai Hendry, on Flickr"><img src="http://farm4.static.flickr.com/3079/2852122333_c3b3366632.jpg" width="500" height="375" alt="The kioskresetstation bug" /></a>

`kioskresetstation=5` as a [[boot]] parameter option will configure
[Webconverger](http://webconverger.com/) to reset after five minutes of
inactivity. This is appropriate in environments where people fail to end their
session by closing all their tabs.

A [bug has been reported whereby Webconverger would not start properly without
a key press after a
reset](http://groups.google.com/group/webc-users/browse_thread/thread/39cf0cca0e1d8038/2381878f7f3b2759).
Trouble is most people didn't experience this bug at all. Sometimes I would see
it. Other times I wouldn't. This is why I removed the bug from the
[[bug_list|TODO]] only to see it reappear later. Very annoying.

In hindsight, the **key press** gives the bug away. Though I spent many days
thinking it was a problem with my
[code](http://git.webconverger.org/?p=webconverger-base.git;a=tree;f=kioskresetstation).
I even wrote my own little C program to detect user inactivity to replace <a
href="http://packages.qa.debian.org/x/xautolock.html">xautolock</a>. Finally I
narrowed it down to Iceweasel, as [Webkit's GtkLauncher](http://packages.qa.debian.org/w/webkit.html) did not have this
problem.

After some [debugging](http://flickr.com/photos/hendry/2852122333/) by removing `nosudo` from the [[boot_paramenters|boot]] I finally found the bug.

The browser was blocking on a read for random numbers!
[/dev/random](http://en.wikipedia.org/wiki//dev/random) is the proper random
number generator in Linux and gives generally quite good numbers. However it
uses entropy like **user input** in order to work. Obvious now isn't it?

(Possible) Solution: I will probably try symlinking /dev/random ->
/dev/urandom, which is less likely to block. I really need find out why Firefox
needs random numbers and what's the risk if the _quality_ of randomness drops.

Update: Mike Hommey has been [working on this
problem](http://bugs.debian.org/cgi-bin/bugreport.cgi?bug=499146) and has a
[patch](https://bugzilla.mozilla.org/attachment.cgi?id=339296) on the [Netscape
portable runtime Debian
package](http://packages.qa.debian.org/n/nspr/news/20080920T084726Z.html) as of
4.7.1-4.
