# Secret docker minimal example

Problem with encore secrets in docker

## run in local

```bash
encore run
```

Reads secrets from `.secrets.local.cue`

## run in docker

```bash
encore build docker --config ./infra.config.json minimum_secret:latest
docker-compose up
```

Reads secrets from `infra.config.json` and `docker-compose.yml`
