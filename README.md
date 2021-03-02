# Solitaire Encryption Algorithm

This repo contains an implementation (in Go) of the Solitaire encryption algorithm developed by Bruce Schneier.

The algorithm works by simulating a deck of playing cards (2 suits in this case) and a set of rules to get keystream values.

To learn more about it read [The Somewhat Simplified Solitaire Encryption Algorithm](http://nifty.stanford.edu/2006/mccann-sssolitaire/SSSolitaire.pdf)

---
### Instructions
First of all if you want to compile this program you need to have Go installed and properly configured.

I'm uploading an executable file for MacOS for demo purposes.

To encode a message:
```
./solitaire -msg="Dr. McCann is insane" -e
```

To decode a message:
```
./solitaire -msg="BYXFTZZLNQBZBLBNWOLE" -d
```

It is imperative that you provide a file named deck.txt formatted as the example given, with numbers 1-28 in an array. To be able to decode a message you need to provide the same order for this array with which the message was encoded. 