# service-monitor

The service-monitor is a service to check the health status of other services using REST APIs. </br>

## Prerequisites
Golang v1.11 or higher </br>
Postgres V9.3 or higher

## Configurations
All the configurable parameters are exposed with the json file available in the path `github.com/arjunajithtp/service-monitor/config/config.json` </br>

The config file path need to be exported to the environment variable `CONFIG_FILE_PATH` </br>

| Variables               | Type   | Description                                                   |
|-------------------------|--------|---------------------------------------------------------------|
| port                    | TEXT   | Web port to expose service-monitor application                |
| monitoringIntervalInMin | NUMBER | The interval in minutes to check the availability of services |
| dbHost                  | TEXT   | DB host                                                       |
| dbPort                  | TEXT   | DB password                                                   |
| dbName                  | TEXT   | DB Name                                                       |
| dbUserName              | TEXT   | Username to access the DB                                     |
| dbPassword              | TEXT   | Password to access the DB                                     |
| Services                | LIST   | The list of services to be monitored                          |

## Report Generation
The report regarding the services health can be generated in JSON form using the `get-status` end point.</br>
The type of generated report can be controlled by passing appropriate url query parameters</br>
### examples: 
`/get-status?fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&status=unavailable`</br>
`/get-status?fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&timeTaken=greater`

## URL Query Parameters

| Parameters | Type     | Description                                                                                                                                                                                                                                                                                                                          |
|------------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| fromDate   | Required | The `from date` in the format yyyy-mm-dd                                                                                                                                                                                                                                                                                               |
| fromTime   | Required | The `from time` in the format HH:MM:SS                                                                                                                                                                                                                                                                                                 |
| toDate     | Required | The `to date` in the format yyyy-mm-dd                                                                                                                                                                                                                                                                                                 |
| toTime     | Required | The `to time` in the format HH:MM:SS                                                                                                                                                                                                                                                                                                   |
| status     | Optional | Flag to generate the report with health status,<br>which is the default report type.<br>It is of type ENUM with accepted values `available` and `unavailable`.<br>available: generate a report for all the available services<br>unavailable: generate a report for all the unavailable services. <br>The default value is available.    |
| timeTaken  | Optional | Flag to generate the report with time take to respond by the available services.<br>It is of type ENUM with accepted values `greater` and `less`.<br>greater: generate a report for all the services whose response time is more than 1 second<br>less: generate a report for all the services whose response time is less than 1 second |

