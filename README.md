# checkfreespace
a simple command line tool that checks the root of the system to see if it has 10% free disk space, if not it returns exit code 2 allowing CI to fail the build.

`checkfreespace`
```
/ on docker
 Free:    149 GB / 227 GB
 Percent: 65.64%
[PASS] Disk OK
```

`checkfreespace -gb 50`
```
/ on M1.lan
 Free:    24.6 GB / 228 GB
 Percent: 10.79%
[FAIL] Free disk space under 50.0GB
```
(exit code 2)

## Options

 `-help` gets you the available options and version

 `-gb 1.5` sets the pass rate if the system has at least 1.5GB of free space

 `-percent 2.5` sets the pass rate to 2.5% free *(default is 10%)*

 `-json` converts output to json format:
`{"Filesystem":"","TotalGb":227,"FreeGb":149,"FreePercentage":65.63876651982379,"Hostname":"docker","Status":"[PASS] Disk OK","Passed":true}`

 `-path /` optionally overrides the default path to check free space on such as `/mnt/usbdrive` or `D:`

## Installation

### Mac Homebrew

`brew install szazeski/tap/checkfreespace`

### Linux (and mac too)

```
wget https://github.com/szazeski/checkfreespace/releases/download/v0.1.0/checkfreespace_1.0.2_$(uname -s)_$(uname -m).tar.gz -O checkfreespace.tar.gz && tar -xf checkfreespace.tar.gz && chmod +x checkfreespace && sudo mv checkfreespace /usr/bin/
```

### Windows

```
Invoke-WebRequest https://github.com/szazeski/checkssl/releases/download/v0.5.0/checkfreespace_1.0.2_Windows_x86_64.zip -outfile checkfreespace.zip; Expand-Archive checkfreespace.zip; echo "if you want, move the file to a PATH directory like WINDOWS folder"
```

# Syscall Issues
- Windows doesn't have `syscall.Statfs`, makes a kernel32.dll call instead.
- Mac has `syscall.Statfs_t.Fstypename` but linux and windows doesn't. It returns the partition type like `apfs`, currently the app just returns a blank string for filesystem.
