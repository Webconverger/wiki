UPDATE: Latest state of Webconverger's custom Firefox extension is at <https://github.com/Webconverger/webconverger-addon/>

The 'Iceweasel-webconverger' package has two small deviations webcnoaddressbar
and webcfullscreen described in [[kiosk]]. Instead of having 3 seperate repositries, you
can consolidate them nicely in one git repository as **branches**. 

I initially used `git-import-orig` to import each individual package into their
own git repository with [git-buildpackage](http://packages.qa.debian.org/g/git-buildpackage.html).

Then from an [iceweasel-webconverger](https://github.com/Webconverger/webconverger-addon/) clone, I fetch in the other git repos as branches, for example:

    x61$ git fetch ssh://au/srv/git/iceweasel-webcnoaddressbar master:webcnoaddressbar

Ok so I end up with:

    x61$ git branch
    * master
      webcfullscreen
      webcnoaddressbar

So lets see the minor changes beween -webcnoaddressbar and the default -webconverger packages:

    x61$ git checkout webcnoaddressbar
    Switched to branch "webcnoaddressbar"
    x61$ git diff --stat master
     chrome.manifest                       |    6 +-
     content/wc.css                        |   14 +---
     content/wc.js                         |   27 +-----
     content/wc.xul                        |   12 ++-
     debian/changelog                      |  156 +--------------------------------
     debian/control                        |    6 +-
     debian/copyright                      |    4 +-
     debian/dirs                           |    2 +
     debian/iceweasel-webconverger.dirs    |    2 -
     debian/iceweasel-webconverger.install |    1 -
     debian/iceweasel-webconverger.links   |    1 -
     debian/install                        |    1 +
     debian/links                          |    1 +
     defaults/preferences/prefs.js         |    1 -
     install.rdf                           |   15 ++--
     15 files changed, 34 insertions(+), 215 deletions(-)

Unfortunately you can see a lot of boring debian/* packaging changes that are
required. Again I should make further efforts to consolidate the 'modes' of a
[[Firefox_kiosk_extension|kiosk]].

Still TODO, nicely generate -webcnoaddressbar and -webcfullscreen packages from the git repositry with `git-buildpackage`. :)

I hope to polish these [[kiosk]] 'digital signage' feature branches so we can see less of [failed windows digital signs](http://www.flickr.com/photos/hendry/2178479733/).
