# WEBSITE CHECKER
This Program was written for bootcamp project. It has two APIs:

* GET /websites : Optionally receives list of websites and returns status of some or all websites stores in memory
* POST /websites : Receives list of websites and stores them in memory for checking status

Additionally, there is a group of workers with configurable size that continuously poll these websites and check their status
All Configurations are stored in properties.yml file.
Control+C can be used to exit at program. Program runs on 127.0.0.1:8080 as per configs

Website Names should not include http:// or https://

## STATUS TYPES

* INIT : Website is added to map but not polled yet
* UP : Website is polled and is up
* DOWN : Website is polled and is up
* INVALID_URI : Website URI is invalid and not polled at all
* DOES_NOT_EXIST : Website is not present in map


## MAKEFILE COMMANDS

* make build
* make run
* make test
* make codeCoverage