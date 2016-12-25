# Web Application Frontend

Simple frontend for webapps.

This project is based on the [Gallium](https://github.com/alexflint/gallium) project.

## Build application

```sh
$ glide install
$ go build -o waf *.go
$ gallium-bundle waf
```

## Run application

```sh
$ WAF_URL=https://github.com/mdouchement open waf.app
```

## License

**MIT**
