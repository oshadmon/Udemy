# Kuiper
[EMQ X Kuiper](https://github.com/emqx/kuiper) - SQL based rule engine, integrated with EdgeX. 
Kuiper is an edge lightweight IoT data analytics / streaming software implemented by Golang, and it can be run at all kinds of resource constrained edge devices. 
* **Source**: The data source of streaming data, such as data from MQTT broker. In EdgeX scenario, the data source is EdgeX message bus, which could be ZeroMQ or MQTT broker.
* **SQL**: SQL is where you specify the business logic of streaming data processing. Kuiper provides SQL-like statements to allow you to extract, filter & transform data.
* **Sink**: Sink is used for sending analysis result to a specified target. For example, send analysis result to another MQTT broker, or an HTTP rest address.


## Links
* Project is by [EMQX](https://www.emqx.io/) 

