from scapy.all import *


send(IP(src='192.168.10.89',dst='192.168.10.89')/UDP(sport=9998))