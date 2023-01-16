# XCrack

### A go based application for offline password attacks

#

## Features

- Hash Cracking with:
    - Brute force
    - Wordlists

- Wordlists:
    - Generation
    - Merging*
    - Cleaning*

*Comming soon

#

[Documentation](https://adzsx.github.io)


Modes:<br>

hash:   Cracks a given hash with either a wordlist or brute force attack (default)<br>
list:   Generated a wordlist based on your preferences<br>
gen:    Generates a hash from a given string<br>
file:	Combine wordlists and generate a new list, with duplicates removed<br>
<br><br>
hash mode:<br>
Syntax:			xcrack (hash) [OPTIONS]<br>
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
Example `xcrack hash -t sha1 -l -n -M 6 -p 6367c48dd193d56ea7b0baad25b19455e529f5ee`<br>
<br>
<br>
list mode:<br>
Syntax:        	xcrack list [OPTIONS]<br>
<br>
Flags:<br>
-p PATH:		The location where the list is created		 	required<br>
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
Example `xcrack list -p ~/Documents/wordlist.txt -l -L -M 6`<br>
<br>
gen mode:<br>
Syntax:			xcrack gen [OPTIONS]<br>
<br>
Flags:<br>
-t TYPE:	Specifies the type of the hash 					default: md5<br>
-p STRING:  Argument will be hashed with TYPE<br>
<br>

#

## License

[MIT](https://choosealicense.com/licenses/mit/)