# checkfreespace

```
./checkfreespace

docker
/ is
 Free:    149 GB / 227 GB
 Percent: 65.64%
[PASS] Disk OK
```


```
Usage of ./checkfreespace:
  -gb float
        a number like 1.2 that will trigger an alert if free space is less than 1.2GB
  -json
        switch to json output
  -percent float
        a number like 2.5 that will trigger an alert if free space is under 2.5 percent (default 10)
  -version
        v0.1 2020-DEC-29
```


# Syscall Issues
- Windows doesn't have `syscall.Statfs`

- Mac has `syscall.Statfs_t.Fstypename` but linux and windows doesn't. It returns the partition type like `apfs`