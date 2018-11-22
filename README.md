# linux-kitty-vim-cjk-input-plugin
Simple script to help you insert CJK words into vim in Kitty terminal

## Requirments
- zenity

## Installation
```
Put CJKInput.vim to your .vimrc
```

## CJKInput.vim
```
function CJKInput()
	let l:cmd = 'zenity --entry --text=CJK-Input 2>/dev/null'
	let l:output = system(l:cmd)
	let l:output = substitute(l:output, '[\r\n]*$', '', '')
	execute 'normal i' . l:output
endfunction
nmap <silent> <leader>i :call CJKInput()<CR>
```
NOTE: dismiss message "Gtk-Message: GtkDialog mapped without a transient parent. This is discouraged." by redirect stderr to /dev/null 

**Thanks for [@MachFour's suggestion](https://github.com/dsewnr/linux-kitty-vim-cjk-input-tool/issues/1).**

## Demo
![](linux-kitty-vim-cjk-input-tool.gif)
