curl -X POST http://127.0.0.1:48080/api/v1/valuedescriptor -H 'Content-Type: text/plain' \
--data '{"name": "humidity", "description": "Ambient humidity in percent", "min": "0",
         "max": "100", "type": "Int64", "uomLabel": "humidity", "defaultValue": "0",
	 "formatting": "%s", "labels": ["environment","humidity"]}'
	
