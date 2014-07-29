from datetime import datetime
import urllib2
import csv
import re

def get_data(date):
    csvdata = fetch_data(date)
    ret = []
    for row in csvdata:
        if ckinv(row) is True:
            ret.append(row)
    return ret


def fetch_data(date):
    """ Fetch data from twse.com.tw
        return list.
    """
    url = "http://www.twse.com.tw/ch/trading/exchange/FMTQIK/FMTQIK2.php?STK_NO=&myear=%(year)d&mmon=%(mon)02d&type=csv" % {
        'year': date.year,
        'mon': date.month}
    cc = urllib2.urlopen(url)
    csv_read = csv.reader(cc)
    return csv_read

def ckinv(oo):
    """ check the value is date or not """
    pattern = re.compile(r"[0-9]{2}/[0-9]{2}/[0-9]{2}")
    b = re.search(pattern, oo[0])
    try:
      b.group()
      return True
    except:
      return False

def utf_8_encoder(unicode_csv_data):
    for line in unicode_csv_data:
        yield line.encode('utf-8')
