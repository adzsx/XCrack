# XCrack
## Description
XCrack is a selection of tools for offline password attacks suck as hash cracking, wordlist generation, wordlist merging, and many more.


## Documentation
>crack:  Crack hashes with brute force or a wordlist
>list:   Generation, cleaning and merging of wordlists 
>hash:   Generate a hash from an input


### Crack mode
Syntax: ``xcrack (crack) [Flags] ``

Flags:  
>-p HASH:   Specify the hashed password required  
>-t TYPE:   specify the hash-TYPE default: md5  
>-n: 	    numbers default  
>-l: 		lowercase letters default  
>-L: 	    uppercase letters  
>-s:        special Characters  
>-c	CHARS:  Only uses CHARS for the password  
  
>-m	LENGTH: min LENGTH of password default: 1
>-M LENGTH: max LENGTH of password default: 8  

>-w	PATH:   uses a wordlist in PATH  


### List mode
Syntax: ``xcrack list [Flags] ``

Flags:  
>-n: 	    numbers default  
>-l: 	    lowercase letters default  
>-L: 	    uppercase letters  
>-s: 	    special Characters  
>-c CHARS:  Only uses CHARS for the password  
  
>-m LENGTH  min LENGTH of password default: 1  
>-M LENGTH  max LENGTH of password default: 8  
  
>-w PATH    Input file for cleaning and merging
>-o PATH:   Output file for wordlist generation, cleaning and merging

### Hash mode
Syntax: xcrack hash [Flags]  

Flags:  
>-t	TYPE:   Specifies the type of the hash default: md5  
>-p STRING: Argument will be hashed with TYPE

##  Contribution
If you want to contribute to the project, feel free to fork this repo with code imrovements. 

If you want to contribute on the long run, contact me at adzsx@proton.me or my [website](https://adzsx.github.io/#contact)

# 
#### [GPL License](https://choosealicense.com/licenses/gpl-3.0/)
