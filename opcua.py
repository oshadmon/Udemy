import argparse 
import datetime
import opcua
import time 


def connect_opcua(conn:str):
    """
    Connect to OPCUA client
    :args: 
        conn:str - OPCUA connection info 
    :param: 
        client:opcua.client.client.Client - connection to OPCUA 
        start:time.time.time - process start time 
        boolean:bool - whether to exit while 
        error_msg:list - record of error messages 
    :return:
        client 
    """
    client = None 
    start = time.time() 
    boolean = False 
    error_msg = [] 
    while boolean == False:
        try:
            client = opcua.Client("opc.tcp://%s/" % conn)
            client.connect()
        except Exeception as e: 
            if e not in error_msg: 
                print('%s - Failed to connect to OPCUA (Error: %s)' % (datetime.datetime.now(), e)) 
                error_msg.append(e) 
            if time.time() > (start * 3605):
                print('%s - FAiled to connect to OPCUA for over an hour' % datetime.datetime())
                status = True 
            else: 
                time.sleep(30) 
        else:
            boolean = True 

    return client


def disconnect_opcua(client)->bool:
    """
    Disconnect from OPCUA
    :args:
        client:opcua.client.client.Client - connection to OPCUA
    :param: 
        status:bool 
    :return:
        status 
    """
    status = True 
    try:
        client.disconnect()
    except Exception as e: 
        print('Faield to disconnect from OPCUA (Error: %s)' % e)
        status = False
    return status 


def get_datalogger(client, tags:list)->(dict, str):
    """
    Exttract data from logger
    :args: 
        client:opcua.client.client.Client - connection to OPCUA 
        tags:list - list of tags to get data from for OPCUA 
    :param: 
        data:dict - data from OPCUA
    :return: 
        data and timestamp
    """
    data = {'timestamp': datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')}
    for tag in tags: 
        try: 
            output = client.get_node("ns=4;s=%s" % tag)
        except Exception as e: 
            print('Failed to get data for %s (Error: %s)' % (tag, e))
            output = None
        if output is not None: 
            try: 
                data[tag] = output.get_value()
            except Exception as e: 
                print('Failed to get value for %s data (Error: %s)' % (tag, e))

    return data


def main(): 
    """
    The following is an example of how to pull data from an OPCUA. Tested against the Ai-Ops krt-DataLogger.service 
    :links: 
    --> Test Tools: https://opcfoundation.org/developer-tools/certification-test-tools/opc-ua-compliance-test-tool-uactt/
    --> OPCUA documentation: https://python-opcua.readthedocs.io/en/latest/
    --> Sample code: https://github.com/FreeOpcUa/python-opcua/tree/master/examples
    :args: 
        conn    OPCUA connection info   [default: 192.168.50.19:4840]
        tags    OPCUA list of tags      [sample list: FIC11_FB.fActualValue,FIC11_FB.fActualValue,FIC11_FB.fSetpointValue,FIC11_FB.FIC11.fOut]
   :param: 
        client:opcua.client.client.Client - connection to the OPCUA 
        tags:list - list of tags based based on user input 
        data:dict - data from DataLogger 
        timestamp:str - timestamp for data from DataLogger 
    """
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument('conn',             type=str,   default='192.168.50.19:4840', help='OPCUA connection info [default: 192.168.50.19:4840]') 
    parser.add_argument('tags',             type=str,   default=None,                 help='OPCUA list of tags [sample list: FIC11_FB.fActualValue,FIC11_FB.fActualValue,FIC11_FB.fSetpointValue,FIC11_FB.FIC11.fOut]')
    args = parser.parse_args() 
    if args.tags == None: 
        print('Tags cannot be an empty string') 
        exit(1) 

    try: 
        tags = list(args.tags.split(','))
    except Exception as e: 
        print('Failed to convert tags into a list of tags (Error: %s)' % e)
        exit(1) 

    client = connect_opcua(args.conn)
    print(get_datalogger(client, tags))
    disconnect_opcua(client) 


if __name__ == '__main__': 
    main() 
