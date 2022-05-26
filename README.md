# GoRabbitMG
## Turorial Reference 
[RabbitMQ Tutorials](https://www.rabbitmq.com/getstarted.html)


This tutorial assumes RabbitMQ is [installed](https://www.rabbitmq.com/download.html) and running on localhost on the standard port (5672). In case you use a different host, port or credentials, connections settings would require adjusting.

### Docker script to run RabbitMQ in local 
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER={username} -e RABBITMQ_DEFAULT_PASS={password} rabbitmq:3.10-management
