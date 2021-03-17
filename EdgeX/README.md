# EdgeX 

The following are notes based on the [EdgeX Tutorial](https://jonamiki.com/wp-content/uploads/2020/08/EdgeX-Foundry-tutorial-ver1.1.pdf) 

## Version
* [EdgeX v1.2 Geneva](https://docs.edgexfoundry.org/1.2/getting-started/quick-start/) - LTS Version used for Lab 
* [EdgeX v2.0 Ireland](https://docs.edgexfoundry.org/2.0/getting-started/quick-start/) - Latest version

## Docker Compose Files
* [docker-compose_step1.yml](docker-compose_step1.yml) - Original Docker-Compose file 
* [docker-compose_hivemq.yml](docker-compose_hivemq.yml) - Docker-Compose file working with HiveMQTT (Topic: `anylogedgex`) 
* [docker-compose_cloudmqtt.yml](docker-compose_cloudmqtt.yml) - Docker-Compose file working with CloudMQTT (Topic: `anylogedgex`) 

## Sections 
* [Install & Setup](Lab_Setup.md) 
* [Add GUI](Lab_GUI.md) 
* [Add Device](Lab_Device.md)
* [Add Random Data Generator](Lab_Random.md) 
* [Export Data](Lab_Export.md) 
* [Grafana](Lab_Grafana.md) 

## Notes
* [Lab Materials](https://github.com/jonas-werner/EdgeX_Tutorial)
* The original lab uses a RPI with DHT11 sensor

