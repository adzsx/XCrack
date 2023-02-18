# XCrack

### A go based application for offline password attacks

#

## Features

* Hash Cracking with:
    * Brute force
    * Wordlists

* Wordlists:
    * Generation
    * Merging
    * Cleaning

#

[Documentation](https://adzsx.github.io)


Modes:<br>

>crack:  Crack hashes with brute force or a wordlist
>list:   Generation, cleaning and merging of wordlists 
>hash:   Generate a hash from an input


<br><br>
crack mode: (default)<br>
Syntax:			xcrack (crack) [OPTIONS]<br>
<br>
Flags:
<br>
-p HASH:	   	Specify the hashed password 					required<br>
-t TYPE:		specify the hash-TYPE 							default: md5<br>
<br>
-n:			numbers											default<br>
-l:			lowercase letters								default<br>
-L:			uppercase letters<br>
-s:			special Characters<br>
-c CHARS:	Only uses CHARS for the password<br>
<br>
-m LENGTH:	min LENGTH of password							default: 1
-M LENGTH:	max LENGTH of password 							default: 8
<br>
-w PATH:	uses a wordlist in PATH<br>
<br>
<br>
list mode:<br>
Syntax:        	xcrack list [OPTIONS]<br>
<br>
Flags:<br>
-o PATH:		The location where the list is created		 	required<br>
New element will be appended<br>
<br>
-n:    		numbers											default<br>
-l:    		lowercase letters								default<br>
-L:    		uppercase letters<br>
-s:    		special Characters<br>
-c CHARS:	Only uses CHARS for the password<br>
<br>
-m LENGTH:  min LENGTH of password							default: 1<br>
-M LENGTH:  max LENGTH of password							default: 8<br>
<br>
-i PATH:    Input file for cleaning and merging
-o PATH:    Output file for merging
<br>
<br>
hash mode:<br>
Syntax:			xcrack gen [OPTIONS]<br>
<br>
Flags:<br>
-t TYPE:	Specifies the type of the hash 					default: md5<br>
-p STRING:  Argument will be hashed with TYPE<br>
<br>

#

## License

[GNU Piblic License](https://choosealicense.com/licenses/gpl-3.0/)
