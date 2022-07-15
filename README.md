# Domotic server

This server is intended to retrieve sample data from sensors, store it and provide an API to request it.

# Samples

Sample categories can be either `temperature` or `humidity`.

# Ports

## MQTT Topics

| Topic                        | Type | Description                                       |
| -                            | -    | -                                                 |
| `discovery`                  | PUB  | Sends a message oftenly to discover new probes    |
| `discover/probes`            | SUB  | Save new probe from id in payload                 |
| `:probeID/samples/:category` | SUB  | Sensors                                           |

## HTTP Routes

| Route                                        | Method | Description                           |
| -                                            | -      | -                                     |
| `/probes`                                    | GET    | List all available probes             |
| `/probes/:probeID/getLatestSample/:category` | GET    | Get last sample for probe in category |
