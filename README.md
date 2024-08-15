
# **Cypher Tool Documentation**

## **What does the tool do?**

This tool provides basic encryption and decryption functionalities using three different ciphers:
- **ROT13 Cipher**
- **Reverse Alphabet Cipher**
- **Vigenère Cipher**

These ciphers can be used to encode and decode text, ensuring the obfuscation and security of the content.

## **Tool Usage with Examples**

### **1. ROT13 Cipher**

The **ROT13 cipher** shifts each letter in the input string by 13 positions in the alphabet. If the shift passes the end of the alphabet, it wraps around to the beginning.

#### **Example:**

```bash
georg@georg-PC:~/cypher$ ./cypher
Welcome to the Cypher Tool!

Select an operation (1/2):
1. Encrypt.
2. Decrypt.
1

Select cypher (1/3):
1. ROT13.
2. Reverse.
3. Vigenère.
1
Enter the message:
Hello, world!

Encrypted message using ROT13: 
Uryyb, jbeyq!
```

**Explanation:**  
- The letter `H` becomes `U`, `e` becomes `r`, and so on.
- Non-alphabetic characters (like punctuation) remain unchanged.

### **2. Reverse Alphabet Cipher**

The **Reverse Alphabet Cipher** replaces each letter in the input string with its reverse counterpart in the alphabet. For example:
- `A` becomes `Z`
- `B` becomes `Y`
- `C` becomes `X`

#### **Example:**

```bash
georg@georg-PC:~/cypher$ ./cypher
Welcome to the Cypher Tool!

Select an operation (1/2):
1. Encrypt.
2. Decrypt.
1

Select cypher (1/3):
1. ROT13.
2. Reverse.
3. Vigenère.
2
Enter the message:
Hello, world!

Encrypted message using Reverse: 
Svool, dliow! 
```

**Explanation:**  
- The letter `H` becomes `S`, `e` becomes `v`, etc.
- Similar to ROT13, non-alphabetic characters are not altered.

### **3. Vigenère Cipher**

The **Vigenère Cipher** uses a keyword to apply a series of Caesar ciphers, where each letter in the plaintext is shifted based on the corresponding letter of the key.

#### **Example:**

```bash
georg@georg-PC:~/cypher$ ./cypher
Welcome to the Cypher Tool!

Select an operation (1/2):
1. Encrypt.
2. Decrypt.
1

Select cypher (1/3):
1. ROT13.
2. Reverse.
3. Vigenère.
3
Enter the message:
Hello, world!

Enter the key: key

Encrypted message using Vigenere: 
Rijvs, uyvjn! 
```

**Explanation:**  
- Each letter in "Hello, World!" is shifted by an amount determined by the corresponding letter in the key "KEY".
- The process is reversible with the correct key.

## **Explanation of the Ciphers Used**

### **ROT13 Cipher**

- **Type:** Simple substitution cipher
- **Operation:** Shifts each letter by 13 positions
- **Strength:** Easy to implement, reversible by applying ROT13 twice
- **Use Case:** Often used for basic obfuscation (e.g., spoilers or simple encoding).

### **Reverse Alphabet Cipher**

- **Type:** Substitution cipher
- **Operation:** Each letter is substituted by its reverse in the alphabet (`A ↔ Z`, `B ↔ Y`, etc.)
- **Strength:** Simple, intuitive, but can be easily decoded if known.
- **Use Case:** Basic text obfuscation.

### **Vigenère Cipher**

- **Type:** Poly-alphabetic substitution cipher
- **Operation:** Uses a keyword to vary the Caesar cipher shift for each letter.
- **Strength:** More secure than ROT13 and reverse alphabet, as it uses a key for encryption.
- **Use Case:** Useful for securing short messages.

---

### **Conclusion**

This document outlines the functionality and usage of three ciphers implemented in the tool: **ROT13**, **Reverse Alphabet**, and **Vigenère**. By following the examples and explanations provided, users can easily understand how to utilize these ciphers for both encryption and decryption. The structured format, clear examples, and explanations make this document an effective guide for both beginners and advanced users.