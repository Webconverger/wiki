[[!meta title="Privacy policy"]]

# Privacy policy

Webconverger is designed to ensure users privacy and [[security]] whilst
surfing the Web, primarily by wiping the slate clean between sessions.
Therefore public users of a Webconverger kiosks should not have to fear they
are leaving any private data behind.

Webconverger does not record your search history or collect any data of your
surfing habits.

The Webconverger operating system sends the following data to our servers on
boot:

* Machine identity, your machine UUID + MAC address
* IP address from where you sent it from
* Version of Webconverger you use

We call this the **ping**:

* [ping client source](https://github.com/Webconverger/webc/blob/master/etc/webc/network-up.d/ping)
* [ping receiver source](https://github.com/Webconverger/ping/blob/master/index.php)

This "ping" is used to determine how many users Webconverger has to warrant
further investment. [Public usage statistics](http://ping.webconverger.org/).

You can **opt out** by appending `noping` to the command line. **noping** means no
collection of the unique machine identity (UUID + MAC address) or version.

## Configuration service

By default (unless **opted out** by specifying `noconfig`), we apply our
[configuration service](http://config.webconverger.com). This configuration
service will use your machine UUID and MAC address as machine identifiers (not
personally identifying), to track and apply configurations. This is how the
company generates revenue in order to support the opensource project.

If you do not have a subscription, your Webconverger instance will be
unconfigured by default, which is usable in many scenarios.

All our sites use [Google Analytics](http://www.google.com/analytics/).

We try to make the privacy policy as clear as possible. If you have any
questions please email privacy AT webconverger.com.

## Default search service

If you have the URL bar enabled, and you do not type in a valid URL, the input
will be passed onto <https://duckduckgo.com/> as search terms.
