
import time

chars = ["a", "b", "c"]
pwlen = 2

password = ["", "", "", ""]

poss = len(chars)**pwlen


n = pwlen-1
counter = [-1, 0, 0, 0]
while counter != [3, 3, 3, 3]:
    counter[0] += 1
    for i, j in enumerate(counter):
        if j > len(chars)-1:
            counter[i] = 0
            counter[i+1] += 1
            continue
    
    for index, value in enumerate(counter):
        password[index] = chars[value]
    
    print(password)