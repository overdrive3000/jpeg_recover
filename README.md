# Recover JPEG files

This is a simple Golang program to show how to recover JPEG files from a raw image of a FAT32 file system.

Basically the program read the FAT32 image loop through each block looking for a JPEG signature. If a signature is found the program fill a buffer with the binary data found and writes the buffer in a file in the hard drive.

This program is based in an exercise from CS50 course http://cs50.edx.org (pset4)
