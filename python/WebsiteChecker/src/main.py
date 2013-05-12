'''
Created on 2013/5/12

@author: frankwang
'''
import httplib
import urllib2
import time
import socket
from threading import Thread

def webchecker(host):
    if host.find("http") != -1:
        host = host.rstrip('\n')
        host = host.strip()
        host = host.replace("http://",'')
        try:
            conn = httplib.HTTPConnection(host,timeout=100)
            conn.request("HEAD", "/")
            r1 = conn.getresponse()
            if r1.status==200 :
                print host+" -->OK"
                return True
            else:
                print host+" -->DOWN "+str(r1.status)+" "+ r1.reason
                return False
        except (httplib.HTTPException, socket.error) as ex:
            print host+" -->DOWN "+"Error: %s" % ex
            return False

if __name__ == '__main__':
    print time.strftime("%Y / %m / %d %H:%M:%S",time.localtime(time.time()))+"\n"
    hosts = []
    alive = 0 
    down = 0
    for line in urllib2.urlopen(URL_LIST_URL)
       if line.find("http") != -1:
            hosts.append(line)
        
    for i in hosts:
        t = Thread(target=webchecker, args=(i,))
        time.sleep(1)
        t.start()
pass
