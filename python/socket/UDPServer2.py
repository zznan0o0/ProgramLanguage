import socket
BUFSIZE = 1024
ip_port = ('0.0.0.0', 9999)
server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)  # udp协议
server.bind(ip_port)
while True:
    data,client_addr = server.recvfrom(BUFSIZE)
    print(client_addr)
    print('server收到的数据', data)
 
    server.sendto(data.upper(),client_addr)

# server.close()