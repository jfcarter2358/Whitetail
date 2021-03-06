# Whitetail

## About
`Whitetail` is a lightweight alternative to ELK for non-intensive applications. It is built with containerization in mind for a non-resource intensive log viewer with basic metadata search capabilites.

Data is stored in `Whitetail` through the use of a `Ceres` database.  The `Ceres` project can be found [here](https://github.com/jfcarter2358/ceres)

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
- [ ] Update manifests
- [x] Update README
- Release 1.0
- [ ] Add content to home page
- [ ] Add analytics page
- [ ] Add Log backup capabilities/config option
- [ ] Add Log backup to S3 capabilities/config option
- [ ] Any other tasks I'll inevitably think of later

## Antler Query Language (AQL)

AQL is a simple query language designed to be used when the standrard filtering (level and service) is not sufficient.  AQL statments are written in nested blocks of binary operations. This means that each operator can only have a singular left and singular right argument. In addition, data is retrieved by prepending your logic with a `SELECTBY` keyword or removed by prepending with a `DELETEBY` keyword. An example query which gets logs of level `INFO` and level `WARN` is as follows:

```
SELECTBY level = INFO OR level = WARN
```

If you want to change the 'OR' statements to include more than just the two levels, you'll wrap the first two up in parenthesis and then OR that with a third filter.

```
SELECTBY ( level = INFO OR level = WARN ) OR level = DEBUG
```

### Filters
The various filters that can be used in AQL statements are as follows (`< text like this is a placeholder >`):
Filter                      | Desrciption
----------------------------|------------
`level = < level >`         | Get logs with level `< level >` (`= < level >` can be replaced with `IN < csv of levels >`)
`service = < service >`     | Get logs from service `< service >` (`= < service >` can be replaced with `IN < csv of services >`)
`year = < year >`           | Get logs wtih a timestamp that has the year `< year >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`month = < month >`         | Get logs wtih a timestamp that has the month `< month >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`day = < day >`             | Get logs wtih a timestamp that has the day `< day >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`hour = < hour >`           | Get logs wtih a timestamp that has the hour `< hour >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`minute = < minute >`       | Get logs wtih a timestamp that has the minute `< minute >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`second = < second >`       | Get logs wtih a timestamp that has the second `< second >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`timestamp = < timestamp >` | Get longs with string timestampe in format `YYYY-MM-DDThh:mm:ss`

### Operators
The various operators are shown below with examples
- __AND__
    - `< left filter > AND < right filter >`
    - Returns logs which satisfy both the left and right filter
- __OR__
    - `< left filter > OR < right filter >`
    - Returns logs which satify the left or right filter
- __NOT__
    - `< left filter > NOT < right filter >`
    - Returns logs from the left filter that are not part of the right filter
- __XOR__
    - `< left filter > XOR < right filter >`
    - Returns logs that are part of the left or right filter but not both
- __LIMIT__
    - Limits the results from a filter to `N` log messages
    - `< left filter > LIMIT < N >`
- __ORDERBY__
    - `< left filter > ORDER_ASCEND < field >`
    - Orders the results of the left filter in ascending order by one of the following fields
        - level
        - service
        - text
        - timestamp
        - year
        - month
        - day
        - hour
        - minute
        - second
- __ORDERDESC__
    - `< left filter > ORDER_DESCEND < field >`
    - Orders the results of the left filter in descending order by one of the following fields
        - level
        - service
        - text
        - timestamp
        - year
        - month
        - day
        - hour
        - minute
        - second

## Configuration

An example of a `whitetail` configuraiton file is as follows:

```json
{
    "http-port": 9001,
    "tcp-port": 9002,
    "udp-port": 9003,
    "basepath": "",
    "database": {
        "url": "http://localhost:9090"
    },
    "logging": {
        "max-age-days": 2,
        "poll-rate": "1h",
        "concise-logger": true,
        "hoverable-long-logger": false
    },
    "branding": {
        "primary_color": {
            "background": "#C3C49E",
            "text": "#000000"
        },
        "secondary_color": {
            "background": "#8F7E4F",
            "text": "#ffffff"
        },
        "tertiary_color": {
            "background": "#524632",
            "text": "#ffffff"
        },
        "INFO_color": "#4F772D",
        "WARN_color": "#E24E1B",
        "DEBUG_color": "#2B50AA",
        "TRACE_color": "#610345",
        "ERROR_color": "#95190C"
    }
}
```

This JSON document holds the information required to run an instance of `Whitetail`. It can be broken down into the following sections.

### Basic Configuration

The basic configuration handles the server itself that `Whitetail` runs. This includes the following values:

Name        | Description
------------|------------
`http-port` | The port to serve the `Whitetail` UI on
`tcp-port`  | The port to listen to TCP logs on
`udp-port`  | The port to listen to UDP logs on
`basepath`  | the basepath to serve the various endpoints on

### Database

The database configuration is held in the `database` key in the configuration file. This defines which database `Whitetail` will use to hold log and index information. Currently, `Whitetail` currently supports a in-container `Sqlite` databse _or_ an external `PostgreSQL` database. 

To configure `Whitetail` to use `Sqlite`, set the `database` section to this:

```json
{
    "database": {
        "url": "http://localhost:9090"
    },
}
```


By default the URL points to a local instsance of `Ceres`, which works for the Kubernetes deployment. You can instead point it at an external instsance via the URL if you wish.

### Logging

Logging configuration mainly concerns itself with the cleanup process to remove old logs, however it does also configure some aspects of the log message formatting.

Name                    | Description
------------------------|------------
`max-age-days`          | How many days to keep logs for (integer)
`poll-rate`             | How often to check for old logs. Is of the form `< number >< time unit >` where valid time units are `ns`, `us`, `ms`, `s`, `m`, `h`
`concise-logger`        | Should the logger name be compaceted for ease of viewing (bool)
`hoverable-long-logger` | Should the logger name be expanded when you hover over it (bool)

### Branding

Branding configuration allows for `Whitetail` to be customized to fit your product that it is being used in conjunction with. You an either change these through the `Settings` page in the UI or through the configuration file.

Name                         | Description
-----------------------------|------------
`primary_color.background`   | Primary branding color
`primary_color.text`         | Color for text over primary branding color
`secondary_color.background` | Secondary branding color
`secondary_color.text`       | Color for text over secondary branding color
`tertiary_color.backgroud`   | Tertiary branding color
`tertiary_color.text`        | Color for text over tertiary branding color
`INFO_color`                 | Color to be used to highligh `INFO` level logs
`WARN_color`                 | Color to be used to highligh `WARN` level logs
`DEBUG_color`                | Color to be used to highligh `DEBUG` level logs
`TRACE_color`                | Color to be used to highligh `TRACE` level logs
`ERROR_color`                | Color to be used to highligh `ERROR` level logs

In addition, you can configure the logo shown in the UI by placing your own logo file at `< whitetail root >/config/custom/logo/logo.png` and you can configure the icon shown in the browser by placing your own icon file at `< whitetail root >/config/custom/icon/favicon.png`

## Container Log Persistence
To persist your logs when running in a containerized environment you have a couple of options. You can either run `Whitetail` with an external (persisted) Ceres database _or_ you can use the available PVC manifests (if you are running in Kubernetes) to backup the ceres data directory

## Contact
`Whitetail` is written by John Carter
If you have any questions or concerns, feel free to contact me at jfcarter2358.at.gmail.com