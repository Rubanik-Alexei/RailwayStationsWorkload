# RailwayStationsWorkload
**Version of server that allows to make calls from browser**
The developing of such thing was inspired to help my friend with their research.

**At this point there are two methods available:**

-Scrapping workload for required station(s) (if the info about workload is provided) with option to store the collected data to redis(With possible idea to store workload with specifying day when it was added and stored for future analyzing as time series).

-Retrieving values from Redis itself for required stations.

**For now this is just a single monolith server**
