# standalone_rev_shells
Build standalone binaries for reverse shells

## Linux builds

* Example below is for deb/ but works for rpm/ as well.
* Make sure to update the ip address and port within the Program.cs files before building the docker image.
```
cd linux/deb

docker build -t deb-standalone .

docker run --rm -it deb-standalone bash
```

* In another another terminal, copy the binary out

```
docker cp 418190c518a9:/opt/network/bin/Debug/net7.0/linux-x64/publish/network .
```

* Use it for your pentest.


## Coming Soon
* Standalone for Windows
* Standalone for ARM
* Standalone for Alpine
* Obfuscation
* Encryption

## Notes
* Remember there is no obfuscation and encryption
* The binary is straight forward, reverse engineering is possible
* Everything is in plaintext over the wire
* Receive permission from system owners before using, do not use on unauthorized systems
* I am not resposible for negligence.

## Sources
* Source code Template used for https://www.revshells.com/, compiled with dotnet core.
