# GoSense

[![Build Status](https://travis-ci.org/netroby/gosense.svg?branch=master)](https://travis-ci.org/netroby/gosense)

Happy New Year! 2016.
A blog web app powered by golang, see demo : https://www.netroby.com
Using MySQL as database to storage data. Using Amazon S3 to handle file uploads.

## Feature


1. AWS S3 file upload storage
2. RSS output
3. Powered by golang
4. Builtin groupcache, amazing fast
5. Less cpu, memory usage than (wordpress etc)

## Install


First clone this repository to you pc/mac/laptop.

```bash
git clone https://github.com/netroby/gosense.git
```
Then rename config.toml.dist to config.toml, and change admin password and aws sdk key, secret

And you must install docker-engine and docker-compose, We using docker to build and running gosense.
We tested gosense with golang 1.5.* , 1.6.*., 1.7.*

if you installed docker-engine  and docker, please run 

```
docker-compose up --build
```

Once you docker up and running, you may access demo via http://127.0.0.1:8080
To login, you need visit http://127.0.0.1:8080/admin/login  (The password will be found in config.toml file)
To create blog , you can visit http://127.0.0.1:8080/admin/addblog

Remember change 127.0.0.1 to your docker hosting machine real ip address.


## License

We using MIT License, so feel free to change anything as you want.

## Donate me please

Your donate will help me improve the code quality. I need money to pay for food, water.

### Bitcoin donate

```
136MYemy5QmmBPLBLr1GHZfkES7CsoG4Qh
```
### Alipay donate
![Scan QRCode donate me via Alipay](https://www.netroby.com/assets/images/alipayme.jpg)

**Scan QRCode donate me via Alipay**


## News

### 2016-08-24

* Now you can run gosense using docker-compose

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


