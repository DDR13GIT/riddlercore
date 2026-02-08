# riddlercore

## Note: You MUST run the commands before doing any commit!
```sh
make init
```

**For commit message follow convention from here: https://www.conventionalcommits.org/en/v1.0.0/#summary**

### Installation
```sh
# clone repository
$ git clone $REPOSITORY
$ cd $REPOSITORY
$ make init

# run necessary service container
$ docker-compose up -d

# run migration
$ make migration-up

# run application; note this requires for a config.yml file. make a config.yml file from example and change accordingly
$ make build-run

```
