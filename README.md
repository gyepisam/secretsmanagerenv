# SecretsManagerEnv

SecretsManagerEnv (smenv) provides a convenient way to launch a processes with
environment variables sourced from @AWS's [Secrets Manager](https://aws.amazon.com/secrets-manager). Smenv is influenced
by [Vaultenv](https://github.com/channable/vaultenv) and Hashicorp's [Envconsul](https://github.com/hashicorp/envconsul).


Smenv is built around the concept of [twelve factor apps](https://12factor.net/config).

## Usage
```bash
smenv -s SECRET [-s SECRET]... [-r AWS_REGION] COMMAND
```

## Example
```bash
smenv -s rails_secrets/prod -s hello_world_secret -r us-east-1 "rails s -b 0.0.0.0"
```
