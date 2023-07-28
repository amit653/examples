
'''
Given an array of strings, group anagrams together
        All inputs will be in lowercase.
        The order of your output does not matter.
Examples:
Input: ["eat", "tea", "tan", "ate", "nat", "bat", "teae", "teaa"] ->["eat","eat","tan","eat"]
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"],
  ["teae"],
  ["teaa"]
]
Input: ["abc", "cafe", "face", "cab",  "goo"]
Output:
[
  ["abc", "cab"],
  ["cafe", "face"],
  ["goo"]
]
'''

anagrams=["eat", "tea", "tan", "ate", "nat", "bat", "teae", "teaa"]
#anagrams=["abc", "cafe", "face", "cab",  "goo"]

grouped_anagrams = {}
for string in anagrams:
    sorted_string=str(sorted(string))
    #print(sorted_string)
    if sorted_string in grouped_anagrams:
        #print("then")
        grouped_anagrams[sorted_string].append(string )
        #print(grouped_anagrams)
    else:
        #print("else")
        grouped_anagrams[sorted_string]=[string]
print(list(grouped_anagrams.values())) #[['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat'], ['teae'], ['teaa']]
lst=[]
for k in grouped_anagrams.values():
    #print(sorted(k))
    lst.append(sorted(k))
print("output= ",lst) # [['ate', 'eat', 'tea'], ['nat', 'tan'], ['bat'], ['teae'], ['teaa']]
