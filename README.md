# Whitetail

## Premise
`Whitetail` is a lightweight alternative to ELK for non-intensive applications. It is built with containerization in mind for a non-resource intensive log viewer with basic metadata search capabilites.

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
- Release 1.0
- [ ] Add content to home page
- [ ] Add analytics page
- [ ] Add Log backup capabilities/config option
- [ ] Add Log backup to S3 capabilities/config option
- [ ] Any other tasks I'll inevitably think of later

## Antler Query Language (AQL)

AQL is a simple query language designed to be used when the standrard filtering (level and service) is not sufficient.  AQL statments are written in nested blocks of binary operations. This means that each operator can only have a singular left and singular right argument. An example query which gets logs of level `INFO` and level `WARN` is as follows:

```
@level:INFO OR @level:WARN
```

If you want to change the 'OR' statements to include more than just the two levels, you'll wrap the first two up in parenthesis and then OR that with a third filter.

```
( @level:INFO OR @level:WARN ) OR @level:DEBUG
```

### Filters
The various filters that can be used in AQL statements are as follows (`< text like this is a placeholder >`):
Filter                 | Desrciption
-----------------------|------------
`level = < level >`     | Get logs with level `< level >` (`= < level >` can be replaced with `IN < csv of levels >`)
`service = < service >` | Get logs from service `< service >` (`= < service >` can be replaced with `IN < csv of services >`)
`year = < year >`       | Get logs wtih a timestamp that has the year `< year >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`month = < month >`     | Get logs wtih a timestamp that has the year `< month >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`day = < day >`         | Get logs wtih a timestamp that has the year `< day >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`hour = < hour >`       | Get logs wtih a timestamp that has the year `< hour >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`minute = < minute >`   | Get logs wtih a timestamp that has the year `< minute >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)
`second = < second >`   | Get logs wtih a timestamp that has the year `< second >` (`=` can be replaced with `<`, `<=`, `>=`, `>`, or `!=`)

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
- __ORDER_ASCEND__
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
- __ORDER_DESCEND__
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
        "sqlite": {
            "path": "./data/whitetail.db"
        }
    },
    "logging": {
        "max-age-days": 2,
        "poll-rate": "1h",
        "concise-logger": true
    },
    "branding": {
        "primary_color": {
            "background": "#C3C49E",
            "text": "#000000"
        },
        "secondary_color": {
            "background": "#8F7E4F",
            "text": "#000000"
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
    "sqlite": {
        "path": "./data/whitetail.db"
    }
}
```
`Sqlite` configuration setting information is as follows

Name   | Description
-------|------------
`path` | The local location to create the `Sqlite` db file

To configure `Whitetail` to use `PostgreSQL`, set the `database` section to this:

```json
{
    "postgres": {
        "host": "localhost",
        "port": 5432,
        "username": "postgres",
        "password": "foobar",
        "database": "whitetail"
    }
}
```

As a note, the database must be created prior to `Whitetail` attempting to connect to it

`PostgreSQL` configuration setting information is as follows

Name       | Description
-----------|------------
`host`     | The hostname of the `PostgreSQL` instalce
`port`     | The port that `PostgreSQL` is running on
`username` | The username for the database
`password` | The password for the database
`database` | The name of the database to use

### Logging

Logging configuration mainly concerns itself with the cleanup process to remove old logs, however it does also configure some aspects of the log message formatting.

Name             | Description
-----------------|------------
`max-age-days`   | How many days to keep logs for (integer)
`poll-rate`      | How often to check for old logs. Is of the form `< number >< time unit >` where valid time units are `ns`, `us`, `ms`, `s`, `m`, `h`
`concise-logger` | Should the logger name be compaceted for ease of viewing (bool)

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
To persist your logs when running in a containerized environment you have a couple of options. You can either run `Whitetail` with an external (persisted) PostgreSQL database, or you can mount a volume at `/whitetail/data` in your container and persist that (if you are using a Sqlite database)

## Contact
`Whitetail` is written by John Carter
If you have any questions or concerns, feel free to contact me at jfcarter2358.at.gmail.com