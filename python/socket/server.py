import socket
import time
import json

HOST = ''
PORT = 9999
BUFSIZ = 1024
ADDR = (HOST, PORT)

#创建套接字
tcpSerSock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
tcpSerSock.bind(ADDR)
#监听端口，最多5人排队
tcpSerSock.listen(5)

print('tcp listem on 9999')

while True:
    print('waiting for connection...')
    # 建立连接
    tcpCliSock, addr = tcpSerSock.accept()  
    print('...connected from:', addr)

    while True:
        data = tcpCliSock.recv(BUFSIZ)
        if not data:
            break

        content = '[%s] %s' % (bytes(time.ctime(), "utf-8"), data)

        print(data)
        print(type(content))
        print(json.loads(data))
        tcpCliSock.send(content.encode("utf-8"))
    
    tcpCliSock.close()


# tcpSerSock.close()