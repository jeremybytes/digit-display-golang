# digit-display-golang
Naive hand-written digit recognition with display applications to show image, prediction, and errors.  

This is a Go port of a project in .NET (F# & C#). Details of that project are available here: [https://github.com/jeremybytes/digit-display](https://github.com/jeremybytes/digit-display).  

This is primarily practice for Go.

Functions:  
* Reading files from the file system
* Training simple nearest-neighbor digit recognizers
    * Manhattan distance
    * Euclidean distance
* Output (pretty bad) ASCII art
* Tracking errors
* Running code in parallel (goroutines)
* Using channels to communicate between functions

Note:  

Running the Euclidean recognizer on 3000 records takes **17-ish seconds**. The equivalent application from the .NET project takes **32-ish seconds** for the same data set. The code is not equivalent (the .NET code has some maps and minBys), but the results are. I'll be exploring this a bit further.
