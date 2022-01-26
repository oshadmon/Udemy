import argparse
import kafka


def connect_admin(servers:str, exception:bool=False)->kafka.admin.client.KafkaAdminClient:
    """
    Connect to Kafka admin consul
    :args:
        servers:str - producer server
        exception:bool - whether or not to print exceptions
    :params:
        admin:kafka.admin.client.KafkaAdminClient - connection to Kafka admin
    :return:
        admin
    """
    try:
        admin = kafka.KafkaAdminClient(bootstrap_servers=servers)
    except Exception as e:
        admin = None
        if exception is True:
            print(f'Failed to connect to admin against servers {servers} (Error: {e})')

    return admin


def delete_topic(admin:kafka.admin.client.KafkaAdminClient, topics:list, timeout:int, exception:bool=False):
    for topic in topics:
        try:
            admin.delete_topics(topic=topic, timeout_ms=timeout)
        except Exception as e:
            if exception is True:
                print(f'Failed to drop topic: {topic} (Error: {e}')


def main():
    """
    Code to remove topic(s) from Kafka
    :url:
        https://kafka-python.readthedocs.io/en/master/apidoc/KafkaAdminClient.html
    :positional arguments:
        servers               comma separated servers to connect against
        topics                comma separated list of topics to drop
    :optional arguments:
        -h, --help            show this help message and exit
        --timeout   TIMEOUT     Milliseconds to wait for topics to be deleted before the broker returns (default: None)
        -e, --exception     EXCEPTION   whether or not to print exception messages (default: False)
    :params:
        admin:kafka.admin.client.KafkaAdminClient - connection to Kafka admin
        servers:list - converted args.servers to list
        topics:list - converted args.topics to list
    """
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument('servers', type=str, default='localhost:9092', help='comma separated servers to connect against')
    parser.add_argument('topics', type=str, default='test-topic', help='comma separated list of topics to drop')
    parser.add_argument('--timeout', type=int, default=None, help='Milliseconds to wait for topics to be deleted before the broker returns')
    parser.add_argument('-e', '--exception', type=bool, nargs='?', const=True, default=False, help='whether or not to print exception messages')
    args = parser.parse_args()

    servers = args.servers.split(",")
    topics = args.topics.split(",")

    admin = connect_admin(servers=servers, exception=args.exception)
    if isinstance(admin, kafka.admin.client.KafkaAdminClient):
        delete_topic(admin=admin, topics=topics, timeout=args.timeout, exception=args.exception)


if __name__ == '__main__':
    main()