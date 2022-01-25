![](https://git.ucode.space/Phil/goshorly/badges/main/pipeline.svg)
# goshorly

An easy self-hosted Link shortener in Golang with Redis <3 [Live-Demo](https://gly.one)

---
## Supported architectures
### Docker
- amd64, arm64
- other versions can be build manually (via docker build)
## Binary Build
- linux (amd64,arm64)
- darwin (amd64,arm64)
- windows (amd64)
- other versions can be build manually (via go build)
---
**WARNING:**
- goshorly is in an early stage, it is not an Final Version! (Pre-Release Status v0.1.X)

If you have an feature request, please do not hesitate to open an issue or merge request.

Available Docker tags:
- https://git.ucode.space/Phil/goshorly/container_registry/1


## Install with Docker
Installation with Docker-Compose (with no reverse proxy / own proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose.yml
nano docker-compose.yml # Change the environment variables to your needs
docker-compose up -d
```

Installation with Docker-Compose (built in proxy / caddy as reverse proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose-proxy.yml
mv docker-compose-proxy.yml docker-compose.yml
nano docker-compose.yml # Change the command line on caddy to your domain & environment variables to your needs
docker-compose up -d
```

## Install with Binary version
- WIP