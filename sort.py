## print largest no
my_list = [17, 3, 11, 5, 11, 9919, 7, 15,45,78,90,2333,789, 23]
largest = my_list[0]
for i in my_list:
    if i>largest:
        largest=i
    
print("largest number",largest)
## print smallest
smallest=my_list[0]
for i in my_list:
    if i<smallest:
        smallest=i
print("smallest number",smallest)

print("## sort the list in ascending order. bubble sort")
my_list = [8, 10, 62, 22, 4]  # list to sort
list_size=len(my_list)
for i in range(0,list_size):
    list_size=list_size-1
    for i in range(0 ,list_size):
        print(i)
        
        if my_list[i]>my_list[i+1]  :
            #swap
            my_list[i],my_list[i+1]=my_list[i+1],my_list[i]
            
print(my_list)
