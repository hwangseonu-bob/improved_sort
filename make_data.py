from random import shuffle

li = [i for i in range(1, 100001)]
shuffle(li)

s = '\n'.join(map(str, li))

with open('data.csv', 'w') as f:
    f.write(s)