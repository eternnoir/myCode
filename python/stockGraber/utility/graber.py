# -*- coding: utf-8 -*-
from datetime import datetime
from model import trad
import urllib2
import csv
import re


def get_data(date):
    csvdata = fetch_data(date)
    ret = []
    for row in csvdata:
        if ckinv(row) is True:
            ret.append(_genNewTrad(row))
    return ret


def _genNewTrad(row):
    dateStr = row[0]
    ye = dateStr.split('/')[0]
    acdate = dateStr.replace( ye, str(int(ye)+ 1911))
    date = datetime.strptime(acdate,'%Y/%m/%d')
    value = float(row[1].replace(',',''))
    delvalue = float(row[2].replace(',',''))
    delnum = float(row[3].replace(',',''))
    taiex = float(row[4].replace(',',''))
    updown = float(row[5].replace(',',''))
    ret = trad.Trad( date, value, delvalue, delnum, taiex, updown)
    return ret


def fetch_data(date):
    url = "http://www.twse.com.tw/ch/trading/exchange/FMTQIK/FMTQIK2.php?STK_NO=&myear=%(year)d&mmon=%(mon)02d&type=csv" % {
        'year': date.year,
        'mon': date.month}
    cc = urllib2.urlopen(url)
    csv_read = csv.reader(cc)
    return csv_read


def ckinv(row):
    """ check the value is date or not """
    r = re.compile(r"[0-9]{2}/[0-9]{2}/[0-9]{2}")
    b = re.search(r, row[0])  # check date
    try:
        b.group()
        return True
    except:
        return False


def utf_8_encoder(unicode_csv_data):
    for line in unicode_csv_data:
        yield line.encode('utf-8')
