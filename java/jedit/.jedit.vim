" to use this file: vim -u .jedit.vim

" disable vi compatibility mode and enable backspace
set nocp
set backspace=indent,eol,start

function Clear()
    :!echo -en "\ec"
endfunction

nnoremap <F5> :call Clear()<CR>:!make<CR>
nnoremap <F6> :call Clear()<CR>:!make run<CR>
nnoremap <F8> :call Clear()<CR>:!make clean<CR>

" sources .vimrc if it exists
if len(findfile(".vimrc", $HOME)) != 0
    source $HOME/.vimrc
endif

filetype indent plugin on

silent e Makefile
silent tabe JEdit.java
silent tabe AnsiEsc.java
silent tabe tty/Termios.java
silent tabe tty/TTYUtil.java
silent tabe tty_TTYUtil.cpp
silent tabe tty_TTYUtil.h
normal gt
normal gt
