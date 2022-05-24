matrix=[]
for i in range(5):
    matrix.append([])
#print (matrix) # [[], [], [], [], []]
    for j in range(5):
        matrix[i].append(j)
print (matrix)  # [[0, 1, 2, 3, 4], [0, 1, 2, 3, 4], [0, 1, 2, 3, 4], [0, 1, 2, 3, 4], [0, 1, 2, 3, 4]]

# using list comprehension
mat = [[1, 2, 3], [4, 5], [6, 7, 8, 9]]
strng=[val for lst in mat for val in lst]
print(strng)
