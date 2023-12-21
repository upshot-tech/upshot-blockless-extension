README for Upshot Blockless Extension
=====================================

This document provides instructions on how to use the Makefile for building the Upshot Blockless Extension.

Prerequisites
-------------

*   Go Programming Environment
*   GNU Make
*   Node.js and npm (for the `example` target)
*   Wget (for the `setup` target)

Makefile Targets
----------------

*   `build`: Compiles the Go source file into a binary.
*   `clean`: Removes the compiled binary and other generated files.
*   `example`: Runs a Node.js build script in the `upshot-function-example` directory.
*   `setup`: Downloads and extracts the appropriate blockless runtime for the current OS and architecture.
*   `test`: Runs a test using the blockless runtime and the built wasm file.

Usage
-----

Run the following commands in your terminal:

*   To compile the source code:
    
        make build
    
*   To clean up generated files:
    
        make clean
    
*   To run the example Node.js project:
    
        make example
    
*   To set up the runtime environment:
    
        make setup
    
*   To run tests:
    
        make test
    

Notes
-----

The `setup` target automatically detects your operating system and CPU architecture to download the correct version of the blockless runtime. This process uses `wget` and `tar` for downloading and extracting the files, respectively.

The `test` target depends on both the `build` and `example` targets, ensuring that all necessary components are built before testing.

Support
-------

For support, ensure that you have all the prerequisites installed and that your environment is set up correctly. If you encounter any issues, check the console output for error messages that may provide more information.