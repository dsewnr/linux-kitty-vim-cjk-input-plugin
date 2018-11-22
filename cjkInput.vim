function CJKInput()
	let l:cmd = 'zenity --entry --text=CJK-Input 2>/dev/null'
	let l:output = system(l:cmd)
	let l:output = substitute(l:output, '[\r\n]*$', '', '')
	execute 'normal i' . l:output
endfunction
nmap <silent> <leader>i :call CJKInput()<CR>
