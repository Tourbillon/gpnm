GPNM
======
Self-hosted golang package name manager like `gopkg.in` service. GPNM will help you to manage your package name with your own domain name and private repositories.

Usage
====

``` shell
$ git clone https://github.com/Tourbillon/gpnm.git
$ cd gpnm && go run main.go start
```

Or you can start it with `docker`

```shell
$ sudo docker pull anbillon/gpnm:latest
$ sudo docker run -d --name=gpnm-test -p 50000:50000 -e DB_PATH=/var/lib/gpnm.db -v ./gpnm.db:/var/lib/gpnm.db hub.anbillon.com/gpnm
```

Then you can access http://localhost:50000/gpnm/package to try it, the default username and password is `admin` and `admin`.The package name will change according to your domain, so no data need to be changed if domain changed.
If everything is ready, you can put this service on internet with `https` and specify `your.domain` to this service. And then you can do something in cmd.
``` shell
$ go get -u your.domain/foo/bar
```


License
======
```text
MIT License

Copyright (C) 2018-present Anbillon Team

This source code is licensed under the MIT license found in the
LICENSE file in the root directory of this source tree.
```
