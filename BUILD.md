# Build `snip`

> The only requirement to build `snip` locally is [Go](https://go.dev/) v1.19 or better. 

To build `snip` you will need to first clone the repository: 

```shell
git clone git@github.com:mchmarny/snip.git
```

Then navigate into the directory and build the distributable locally: 

```shell
make cli
```

The resulting binary can be found in `bin` directory or just run locally like this:

```shell
bin/snip
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.
