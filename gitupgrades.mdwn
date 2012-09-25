*These are some draft notes for implementing a new upgrade and config system using git-fs. This is a work-in-progress.*

There is the concept of "reboot-required-parameters". These are
parameters in `cmdline` that require a reboot to take effect. Initially,
these will only be `git-revision` and `boot_append`, but perhaps there will
be more later.

 - initrd:
   - If a git-revision value is present on cmdline, it is passed to git-fs.
     Otherwise, git-fs just mounts whatever HEAD points to. The latter
     happens when the image is ro, or a rw image is booted for the first
     time.

 - `live-config.sh` / `updates.sh`:
   - *If `webc-cmdline=new` is on the real cmdline,
     copy `/live/image/live/webc-cmdline-new` to `/etc/webc/cmdline` (in
     case downloading the new config fails)*
   - *Else,* copy `/live/image/live/webc-cmdline` (if any) to
     `/etc/webc/cmdline` (in case downloading the new config fails)
   - Download the new config to `/etc/webc/cmdline`
   - If `/live/image` is not writable, bail out, we're done.
   - *Run git fetch to fetch the `git-revision` from the new cmdline (no-op if
     this revision is already available, but still important to know if
     the boot is succesful)*
   - *If fetching the new cmdline **and** the git fetch worked, this is
     a successful boot and we need to do the bootloader magic required to
     make this the default boot option (and clean up the old default config, kernel and initrd as well, probably)*
   - *If any reboot-required-parameters changed, and the new
     `git-revision` is available in /.git (i.e. the git fetch worked):*
     - *Extract the kernel and initrd from the new `git-revision` and
       store them as `/live/image/live/vmlinuz-$(git-revision)` and
       `/live/image/live/initrd-$(git-revision)`*
     - *Generate a new bootloader config (appending the `boot_append`
       and `git-revision` value from the new `cmdline` and add
       `webc-cmdline=new` to the kernel cmdline) and store that as
       `/live/image/boot/boot-new.cfg`*
     - *Copy the new cmdline (`/etc/webc/cmdline`) to
       `/live/image/live/webc-cmdline-new`*
     - *Set up the bootloader to boot from this new config (extlinux has
       support for booting a particular boot entry **once**, perhaps
       syslinux as well?)*
     - *If auto_reboot is set in the **new** cmdline, reboot*
   - *Else,* If any other parameters were changed:
     - Copy the new cmdline (`/etc/webc/cmdline`) to
       `/live/image/live/webc-cmdline`
