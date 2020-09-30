# Using consul watch

This project is a study about consul watch: https://www.consul.io/docs/dynamic-app-config/watches

## Running this project

Start docker of project that contains the configuration of watch

```
docker-compose up
```
config file path : ./config/consul/config.json

Start the project

Mac
```
make run-ped-consul-watch-darwin
```

Linux
```
make run-ped-consul-watch-linux
```

After this run 
```
./init_config.sh http://localhost:8500
```
to configure consul with org mgm



