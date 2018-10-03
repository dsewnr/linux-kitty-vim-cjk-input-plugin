# linux-kitty-vim-cjk-input-tool
Simple tool to help you insert CJK words to vim in Kitty terminal

## Requirments
- go
- zenity
- xclip

## Installation
```
$ git clone https://github.com/dsewnr/linux-kitty-vim-cjk-input-tool.git
$ cd linux-kitty-vim-cjk-input-tool/cjkInput
$ go build -o cjkInput main.go
$ mv cjkInput /usr/local/bin
```
And put cjkInput.vim to your .vimrc
```
function CJKInput()
    let l:cmd = 'cjkInput'
    let l:output = system(l:cmd)
    let l:output = substitute(l:output, '[\r\n]*$', '', '')
    execute 'normal i' . l:output
endfunction
nmap <silent> \i :call CJKInput()<CR>
```

## Demo
![](linux-kitty-vim-cjk-input-tool.gif)
