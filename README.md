![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)

# In-Memory Key-Value Store

Project capable of performing basic "redis" functions

---

## Installation

```bash
git clone git@github.com:SenRecep/redisclone.git

export BUILD_INFO="$(git rev-parse HEAD)-$(go env GOOS)-$(go env GOARCH)" 

docker build --build-arg="BUILD_INFORMATION=${BUILD_INFO}" -t senrecep/redisclone:latest . 

SERVER_ENV="production" LOG_LEVEL="INFO" docker run --cpus="2" --env SERVER_ENV --env LOG_LEVEL -p 8000:8000 --name redisclone  senrecep/redisclone:latest
```


---

## Contributor(s)

* [recepsen](https://github.com/SenRecep) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/SenRecep/redisclone/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com//redisclone/blob/main/CODE_OF_CONDUCT.md