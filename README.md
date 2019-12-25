# EFORE opus_snmp 


Read SNMP params, and returns a json object.

Programm flags:

-ip - controller ip address (defaut value "localhost");

-port - controller SNMP tcp port (defaut value 161);

-com - controller SNMP community  (defaut value Public);

-type - type of MIB sub-tree (defaut value 1):
- 1 - batteryControllerMeasurement (BCM) sub-tree
- 2 - batteryControllerOperation (BCO) sub-tree
- 3 - vidi sub-tree **(incomplete implementation for today)**
- 4 - batteryControllerAlarm sub-tree **(incomplete implementation for today)**

**Example:**

`opus_snmp -ip=192.168.10.10 -port=1161 -com=public -type=2`