from scapy.all import *
send(IP(src='192.168.10.222',dst='192.168.10.89')/UDP(sport=9999, dport=9999)/b'test')