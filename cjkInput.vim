function CJKInput()
    let l:cmd = 'cjkInput'
    let l:output = system(l:cmd)
    let l:output = substitute(l:output, '[\r\n]*$', '', '')
    execute 'normal i' . l:output
endfunction
nmap <silent> \i :call CJKInput()<CR>
