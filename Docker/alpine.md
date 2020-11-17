# Alpine Images 
[Alpine Linux](https://alpinelinux.org/) is a Linux distribution built around [musl libc](https://www.musl-libc.org/) and [BusyBox](https://www.busybox.net/). The image is only 5 MB in size and has access to a package repository that is much more complete than other BusyBox based images. This makes Alpine Linux a great image base for utilities and even production applications. 

[Read more about Alpine Linux](https://alpinelinux.org/about/) and you can see how their mantra fits in right at home with Docker images. 

## Other Info 
* [Alpine on Docker Hub](https://hub.docker.com/_/alpine/)
* Has no bash interface (by default) due to size 
* `docker pull alpine` 
* Start Alpine (with bash) - `docker container run -it alpine sh `
* Use `apk` within  alpine to install bash and likes of it

