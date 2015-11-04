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
git clone https://github.com/netroby/gosense.git
./up.sh
```

Once you docker up and running, you may access demo via http://127.0.0.1:8080

And if you want reload gosense, just type docker restart gosense

## License

MIT License

## Donate me please

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


