# RailwayStationsWorkload
**Attempt to migration from monolith server to closely coupled services, where gateway make calls to services via protobuf clients**

*Now requests to api will cause saving excel file with workload for requested stations filtered by days of the week*

Graphic representaion of service structure
![alt text](https://github.com/Rubanik-Alexei/RailwayStationsWorkload/blob/microVersion/graph.jpg)
- Gateway service parsing form from request, checking for errors like unemptiness and correct tags and then sends gRPC request to corresponding service:
  - if "/wl" request - sends request to workload service
  - if "/db" request - sends request to redis service
- Workload service actually scrapping workload for requested stations and returning collected data to gateway and optionally also sending it as gRPC request to redis service.
- Redis service searching data in Redis for requested stations in "/db" request and returning it to gateway and storing data from workload service requests into Redis.
