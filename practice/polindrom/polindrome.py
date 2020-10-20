
# coding=utf-8
'''
Given a string, write a python function to check if it is palindrome or not. A string is said to be palindrome if the reverse of the string is the same as string. For example, “radar” is a palindrome, but “radix” is not a palindrome.
Examples: 
 

Input : malayalam
Output : Yes

Input : geeks
Output : No
'''

import sys


def ispolindrome(string):
    cleaned = "".join([i.lower() for i in string if i.isalnum()])
    
    return (cleaned == cleaned[::-1])


def main():
    string = sys.argv[1]
    #string = "---    a   , ba###   !!"
    
    print(ispolindrome(string))
    


if __name__ == "__main__":
    sys.exit(main())
