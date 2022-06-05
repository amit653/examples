#val=0
# Program to print denomminators of a amount in Rupees
rem=int(input('Enter any amount less than 10 million : '))

st=''
flag=True
def rem_check(rem):

    l=[1,2,5,10,20,50,100,200,500,2000]
    lst=[i for i in l if rem>=i ]
    #print(lst)
    mx=1
    for i in lst:
        if mx<i:
            mx=i
    return mx

#flag=True
while(flag):
 
    if  rem<=2:
        #print("sum ", val)
        flag=False
        if rem==0:
            print(f'\nI will return { st }  ')
        else:
            print(f'{ st } And {rem} rupees! \n')
    else :
        
        no=rem
        mx=rem_check(no)
        #print("mx ",mx)
        rem= no % mx #r 6  q,val= q* i  ,16%10 =6
        #print("rem",rem)
        q= no// mx # q 1
        #print("q",q)
        #val= q*mx  #1*10
        st=st+str(q) +', note of '+str(mx)+' Rupees .  \n'
       
    

