input = open('input1.dat', 'r')
lines = input.readlines()

depths = []
for l in lines:
    depths.append(l.strip())

print('len depths', len(depths))
c = 0

for i in range(1, len(depths)):
    if int(depths[i]) > int(depths[i-1]):
        c+=1
        print(depths[i], 'increased')
    else:
        print(depths[i], 'decreased')

print(c)