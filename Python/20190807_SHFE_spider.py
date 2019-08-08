#!/usr/bin/python3
# encoding:utf-8

import os
import asyncio
import aiohttp
import json
import pandas as pd
from datetime import datetime, timedelta
import queue


class MySpider(object):
    def __init__(self, urlQ: queue.Queue):
        self.urlQ = urlQ

    async def fetch(self, url, name):
        async with asyncio.Semaphore(5):  # 限制并发数为5个
            async with aiohttp.ClientSession() as response:
                async with response.get(url) as html:
                    print("[%s] is getting %s" % (name, url))
                    text = await html.text(encoding="utf-8")
                    print("[%s] finish %s" % (name, url))
                    self.parse(url, text, url[-12:-4])

    def parse(self, url: str, text: str, the_date: str):
        try:
            dic = json.loads(text)
        except:
            return

        if dic:
            df = pd.DataFrame(dic['o_cursor'])

            df_ok = pd.DataFrame()
            df_ok['INSTRUMENTID'] = df['INSTRUMENTID'].str.strip()
            df_ok['PRODUCTSORTNO'] = df['PRODUCTSORTNO']
            df_ok['trade_name'] = df['PARTICIPANTABBR1'].str.strip()
            df_ok['trade_amount'] = df['CJ1']
            df_ok['trade_change'] = df['CJ1_CHG']
            df_ok['open_name'] = df['PARTICIPANTABBR2'].str.strip()
            df_ok['open_amount'] = df['CJ2']
            df_ok['open_change'] = df['CJ2_CHG']
            df_ok['close_name'] = df['PARTICIPANTABBR3'].str.strip()
            df_ok['close_amount'] = df['CJ3']
            df_ok['close_change'] = df['CJ3_CHG']
            path_dir = "C:/Users/lewin/mydata/SHFE/%s" % the_date
            if not os.path.isdir(path_dir):
                os.mkdir(path_dir)

            set_product = set(df_ok['PRODUCTSORTNO'])
            dic_name = {10: "铜", 20: "铝", 30: "锌", 40: "铅", 45: "镍", 48: "锡", 50: "黄金", 60: "白银", 70: "螺纹钢",
                        85: "热轧卷板", 92: "燃料油", 95: "石油沥青", 100: "天然橡胶", 110: "纸浆"}
            print("Saving files...")
            for product in set_product:
                path_file = os.path.join(path_dir, "%s_%s.csv" % (the_date, dic_name[product]))
                df_ok[df_ok['PRODUCTSORTNO'] == product].to_csv(path_file, encoding="utf-8", index=False)
                # print("saved file: %s" % path_file)

    async def main(self, name):
        print("Start main!!!!", self.urlQ.qsize())
        while not self.urlQ.empty():
            await self.fetch(self.urlQ.get(), name)


urlQ = queue.Queue()
since = datetime.strptime('20190101', '%Y%m%d')
today = datetime.strptime('20190501', '%Y%m%d')
while since <= today:
    urlQ.put("http://www.shfe.com.cn/data/dailydata/kx/pm%s.dat" % since.strftime("%Y%m%d"))
    since += timedelta(1)
print("Urls ok!!")

loop = asyncio.get_event_loop()

my_spider = MySpider(urlQ)
tasks = [asyncio.ensure_future(my_spider.main(i)) for i in range(20)]
loop.run_until_complete(asyncio.wait(tasks))
loop.run_until_complete(asyncio.sleep(0))
loop.close()
