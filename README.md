<div id="top"></div>

<!-- Service Demo in Go and Python -->
<br />
<div align="center">
  <a href="[https://github.com/Sholontla](https://github.com/Sholontla/demo-service/edit/main](https://github.com/Sholontla/demo-service/edit/main/)">
  </a>

<h3 align="center">Service Demo in Go and Python</h3>

  <p align="center">
   This Structure Demo is use for testing and demostration to understad Demo Structure Service please visit the diagram: https://github.com/Sholontla/demo-service/blob/main/demo-service-diagram.pdf 
    <br />
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About The Project

The Main Project is structure by:
Golang as main programming language
Python as HTTP client to make request to the service prodcucer.


<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

this service is made with Python.
  . python_http_client (only th python http_client)
  
- Golang
this services are dev with Go.
 . admin
 . files_server (batch file processing)
 . logger (logger service with Prometheus, grafana, promtail, loki)
 . producer (producer and gRPC client)
 . service_consumer1 (producer and gRPC server 1)
 . service_consumer2 (consumer and gRPC server 2)
 . service_consumer3 (consumer and gRPC server 3)
 . service_consumer4 (consumer and gRPC server 4)
 . service_consumer5 (consumer and gRPC server 5)

- tools / frameworks
  . fiber
  . gRPC 
- Docker
- Docker - Compose

O/I
- Windows
- Linux

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Services:

1. End Points:

### Endpoints

- localhost:8004/service/producer/register1
- localhost:8004/service/producer/register2
- localhost:8004/service/producer/register3
- localhost:8004/service/producer/register4
- localhost:8004/service/producer/register5

2. Data Structure:
{
  "producer_id": "producer1",
  "producer_message": {
      "producer_message_id": "producer1",
      "host": "producer1",
      "client": "producer1",
      "ip": "producer1",
      "port": "producer1",
      "data_producer": {
          "data_producer_id": "producer1",
          "product": "producer1",
          "name": "producer1",
          "category": "producer1",
          "sub_category": "producer1",
          "price": 12.32,
          "quantity": 3,
          "supplier": "producer1",
          "description": "producer1",
          "gender": "producer1"
      },
      "information_created_at": "producer1"
  },
  "producer_service_area": "producer1",
  "producer_created_at": "producer1"
}

. What is missing:
  . the HTTP server to send the certs for gRPC client and server
  . Makfile tu assembly the demo service
  . Admin service to manage the structure project
  . Add to the strcuture the Batch fileServer
  . Add to the strcucture the cloud service
  
## License

For testing and demostrations purposes.

<!-- CONTACT -->

## Contact

Gerardo Ruiz Bustani - solbustani@gmail.com - 442 488 6193

Project Link: https://github.com/Sholontla/demo-service/edit/main/
