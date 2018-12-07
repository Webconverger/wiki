[[!meta title="Checklist for making a release"]]

DEPRECATED: Favouring <https://travis-ci.org/Webconverger/webc/>

1. Tag [webc](https://github.com/Webconverger/webc) and [Debian Live config](https://github.com/Webconverger/Debian-Live-config) - `git tag -s 13.0 && git push origin 13.0`. Permission problems? `sudo chown -R $USER:root .git`
* Build it `gb# /srv/www/build.webconverger.org/webc-build.sh`
* [[Test|testing]] it
* Just debug.log [[ping|privacy]], check clients config fetch
* On gb, run `aws s3 cp --acl public-read --storage-class REDUCED_REDUNDANCY webc-35.0.iso s3://webcdl/`
* Update houston.dreamhost.com:/home/hendry/dl2.webconverger.com/index.php
* Back up old release to http://archive.webconverger.com/
* Blog release notes, for [example](http://webconverger.org/blog/2015/Webconverger_32_release/)
* Make sure <https://webconverger.com/> links to latest release

* Mail [webc-users](https://github.com/Webconverger/dotCom/blob/master/list/maillist), linking blog entry, for [example](https://groups.google.com/forum/#!topic/webc-users/6rDiV-Sjhf8/discussion)
* Log into [twitter](http://twitter.com/#!/webconverger) and tweet linking the release notes, for [example](http://twitter.com/webconverger/status/205963007925301248)
* Email distrowatch distro@distrowatch.com with release notes
* [Update Wikipedia](http://en.wikipedia.org/wiki/Webconverger)
* Send [Webconverger newsletter](https://github.com/Webconverger/dotCom/blob/master/list/maillist)
