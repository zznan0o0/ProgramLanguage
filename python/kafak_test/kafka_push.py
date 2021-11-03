
from kafka import KafkaProducer

producer = KafkaProducer()  # 连接kafka

msg = "Hello World".encode('utf-8')  # 发送内容,必须是bytes类型
producer.send('mytest', msg)  # 发送的topic为test
producer.close()

print("success")

# from kafka import KafkaConsumer

# consumer = KafkaConsumer('mytest')
# for msg in consumer:
#     print(msg)
#     recv = "%s:%d:%d: key=%s value=%s" % (msg.topic, msg.partition, msg.offset, msg.key, msg.value)
#     print(recv)
