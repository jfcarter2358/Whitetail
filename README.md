# Whitetail

## Premise
Whitetail is a lightweight alternative to ELK for non-intensive applications. It is built with containerization in mind for a non-resource intensive log viewer with basic search capabilites.

## TO DO
- [x] Add settings page
- [x] Add docker file
- [x] Add kubernetes manifest
- [x] Add query language so that it's not just a keyword search on the logs page
- [x] Add a log refresh button on the logs page
- [x] Add filtered by log level
- [ ] Add log age cleanup
- [x] Add configurable branding
- [x] Add 'Reset to Default' buittons in settings page
- [ ] Add configuration section to README
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

#### Filters
The various filters that can be used in AQL statements are as follows (`< text like this is a placeholder >`):
Filter | Desrciption
---|---
`@level:< level >` | Get logs with level `< level >`
`@service:< service >` | Get logs from service `< service >`
`@year:< year >` | Get logs wtih a timestamp that has the year `< year >`
`@month:< month >` | Get logs wtih a timestamp that has the year `< month >`
`@day:< day >` | Get logs wtih a timestamp that has the year `< day >`
`@hour:< hour >` | Get logs wtih a timestamp that has the year `< hour >`
`@minute:< minute >` | Get logs wtih a timestamp that has the year `< minute >`
`@second:< second >` | Get logs wtih a timestamp that has the year `< second >`
`< word >`  | Get logs whose message includes `< word >`
`@all` | Get all logs

#### Operators
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
        - Level
        - Service
        - Text
        - Timestamp
- __ORDER_DESCEND__
    - `< left filter > ORDER_DESCEND < field >`
    - Orders the results of the left filter in descending order by one of the following fields
        - Level
        - Service
        - Text
        - Timestamp

## Container Log Persistence
To persist your logs when running in a containerized environment you have a couple of options. You can either run Whitetail with an external (persisted) PostgreSQL database, or you can mount a volume at `/whitetail/data` in your container and persist that (if you are using a Sqlite database)

## Contact
Whitetail is written by John Carter
If you have any questions or concerns about Whitetail, feel free to contact me at jfcarter2358.at.gmail.com