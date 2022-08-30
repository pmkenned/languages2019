#!/bin/bash
#
# special: ' " \ $ ` * ~ ? < > ( ) ! | & ; space newline
# " quotes every character except $ ` \ ! * @
#
# REDIRECTION
# < file : close fd 0 and open file for reading
# > file : close fd 1 and open file for writing
# can be put almost anywhere in the command
# example: foo bar 35 <notes.txt >/dev/null
#
# PIPING
# commandA | commandB
#
# pipeline: one or multiple commands separated by |
# command-list: one or more pipelines separated and terminated by ; & or
# newline
# exit status of a command-list is the exit status of the last executed command
# foo | bar ; fizz | buzz
# \pipeline/  \pipeline/
#
# builtins: about 70
# help, cd, echo, type, export, return, {
#
# export foo: marks a variable as an environment variable
#
# bash precedence:
# 1. function call (separate namespace from variables)
# 2. built-in command
# 3. process command
#
# brace expansion: preamble{comma-separated-list}postscript
#   foo{apple,banana}bar => fooapplebar foobananabar
#   {35,14,hi}bar => 35bar 14bar hibar
#   ~: expanded to home directory
#
# command substitution:
#   $(command): must be used for nesting
#   `command`
#
# arithmetic substitution:
#   $((expression)
#   $((3 + $foo))

# filename expansion:
#   *: match any characters
#   ?: match any single character
#
# precedence for expansions:
# 1. brace
# 2. tilde
# 3. variable, arithmetic, command (evaluated inside out)
# 4. filename
#
# execute in subshell:
# (command-list)
# example: (cd /; ls -la;)
#
# execute in current shell:
# { cd /; ls -la }
# NOTE: { is a builtin, so there must be a space
# can be used to redirect output of all commands
# exit status is that of last command
#
# &: puts pipeline in background
# 1. shell doesn't wait for the pipelnie
# 2. pipeline runs in a subshell
# 3. the pipeline may not read the terminal
# 4. the pipeline may (or may not) write to the terminal
# & can be used to run {} command-lists in background
#
# SESSIONS
# shell starts a new session, starts out with one job (the shell) when the
# shell runs a pipeline in the foreground, the processes of that pipeline run
#   as members of the existing job
# when a subshell is created, that subshell runs as a new, separate job
# each session has a controlling terminal
# only one job is running in the foreground
# when background jobs try to read stdin, they are sent the SIGTTIN signal
# ctrl-z sends SIGTSTP to processes of the foreground job; also sends SIGCONT
#   to processes in the background job and moves that job to the foreground
# builtins for job control:
#   jobs: list background jobs and their numbers
#   bg: send SIGCONT to background job
#   fg: send SIGCONT to background job and move it to the foreground
#
# SOURCE

x=3
printf '$x\n'
printf "$x\n"
printf "${x}\n"

# if command-list then
# fi
# 
# while command-list do
# done

function foo {
    echo $1
    ls -la
    return 1
}

foo 15
echo $?

echo $((3 + $x))
