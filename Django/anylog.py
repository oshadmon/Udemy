import json
import requests



class AnyLogConnection:
    """
    Missing functions:
        - get query log
        - get rest
        - get rest log
        - query status
        - get streaming
        - get operator
        - get rows count [where dbms group]
    """
    def __init__(self, conn:str, auth:tuple=(), timeout:int=30):
        """
        Connection information for AnyLog
        :params:
            conn:str - REST IP & Port
            auth:tuple - Username/password (optional)
            timeout:int - timeout
        """
        self.conn = conn
        self.auth = auth
        self.timeout = timeout

    def get_cmd(self, command:str, query:bool=False)->str:
        """
        Execute GET command
        :args:
            command:str - command to execute:
            query:bool - whether or not command is of type query (ie remote)
        :params:
            output:str - content to print
            headers:dict - REST headers information
        :return:

        """
        headers = {
            'command': command,
            'User-Agent': 'AnyLog/1.23'
        }

        if query is True:
            headers['destination'] = 'network'

        try:
            r = requests.get(url='http://%s' % self.conn, headers=headers, auth=self.auth, timeout=self.timeout)
        except Exception as e:
            output = "Failed to execute '%s' against '%s' (Error: %s)" % (headers['command'], self.conn, e)
        else:
            if int(r.status_code) == 200:
                try:
                    output = r.json()
                except Exception as e:
                    try:
                        output = r.text
                    except Exception as e:
                        output = "Failed to extract results for '%s' (Error: %s)" % (headers['command'], e)
            else:
                output = "Failed to execute '%s' against '%s' (Network Error: %s)" % (headers['command'], self.conn,
                                                                                      r.status_code)
        if 'Content-type: text/json' in output:
            output = output.split('Content-type: text/json')[1]
        elif 'Content-type: text' in output:
            output = output.split('Content-type: text')[1]

        return output

    def get_status(self)->str:
        """
        Execute `get status` command
        :return:
            raw results from get_cmd
        """
        return self.get_cmd(command='get status')

    def get_event_log(self)->str:
        """
        Execute `get event log` command
        :params:
            resuls:str - raw content from get_cmd
            output:str - results to print on screen
        :return:
            return results to print on screen
        """
        results = self.get_cmd(command='get event log where format=json')
        result = 'ID: %s | Time: %s | Type: %s | Text: %s'
        output = ''
        for row in json.loads(results):
            output += result % (row['ID'], row['Time'], row['Type'], row['Text']) + '<br/>'

        return output

    def get_error_log(self)->str:
        """
        Execute `get status` command
        :params:
            resuls:str - raw content from get_cmd
            output:str - results o print on screen
        :return:
            return results to print on screen
        """
        results = self.get_cmd(command='get error log where format=json')
        result = 'ID: %s | Time: %s | Type: %s | Text: %s'
        output = ''
        for row in json.loads(results):
            output += result % (row['ID'], row['Time'], row['Type'], row['Text']) + '<br/>'

        return output



