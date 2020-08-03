def ShootBullet(obj1,obj2):#obj1,obj2 = []GameObject
    obj = []
    pobj = []
    if len(obj1) < len(obj2):
        for i in range(0,len(obj1) - 1):
            obj[i] = obj1[i]
        for i in range(0,len(obj2) - 1):
            pobj[i] = obj1[i]
    if len(obj2) < len(obj1):
        for i in range(0,len(obj2) - 1):
            obj[i] = obj2[i]
        for i in range(0,len(obj1) - 1):
            pobj[i] = obj2[i]
    for i in pobj:
        for j in obj:
            i.x = j.x
            i.y = j.y
class object:
    def __init__(self,x,y):
        self.x = x
        self.y = y
objtest1 = [object(10,10),object(10,10),object(0,0)]
objtest2 = [object(100,100),object(100,100)]
ShootBullet(objtest1,objtest2)
for i in range(0,len(objtest1) - 1):
    print(objtest1[i].x)
    print(objtest1[i].y)
for i in range(0,len(objtest2) - 1):
    print(objtest2[i].x)
    print(objtest2[i].y)
