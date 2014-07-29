# -*- coding: utf-8 -*-
import model
from sqlalchemy import Column, Boolean,Integer, Float, DateTime

class Trad(model.Base):
    __tablename__ = "trad"
    id = Column(Integer, primary_key=True)
    date = Column(DateTime)
    value = Column(Float)
    dealvalue = Column(Float)
    dealnum = Column(Float)
    taiex = Column(Float)
    updown = Column(Float)
    isup = Column(Boolean)

    def __init__(self,date,value,dealvalue,dealnum,taiex,updown):
        self.date = date
        self.value = value
        self.dealvalue = dealvalue
        self.delnum = dealnum
        self.taiex = taiex
        self.updown = updown
        if updown < 0:
            self.isup = False
        else:
            self.isup = True

