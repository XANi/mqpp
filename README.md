== MQPP ==

MQ pretty printer, for your MQ debugging needs.

Compilation: `make` for release version, `go get -u github.com/XANi/mqpp` to just get the latest

Usage:

    mqpp --mqtt-url tcp://mqtt:mqtt@hostname
    mqpp --amqp-url amqp://guest:guest@hostname
    mqpp --mqtt-url tcp://mqtt:mqtt@hostname --topic 'events/#'

AMQP `.` are replaced with `/` to make display of it be consistent with the MQTT. 
It is also possible to subscribe to both: 

    ./mqpp --amqp-url amqp://admin:admin@cthulhu --mqtt-url tcp://mqtt:mqtt@cthulhu:1883 --topic 'events/state/in/mpower_socket1' 
    [N] main[main.func1] For usage run with --help
    [N] main[Get.func1] connected to amqp:amqp://admin:admin@cthulhu
    [N] main[Get.func1] connected to mqtt:tcp://mqtt:mqtt@cthulhu:1883
    ^---                                               : connected to AMQP, consuming messages
    events/state/in/mpower_socket1                     : ON
    [amq.topic]/events/state/in/mpower_socket1         : ON
    [amq.topic]/events/state/in/mpower_socket1         : ON
    events/state/in/mpower_socket1                     : ON
    events/state/in/mpower_socket1                     : ON
    [amq.topic]/events/state/in/mpower_socket1         : ON

Add `--time-format` (using golang time formatting `Mon Jan 2 15:04:05 MST 2006`), or put `iso` for standard time format or `ts` for just the hour/minute/second part.
Note that 

Note that MQTT does **NOT** have timestamps in actual message so it will just be TS of receiving the message.

And in case of AMQP it is **OPTIONAL** so if app doesn't set it it will also just be a delivery date

Timestams that came from message are marked *green* while ones added on message receive are *orange*