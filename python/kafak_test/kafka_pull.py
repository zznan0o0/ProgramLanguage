# from kafka import KafkaConsumer

# consumer = KafkaConsumer('mytest')
# for msg in consumer:
#     print(msg)
#     recv = "%s:%d:%d: key=%s value=%s" % (msg.topic, msg.partition, msg.offset, msg.key, msg.value)
#     print(recv)

from kafka import KafkaConsumer
import time

def start_consumer():
    consumer = KafkaConsumer('mytest', bootstrap_servers = ['localhost:9092'])
    print(1111)
    for msg in consumer:
        print(msg)
        print("topic = %s" % msg.topic) # topic default is string
        print("partition = %d" % msg.offset)
        print("value = %s" % msg.value.decode()) # bytes to string
        print("timestamp = %d" % msg.timestamp)
        print("time = ", time.strftime("%Y-%m-%d %H:%M:%S", time.localtime( msg.timestamp/1000 )) )


start_consumer()