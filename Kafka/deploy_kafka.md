# Deploy Kafka 

## Command Line 
0. Go to `/usr/local/kafka/`
```bash
cd /usr/local/kafka/
```

1. create topic with a single partition with single replica
   * Partition - the number of brokers to split the data into 
   * Replica - the number of copies of the data to be created
```bash 
bin/kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
```

2. View list of topics
```bash
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

3. Initiate producer 
```bash
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic  test-topic
```

4. In a new terminal, initiate consumer
```bash
cd /usr/local/kafka
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic  test-topic --from-beginning
```

## Python 

The Python code uses [kafka-python](https://kafka-python.readthedocs.io/) 
```bash
# via pip
pip3 install kafka-python
# via apt
Sudo apt-get -y install python3-kafka
```

* [consumer](consumer.py) - this is the code that reads content sent into Kafka
```bash 
python3 consumer.py test-topic localhost:9092 --timeout -1 -e
```
* [producer](producer.py) - this is the code that writes content into Kafka
```bash 
python3 producer.py localhost:9092 test-topic '{"hello": "world"}' -e
```
