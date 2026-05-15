#1)
lags = [2, 5, 12, 8, 15]
#Op :['Replica 3 lag 12s exceeds threshold!', 'Replica 5 lag 15s exceeds threshold!']
op=[]
for i,lst in enumerate(lags,start=1):
  print(i,lst)
   
  if i==3  :
    print(i)
    op.append(f"replica{i} lags {lst}s")
  elif i==5:
    op.append(f"replica{i} lags {lst}s")
    print(i)
print(op)
print("==========")
#2
incidents = [
    {"severity": "high", "issue": "Primary DB down"},
    {"severity": "medium", "issue": "Replication lag"},
    {"severity": "high", "issue": "Data corruption"},
]
#Op :{'high': 2, 'medium': 1}
dct={}
for lst in  incidents:
  
  if lst["severity"] in dct:
    dct[lst["severity"]]=dct[lst["severity"]]+1
  else:
    dct[lst["severity"]]=1
print (dct)  
print("==========")
#3)
rides = [
    {"user": "Alice", "fare": 25.5},
    {"user": "Bob", "fare": 40.0},
    {"user": "Alice", "fare": 15.0}
]
# op: {'Alice': 40.5, 'Bob': 40.0} , aggregate totals per user
dct={}
for lst in  rides:
 if lst["user"] in dct:
   dct[lst["user"]]=dct[lst["user"]]+lst["fare"]
 else:
   dct[lst["user"]]=lst["fare"]
print (f"#########{dct}##########")
#4)
# max and min number
numbers = [4, 12, 1, 23, 7,44,0]
max=numbers[0]
min=numbers[0]
for i in numbers:
  if i> max:
    max=i
  if i<min:
    min=i
print(f'Input {numbers} has maxno:{max} and  minno:{min}')
#5)
'''file a.txt
a 10
b 20
a 40
b 10
'''
#O/p : sum (col2) group by col1   i.e  {'a':50,'b':30}
# awk way : awk '{sum[$1]+=$2} END {for (k in sum) print k, sum[k]}' a.txt
ls=[]
with open("a.txt","r") as f:
  for i in f:
    k,v=i.strip().split()
    #print (k,v)
    ls.append((k,int(v)))
print(ls) #[('a', 10), ('b', 20), ('a', 40), ('b', 10)]
dct={}
for i in ls:
  print(i[0],i[1])
  if i[0] in dct:
    dct[i[0]]=dct[i[0]]+i[1]
  else:
    dct[i[0]]=i[1]
  #print(f'{i[0]},{i[1]}')
print(dct)
