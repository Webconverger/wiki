---
{
  "title": "Webconverger service level",
  "permalink": "/sla/"
}
---

Email: support@webconverger.com

Our current support hours are 9:00-17:00 (UTC+8) and we endeavour to respond to
your request within four working hours. For example an email sent to us at 1PM
(UTC+12) would reach us at 9AM (UTC+8), and we try respond at the latest by 5PM
(UTC+12). An email sent to us at 9AM (UTC+12), would be received at 5AM
(UTC+8), and would have a response by 5PM (UTC+12).

# How we handle issues

Once we have a confirmed reproducible issue from you, we will log the issue
upon our public [bug tracking
system](https://github.com/Webconverger/webc/issues), however it will be
anonymised, not to mention any names.

For issues that cannot be anonymised, such as a billing issue, we log them in
an internal bug tracker.

For paying customers, we prioritise engineering time to the issue to get it
fixed.

# Deployment support

Please verify you can reach <http://config.webconverger.com> which is the
configuration server your Webconverger will download its configuration from. An
installed webconverger will cache its configuration, so once deployed, it will
not need further access to http://config.webconverger.com. However if you
change the configuration, your install will need to access
<http://config.webconverger.com> in order to download the new configuration.

Typical steps to trouble shoot problems:

1. Reboot the machine
2. Re-install the machine with the [latest update](https://build.webconverger.com/latest.iso)
3. Try the older [release snapshot](https://build.webconverger.com/), since you might be suffering with a regression
4. Log machine identity and try to reproduce the problem. Our [bug tracking](https://github.com/Webconverger/webc/issues) is transparent (we don't name customers)

In any case, please email support@webconverger.com with the machine identity and problem description.

To provide us extra debugging information, we may ask you to please append the
keyword "support" to your configuration in order to send us logs and reboot the
machine.  http://support.webconverger.com/[MACHINE_ID]

# Long term support

Once Webconverger is installed and working the two classes of issues you should expect to see are:

1. Hardware fault - you need to replace the hardware, note a new machine identity requires a new seat
2. Update regression - we can roll you back to an earlier version or help you resolve any issue

Hardware faults are likely to be straightforward to identify if you are using
the same hardware across your deployment. In this case replacing the hardware
(with ideally exactly the same hardware) and re-installing Webconverger is the
solution.

Since Webconverger installs will by default update to the latest version with
the latest browser, it is possible you might see a regression. For example the
Web browser's page layout behaviour might change, since a new browser update
might remove or change some feature. This is very rare and we actively support
legacy applications by having a policy of avoiding any changes to the default
browsing experience, so you can expect to have a consistent Webconverger
behaviour.

If browser updates become a problem, one option is to use the [Update
API](http://webconverger.org/upgrade/) to pin your Webconverger to a previous
release which has been established to work reliably with your deployment
hardware and Websites.
