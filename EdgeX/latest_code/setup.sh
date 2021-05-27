echo "setting up device profile"
curl http://localhost:48081/api/v1/deviceprofile/uploadfile -F "file=@modbus.test.device.profile.yml"
echo "setting up device"
curl http://localhost:48081/api/v1/device -H "Content-Type:application/json" -X POST \
  -d '{ 
   "name" :"Modbus TCP test device",
   "description":"This device is a product for monitoring and controlling digital inputs and outputs over a LAN.",
   "adminState":"UNLOCKED",
   "operatingState":"ENABLED",
   "protocols":{
      "modbus-tcp":{
         "Address" : "edgex-modbus-simulator",
         "Port" : "1502",
         "UnitID" : "1"
      }
   },
   "labels":[ 
      "air conditioner",
      "modbus TCP"
   ],
   "service":{"name":"edgex-device-modbus"},
   "profile":{"name":"Test.Device.Modbus.Profile"},
   "autoEvents":[ 
      { 
         "frequency":"10s",
         "onChange":false,
         "resource":"HVACValues"
      }
   ]
}'
echo "done"
