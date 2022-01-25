import argparse
import kafka 
from kafka.errors import KafkaError


def connect_producer(servers:str, exception:bool=False)->kafka.producer.kafka.KafkaProducer:
    """
    Connect to Kafka producer
    :args: 
        servers:str - producer server
        exception:bool - whether or not to print exceptions
    :params: 
        producer:kafka.producer.kafka.KafkaProducer - connection to producer
    :return: 
        producer
    """
    try: 
        producer = kafka.KafkaProducer(bootstrap_servers=servers)
    except Exception as e: 
        if exception is True: 
            print(f'Failed to connect to producer {servers} (Error: {e})')
        producer = None 
        
    return producer 


def publish_data(producer:kafka.producer.kafka.KafkaProducer, topic:str, content:str, exception:bool=False): 
    """
    Publish data to kafka
    :args: 
        producer:kafka.producer.kafka.KafkaProducer - connection to producer
        topic:str - topic to publish against
        content:str - content to publish
        exception:bool - whether or not to print exceptions
    :params: 
        future:kafka.producer.future.FutureRecordMetadata - publish result 
        record_metadata:kafka.producer.future.RecordMetadata 
    """
    # convert string to bytes
    try: 
        encode_content = content.encode()
    except Exception as e: 
        if exception is True: 
            print(f'Failed to encode {content} (Error: {e}') 
    else: 
        try: 
            future = producer.send(topic, value=encode_content)
        except KafkaError as e: 
            if exception is True: 
                print(f'Failed to publish against topic {topic} (Error: {e})') 
        else: 
            try:
                record_metadata = future.get(timeout=100)
            except KafkaError as e:
                if exception is True: 
                    print(f'Failed to record metadata (Error: {e})')
            else: 
                if isinstance(record_metadata, kafka.producer.future.RecordMetadata) is True: 
                    print('Success') 

def main(): 
    """
    Sample main for kafka publisher
    :url: 
        https://kafka-python.readthedocs.io/en/master/usage.html#kafkaproducer
    :positional arguments:
        servers               comma seperated servers to connect against
        topic                 topic to connect against
        content               content to publish
    :optional arguments:
        -h, --help            show this help message and exit
        -e, --exception       whether or not to print exception messages (default: False) 
    :params: 
        producer:kafka.producer.kafka.KafkaProducer - connection to producer
        servers:list - list of servers based on args.servers 
    """
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument('servers', type=str, default='localhost:9092', help='comma seperated servers to connect against') 
    parser.add_argument('topic', type=str, default='test-topic', help='topic to connect against') 
    parser.add_argument('content', type=str, default=None, help='content to publish') 
    parser.add_argument('-e', '--exception', type=bool, nargs='?', const=True, default=False, help='whether or not to print exception messages') 
    args = parser.parse_args() 
    
    servers = args.servers.split(',') 
    
    # connect to producer
    producer = connect_producer(servers=servers, exception=args.exception)
    if isinstance(producer, kafka.producer.kafka.KafkaProducer) and args.content is not None: 
        publish_data(producer=producer, topic=args.topic, content=args.content, exception=args.exception) 


if __name__ == '__main__':
    main() 


