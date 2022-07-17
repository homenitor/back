# Domotic server

This server is intended to retrieve sample data from sensors, store it and provide an API to request it.

# Samples

Sample categories can be either `temperature` or `humidity`.


# Launch

## Commands

| Command                 | Description                              |
| ----------------------- | ---------------------------------------- |
| `make run-dependencies` | Docker pull and run all dependencies     |
| `make run-app`          | Build and run app                        |
| `make run`              | Launch `run-dependencies` then `run-app` |
| `make test`             | Run all tests                            |
| `make clean`            | Clean all files                          |

## Environment variables

| Variable                  | Default value             |
| ------------------------- | ------------------------- |
| `DISCOVERY_PERIOD`        | 1m                        |
| `LOG_LEVEL`               | info                      |
| `MQTT_HOST`               | 127.0.0.1                 |
| `MQTT_PORT`               | 1883                      |
| `MQTT_QUALITY_OF_SERVICE` | 2                         |
| `MONGODB_URI`             | mongodb://localhost:27017 |

# Ports

## MQTT Topics

| Topic                        | Type | Description                                        |
| ---------------------------- | ---- | -------------------------------------------------- |
| `discovery`                  | PUB  | Sends a message oftenly to discover new probes     |
| `discover/probes`            | SUB  | Save new probe from id in payload                  |
| `:probeID/samples/:category` | SUB  | Save sample for the probe in category, timestamped |

## HTTP Routes

| Route                                        | Method | Description                                  |
| -------------------------------------------- | ------ | -------------------------------------------- |
| `/probes`                                    | GET    | List all available probes                    |
| `/probes/:probeID/getLatestSample/:category` | GET    | Get last sample for probe in category        |
| `/samples/:category`                         | GET    | List all samples for category for each probe |
