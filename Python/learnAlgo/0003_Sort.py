import random, time


def Sort_Select(li:list):
    """
    1万个数字2.5秒
    5万个数字84.3秒时间,22.8M内存。差不多是平方级的。
    """
    minindex = 0
    length = len(li)
    i, ii = 0,0
    for i in range(length):
        minindex = i
        for ii in range(i+1, length):
            if li[ii] < li[minindex]:
                minindex = ii
        li[i], li[minindex] = li[minindex], li[i]


def Is_Sorted(li:list):
    for i in range(len(li)-1):
        if li[i] > li[i+1]:
            return False
    return True

def Generate_int_list(num:int):
    data = []
    for i in range(num):
        data.append(random.randint(-10000000,10000000))
    return data


if __name__ == '__main__':
    data = Generate_int_list(10000)
    t1 = time.time()
    Sort_Select(data)
    t2 = time.time()
    print("%s, use %s seconds"%(Is_Sorted(data), t2-t1))
