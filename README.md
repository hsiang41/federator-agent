## federatorai-agent

### Agent of Federator.ai cloud services

Federatorai-agent is a corn job service, which is added a scheduled job. Current has 
transmitted metrics of the Prometheus job and Write a recommendation job.

### Scheduled job library

| Job Library name | Path (library path                    | Configuration                                       |
|------------------|:-------------------------------------:|:---------------------------------------------------:|
| datapipe_input   | /lib/inputlib/datapipe.so             | /etc/alameda/federatorai-agent/input/datapipe.toml  |
| datapipe_writer  | lib/outputlib/datapipe_recommender.so | /etc/alameda/federatorai-agent/output/datapipe.toml |


#### datapipe.so

Transmitted metrics from the prometheus to the Alameda server.

#### datapipe_recommender.so

Periodical query the resource recommendation from the Alameda API server and write into the local customer resource 
recommendation.

### Agent schedule configuration
```apple js
[log]
 set-logcallers = true
 output-level = "debug" # debug, info, warn, error, fatal, none

[input_jobs]
    [input_jobs.datapipe_input]
    name = "datapipe_input"
    schedule-spec = "*/5 * * * * ?"
    lib-path = "/lib/inputlib/datapipe.so"
    lib-configuration = "/etc/alameda/federatorai-agent/input/datapipe.toml"

[output_jobs]
    [output_jobs.datapipe_output]
    name = "datapipe_output"
    schedule-spec = "*/30 * * * * ?"
    lib-path = "lib/outputlib/datapipe_recommender.so"
    lib-configuration = "/etc/alameda/federatorai-agent/output/datapipe.toml"
```
**schedule-spec**="*/5 * * * * ?"  
Context format is "`Seconds` `Minutes` `Hours` `Day of mounth` `Month` `Day of week`".  

| name       | required | allowed value   | allowed special character |
|:-----------|:--------:|:---------------:|:-------------------------:|
|Seconds     | yes      | 0 - 59          | */,-                      |
|Minutes     | yes      | 0 - 59          | */,-                      |
|Hours       | yes      | 0 - 23          | */,-                      |
|Day of month| yes      | 1 - 31          | */,-?                     |
|Month       | yes      | 1-12 or JAN-DEC | */,-                      |
|Day of week | no       | 0-6 or SUN-SAT  | */,-?                     |  

Special Characters

Asterisk ( * )  
The asterisk indicates that the cron expression will match for all values of the field; e.g., using an asterisk in the 5th field (month) would indicate every month.

Slash ( / )  
Slashes are used to describe increments of ranges. For example 3-59/15 in the 1st field (minutes) would indicate the 3rd minute of the hour and every 15 minutes thereafter. The form "*\/..." is equivalent to the form "first-last/...", that is, an increment over the largest possible range of the field. The form "N/..." is accepted as meaning "N-MAX/...", that is, starting at N, use the increment until the end of that specific range. It does not wrap around.

Comma ( , )  
Commas are used to separate items of a list. For example, using "MON,WED,FRI" in the 5th field (day of week) would mean Mondays, Wednesdays and Fridays.

Hyphen ( - )  
Hyphens are used to define ranges. For example, 9-17 would indicate every hour between 9am and 5pm inclusive.

Question mark ( ? )  
Question mark may be used instead of '*' for leaving either day-of-month or day-of-week blank.  

**lib-path**  
Config the Cron job triggered library path.

**lib-configuration**  
Config the Cron job triggered library configration path.  


### How to build the docker image  

`make docker-build`

Output docker image name is "federatorai-agent:latest"  
