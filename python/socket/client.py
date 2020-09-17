import socket

HOST = '127.0.0.1'  # or 'localhost'
PORT = 9999
BUFSIZ = 1024
ADDR = (HOST, PORT)

tcpCliSock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
tcpCliSock.connect(ADDR)

while True:
    data = input('> ')
    if not data:
        break

    tcpCliSock.send(data.encode("utf-8"))
    data = tcpCliSock.recv(BUFSIZ)
    if not data:
        break
    print(data.decode("utf-8"))

tcpCliSock.close()
