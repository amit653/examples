
def duplicate(str):
    dct={}
    count=1
    for i in str:
        if i in dct:
            dct[i]=dct[i]+count
        else:
            dct[i]=count
    print (dct)
    # Read dictionary and remove key count>1 using appending a emmpty list
    lst=[]
    for i,j in dct.items():
        if j==1:
            lst.append(i)
            print(i,j)
    print(''.join(lst))
    
    ## alternate using list comrehension
    
    #print(''.join([ i for i,j in dct.items() ]))
    print(''.join([ i for i,j in dct.items() if j==1]))

str=input(" string -")
duplicate(str)

###
#f = [lambda x, y: x+y]
#f[0](5,4)

