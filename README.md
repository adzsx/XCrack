# XCrack
## Description
XCrack is a selection of tools for offline password attacks suck as hash cracking, wordlist generation, wordlist merging, and many more.


## Documentation

Modes:  
* crack: Crack hashes with brute force or a wordlist
* list: Generation, cleaning and merging of wordlists
* hash: Generate a hash from an input


  
### Crack mode
Syntax: ``xcrack (crack) [Flags] ``
  
Flags:  
> -p:	HASH: Specify the hashed password required  
>	-t:	TYPE: specify the hash-TYPE default: md5  

>-n: 	numbers default  
>-l: 		lowercase letters default  
>-L: 	uppercase letters  
>-s: 	special Characters  
>-c: 	CHARS: Only uses CHARS for the password  
  
>-m:	LENGTH: min LENGTH of password default: 1 -M LENGTH: max LENGTH of password default: 8  
>-w 	PATH: uses a wordlist in PATH  

Example: ``xcrack -p fc24957e58927ade2522e35c23411de4f07f473e -t sha1 -l -M 6``
  
  
### List mode
Syntax: ``xcrack list [Flags] ``
  
Flags:  
>-n: 	numbers default  
>-l: 		lowercase letters default  
>-L: 	uppercase letters  
>-s: 	special Characters  
>-c: 	CHARS: Only uses CHARS for the password  
  
>-m 	LENGTH: min LENGTH of password default: 1  
>-M 	LENGTH: max LENGTH of password default: 8  
  
>-i: 		PATH: Input file for cleaning and merging
>-o:	PATH: Path for output file

Example: ``xcrack list -l -L -M 6 -o ./wordlist.txt``
  
### Hash mode
Syntax: xcrack hash [Flags]  
  
Flags:  
>-t: 	TYPE: Specifies the type of the hash default: md5  
>-p: 	STRING: Argument will be hashed with TYPE

Example: ``xcrack hash -f "xcrack" -t sha1``

##  Contribution
If you want to contribute to the project, feel free to open an issue on this repo with code imrovements or feature ideas. 

If you want to contribute on the long run, write me an E-Mail at adzsx@proton.me or from my [website](https://adzsx.github.io/#contact)

# 
#### [GPL License](https://choosealicense.com/licenses/gpl-3.0/)
