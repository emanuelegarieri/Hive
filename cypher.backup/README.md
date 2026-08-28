# Cypher Tool

## _Overview_

### This command-line tool allows users to encrypt and decrypt messages using different cypher techniques. The user can choose whether to encrypt or decrypt a message, select a cypher, enter a message, and receive the processed result.

**_The tool supports the following cyphers:_**

* ROT13
* Reverse Alphabet Cypher
* Chinese Character cypher (Custom cypher)

**_Non-alphabetic characters such for example numbers, spaces, and punctuation marks remain unchanged during encryption and decryption._**

## How to Use the Tool

* Run the program.

* Choose an operation:
    * Type 1 to encrypt
    * Type 2 to decrypt

* Choose a cypher:

    * Type 1 for ROT13
    * Type 2 for Reverse Alphabet
    * Type 3 for Chinese Character
* Enter the message.

* The program displays the result.

Example of the inputs:

```
Input: 1
Operation: 1 is ROT13
cypher: 1
```
| Input   | Output    |
| ------ | ----- |
| Hello World | Uryyb Jbeyq |

## Cypher Descriptions:

#### ROT13 
* _Rot13 shifts every letter 13 positions forward in the alphabet._
* _Applying ROT13 twice or decrypting back the message returns the original message._

Examples:
```
A → N
B → O
N → A
Z → M

Example:
HELLO → URYYB
```



#### Reverse Alphabet Cypher
* _This cypher mirrors the alphabet._ 
* _Applying the cypher twice ror decrypting back the message returns the original message_


Examples:
```
A → Z
B → Y
C → X
X → C
Y → B
Z → A

Example:
HELLO → SVOOL
```



**Chinese Character cypher**

* _This custom cypher converts Latin alphabet characters into predefined Chinese Unicode characters._
* _Lowercase letters are mapped starting from the character 淡._
* _Uppercase letters are mapped starting from the character 中._
* _During decryption, the Chinese characters are converted back to their original Latin letters._
* _This cypher is fully reversible and supports both encryption and decryption._


Example:
```
a → 淡
b → 淡 + 1
A → 中
B → 中 + 1
```

## **Usage**
```
$ ./cyphertool


Welcome to the encryption tool!
---

Choose what you want to do:

- Type 1 if you want to encrypt
- Type 2 if you want to decrypt

You choosed to: Encrypt
---

Choose a cypher:

- Type 1 if you want to use Rot13
- Type 2 if you want to use Reverse Alphabet
- Type 3 if you want to use Chinese Converter

You choosed: Rot13
---

Enter your text:
Hello World!

Here is your result:
Uryyb Jbeyq!
```

## Input Validation
*  _The program validates all user inputs._
* _Invalid menu selections are rejected._
* _Empty messages are not accepted._
* _The user is repeatedly prompted until a valid option is entered._
_Leading and trailing whitespace is removed before validation using.
```
strings.TrimSpace().
```



### Technologies Used
#### The program was developed in Go and uses the following standard packages:

* _fmt_ for user interaction and output.
* _bufio_ for reading user input.
* _os_ for accessing standard input.
* _strings_ for processing and validating text input.

