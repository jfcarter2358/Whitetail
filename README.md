# Whitetail

## About
`Whitetail` is a lightweight alternative to ELK for non-intensive applications. It is built with containerization in mind for a non-resource intensive log viewer with basic metadata search capabilities.

Data is stored in `Whitetail` through the use of a `CeresDB` database.  The `CeresDB` project can be found [here](https://github.com/jfcarter2358/ceresdb)

## TO DO
- [x] Add settings page
- [x] Add docker file
- [x] Add kubernetes manifest
- [x] Add query language so that it's not just a keyword search on the logs page
- [x] Add a log refresh button on the logs page
- [x] Add filtered by log level
- [x] Add log age cleanup
- [x] Add configurable branding
- [x] Add 'Reset to Default' buittons in settings page
- [x] Add configuration section to README
- [x] Make `PostgreSQL` database name configurable
- [x] Add spinner for query
- [x] Make list of services repopulate on click
- [x] Break logs and query out into separate pages
- [x] Improve logs page performance
- [x] Add logger name tooltip
- [x] Add query page AQL error box
- [x] Add `${time unit} > ${value}` indices
- [x] Add `${time unit} >= ${value}` indices
- [x] Add `${time unit} <$ {value}` indices
- [x] Add `${time unit} <= ${value}` indices
- [x] Make `[enter]` on query screen execute query
- [x] Optimize AST query
- [x] Remove UUIDs from indices on delete
- [x] Move backend database to use Ceres
- [x] Online documentation
- [x] Update manifests
- [x] Update README
- Release 1.0.0
- [x] Add files endpoint to upload files and have them stored as logs
- [x] Update README
- [ ] Add content to home page
- [ ] Add analytics page
- [ ] Update docs content and host on readthedocs
- [ ] Update tests
- Release 1.3.0
- [ ] Add Log backup capabilities/config option
- [ ] Add Log backup to S3 capabilities/config option

## Antler Query Language (AQL)

For more information on AQL, see the documentation here: [CeresDB AQL documentation](https://ceresdb.readthedocs.io/en/latest/querying.html)

## Configuration

Whitetail can be configured via the following environment variables:

Variable                          | Default value | Description
----------------------------------|---------------|------------
WHITETAIL_HTTP_PORT               | `9001` | Port to expose UI
WHITETAIL_TCP_PORT                | `9002` | Port to expose TCP listener
WHITETAIL_UDP_PORT                | `9003` | Port to expose UDP listener
WHITETAIL_BASE_PATH               | `""` | UI base path to serve on
WHITETAIL_DB                      | `{"username":"ceresdb","password":"ceresdb","name":"whitetail","port":7437,"host":"ceresdb"}` | Configure CeresDB database connection
WHITETAIL_LOGGING                 | `{"max-age-days":2,"poll_rate":"1h","concise_logger":true,"hoverable_long_logger":false}` | Configure log retention, polling, and display
WHITETAIL_BRANDING                | `{"primary_color":{"background":""#C3C49E","text":"#000000"},"secondary_color":{"background":"#8F7E4F","text":"#FFFFFF"},"tertiary_color":{"background":"#524632","text":"#FFFFFF"},"INFO_color":"#4F772D","WARN_color":"#E24E1B","DEBUG_color":"#2B50AA","TRACE_color":"#610345","ERROR_color":"#95190C"}` | color scheme for interface and log levels
WHITETAIL_PRINT_ELEVATED_MESSAGES | `false` | Print logs of level `ERROR` and `WARN` to stdout during operation

### WHITETAIL_DB

Name       |
-----------|
`username` | Username to authenticate with CeresDB
`password` | Password to authenticate with CeresDB
`name`     | Name of the database to create
`port`     | CeresDB port to connect to
`host`     | CeresDB host to connect to

### WHITETAIL_LOGGING

Logging configuration mainly concerns itself with the cleanup process to remove old logs, however it does also configure some aspects of the log message formatting.

Name                    | Description
------------------------|------------
`max-age-days`          | How many days to keep logs for (integer)
`poll-rate`             | How often to check for old logs to clean them up. Is of the form `< number >< time unit >` where valid time units are `ns`, `us`, `ms`, `s`, `m`, `h`
`concise-logger`        | Should the logger name be compacted for ease of viewing (bool)
`hoverable-long-logger` | Should the logger name be expanded when you hover over it (bool)

### WHITETAIL_BRANDING

Branding configuration allows for `Whitetail` to be customized to fit your product that it is being used in conjunction with. You an either change these through the `Settings` page in the UI or through the configuration file.

Name                         | Description
-----------------------------|------------
`primary_color.background`   | Primary branding color
`primary_color.text`         | Color for text over primary branding color
`secondary_color.background` | Secondary branding color
`secondary_color.text`       | Color for text over secondary branding color
`tertiary_color.backgroud`   | Tertiary branding color
`tertiary_color.text`        | Color for text over tertiary branding color
`INFO_color`                 | Color to be used to highlight `INFO` level logs
`WARN_color`                 | Color to be used to highlight `WARN` level logs
`DEBUG_color`                | Color to be used to highlight `DEBUG` level logs
`TRACE_color`                | Color to be used to highlight `TRACE` level logs
`ERROR_color`                | Color to be used to highlight `ERROR` level logs

In addition, you can configure the logo shown in the UI by placing your own logo file at `< whitetail root >/config/custom/logo/logo.png` and you can configure the icon shown in the browser by placing your own icon file at `< whitetail root >/config/custom/icon/favicon.png`

## Container Log Persistence

To persist your logs when running in a containerized environment you have a couple of options. You can either run `Whitetail` with an external (persisted) Ceres database _or_ you can persist the `/home/ceresdb/data` directory in your CeresDB deployment

## Contact

`Whitetail` is written by John Carter

If you have any questions or concerns, feel free to contact me at jfcarter2358.at.gmail.com
