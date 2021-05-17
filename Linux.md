

```
/	The top-level directory is the root filesystem and contains all of the files required 
to boot the operating system before other filesystems are mounted as well as the files required to boot the other filesystems. 
After boot, all of the other filesystems are mounted at standard mount points as subdirectories of the root.

/bin	Contains essential command binaries.
/boot	Consists of the static bootloader, kernel executable, and files required to boot the Linux OS.
/dev	Contains device files to facilitate access to every hardware device attached to the system.
/etc	Local system configuration files. Configuration files for installed applications may be saved here as well.
/home	Each user on the system has a subdirectory here for storage.
/lib	Shared library files that are required for system boot.
/media	External removable media devices such as USB drives are mounted here.
/mnt	Temporary mount point for regular filesystems.
/opt	Optional files such as third-party tools can be saved here.
/root	The home directory for the root user.
/sbin	This directory contains executables used for system administration (binary system files).
/tmp	The operating system and many programs use this directory to store temporary files. This directory is generally cleared upon system boot and may be deleted at other times without any warning.
/usr	Contains executables, libraries, man files, etc.
/var	This directory contains variable data files such as log files, email in-boxes, web application related files, cron files, and more.
```

Linux Commands
```
whoami - Displays current username.
id - Returns users identity
hostname - Sets or prints the name of current host system.
uname - Prints operating system name.
pwd - Returns working directory name.
ifconfig - The ifconfig utility is used to assign or to view an address to a network interface and/or configure network interface parameters.
ip - Ip is a utility to show or manipulate routing, network devices, interfaces and tunnels.
netstat - Shows network status.
ss - Another utility to investigate sockets.
ps - Shows process status.
who - Displays who is logged in.
env - Prints environment or sets and executes command.
lsblk - Lists block devices.
lsusb - Lists USB devices
lsof - Lists opened files.
lspci - Lists PCI devices.
```

```
dpkg	The dpkg is a tool to install, build, remove, and manage Debian packages. The primary and more user-friendly front-end for dpkg is aptitude.
apt	Apt provides a high-level command-line interface for the package management system.
aptitude	Aptitude is an alternative to apt and is a high-level interface to the package manager.
snap	Install, configure, refresh, and remove snap packages. Snaps enable the secure distribution of the latest apps and utilities for the cloud, servers, desktops, and the internet of things.
gem	Gem is the front-end to RubyGems, the standard package manager for Ruby.
pip	Pip is a Python package installer recommended for installing Python packages that are not available in the Debian archive. It can work with version control repositories (currently only Git, Mercurial, and Bazaar repositories), logs output extensively, and prevents partial installs by downloading all requirements before starting installation.
git	Git is a fast, scalable, distributed revision control system with an unusually rich command set that provides both high-level operations and full access to internals.
```

-h : single '-' for letter (which can be concatenated)
--help : double '--' for whole word
-l : 'l' is longform version like 'ls -l'
'man' : manual for following app
'locate' xyz : locate any reference to the following keyword
'whereis' xyz : Find a binary with details
[], *, ., ',', ? : Wildcards
find / -type f -name apache2 : Find dir (/) options (-type) expression (-name) (.)
ps : processes 'ps aux' - all proccesses running
cat : will show contents of file
cat > : will create a file (ctrl + d to exit write mode)
cat >> : will appened the file
touch : will create followed by new filename
mkdir : will create followed by new directory name
rmdir : remove directory (unless not empty, use rmdir -r)
cp : to copy filename
mv filename newfilename : Moves a file and gives new name
head /1/2/3 pathway : Will give you the head of a file
tail /1/2/3 : will give you the tail of the file
tail -20 /a/b/c : will give you last 20 lines
nl /x/y/z : numberd lines
nl /a/b/c | grep output : numbered line keyword 'output'

hostname : Provides name of current user
uptime : Displays system uptime
w : user loggin summary
users : List of current users
ls -la : ls by line
more [filename] : view contents of file, pg by pg
sort [filename] : prints file in sorted order
free : Displays total free, used, and swap memory

top :  task manager program