import mechanize
import cookielib
import sys
from BeautifulSoup import BeautifulSoup
import html2text
import urllib

def getProjectNum(sou):
	msg = str(sou.findAll('input', attrs={'type': 'radio'})[0])
	items = msg.strip().split(' ')
	num = items[4].strip().split('\"')
	return num[1]
def getSignoutNum(sou):
	msg = str(sou.findAll('input', attrs={'type': 'radio'})[0])
	items = msg.strip().split(' ')
	num = items[4].strip().split('\"')
	return num[1]

# Browser
br = mechanize.Browser()

# Cookie Jar
cj = cookielib.LWPCookieJar()
br.set_cookiejar(cj)

# Browser options
br.set_handle_equiv(True)
#br.set_handle_gzip(True)
br.set_handle_redirect(True)
br.set_handle_referer(True)
br.set_handle_robots(False)

# Follows refresh 0 but not hangs on refresh > 0
br.set_handle_refresh(mechanize._http.HTTPRefreshProcessor(), max_time=1)

# Want debugging messages?
#br.set_debug_http(True)
#br.set_debug_redirects(True)
#br.set_debug_responses(True)

# User-Agent (this is cheating, ok?)
br.addheaders = [('User-agent', 'Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.1) Gecko/2008071615 Fedora/3.0.1-1.fc9 Firefox/3.0.1')]


if __name__ == '__main__':
	print 'Try to lgin'
	action = sys.argv[1] 
	id = sys.argv[2]
	pw = sys.argv[3]
	br.open('http://wallaby.cc.ncu.edu.tw/login').read()
	br.select_form(nr=0)
 	br.form['j_username'] = id
 	br.form['j_password'] = pw
 	br.submit()
 	br.open('http://140.115.182.62/PartTime/parttime.php/signin').read()
 	print br.title()
 	if  br.title() is "Welcome to Simple Portal":
 		print 'login faild'
 	else:
 		print 'login susscess'
	
	if action == 'signin':
 		soup = BeautifulSoup(br.open('http://140.115.182.62/PartTime/parttime.php/signin').read())
 		num = getProjectNum(soup)
 		print 'project name:'+num
 		parameters = {'signin' : num,
			'submit' : '%E9%80%81%E5%87%BA'
		}
		data = urllib.urlencode(parameters)
		br.open('http://140.115.182.62/PartTime/parttime.php/signin',data)
	elif action == 'signout':
		soup = BeautifulSoup(br.open('http://140.115.182.62/PartTime/parttime.php/signout').read())
		signoutNum =  getSignoutNum(soup)
		
		parameters = {'signout' : signoutNum,
			'submit' : '%E9%80%81%E5%87%BA'
		}
    	data = urllib.urlencode(parameters)
    	br.open('http://140.115.182.62/PartTime/parttime.php/signout',data)
        
        
