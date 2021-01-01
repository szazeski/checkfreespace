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

`checkfreespace -help` gets you the available options and version

`checkfreespace -gb 1` sets the pass rate if the system has at least 1GB of free space

`checkfreespace -percent 2` sets the pass rate to 2% free (default is 10%)

`checkfreespace -json` converts output to json format
```
{"Filesystem":"","TotalGb":227,"FreeGb":149,"FreePercentage":65.63876651982379,"Hostname":"docker","Status":"[PASS] Disk OK","Passed":true}
```

## Installation

### Linux 64-bit
```
wget https://github.com/szazeski/checkfreespace/releases/download/v0.1/checkfreespace-linux-64 && chmod +x checkfreespace-linux-64 && sudo mv checkfreespace-linux-64 /usr/bin/checkfreespace 
```

### Linux 32-bit
```
wget https://github.com/szazeski/checkfreespace/releases/download/v0.1/checkfreespace-linux-32 && chmod +x checkfreespace-linux-32 && sudo mv checkfreespace-linux-32 /usr/bin/checkfreespace 
```

### Linux ARM
```
wget https://github.com/szazeski/checkfreespace/releases/download/v0.1/checkfreespace-linux-arm && chmod +x checkfreespace-linux-arm && sudo mv checkfreespace-linux-arm /usr/bin/checkfreespace 
```

### Mac 64-bit
```
wget https://github.com/szazeski/checkfreespace/releases/download/v0.1/checkfreespace-mac && chmod +x checkfreespace-mac && sudo mv checkfreespace-mac /usr/bin/checkfreespace 
```

# Syscall Issues
- Windows doesn't have `syscall.Statfs`

- Mac has `syscall.Statfs_t.Fstypename` but linux and windows doesn't. It returns the partition type like `apfs`
