# -*- coding: utf-8 -*-
import os
import shutil
import time
import argparse

parser = argparse.ArgumentParser(description='FileKiller, help you to remove dir or path before N days ')
parser.add_argument('-p', action="store", dest="path")
parser.add_argument('-d', action="store", dest="days", type= int)

def scanPath(path,deadtime):
    for f in os.listdir(path):
        if f.startswith('.') or f.startswith('_'):
            continue
        f = os.path.join(path, f)
        if os.path.isdir(f) or os.path.isfile(f):
            (mode, ino, dev, nlink, uid, gid, size, atime, mtime, ctime) = os.stat(f)
            if ctime < deadtime:
                print 'Remove '+f
                removeFileOrDic(f)

def removeFileOrDic(file):
    if os.path.isdir(file):
        shutil.rmtree(file)
    elif os.path.isfile(file):
        os.remove(file)
    else:
        print 'err type'

def main():
    args = parser.parse_args()
    days = args.days
    path = args.path
    now = time.time()
    deadtime = now - days * 86400
    scanPath(path,deadtime)

if __name__ == '__main__':
    main()