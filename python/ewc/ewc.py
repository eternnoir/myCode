# -*- coding: utf-8 -*-
__author__ = 'frank'

import csv
import argparse
from operator import itemgetter, attrgetter

parser = argparse.ArgumentParser(description='For EWC Use')
parser.add_argument('-e', action="store", dest="empl")
parser.add_argument('-w', action="store", dest="ewcm")

class employee(object):
    def __init__(self,enum,cname,ename,co,dep1,dep2,leader):
        self.enum = enum
        self.cname = unicode(cname, "utf-8")
        self.ename = ename
        self.co =co
        self.dep1 = dep1
        self.dep2 = dep2
        self.leader = leader

class ewcMember(object):
    def __init__(self,name,job,deps):
        self.name = unicode(name, "utf-8")
        self.job = job
        self.deps = deps

    def checkDep(self,dep):
        ret = False
        for d in self.deps:
            if (d in dep):
                ret = True
        return ret
def convertCsvToEmployee(data,ewList):
    ret = []
    for row in csv.reader(data, delimiter=','):
        depWithoutYear = row[4].split('_')[0]
        dep2WithoutYear = row[5].split('_')[0]
        emp = employee(row[0],row[1],row[2],row[3],depWithoutYear,dep2WithoutYear,findLeader(depWithoutYear,ewList))
        ret.append(emp)
    return ret

def convertCsvToEwc(data):
    ret = []
    for row in csv.reader(data, delimiter=','):
        depList = row[0].split('+')
        ew = ewcMember(row[2],row[1],depList)
        ret.append(ew)
    return ret

def findLeader(dep,ewList):
    ret = u''
    for ew in ewList:
        if ew.checkDep(dep):
            ret+=ew.name+u';'
    return ret

def printResultToCsv(cmlist):
    strherder = unicode("員編,中文姓名,英文姓名,處級單位,部級單位,負責福委",'utf-8')
    print(strherder.encode('utf-8'))
    for em in cmlist:
        str = u'%s,%s,%s,%s,%s,%s' %(em.enum,em.cname,em.ename,em.dep1,em.dep2,em.leader)
        print(str.encode("utf-8"))

def main():
    args = parser.parse_args()
    emPath = args.empl
    ewPath = args.ewcm
    emCsvData = open(emPath, 'rb')
    ewCsvData = open(ewPath, 'rb')
    ewList = convertCsvToEwc(ewCsvData)
    emList =  convertCsvToEmployee(emCsvData,ewList)
    emList = sorted(emList, key=attrgetter('dep1'))
    printResultToCsv(emList)

if __name__ == "__main__":
    main()