# Devices 

## Background info
* **Definition**: A  device could be any type of edge appliance which is generating or forwarding data
* Registering a device is required because
   * Become aware of the existence of the device
   * Able to recieve data 
   * Able to control device (if supported)

## Sensor Cluster 
* A device container one or more sets of sensors. The example is for a [DHT11](https://learn.adafruit.com/dht/overview) Temp/Humidity sensor

1. Check what devices and sensors are already connected
```
curl http://127.0.0.1:48080/api/v1/valuedescriptor | jq
```

2. Add value description for humidty sensor 
```
curl -X POST http://127.0.0.1:48080/api/v1/valuedescriptor -H 'Content-Type: text/plain' \
--data '{"name": "humidity", "description": "Ambient humidity in percent", "min": "0",
         "max": "100", "type": "Int64", "uomLabel": "humidity", "defaultValue": "0",
	 "formatting": "%s", "labels": ["environment","humidity"]}'
```

3. Add value description for temperature sensor 
```
curl -X POST http://127.0.0.1:48080/api/v1/valuedescriptor -H 'Content-Type: text/plain' \
--data '{"name": "temperature", "description": "Ambient temperature in Celsius", "min": "-50",
	 "max": "100", "type": "Int64", "uomLabel": "temperature", "defaultValue": "0",
	 "formatting": "%s", "labels": ["environment", "temperature"]}'
```

4. Add a new device 
```
curl -X POST http://127.0.0.1:48081/api/v1/device -H 'Content-Type: text/plain' \
--data '{"name": "Temp_and_Humidity_sensor_cluster_01", "description": "Raspberry Pi sensor cluster", 
         "adminState": "unlocked", "operatingState": "enabled", "protocols": { "example": {
         "host": "dummy", "port": "1234", "unitID": "1"}}, "labels": ["Humidity sensor", 
         "Temperature sensor", "DHT11"], "location": "Tokyo", "service": {
         "name": "edgex-device-rest"}, "profile": {"name": "SensorCluster"}}'

```

**Note**: Instead of running these commands, user can run [device creation scripts](https://github.com/jonas-werner/EdgeX_Tutorial/tree/master/deviceCreation) 
```
python3 createSensorCluster.py -ip 127.0.0.1

Result for createValueDescriptors #1: <Response [200]> - Message: b62f3e3b-e90a-4a05-be1f-a1c8c4cf2df9
Result for createValueDescriptors #2: <Response [200]> - Message: fb27d3ea-407a-46b7-b36c-981761171bbc
Result of uploading device profile: <Response [200]> with message 10263a43-99fc-4d13-9809-8cb64a6f1eb1
Result for addNewDevice: <Response [404]> - Message: no item found for supplied key: invalid device service: edgex-device-rest 
```
