# Install as Service
The following install process is done on an Ubuntu 20.04 LTS using directions provided by [TecAdmin](https://tecadmin.net/how-to-install-apache-kafka-on-ubuntu-20-04/#:~:text=1%20How%20to%20Install%20Apache%20Kafka%20on%20Ubuntu,Step%205%20%E2%80%93%20Create%20a%20Topic%20in%20Kafka) (last updated February 2022)

1. Update + Install Java
```shell
sudo apt-get -y update 
sudo apt-get -y install default-jdk
java --version 
```

2. From [Apache](https://kafka.apache.org/downloads) download the latest version of _Kafka_ & untar 
```shell
wget https://dlcdn.apache.org/kafka/3.1.0/kafka_2.13-3.1.0.tgz
tar -xzvf kafka_2.13-3.1.0.tgz
```

3. Move new kafka directory into `/usr/local/kafaka`
```shell
sudo kafka_2.13-3.1.0 /usr/local/kafka
```

4. Create [Zookeeper](https://zookeeper.apache.org/) service file - centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services. 
```yaml
# file name: /etc/systemd/system/zookeeper.service
[Unit]
Description=Apache Zookeeper server
Documentation=http://zookeeper.apache.org
Requires=network.target remote-fs.target
After=network.target remote-fs.target

[Service]
Type=simple
ExecStart=/usr/local/kafka/bin/zookeeper-server-start.sh /usr/local/kafka/config/zookeeper.properties
ExecStop=/usr/local/kafka/bin/zookeeper-server-stop.sh
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
```
5. Create Kafka service file

   * Locate `java-${VERSION}-openjdk-${PLATFORM}` directory
```bash
anylog@anylog-2004:~$ update-alternatives --list java
/usr/lib/jvm/java-11-openjdk-amd64/bin/java
```
   * Create service file - make sure to update JAVA_HOME path
```yaml

# file name: /etc/systemd/system/kafka.service
[Unit]
Description=Apache Kafka Server
Documentation=http://kafka.apache.org/documentation.html
Requires=zookeeper.service

[Service]
Type=simple
Environment="JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64" #<-- line should be updated based on the location of java
ExecStart=/usr/local/kafka/bin/kafka-server-start.sh /usr/local/kafka/config/server.properties
ExecStop=/usr/local/kafka/bin/kafka-server-stop.sh

[Install]
WantedBy=multi-user.target
```

6. Apply changes & update
```bash
sudo systemctl daemon-reload
sudo apt-get -y update
```

7. Start & validate Zookeeper
```bash
sudo systemctl start zookeeper
sudo systemctl status zookeeper
```