#!/usr/bin/python
# -*- coding: utf-8 -*-

import os
import sys
from PIL import Image

def cropOneOrb(sourceImg):
    w , h = sourceImg.size
    oneBall = w/6
    for i in range(5):
        for j in range(6):
            cx = oneBall*j;
            cy = oneBall*i;
            img = sourceImg.crop((cx,cy,cx+oneBall,cy+oneBall))
            img.save(str(j+1)+str(i+1)+".png")

if __name__ == '__main__':
    image = Image.open(sys.argv[1])
    cropOneOrb(image)

