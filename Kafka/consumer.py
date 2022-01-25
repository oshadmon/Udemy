import argparse 
import json 
import kafka

def connect_consumer(topic:str, servers:str, auto_offset_reset:str='earliest', enable_auto_commit:bool=True, consumer_timeout_ms:int=1000, exception:bool=False)->kafka.consumer.group.KafkaConsumer:
    """
    Connect to a kafka consumer instance
    :urls: 
        full list of options: https://kafka-python.readthedocs.io/en/master/apidoc/KafkaConsumer.html 
    :args:
        topic - topic(s) to connect against
        servers:str - either a string or list of string that the consumer should contact to bootstrap initial cluster metadata
        auto_offset_reset:str - A policy for resetting offsets on OffsetOutOfRange errors: ‘earliest’ will move to the oldest available message, ‘latest’ will move to the most recent.
        enable_auto_commit:bool - If True , the consumer’s offset will be periodically committed in the background. Default: True.
        consumer_timeout_ms:int - Number of milliseconds to block during message iteration before raising StopIteration (i.e., ending the iterator). Default block forever [float(‘inf’)].
        exception:bool - whether or not to print exception 
    :params: 
        consumer:kafka.consumer.group.KafkaConsumer - connection to kafka consumer
    :return: 
        consumer
    """
    try: 
        consumer = kafka.KafkaConsumer(topic, bootstrap_servers=servers, auto_offset_reset=auto_offset_reset, enable_auto_commit=enable_auto_commit, consumer_timeout_ms=consumer_timeout_ms)
    except Exception as e: 
        if exception is True: 
            print(f'Failed to connect to consumer {servers} against topic {topic} (Error: {e})') 
        consumer = None 

    return consumer 


def read_messages(consumer, exception:bool=False): 
    """
    Read messages published to consumer
    :args:
        consumer:kafka.consumer.group.KafkaConsumer - connection to kafka consumer        
    :params:
        message:ConsumerRecord - record provided 
            value - value object from message
    :sample-message:
        ConsumerRecord(topic='test-topic', partition=0, offset=0, timestamp=1643069195255, timestamp_type=0, key=None, value=b'Welcome to kafka', headers=[], checksum=None, serialized_key_size=-1, serialized_value_size=16, serialized_header_size=-1)
    """
    for message in consumer: 
        try: 
            value = message.value
        except Exception as e: 
            if exception is True: 
                print(f'Failed to extract value from message')
        else: 
            send_to_anylog(value=value, exception=exception)


def send_to_anylog(value:bytes=None, exception:bool=False): 
    """
    Convert data from JSON to dict & print results
    :args: 
        value:bytes
    :params: 
        utf_value:str - converted value into UTF-8 format
        dict_value:dict - utf_value converted into a dictionary
    :print: 
        actual value & value type
    """
    # convert to UTF-8 format
    try: 
        utf_value = value.decode('utf-8')
    except Exception as e:
        print(f'Failed to convert value into UTF-8 (Error: {e}') 
    else: 
        try: 
            dict_value = json.loads(utf_value)
        except Exception as e: 
            print(utf_value, type(utf_value))
        else: 
            print(dict_value, type(dict_value)) 


def main(): 
    """
    Sample main for kafka consumer
    :url: 
        https://kafka-python.readthedocs.io/en/master/usage.html#kafkaconsumer 
    :positional arguments:
        topic                 topic to connect against
        servers               comma seperated servers to connect against
    :optional arguments:
        -h, --help            show this help message and exit
        --offset    {earliest,latest}    policy for resetting offsets on OffsetOutOfRange (default: earliest)
        --disable-auto-commit            if set, consumer’s offset will not be committed in the background (default: True)
        --timeout TIMEOUT                Number of milliseconds to block during message iteration before raising StopIteration (setting -1 will wait forever) (default: 1000)
        -e, --exception                  whether or not to print exception messages (default: False)
    :params: 
        consumer:kafka.consumer.group.KafkaConsumer - connection to consumer process 
        servers:list - list of servers based on args.servers 
    """
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument('topic', type=str, default='test-topic', help='topic to connect against') 
    parser.add_argument('servers', type=str, default='localhost:9092', help='comma seperated servers to connect against') 
    parser.add_argument('--offset', type=str, choices=['earliest', 'latest'], default='earliest', help='policy for resetting offsets on OffsetOutOfRange')
    parser.add_argument('--disable-auto-commit', type=bool, nargs='?', const=False, default=True, help='if set, consumer’s offset will not be committed in the background')
    parser.add_argument('--timeout', type=int, default=1000, help='Number of milliseconds to block during message iteration before raising StopIteration (setting -1 will wait forever)') 
    parser.add_argument('-e', '--exception', type=bool, nargs='?', const=True, default=False, help='whether or not to print exception messages') 
    args = parser.parse_args() 
    
    servers = args.servers.split(',') 
    
    # connect to consumer 
    consumer = connect_consumer(topic=args.topic, servers=servers, auto_offset_reset=args.offset, enable_auto_commit=args.disable_auto_commit, consumer_timeout_ms=args.timeout, exception=args.exception)

    # validate consumer connected successfully, if so read messages 
    if isinstance(consumer, kafka.consumer.group.KafkaConsumer):
        read_messages(consumer=consumer) 


if __name__ == '__main__': 
    main()
