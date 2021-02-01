# Whitetail

## Premise
Whitetail is a lightweight alternative to ELK for non-intensive applications. It is built with containerization in mind for a non-resource intensive log viewer with basic search capabilites.

## TO DO
- [ ] Add settings page
- [x] Add docker file
- [x] Add kubernetes manifest
- [ ] Add query language so that it's not just a keyword search on the logs page
- [x] Add a log refresh button on the logs page
- [x] Add filtered by log level
- [ ] Add log age cleanup
- [ ] Add configurable branding
- [ ] Add configuration section to README
- [ ] Any other tasks I'll inevitably think of later

## Container Log Persistence
To persist your logs when running in a containerized environment you have a couple of options. You can either run Whitetail with an external (persisted) PostgreSQL database, or you can mount a volume at `/whitetail/data` in your container and persist that (if you are using a Sqlite database)

## Contact
Whitetail is written by John Carter
If you have any questions or concerns about Whitetail, feel free to contact me at jfcarter2358.at.gmail.com