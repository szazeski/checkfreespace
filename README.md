# checkfreespace
a simple command line tool that checks the root of the system to see if it has 10% free disk space, if not it returns exit code 2 allowing CI to fail the build.

`checkfreespace`
```
/ on docker
 Free:    149 GB / 227 GB
 Percent: 65.64%
[PASS] Disk OK
```

## Options

 `-help` gets you the available options and version

 `-gb 1.5` sets the pass rate if the system has at least 1.5GB of free space

 `-percent 2.5` sets the pass rate to 2.5% free *(default is 10%)*

 `-json` converts output to json format:
`{"Filesystem":"","TotalGb":227,"FreeGb":149,"FreePercentage":65.63876651982379,"Hostname":"docker","Status":"[PASS] Disk OK","Passed":true}`

 `-path /` optionally overrides the default path to check free space on such as `/mnt/usbdrive` or `D:`

## Installation

### Linux 64-bit
`wget -O checkfreespace https://get.checkcli.com/checkfreespace/linux/64 && chmod +x checkfreespace && sudo mv checkfreespace /usr/bin/checkfreespace`

### Linux 32-bit
`wget -O checkfreespace https://get.checkcli.com/checkfreespace/linux/32 && chmod +x checkfreespace && sudo mv checkfreespace /usr/bin/checkfreespace`

### Linux ARM64
`wget -O checkfreespace https://get.checkcli.com/checkfreespace/linux/arm64 && chmod +x checkfreespace && sudo mv checkfreespace /usr/bin/checkfreespace`

### Linux ARM
`wget -O checkfreespace https://get.checkcli.com/checkfreespace/linux/arm && chmod +x checkfreespace && sudo mv checkfreespace /usr/bin/checkfreespace`

### Mac Intel 64-bit
(the app isn't signed yet, so run the app in finder to accept the Gatekeeper dialog by right clicking on it and selecting open)

`curl -O -L https://github.com/szazeski/checkfreespace/releases/download/v1.0.0/checkfreespace-darwin-amd64 && chmod +x checkfreespace-darwin-amd64`

Right click and open the app to approve gatekeeper for an unsigned app, then `mv checkfreespace-darwin-amd64 /usr/local/bin/checkfreespace`

### Mac ARM
(the app isn't signed yet, so run the app in finder to accept the Gatekeeper dialog by right clicking on it and selecting open)

`curl -O -L https://github.com/szazeski/checkfreespace/releases/download/v1.0.0/checkfreespace-darwin-arm64 && chmod +x checkfreespace-darwin-arm64`

Right click and open the app to approve gatekeeper for an unsigned app, then `mv checkfreespace-darwin-arm64 /usr/local/bin/checkfreespace`


### Windows
Download the proper file from the release section and save it in the `C:\Windows` folder if you want it in the system PATH.

# Syscall Issues
- Windows doesn't have `syscall.Statfs`, makes a kernel32.dll call instead.
- Mac has `syscall.Statfs_t.Fstypename` but linux and windows doesn't. It returns the partition type like `apfs`, currently the app just returns a blank string for filesystem.
