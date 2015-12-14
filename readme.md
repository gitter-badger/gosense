# GoSense

A blog web app powered by golang, see demo : https://www.netroby.com


## Install

```
go get github.com/tools/godep
go get -u github.com/netroby/gosense
godep go build
./gosense # windows gosense.exe
```
Or you can install docker, then run docker container 

```
# to install docker on your platform
# wget -qO- get.docker.com | sudo sh
git clone https://github.com/netroby/gosense.git
./up.sh
```
Make sure your docker verision 1.9.1+
```
$ docker version
Client:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:20:08 UTC 2015
 OS/Arch:      linux/amd64

Server:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:20:08 UTC 2015
 OS/Arch:      linux/amd64

```

Once you docker up and running, you may access demo via http://127.0.0.1:8080


## Graceful restart 

And if you want reload gosense, just  run following command

```
docker kill -s HUP gosense
```
Or may you want to rebuild binary and graceful reload ?

```
git pull --rebase
./graceful-restart.sh
```

## License

MIT License

## Donate me please

### Bitcoin donate

```
136MYemy5QmmBPLBLr1GHZfkES7CsoG4Qh
```
### Alipay donate
![Scan QRCode donate me via Alipay](https://www.netroby.com/assets/images/alipayme.jpg)

**Scan QRCode donate me via Alipay**


## News

### 2015-10-31

* Add Cache for list and view page, now performance got better

### 2015-10-30

* Now gosense handle time better than before
* The search by keyword now working correct

### 2015-10-29

* The docker scripts now working perfect, you can easier deploy gosense 

### 2015-10-28

* Fix bug in add blog
* Now we can add, edit blog
* Fix SQL load


