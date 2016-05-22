# Routinator

This project is to help you setup and deploy an OpenBSD
router/firewall. It provides all of the basic configuration you will
need to get started. You just have to supply the details you want and
it will take care of the rest.

### What you get

The installer will put the base set of packages in place and download
the system source code. It will also download the routinator binary,
sample configuration, and templates. Once the installer is complete,
you just need to run the routinator program located in the setup
directory and it will create and install all the configuration files
necessary to get your router up and running. You will need to edit the
`firewall.example.json` file and add your system details and desired
network configuration. The routinator program will take care of the
rest.


### Getting Started

Once you have OpenBSD installed, run the following commands:

```sh
$ ftp https://github.com/abedra/routinator/releases/download/0.0.3/install.sh
$ sh install.sh
```

This will take some time depending on your internet connection.  After
it is done you will have the following.

* The OpenBSD source tree in place and ready to be updated.
* The routinator binary and supporting files in the `setup` folder

### Configuring your firewall

Before running the routinator program, make sure to take a look at
`firewall.example.json` and make sure it is properly configured. These
values will be placed into the proper configuration files. Be sure to
make note of the interface names and enter them correctly. The values
should be pretty self explanatory. Once you are ready just run the
program.

```
$ cd setup
$ ./routinator
```

Once the program finishes all configuration files will be moved into
place. Simply reboot your system and you will be up and running.


### Additional resources

In `/root/bin` there are a few script to make keeping your system up
to date easy. They are put on the `PATH` for the root user by the
routinator program but you can invoke them directly if you need to as
long as you have the appropriate permissions. As long as you entered
the correct version and architecture in the json configuration these
scripts won't need any modification.

##### Updating the sources

```sh
$ /root/bin/update
```

##### Recompiling the kernel

```sh
$ /root/bin/recompile_kernel
$ reboot
```

##### Updating the system binaries

```sh
$ /root/bin/recompile_system
```

### Additional concerns

This provides a base router and firewall with a dhcp server and simple
NAT. The firewall config located in `/etc/pf.conf` controls the
firewall rules. You should take time to learn more about pf and add to
these rules to make sure your network is as safe and productive as
possible. While these rules are fine to get started and provide a
secure enough base layer, more specific configuration should be
considered before feeling completely comfortable with the security of
your setup. You should also consider changing the port that ssh runs
on. If you want a great reference on pf and all of the things it can
do, I recommend reading the
[Book of PF, 3rd Edition](https://www.nostarch.com/pf3).

You may also want to add static IP assignments. You can do so in
`/etc/dhcpd.conf`. If you want to learn more take a look at
[this example](http://www.openbsd.org/faq/pf/example1.html).
