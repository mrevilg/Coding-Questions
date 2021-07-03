# Daily Coding for Python
# Pushing each concept as far as I can

print("Hello World!")
something = input("What is your name? ")
age = input("How old are you? ")
byear = input("What year were you born in? ")
now = input("What year is it? ")

print(something)
print(age)
print(int(now) - int(byear))

for i in range(1, 35):
    if i % 15 == 0:
        print("Fizz Buzz") 
    elif i % 3 == 0:
        print("Fizz")
    elif i % 5 == 0:
        print("Buzz")
    else: 
        print(i)    