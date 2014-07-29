# -*- coding: utf-8 -*-
import utility.graber
import model
from sqlalchemy import create_engine
from datetime import datetime
from dateutil.relativedelta import relativedelta
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

Base = model.Base

engine = None
def StoreTAIEXDailyToDb(connectionStr,fromDate,endDate):
    engine = create_engine(connectionStr)
    Base.metadata.create_all(engine)
    while True:
        if fromDate>endDate:
            break
        else:
            fromDate = fromDate + relativedelta(months=1)
        trads = utility.graber.get_data(fromDate)
        for trad in trads:
            __storeObjectToDb(engine,trad)

def __storeObjectToDb(engine,object):
    Session = sessionmaker(bind=engine)
    session = Session()
    results = session.query(model.trad.Trad).filter_by(date=object.date).all()
    if len(results)>0:
        session.rollback()
        return
    print object.date
    session.add(object)
    session.commit()

if __name__ == "__main__":
    d1 = datetime(1990,01,01)
    d2 = datetime.today()
    StoreTAIEXDailyToDb('sqlite:///./stock.db',d1,d2)
