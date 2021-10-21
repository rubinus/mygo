cmake_minimum_required(VERSION 3.20)
project(mygo C)

set(CMAKE_C_STANDARD 11)

add_executable(mygo
        c/hello/c.c)
