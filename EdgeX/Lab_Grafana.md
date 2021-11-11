# Grafana 

The following process is based the idea that [Exporting Data](Lab_Export.md#hivemqtt) in [docker-compose.yml](docker-compose_hivemqtt.yml)
## Install & Configure New 
1. Download & run [Mosquitto](https://mosquitto.org/) 
```
docker pull eclipse-mosquitto
docker run --name mosquitto -d -p 1883:1883 -p 9001:9001 eclipse-mosquitto
```

2. Download & Install [InfluxDB](https://www.influxdata.com/products/influxdb-cloud/?utm_source=bing&utm_medium=cpc&utm_campaign=2020-09-03_Cloud_Traffic_Brand-InfluxDB_NA&utm_term=influxdb) 
```
docker pull influxdb
docker run -d --name influxdb -p 8086:8086 -e INFLUXDB_DB=sensordata -e INFLUXDB_ADMIN_USER=root -e INFLUXDB_ADMIN_PASSWORD=pass -e INFLUXDB_HTTP_AUTH_ENABLED=true influxdb
```

3. Download & Install [Grafana](https://grafana.com/)
```
docker pull grafana/grafana
docker run -d --name=grafana -p 3000:3000 grafana/grafana
```
**Note**: Depending on firewall port 3000 may need to open 

4. From GitHub download [messenger](https://github.com/jonas-werner/EdgeX_Tutorial/tree/master/messenger) directory
```
git clone https://github.com/jonas-werner/EdgeX_Tutorial
cd EdgeX_Tutorial/messenger
```

5. Update _app.py_ with machine IP (don't use `127.0.0.1`) 
```
sed -i "s/<edgex ip>/139.162.205.95/g" app.py
```

6. Docker Build & Run messenger
```
docker build . -t messenger
docker run -d --name messenger messenger:latest
```

7. Check new containers are running 
```
for container in mosquitto influx grafana messenger ; do docker ps -a | grep ${container} ; done
```

## Redirect to Local EdgeX 
0. Go back to _HOME_ directory 
```
cd $HOME/edgex
```

1. Stop EdgeX if running 
```
docker-compose down -v 
```

2. Update Broker IP in _app-service-mqtt_  
```
sed -i "s/broker.hivemq.com/${YOUR_IP}/g" docker-compose.yml
```

3. Start EdgeX with changes 
```
docker-compose pull 
docker-compose up -d 
```

## Grafana 
1. Login to Grafana - http://127.0.0.1:3000 (username: admin | password: admin) 

2. Configure _Data Source_ to support Influx 
   * **URL**: http://127.0.0.1:8086   
   * **Database**: root 
   * **Password**: pass 

3. Create New _Dashboard_ & User the new _InfluxDB_ 
