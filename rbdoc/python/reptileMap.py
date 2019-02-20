# -*- coding:UTF-8 -*-
import requests
import time
import xlwt
import os
import sys
from apscheduler.schedulers.blocking import BlockingScheduler
from apscheduler.triggers.cron import CronTrigger


class LocaDiv(object):
    def __init__(self, loc_all):
        self.loc_all = loc_all

    def lat_all(self):
        lat_sw = float(self.loc_all.split(',')[0])
        lat_ne = float(self.loc_all.split(',')[2])
        lat_list = []
        for i in range(0, int((lat_ne - lat_sw + 0.0001) / 0.05)):  # 0.1为网格大小，可更改
            lat_list.append(lat_sw + 0.05 * i)  # 0.05
        lat_list.append(lat_ne)
        return lat_list

    def lng_all(self):
        lng_sw = float(self.loc_all.split(',')[1])
        lng_ne = float(self.loc_all.split(',')[3])
        lng_list = []
        for i in range(0, int((lng_ne - lng_sw + 0.0001) / 0.05)):  # 0.1为网格大小，可更改
            lng_list.append(lng_sw + 0.05 * i)  # 0.1为网格大小，可更改
        lng_list.append(lng_ne)
        return lng_list

    def ls_com(self):
        l1 = self.lat_all()
        l2 = self.lng_all()
        ab_list = []
        for i in range(0, len(l1)):
            a = str(l1[i])
            for i2 in range(0, len(l2)):
                b = str(l2[i2])
                ab = a + ',' + b
                ab_list.append(ab)
        return ab_list

    def ls_row(self):
        l1 = self.lat_all()
        l2 = self.lng_all()
        ls_com_v = self.ls_com()
        ls = []
        for n in range(0, len(l1) - 1):
            for i in range(0 + len(l1) * n, len(l2) + (len(l2)) * n - 1):
                a = ls_com_v[i]
                b = ls_com_v[i + len(l2) + 1]
                ab = a + ';' + b
                ls.append(ab)
        return ls


def reptileMap(key):
    print('key='+key)
    print('[info]开始爬取数据...')
    startTime = time.time()
    loc0 = LocaDiv('116.208904,39.747315,116.5536,40.025504')
    locs = loc0.ls_row()
    date = time.strftime("%Y%m%d-%H")
    path = 'ReptileMap'+date+'.xls'
    # 删除旧文件
    if os.path.exists(path):
       os.remove(path)
    # file = open(path+'\\data_xy.txt', 'a', encoding='utf-8')
    # file2 = open(path+'\\data.txt', 'a', encoding='utf-8')
    count = 0
    workbook = xlwt.Workbook()
    sheet1 = workbook.add_sheet('student')
    keys1 = ['angle', 'direction', 'lcodes', 'name', 'polyline', 'speed', 'status']
    for i in range(0, len(keys1)):
        sheet1.write(0, i, keys1[i])       # 写入表头
    for loc in locs:
        pa = {
            'key': str(key),
            # 'level': 6,                   # 道路等级为6，即返回的道路路况等级最小到无名道路这一级别
            'rectangle': str(loc),          # 矩形区域
            'extensions': 'all'
            # 'output': 'JSON'
        }
        print('[info]探测区块：'+loc)
        obj = requests.get('http://restapi.amap.com/v3/traffic/status/rectangle?', params=pa)
        data = obj.json()
        if data['status'] == '0':
            print('[info]'+str(data))
            print('[warn]请求参数错误')
            break

        for road in data['trafficinfo']['roads']:
            count = count+1

            rangle = road['angle'] if 'angle' in road else ''
            rdirection = road['direction'] if 'direction' in road else ''
            rlcodes = road['lcodes'] if 'lcodes' in road else ''
            rname = road['name'] if 'name' in road else ''
            rpolyline = road['polyline'] if 'polyline' in road else ''
            rspeed = road['speed'] if 'speed' in road else '0'
            rstatus = road['status'] if 'status' in road else ''

            sheet1.write(count, 0, rangle)
            sheet1.write(count, 1, rdirection)
            sheet1.write(count, 2, rlcodes)
            sheet1.write(count, 3, rname)
            sheet1.write(count, 4, rpolyline)
            sheet1.write(count, 5, rspeed)
            sheet1.write(count, 6, rstatus)

            # 判断有无路线，有些没有，直接取值会报错
            #if 'polyline' in road:
            #rpolyline = road['polyline'].split(";")
            #for i in range(0, len(rpolyline)):
            #roadloc = rname+','+rdirection+','+rpolyline[i]+','+rname+'-'+rdirection+'\n'
            #file.write(roadloc)

        time.sleep(1)    # 间隔1s执行一次分块请求，避免并发度高被限制
    workbook.save(path)
    endTime = time.time()
    print('[info]数据爬取完毕，用时%.2f秒' % (endTime-startTime))
    print('[info]数据存储路径：'+os.path.abspath('.')+'\\'+path)

def test(key):
    print(key)

def startWork(key, loop):
    reptileMap(key)
    if loop:
        scheduler = BlockingScheduler()
        # 周一到周五,每小时执行一次   每5秒second='*/5'
        trigger = CronTrigger(day_of_week='0-6', hour='0-23')
        scheduler.add_job(reptileMap, trigger, args=(key,))
        # 周六到周日,24
        # scheduler.add_job(test, 'cron', day_of_week='5-6', hour='0-23')
        scheduler.start()
        print('开始实时取数据...')

if __name__ == '__main__':
    print('**************************EXE执行方式**************************************')
    print('[info]执行命令(key缺省)案例：reptileMap.exe -key 0b1804994cd63974f873a29a269d65e7 -loop')
    print('**************************脚本执行方式**************************************')
    print('[info]请输入高德地图key,缺省时将使用开发者预留key（保存5天，5天后预留key删除）')
    print('[info]执行命令(带key)案例：python reptileMap.py -key 0b1804994cd63974f873a29a269d65e7')
    print('[info]执行命令(key缺省)案例：python reptileMap.py')
    print('[info]执行命令(循环执行,每小时爬一次)案例：python reptileMap.py -key 0b1804994cd63974f873a29a269d65e7 -loop')
    print('[info]缺省-loop时默认只爬一次')
    param = sys.argv
    key = '0b1804994cd63974f873a29a269d65e7'
    loop = False
    for i in range(0, len(param)):
        if param[i] == '-key':
            key = param[i+1]
        if param[i] == '-loop':
            loop = True
    startWork(key, loop)
