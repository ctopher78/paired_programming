'''
Given a string, write a python function to check if it is palindrome or not. A string is said to be palindrome if the reverse of the string is the same as string. For example, “radar” is a palindrome, but “radix” is not a palindrome.
Examples: 
 

Input : malayalam
Output : Yes

Input : geeks
Output : No
'''

import sys


def ispalidrome(text):
    cleaned = "".join([i.lower() for i in text if i.isalnum()])
    return (text == cleaned[::-1])


def main():
    try:
        text = sys.argv[1]
    except IndexError:
        sys.exit("No argument passed")

    print(ispalidrome(text))



if __name__ == "__main__":
    sys.exit(main())